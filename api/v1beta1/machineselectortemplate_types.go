/*
Copyright © 2022 - 2025 SUSE LLC

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

type MachineInventorySelectorTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineInventorySelectorTemplateSpec `json:"spec,omitempty"`
}

type MachineInventorySelectorTemplateSpec struct {
	// Template machine inventory selector template.
	Template MachineInventorySelector `json:"template"`
}

// +kubebuilder:object:root=true

// MachineInventorySelectorTemplateList contains a list of MachineInventorySelectorTemplates.
type MachineInventorySelectorTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineInventorySelectorTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MachineInventorySelectorTemplate{}, &MachineInventorySelectorTemplateList{})
}
