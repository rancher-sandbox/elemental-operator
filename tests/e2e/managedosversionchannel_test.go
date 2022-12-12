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

package e2e_test

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	provv1 "github.com/rancher/elemental-operator/pkg/apis/elemental.cattle.io/v1beta1"
	fleet "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	http "github.com/rancher-sandbox/ele-testhelpers/http"
	kubectl "github.com/rancher-sandbox/ele-testhelpers/kubectl"

	"github.com/rancher/elemental-operator/tests/catalog"
)

var _ = Describe("ManagedOSVersionChannel e2e tests", func() {
	var k *kubectl.Kubectl

	AfterEach(func() {
		err := k.Delete("managedosversionchannel", "--all", "--force", "--wait", "-n", fleetNamespace)
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Create ManagedOSVersions", func() {
		BeforeEach(func() {
			k = kubectl.New()
		})
		It("Reports failure events", func() {

			By("Create an invalid ManagedOSVersionChannel")
			ui := catalog.NewManagedOSVersionChannel(
				"invalid",
				"",
				"",
				map[string]interface{}{"uri": "http://" + e2eCfg.BridgeIP + ":9999"},
				nil,
			)

			err := k.ApplyYAML(fleetNamespace, "invalid", ui)
			Expect(err).ShouldNot(HaveOccurred())
			defer k.Delete("managedosversionchannel", "-n", fleetNamespace, "invalid")

			By("Check that reports event failure")
			Eventually(func() string {
				r, _ := kubectl.Run("describe", "-n", fleetNamespace, "managedosversionchannel", "invalid")

				return r
			}, 1*time.Minute, 2*time.Second).Should(
				ContainSubstring("spec.Type can't be empty"),
			)
		})

		It("creates a list of ManagedOSVersion from a JSON server", func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			versions := []provv1.ManagedOSVersion{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "v1"},
					Spec: provv1.ManagedOSVersionSpec{
						Version:    "v1",
						Type:       "container",
						MinVersion: "0.0.0",
						Metadata: &fleet.GenericMap{
							Data: map[string]interface{}{
								"upgradeImage": "registry.com/repository/image:v1",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "v2"},
					Spec: provv1.ManagedOSVersionSpec{
						Version:    "v2",
						Type:       "container",
						MinVersion: "0.0.0",
						Metadata: &fleet.GenericMap{
							Data: map[string]interface{}{
								"upgradeImage": "registry.com/repository/image:v2",
							},
						},
					},
				},
			}

			b, err := json.Marshal(versions)
			Expect(err).ShouldNot(HaveOccurred())

			http.Server(ctx, e2eCfg.BridgeIP+":9999", string(b))

			By("Create a ManagedOSVersionChannel")
			ui := catalog.NewManagedOSVersionChannel(
				"testchannel",
				"json",
				"10m",
				map[string]interface{}{"uri": "http://" + e2eCfg.BridgeIP + ":9999"},
				nil,
			)

			err = k.ApplyYAML(fleetNamespace, "testchannel", ui)
			Expect(err).ShouldNot(HaveOccurred())
			defer k.Delete("managedosversionchannel", "-n", fleetNamespace, "testchannel")

			r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersionChannel", "testchannel", `jsonpath={.spec.type}`)

			Expect(string(r)).To(Equal("json"))

			By("Check new ManagedOSVersions are created")
			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersion", "v1", `jsonpath={.spec.metadata.upgradeImage}`)
				return string(r)
			}, 5*time.Minute, 2*time.Second).Should(
				Equal("registry.com/repository/image:v1"),
			)

			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersion", "v2", `jsonpath={.spec.metadata.upgradeImage}`)

				return string(r)
			}, 1*time.Minute, 2*time.Second).Should(
				Equal("registry.com/repository/image:v2"),
			)

			err = k.Delete("managedosversionchannel", "-n", fleetNamespace, "testchannel")
			Expect(err).ShouldNot(HaveOccurred())

			By("Check ManagedOSVersions are deleted on channel clean up")
			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersion", "v2", `jsonpath={}`)

				return string(r)
			}, 1*time.Minute, 2*time.Second).Should(
				Equal(""),
			)
		})

		It("creates a list of ManagedOSVersion from a custom hook", func() {

			versions := []provv1.ManagedOSVersion{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "foo"},
					Spec: provv1.ManagedOSVersionSpec{
						Version:    "v1",
						Type:       "container",
						MinVersion: "0.0.0",
						Metadata: &fleet.GenericMap{
							Data: map[string]interface{}{
								"upgradeImage": "registry.com/repository/image:v1",
							},
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "bar"},
					Spec: provv1.ManagedOSVersionSpec{
						Version:    "v2",
						Type:       "container",
						MinVersion: "0.0.0",
						Metadata: &fleet.GenericMap{
							Data: map[string]interface{}{
								"upgradeImage": "registry.com/repository/image:v2",
							},
						},
					},
				},
			}

			b, err := json.Marshal(versions)
			Expect(err).ShouldNot(HaveOccurred())

			By("Create a ManagedOSVersionChannel")
			ui := catalog.NewManagedOSVersionChannel(
				"testchannel2",
				"custom",
				"10m",
				map[string]interface{}{
					"image":      "opensuse/tumbleweed",
					"command":    []string{"/bin/bash", "-c", "--"},
					"mountPath":  "/output",      // This defaults to /data
					"outputFile": "/output/data", // This defaults to /data/output
					"args":       []string{fmt.Sprintf("echo '%s' > /output/data", string(b))},
				},
				nil,
			)

			err = k.ApplyYAML(fleetNamespace, "testchannel2", ui)
			Expect(err).ShouldNot(HaveOccurred())
			defer k.Delete("managedosversionchannel", "-n", fleetNamespace, "testchannel2")

			r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersionChannel", "testchannel2", `jsonpath={.spec.type}`)

			Expect(string(r)).To(Equal("custom"))

			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersion", "foo", `jsonpath={.spec.metadata.upgradeImage}`)
				return string(r)
			}, 2*time.Minute, 2*time.Second).Should(
				Equal("registry.com/repository/image:v1"),
			)

			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersion", "bar", `jsonpath={.spec.metadata.upgradeImage}`)
				return string(r)
			}, 2*time.Minute, 2*time.Second).Should(
				Equal("registry.com/repository/image:v2"),
			)
		})

		It("on a broken a channel it stops on failed sync ready reason", func() {

			By("Create a ManagedOSVersionChannel with wrong content")
			ui := catalog.NewManagedOSVersionChannel(
				"testchannel2",
				"custom",
				"10m",
				map[string]interface{}{
					"image":   "opensuse/tumbleweed",
					"command": []string{"/bin/bash", "-c", "--"},
					"args":    []string{fmt.Sprintf("echo '%s' > /output/data", string("wrong content"))},
				},
				nil,
			)

			err := k.ApplyYAML(fleetNamespace, "testchannel2", ui)
			Expect(err).ShouldNot(HaveOccurred())
			defer k.Delete("managedosversionchannel", "-n", fleetNamespace, "testchannel2")

			r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersionChannel", "testchannel2", `jsonpath={.spec.type}`)

			Expect(string(r)).To(Equal("custom"))

			Eventually(func() string {
				r, _ := kubectl.GetData(fleetNamespace, "ManagedOSVersionChannel", "testchannel2", `jsonpath={.status.conditions[0].status}`)
				fmt.Println(string(r))
				return string(r)
			}, 2*time.Minute, 2*time.Second).Should(
				Equal("True"),
			)
		})
	})
})
