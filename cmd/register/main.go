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

package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twpayne/go-vfs/v4"
	"gopkg.in/yaml.v3"

	elementalv1 "github.com/rancher/elemental-operator/api/v1beta1"
	"github.com/rancher/elemental-operator/pkg/install"
	"github.com/rancher/elemental-operator/pkg/log"
	"github.com/rancher/elemental-operator/pkg/register"
	"github.com/rancher/elemental-operator/pkg/version"
)

const (
	defaultStatePath                = "/oem/registration/state.yaml"
	defaultConfigPath               = "/oem/registration/config.yaml"
	defaultLiveConfigPath           = "/run/initramfs/live/livecd-cloud-config.yaml"
	registrationUpdateSuppressTimer = 24 * time.Hour
)

var (
	cfg        elementalv1.Config
	debug      bool
	configPath string
	statePath  string
)

var (
	errEmptyRegistrationURL = errors.New("registration URL is empty")
)

func main() {
	cmd := newCommand(vfs.OSFS)
	if err := cmd.Execute(); err != nil {
		log.Fatalf("FATAL: %s", err)
	}
}

func newCommand(fs vfs.FS) *cobra.Command {
	installer := install.NewInstaller(fs)
	cmd := &cobra.Command{
		Use:   "elemental-register",
		Short: "Elemental register command",
		Long:  "elemental-register registers a node with the elemental-operator via a config file or flags",
		RunE: func(_ *cobra.Command, args []string) error {
			// Initialize Config
			initConfig()
			if err := viper.Unmarshal(&cfg); err != nil {
				return fmt.Errorf("decoding configuration: %w", err)
			}
			// Version subcommand
			if viper.GetBool("version") {
				log.Infof("Register version %s, commit %s, commit date %s", version.Version, version.Commit, version.CommitDate)
				return nil
			}
			// Determine if registration should execute or skip a cycle
			stateHandler := register.NewFileStateHandler(fs, statePath)
			if skip, err := shouldSkipRegistration(stateHandler, installer); err != nil {
				return fmt.Errorf("determining if registration should run: %w", err)
			} else if skip {
				log.Info("Nothing to do")
				return nil
			}
			// Validate CA
			caCert, err := getRegistrationCA(fs, cfg)
			if err != nil {
				return fmt.Errorf("validating CA: %w", err)
			}
			// Register
			client := register.NewClient(stateHandler)
			data, err := client.Register(cfg.Elemental.Registration, caCert)
			if err != nil {
				return fmt.Errorf("registering machine: %w", err)
			}
			// Validate remote config
			log.Debugf("Fetched configuration from manager cluster:\n%s\n\n", string(data))
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return fmt.Errorf("parsing returned configuration: %w", err)
			}
			// Install
			if !installer.IsSystemInstalled() {
				log.Info("Installing Elemental")
				return installer.InstallElemental(cfg)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&cfg.Elemental.Registration.URL, "registration-url", "", "Registration url to get the machine config from")
	cmd.Flags().StringVar(&cfg.Elemental.Registration.CACert, "registration-ca-cert", "", "File with the custom CA certificate to use against he registration url")
	cmd.Flags().BoolVar(&cfg.Elemental.Registration.EmulateTPM, "emulate-tpm", false, "Emulate /dev/tpm")
	cmd.Flags().Int64Var(&cfg.Elemental.Registration.EmulatedTPMSeed, "emulated-tpm-seed", 1, "Seed for /dev/tpm emulation")
	cmd.Flags().BoolVar(&cfg.Elemental.Registration.NoSMBIOS, "no-smbios", false, "Disable the use of dmidecode to get SMBIOS")
	cmd.Flags().StringVar(&cfg.Elemental.Registration.Auth, "auth", "tpm", "Registration authentication method")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	cmd.Flags().StringVar(&configPath, "config-path", "", "The full path of the elemental-register config")
	cmd.Flags().StringVar(&statePath, "state-path", "", "The full path of the elemental-register config")
	cmd.PersistentFlags().BoolP("version", "v", false, "print version and exit")
	_ = viper.BindPFlag("version", cmd.PersistentFlags().Lookup("version"))
	//Set Defaults
	viper.SetDefault("state-path", defaultStatePath)
	if installer.IsSystemInstalled() {
		viper.SetDefault("config-path", defaultConfigPath)
	} else {
		viper.SetDefault("config-path", defaultLiveConfigPath)
	}
	return cmd
}

func initConfig() {
	if debug {
		log.EnableDebugLogging()
	}

	log.Infof("Register version %s, commit %s, commit date %s", version.Version, version.Commit, version.CommitDate)

	viper.SetConfigFile(configPath)
}

func shouldSkipRegistration(stateHandler register.StateHandler, installer install.Installer) (bool, error) {
	if !installer.IsSystemInstalled() {
		return false, nil
	}
	state, err := stateHandler.Load()
	if err != nil {
		return false, fmt.Errorf("loading registration state")
	}
	return state.HasLastUpdateElapsed(registrationUpdateSuppressTimer), nil
}

func getRegistrationCA(fs vfs.FS, config elementalv1.Config) ([]byte, error) {
	registration := config.Elemental.Registration

	if registration.URL == "" {
		return nil, errEmptyRegistrationURL
	}
	/* Here we can have a file path or the cert data itself */
	if _, err := fs.Stat(registration.CACert); err == nil {
		log.Info("CACert passed as a file")
		return fs.ReadFile(registration.CACert)
	}
	if registration.CACert == "" {
		log.Warning("CACert is empty")
	}
	return []byte(registration.CACert), nil
}
