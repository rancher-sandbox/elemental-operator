/*
Copyright © 2022 - 2023 SUSE LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	values "github.com/rancher/wrangler/pkg/data"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	elementalv1 "github.com/rancher/elemental-operator/api/v1beta1"
	"github.com/rancher/elemental-operator/pkg/hostinfo"
	"github.com/rancher/elemental-operator/pkg/register"
)

type LegacyConfig struct {
	Elemental   elementalv1.Elemental  `yaml:"elemental"`
	CloudConfig map[string]interface{} `yaml:"cloud-config,omitempty"`
}

var (
	sanitize         = regexp.MustCompile("[^0-9a-zA-Z_]")
	sanitizeHostname = regexp.MustCompile("[^0-9a-zA-Z]")
	doubleDash       = regexp.MustCompile("--+")
	start            = regexp.MustCompile("^[a-zA-Z0-9]")
	errValueNotFound = errors.New("value not found")
)

func (i *InventoryServer) apiRegistration(resp http.ResponseWriter, req *http.Request) error {
	var err error
	var registration *elementalv1.MachineRegistration

	// get the machine registration relevant to this request
	if registration, err = i.getMachineRegistration(path.Base(req.URL.Path)); err != nil {
		http.Error(resp, err.Error(), http.StatusNotFound)
		return err
	}

	if !websocket.IsWebSocketUpgrade(req) {
		logrus.Debug("got a plain HTTP request: send unauthenticated registration")
		if err = i.unauthenticatedResponse(registration, resp); err != nil {
			logrus.Error("error sending unauthenticated response: ", err)
		}
		return err
	}

	// upgrade to websocket
	conn, err := upgrade(resp, req)
	if err != nil {
		logrus.Error("failed to upgrade connection to websocket: %w", err)
		return err
	}
	defer conn.Close()
	logrus.Debug("connection upgraded to websocket")

	// attempt to authenticate the machine, if err, authentication has failed
	inventory, err := i.authMachine(conn, req, registration.Namespace)
	if err != nil {
		logrus.Error("authentication failed: ", err)
		return err
	}
	// no error and no inventory: Auth header is missing or unrecognized
	if inventory == nil {
		logrus.Warn("websocket connection: unrecognized authentication method")
		if err = i.unauthenticatedResponse(registration, resp); err != nil {
			logrus.Error("error sending unauthenticated response: ", err)
		}
		return err
	}
	logrus.Debug("attestation completed")

	if err = register.WriteMessage(conn, register.MsgReady, []byte{}); err != nil {
		logrus.Error("cannot finalize the authentication process: %w", err)
		return err
	}

	isNewInventory := inventory.CreationTimestamp.IsZero()
	if isNewInventory {
		initInventory(inventory, registration)
	}

	if err = i.serveLoop(conn, inventory, registration); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (i *InventoryServer) unauthenticatedResponse(registration *elementalv1.MachineRegistration, writer io.Writer) error {
	if registration.Spec.Config == nil {
		registration.Spec.Config = &elementalv1.Config{}
	}

	mRegistration := registration.Spec.Config.Elemental.Registration

	return yaml.NewEncoder(writer).Encode(elementalv1.Config{
		Elemental: elementalv1.Elemental{
			Registration: elementalv1.Registration{
				URL:             registration.Status.RegistrationURL,
				CACert:          i.getRancherCACert(),
				EmulateTPM:      mRegistration.EmulateTPM,
				EmulatedTPMSeed: mRegistration.EmulatedTPMSeed,
				NoSMBIOS:        mRegistration.NoSMBIOS,
			},
		},
	})
}

func (i *InventoryServer) writeMachineInventoryCloudConfig(conn *websocket.Conn, protoVersion register.MessageType, inventory *elementalv1.MachineInventory, registration *elementalv1.MachineRegistration) error {
	sa := &corev1.ServiceAccount{}

	if err := i.Get(i, types.NamespacedName{
		Name:      registration.Status.ServiceAccountRef.Name,
		Namespace: registration.Status.ServiceAccountRef.Namespace,
	}, sa); err != nil {
		return fmt.Errorf("failed to get service account: %w", err)
	}

	if len(sa.Secrets) == 0 {
		return fmt.Errorf("no secrets associated to the %s/%s service account", sa.Namespace, sa.Name)
	}

	secret := &corev1.Secret{}
	err := i.Get(i, types.NamespacedName{
		Name:      sa.Secrets[0].Name,
		Namespace: sa.Namespace,
	}, secret)

	if err != nil || secret.Type != corev1.SecretTypeServiceAccountToken {
		return fmt.Errorf("failed to get secret: %w", err)
	}

	serverURL, err := i.getValue("server-url")
	if err != nil {
		return fmt.Errorf("failed to get server-url: %w", err)
	}
	if serverURL == "" {
		return fmt.Errorf("server-url is not set")
	}

	if registration.Spec.Config == nil {
		registration.Spec.Config = &elementalv1.Config{}
	}

	elementalConf := elementalv1.Elemental{
		Registration: elementalv1.Registration{
			URL:    registration.Status.RegistrationURL,
			CACert: i.getRancherCACert(),
		},
		SystemAgent: elementalv1.SystemAgent{
			URL:             fmt.Sprintf("%s/k8s/clusters/local", serverURL),
			Token:           string(secret.Data["token"]),
			SecretName:      inventory.Name,
			SecretNamespace: inventory.Namespace,
		},
		Install: registration.Spec.Config.Elemental.Install,
	}

	cloudConf := registration.Spec.Config.CloudConfig

	// legacy register client (old ISO): we should send serialization of legacy (pre-kubebuilder) config.
	if protoVersion == register.MsgUndefined {
		logrus.Debug("Detected old register client: sending legacy CloudConfig serialization.")
		return sendLegacyConfig(conn, elementalConf, cloudConf)
	}

	config := elementalv1.Config{
		Elemental:   elementalConf,
		CloudConfig: cloudConf,
	}

	// If client does not support MsgConfig we send back the config as a
	// byte-stream without a message-type to keep backwards-compatibility.
	if protoVersion < register.MsgConfig {
		logrus.Debug("Client does not support MsgConfig, sending back raw config.")

		writer, err := conn.NextWriter(websocket.BinaryMessage)
		if err != nil {
			return err
		}
		defer writer.Close()

		return yaml.NewEncoder(writer).Encode(config)
	}

	logrus.Debug("Client supports MsgConfig, sending back config.")

	// If the client supports the MsgConfig-message we use that.
	data, err := yaml.Marshal(config)
	if err != nil {
		logrus.Errorf("error marshalling config: %v", err)
		return err
	}

	return register.WriteMessage(conn, register.MsgConfig, data)
}

func (i *InventoryServer) getRancherCACert() string {
	cacert, err := i.getValue("cacerts")
	if err != nil {
		logrus.Errorf("Error getting cacerts: %s", err.Error())
	}

	if cacert == "" {
		if cacert, err = i.getValue("internal-cacerts"); err != nil {
			logrus.Errorf("Error getting internal-cacerts: %s", err.Error())
			return ""
		}
	}
	return cacert
}

func replaceStringData(data map[string]interface{}, name string) (string, error) {
	str := name
	result := &strings.Builder{}
	for {
		i := strings.Index(str, "${")
		if i == -1 {
			result.WriteString(str)
			break
		}
		j := strings.Index(str[i:], "}")
		if j == -1 {
			result.WriteString(str)
			break
		}

		result.WriteString(str[:i])
		obj := values.GetValueN(data, strings.Split(str[i+2:j+i], "/")...)
		if str, ok := obj.(string); ok {
			result.WriteString(str)
		} else {
			return "", errValueNotFound
		}
		str = str[j+i+1:]
	}

	resultStr := sanitize.ReplaceAllString(result.String(), "-")
	resultStr = doubleDash.ReplaceAllString(resultStr, "-")
	if !start.MatchString(resultStr) {
		resultStr = "m" + resultStr
	}
	if len(resultStr) > 58 {
		resultStr = resultStr[:58]
	}
	return resultStr, nil
}

func (i *InventoryServer) serveLoop(conn *websocket.Conn, inventory *elementalv1.MachineInventory, registration *elementalv1.MachineRegistration) error {
	protoVersion := register.MsgUndefined

	for {
		var data []byte

		msgType, data, err := register.ReadMessage(conn)
		if err != nil {
			return fmt.Errorf("websocket communication interrupted: %w", err)
		}
		replyMsgType := register.MsgReady
		replyData := []byte{}

		switch msgType {
		case register.MsgVersion:
			protoVersion, err = decodeProtocolVersion(data)
			if err != nil {
				return fmt.Errorf("failed to negotiate protol version: %w", err)
			}
			logrus.Infof("Negotiated protocol version: %d", protoVersion)
			replyMsgType = register.MsgVersion
			replyData = []byte{byte(protoVersion)}
		case register.MsgSmbios:
			err = updateInventoryFromSMBIOSData(data, inventory, registration)
			if err != nil {
				return fmt.Errorf("failed to extract labels from SMBIOS data: %w", err)
			}
			logrus.Debugf("received SMBIOS data - generated machine name: %s", inventory.Name)
		case register.MsgLabels:
			if err := mergeInventoryLabels(inventory, data); err != nil {
				return err
			}
		case register.MsgGet:
			return i.handleGet(conn, protoVersion, inventory, registration)
		case register.MsgSystemData:
			err = updateInventoryFromSystemData(data, inventory, registration)
			if err != nil {
				return fmt.Errorf("failed to extract labels from system data: %w", err)
			}

		default:
			return fmt.Errorf("got unexpected message: %s", msgType)
		}
		if err := register.WriteMessage(conn, replyMsgType, replyData); err != nil {
			return fmt.Errorf("cannot complete %s exchange", msgType)
		}
	}
}

func (i *InventoryServer) handleGet(conn *websocket.Conn, protoVersion register.MessageType, inventory *elementalv1.MachineInventory, registration *elementalv1.MachineRegistration) error {
	var err error

	inventory, err = i.commitMachineInventory(inventory)
	if err != nil {
		if writeErr := writeError(conn, protoVersion, register.NewErrorMessage(err)); writeErr != nil {
			logrus.Errorf("Error reporting back error to client: %v", writeErr)
		}

		return err
	}

	if err := i.writeMachineInventoryCloudConfig(conn, protoVersion, inventory, registration); err != nil {
		if writeErr := writeError(conn, protoVersion, register.NewErrorMessage(err)); writeErr != nil {
			logrus.Errorf("Error reporting back error to client: %v", writeErr)
		}

		return fmt.Errorf("failed sending elemental cloud config: %w", err)
	}

	logrus.Debug("Elemental cloud config sent")

	return nil
}

func sendLegacyConfig(conn *websocket.Conn, elementalConf elementalv1.Elemental, cloudConf map[string]runtime.RawExtension) (err error) {
	legacyCloudConf := make(map[string]interface{})
	for cloudKey, cloudData := range cloudConf {
		var data interface{}
		if err := json.Unmarshal(cloudData.Raw, &data); err != nil {
			logrus.Warnf("Error unmarshalling key %s: %s", cloudKey, err.Error())
			logrus.Debugf("Unmarshalling of '%s' failed", &cloudData.Raw)
			continue
		}
		legacyCloudConf[cloudKey] = data
	}

	config := LegacyConfig{
		Elemental:   elementalConf,
		CloudConfig: legacyCloudConf,
	}

	writer, err := conn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return err
	}
	defer writer.Close()

	return yaml.NewEncoder(writer).Encode(config)
}

// writeError reports back an error to the client if the negotiated protocol
// version supports it.
func writeError(conn *websocket.Conn, protoVersion register.MessageType, errorMessage register.ErrorMessage) error {
	if protoVersion < register.MsgError {
		logrus.Debugf("client does not support reporting errors, skipping")
		return nil
	}

	data, err := yaml.Marshal(errorMessage)
	if err != nil {
		return fmt.Errorf("error marshalling error-message: %w", err)
	}

	return register.WriteMessage(conn, register.MsgError, data)
}

func decodeProtocolVersion(data []byte) (register.MessageType, error) {
	protoVersion := register.MsgUndefined

	if len(data) != 1 {
		return protoVersion, fmt.Errorf("unknown format: %v (%s)", data, data)
	}
	clientVersion := register.MessageType(data[0])
	protoVersion = clientVersion
	if clientVersion != register.MsgLast {
		logrus.Debugf("elemental-register (%d) and elemental-operator (%d) protocol versions differ", clientVersion, register.MsgLast)
		if clientVersion > register.MsgLast {
			protoVersion = register.MsgLast
		}
	}

	return protoVersion, nil
}

// updateInventoryFromSMBIOSData() updates mInventory Name and Labels from the MachineRegistration and the SMBIOS data
func updateInventoryFromSMBIOSData(data []byte, mInventory *elementalv1.MachineInventory, mRegistration *elementalv1.MachineRegistration) error {
	smbiosData := map[string]interface{}{}
	if err := json.Unmarshal(data, &smbiosData); err != nil {
		return err
	}
	// Sanitize any lower dashes into dashes as hostnames cannot have lower dashes, and we use the inventory name
	// to set the machine hostname. Also set it to lowercase
	name, err := replaceStringData(smbiosData, mInventory.Name)
	if err != nil {
		if errors.Is(err, errValueNotFound) {
			logrus.Warningf("Value not found: %v", mInventory.Name)
			name = "m"
		} else {
			return err
		}
	}

	mInventory.Name = strings.ToLower(sanitizeHostname.ReplaceAllString(name, "-"))
	logrus.Debug("Adding labels from registration")
	// Add extra label info from data coming from smbios and based on the registration data
	if mInventory.Labels == nil {
		mInventory.Labels = map[string]string{}
	}
	for k, v := range mRegistration.Spec.MachineInventoryLabels {
		parsedData, err := replaceStringData(smbiosData, v)
		if err != nil {
			if errors.Is(err, errValueNotFound) {
				logrus.Debugf("Value not found: %v", v)
				continue
			}
			logrus.Errorf("Failed parsing smbios data: %v", err.Error())
			return err
		}

		logrus.Debugf("Parsed %s into %s with smbios data, setting it to label %s", v, parsedData, k)
		mInventory.Labels[k] = strings.TrimSuffix(strings.TrimPrefix(parsedData, "-"), "-")
	}
	return nil
}

// updateInventoryFromSystemData creates labels in the inventory based on the hardware information
func updateInventoryFromSystemData(data []byte, inv *elementalv1.MachineInventory, reg *elementalv1.MachineRegistration) error {
	systemData := &hostinfo.HostInfo{}
	if err := json.Unmarshal(data, &systemData); err != nil {
		return err
	}
	logrus.Info("Adding labels from system data")

	memory := map[string]interface{}{}
	if systemData.Memory != nil {
		memory["Total Physical Bytes"] = strconv.Itoa(int(systemData.Memory.TotalPhysicalBytes))
	}

	// Both checks below is due to ghw not detecting aarch64 cores/threads properly, so it ends up in a label
	// with 0 valuie, which is not useful at all
	// tracking bug: https://github.com/jaypipes/ghw/issues/199
	cpu := map[string]interface{}{}
	if systemData.CPU != nil {
		if systemData.CPU.TotalCores > 0 {
			cpu["Total Cores"] = strconv.Itoa(int(systemData.CPU.TotalCores))
		}

		if systemData.CPU.TotalThreads > 0 {
			cpu["Total Threads"] = strconv.Itoa(int(systemData.CPU.TotalThreads))
		}

		// This should never happen but just in case
		if len(systemData.CPU.Processors) > 0 {
			// Model still looks weird, maybe there is a way of getting it differently as we need to sanitize a lot of data in there?
			// Currently, something like "Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz" ends up being:
			// "Intel-R-Core-TM-i7-7700K-CPU-4-20GHz"
			cpu["Model"] = sanitizeString(systemData.CPU.Processors[0].Model)
			cpu["Vendor"] = sanitizeString(systemData.CPU.Processors[0].Vendor)
			// Capabilities available here at systemData.CPU.Processors[X].Capabilities
		}
	}

	gpu := map[string]interface{}{}
	// This could happen so always check.
	if systemData.GPU != nil && len(systemData.GPU.GraphicsCards) > 0 && systemData.GPU.GraphicsCards[0].DeviceInfo != nil {
		gpu["Model"] = sanitizeString(systemData.GPU.GraphicsCards[0].DeviceInfo.Product.Name)
		gpu["Vendor"] = sanitizeString(systemData.GPU.GraphicsCards[0].DeviceInfo.Vendor.Name)
	}

	network := map[string]interface{}{}
	if systemData.Network != nil {
		network["Number Interfaces"] = strconv.Itoa(len(systemData.Network.NICs))
		for _, iface := range systemData.Network.NICs {
			network[iface.Name] = map[string]interface{}{
				"Name":      iface.Name,
				"IsVirtual": strconv.FormatBool(iface.IsVirtual),
				// Capabilities available here at iface.Capabilities
				// interesting to store anything in here or show it on the docs? Difficult to use it afterwards as its a list...
			}
		}
	}

	block := map[string]interface{}{}
	if systemData.Block != nil {
		block["Number Devices"] = strconv.Itoa(len(systemData.Block.Disks)) // This includes removable devices like cdrom/usb
		for _, disk := range systemData.Block.Disks {
			block[disk.Name] = map[string]interface{}{
				"Size":               strconv.Itoa(int(disk.SizeBytes)),
				"Name":               disk.Name,
				"Drive Type":         disk.DriveType.String(),
				"Storage Controller": disk.StorageController.String(),
				"Removable":          strconv.FormatBool(disk.IsRemovable),
			}
			// Vendor and model also available here, useful?
		}
	}

	labels := map[string]interface{}{}
	labels["System Data"] = map[string]interface{}{
		"Memory":        memory,
		"CPU":           cpu,
		"GPU":           gpu,
		"Network":       network,
		"Block Devices": block,
	}

	// Also available but not used:
	// systemData.Product -> name, vendor, serial,uuid,sku,version. Kind of smbios data
	// systemData.BIOS -> info about the bios. Useless IMO
	// systemData.Baseboard -> asset, serial, vendor,version,product. Kind of useless?
	// systemData.Chassis -> asset, serial, vendor,version,product, type. Maybe be useful depending on the provider.
	// systemData.Topology -> CPU/memory and cache topology. No idea if useful.

	logrus.Debugf("Parsing labels from System Data")

	if inv.Labels == nil {
		inv.Labels = map[string]string{}
	}

	for k, v := range reg.Spec.MachineInventoryLabels {
		logrus.Debugf("Parsing: %v : %v", k, v)

		parsedData, err := replaceStringData(labels, v)
		if err != nil {
			if errors.Is(err, errValueNotFound) {
				logrus.Debugf("Value not found: %v", v)
				continue
			}
			logrus.Errorf("Failed parsing system data: %v", err.Error())
			return err
		}

		logrus.Debugf("Parsed %s into %s with system data, setting it to label %s", v, parsedData, k)
		inv.Labels[k] = strings.TrimSuffix(strings.TrimPrefix(parsedData, "-"), "-")
	}

	return nil
}

// sanitizeString will sanitize a given string by:
// replacing all invalid chars as set on the sanitize regex by dashes
// removing any double dashes resulted from the above method
// removing prefix+suffix if they are a dash
func sanitizeString(s string) string {
	s1 := sanitize.ReplaceAllString(s, "-")
	s2 := doubleDash.ReplaceAllString(s1, "-")
	return strings.TrimSuffix(strings.TrimPrefix(s2, "-"), "-")
}

func mergeInventoryLabels(inventory *elementalv1.MachineInventory, data []byte) error {
	labels := map[string]string{}
	if err := json.Unmarshal(data, &labels); err != nil {
		return fmt.Errorf("cannot extract inventory labels: %w", err)
	}
	logrus.Debugf("received labels: %v", labels)
	logrus.Warn("received labels from registering client: no more supported, skipping")
	if inventory.Labels == nil {
		inventory.Labels = map[string]string{}
	}
	return nil
}
