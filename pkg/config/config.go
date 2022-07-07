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

type RancherOS struct {
	Install Install `json:"install,omitempty"`
}

type Install struct {
	Automatic          bool   `json:"automatic,omitempty"`
	Firmware           string `json:"firmware,omitempty"`
	Device             string `json:"device,omitempty"`
	ConfigURL          string `json:"configUrl,omitempty"`
	ISOURL             string `json:"isoUrl,omitempty"`
	ContainerImage     string `json:"containerImage,omitempty"`
	PowerOff           bool   `json:"powerOff,omitempty"`
	Reboot             bool   `json:"reboot,omitempty"`
	NoFormat           bool   `json:"noFormat,omitempty"`
	Debug              bool   `json:"debug,omitempty"`
	TTY                string `json:"tty,omitempty"`
	ServerURL          string `json:"-"`
	Token              string `json:"-"`
	Role               string `json:"-"`
	Password           string `json:"password,omitempty"`
	RegistrationURL    string `json:"registrationUrl,omitempty"`
	RegistrationCACert string `json:"registrationCaCert,omitempty"`
	EjectCD            bool   `json:"ejectCD,omitempty"`
}

type Config struct {
	SSHAuthorizedKeys []string               `json:"ssh_authorized_keys,omitempty"`
	RancherOS         RancherOS              `json:"rancheros,omitempty"`
	Data              map[string]interface{} `json:"-"`
}

type YipConfig struct {
	Stages   map[string][]Stage `json:"stages,omitempty"`
	Rancherd Rancherd           `json:"rancherd,omitempty"`
}

type Stage struct {
	Users map[string]User `json:"users,omitempty"`
}

type Rancherd struct {
	Server string `json:"server,omitempty"`
	Role   string `json:"role,omitempty"`
	Token  string `json:"token,omitempty"`
}

type User struct {
	Name              string   `json:"name,omitempty"`
	PasswordHash      string   `json:"passwd,omitempty"`
	SSHAuthorizedKeys []string `json:"ssh_authorized_keys,omitempty"`
}
