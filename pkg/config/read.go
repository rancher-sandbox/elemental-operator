/*
Copyright © 2022 SUSE LLC

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

package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rancher-sandbox/go-tpm"
	"github.com/rancher/elemental-operator/pkg/dmidecode"
	values "github.com/rancher/wrangler/pkg/data"
	"github.com/rancher/wrangler/pkg/data/convert"
	schemas2 "github.com/rancher/wrangler/pkg/schemas"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

const authorizedKeysKey = "ssh_authorized_keys"

var (
	defaultMappers = schemas2.Mappers{
		NewToMap(),
		NewToSlice(),
		NewToBool(),
		&FuzzyNames{},
	}
	schemas = schemas2.EmptySchemas().Init(func(s *schemas2.Schemas) *schemas2.Schemas {
		s.AddMapper("config", defaultMappers)
		s.AddMapper("elemental", defaultMappers)
		s.AddMapper("install", defaultMappers)
		return s
	}).MustImport(Config{})
	schema = schemas.Schema("config")
)

// ToEnv converts the config into a slice env.
// The configuration fields are prefixed with "_COS"
// to allow installation parameters to be set in the cos.sh script:
// e.g. https://github.com/rancher-sandbox/cOS-toolkit/blob/affc831b76d50298bbbbe637f31c81c52c5489b8/packages/backports/installer/cos.sh#L698
func ToEnv(cfg Config) ([]string, error) {
	// Pass only the elemental values, as those are the ones used by the installer/cli
	// Ignore the Data as that is cloud-config stuff
	data, err := convert.EncodeToMap(&cfg.Elemental)
	if err != nil {
		return nil, err
	}

	return mapToEnv("ELEMENTAL_", data), nil
}

// it's a mapping of how config env option should be transliterated to the elemental CLI
var defaultOverrides = map[string]string{
	"ELEMENTAL_INSTALL_CONFIG_URL": "ELEMENTAL_INSTALL_CLOUD_INIT",
	"ELEMENTAL_INSTALL_POWEROFF":   "ELEMENTAL_POWEROFF",
	"ELEMENTAL_INSTALL_DEVICE":     "ELEMENTAL_INSTALL_TARGET",
	"ELEMENTAL_INSTALL_SYSTEM_URI": "ELEMENTAL_INSTALL_SYSTEM",
	"ELEMENTAL_INSTALL_DEBUG":      "ELEMENTAL_DEBUG",
}

func envOverrides(keyName string) string {
	for k, v := range defaultOverrides {
		keyName = strings.ReplaceAll(keyName, k, v)
	}
	return keyName
}

func mapToEnv(prefix string, data map[string]interface{}) []string {
	var result []string
	for k, v := range data {
		keyName := strings.ToUpper(prefix + convert.ToYAMLKey(k))
		keyName = strings.ReplaceAll(keyName, "-", "_")
		// Apply overrides needed to convert between configs types
		keyName = envOverrides(keyName)

		if data, ok := v.(map[string]interface{}); ok {
			subResult := mapToEnv(keyName+"_", data)
			result = append(result, subResult...)
		} else {
			result = append(result, fmt.Sprintf("%s=%v", keyName, v))
		}
	}
	return result
}

func readFileFunc(path string) func() (map[string]interface{}, error) {
	return func() (map[string]interface{}, error) {
		return readFile(path)
	}
}

func readNested(data map[string]interface{}, overlay bool) (map[string]interface{}, error) {
	var (
		nestedConfigFiles = convert.ToStringSlice(values.GetValueN(data, "elemental", "install", "configUrl"))
		funcs             []reader
	)

	if overlay {
		funcs = append(funcs, func() (map[string]interface{}, error) {
			return data, nil
		})
	}

	for _, nestedConfigFile := range nestedConfigFiles {
		funcs = append(funcs, readFileFunc(nestedConfigFile))
	}

	if !overlay {
		funcs = append(funcs, func() (map[string]interface{}, error) {
			return data, nil
		})
	}

	return merge(funcs...)
}

func readFile(path string) (result map[string]interface{}, _ error) {
	result = map[string]interface{}{}

	switch {
	case strings.HasPrefix(path, "http://"):
		fallthrough
	case strings.HasPrefix(path, "https://"):
		resp, err := http.Get(path)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		buffer, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}

		return result, yaml.Unmarshal(buffer, &result)
	case strings.HasPrefix(path, "tftp://"):
		return tftpGet(path)
	}

	f, err := ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	if err := yaml.Unmarshal(f, &data); err != nil {
		return nil, err
	}

	return readNested(data, false)
}

type reader func() (map[string]interface{}, error)

func merge(readers ...reader) (map[string]interface{}, error) {
	d := map[string]interface{}{}
	for _, r := range readers {
		newData, err := r()
		if err != nil {
			return nil, err
		}
		if err := schema.Mapper.ToInternal(newData); err != nil {
			return nil, err
		}
		d = values.MergeMapsConcatSlice(d, newData)
	}
	return d, nil
}

func readConfigMap(ctx context.Context, cfg string, includeCmdline bool) (map[string]interface{}, error) {
	var (
		data map[string]interface{}
		err  error
	)

	if includeCmdline {
		data, err = merge(readCmdline, readFileFunc(cfg))
		if err != nil {
			return nil, err
		}
	} else {
		data, err = merge(readFileFunc(cfg))
		if err != nil {
			return nil, err
		}
	}

	if cfg != "" {
		values.PutValue(data, cfg, "elemental", "install", "configUrl")
	}

	return updateData(ctx, data)
}

func updateData(ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	registrationURL := convert.ToString(values.GetValueN(data, "elemental", "registration", "url"))
	registrationCA := convert.ToString(values.GetValueN(data, "elemental", "registration", "ca_cert"))
	emulatedTPM := convert.ToString(values.GetValueN(data, "elemental", "registration", "emulateTPM"))
	emulatedSeed := convert.ToString(values.GetValueN(data, "elemental", "registration", "emulatedTPMSeed"))
	noSmbios := convert.ToString(values.GetValueN(data, "elemental", "registration", "noSMBIOS"))
	isoURL := convert.ToString(values.GetValueN(data, "elemental", "install", "iso"))
	ContainerImage := convert.ToString(values.GetValueN(data, "elemental", "install", "system-uri"))
	// Can't set both values at the same time
	if isoURL != "" && ContainerImage != "" {
		return nil, fmt.Errorf("cant set both elemental.install.iso and elemental.install.system-uri in /proc/cmdline or in MachineRegistration both .spec.cloudConfig.elemental.install.iso and .spec.cloudConfig.elemental.install.system-uri")
	}

	if registrationURL != "" {
		for {
			select {
			case <-ctx.Done():
				return data, nil
			default:
				newData, err := returnRegistrationData(registrationURL, registrationCA, emulatedTPM, emulatedSeed, noSmbios)
				if err == nil {
					newISOURL := convert.ToString(values.GetValueN(newData, "elemental", "install", "iso"))
					newContainerImage := convert.ToString(values.GetValueN(newData, "elemental", "install", "system-uri"))
					// If values from registration server are empty, override the old values into the new data
					if newISOURL == "" {
						values.PutValue(newData, isoURL, "elemental", "install", "iso")
					}
					if newContainerImage == "" {
						values.PutValue(newData, ContainerImage, "elemental", "install", "system-uri")
					}
					// Return the data obtained from the registration url with our full data
					return newData, nil
				}
				logrus.Errorf("failed to read registration URL %s, retrying: %v", registrationURL, err)
				time.Sleep(15 * time.Second)
			}
		}
	}

	return data, nil
}

func ToFile(cfg Config, output string) error {
	data, err := ToBytes(cfg)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(output, data, 0600)
}

func ToBytes(cfg Config) ([]byte, error) {
	var (
		data map[string]interface{}
		err  error
	)
	if len(cfg.Data) > 0 {
		data = values.MergeMaps(nil, cfg.Data)
	} else {
		data, err = convert.EncodeToMap(cfg)
		if err != nil {
			return nil, err
		}
		if len(cfg.Elemental.Install.SSHKeys) > 0 {
			data[authorizedKeysKey] = cfg.Elemental.Install.SSHKeys
		}
	}
	values.RemoveValue(data, "install")
	values.RemoveValue(data, "elemental", "install")
	values.RemoveValue(data, "elemental", "registration")
	values.RemoveValue(data, "elemental", "system_agent")
	values.RemoveValue(data, "data") // Remove data here, as we are merging the data into the cfg directly, we want to output a valid cloud-config
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}

	return append([]byte("#cloud-config\n"), bytes...), nil
}

func ReadConfig(ctx context.Context, cfg string, includeCmdline bool) (result Config, err error) {
	data, err := readConfigMap(ctx, cfg, includeCmdline)
	if err != nil {
		return result, err
	}

	if err := convert.ToObj(data, &result); err != nil {
		return result, err
	}

	result.Data = data
	return result, nil
}

func returnRegistrationData(url, ca, emulatedTPM, TPMSeed, noSmbios string) (map[string]interface{}, error) {
	opts := []tpm.Option{tpm.WithCAs([]byte(ca)), tpm.AppendCustomCAToSystemCA}

	// can be explicitly disabled for testing purposes
	if noSmbios != "true" {
		smbios, err := getSMBiosHeaders()
		if err != nil {
			return nil, err
		}
		opts = append(opts, tpm.WithHeader(smbios))
	}

	if emulatedTPM == "true" {
		logrus.Info("TPM Emulation enabled")
		opts = append(opts, tpm.Emulated)
		if TPMSeed != "" {
			s, err := strconv.ParseInt(TPMSeed, 10, 64)
			if err == nil {
				logrus.Infof("TPM Emulation Seed '%s'", TPMSeed)
				opts = append(opts, tpm.WithSeed(s))
			}
		} else {
			opts = append(opts, tpm.WithSeed(1))
		}
	}

	data, err := tpm.Get(url, opts...)
	if err != nil {
		return nil, err
	}

	logrus.Infof("Retrieved config from registrationURL: %s", data)
	result := map[string]interface{}{}
	return result, json.Unmarshal(data, &result)
}

func getSMBiosHeaders() (http.Header, error) {
	smbios, err := dmidecode.Decode()
	if err != nil {
		return nil, err
	}
	smbiosData, err := json.Marshal(smbios)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Set("X-Cattle-Smbios", base64.StdEncoding.EncodeToString(smbiosData))
	return header, nil
}

func readCmdline() (map[string]interface{}, error) {
	//supporting regex https://regexr.com/4mq0s
	parser, err := regexp.Compile(`(\"[^\"]+\")|([^\s]+=(\"[^\"]+\")|([^\s]+))`)
	if err != nil {
		return nil, nil
	}

	procCmdLine := os.Getenv("PROC_CMDLINE")
	if procCmdLine == "" {
		procCmdLine = "/proc/cmdline"
	}
	bytes, err := ioutil.ReadFile(procCmdLine)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	data := map[string]interface{}{}
	for _, item := range parser.FindAllString(string(bytes), -1) {
		parts := strings.SplitN(item, "=", 2)
		value := "true"
		if len(parts) > 1 {
			value = strings.Trim(parts[1], `"`)
		}
		keys := strings.Split(strings.Trim(parts[0], `"`), ".")
		existing, ok := values.GetValue(data, keys...)
		if ok {
			switch v := existing.(type) {
			case string:
				values.PutValue(data, []string{v, value}, keys...)
			case []string:
				values.PutValue(data, append(v, value), keys...)
			}
		} else {
			values.PutValue(data, value, keys...)
		}
	}

	if err := schema.Mapper.ToInternal(data); err != nil {
		return nil, err
	}

	return readNested(data, true)
}
