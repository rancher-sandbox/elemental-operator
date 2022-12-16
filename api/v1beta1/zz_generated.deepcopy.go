//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	upgrade_cattle_iov1 "github.com/rancher/system-upgrade-controller/pkg/apis/upgrade.cattle.io/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleTarget) DeepCopyInto(out *BundleTarget) {
	*out = *in
	in.BundleDeploymentOptions.DeepCopyInto(&out.BundleDeploymentOptions)
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterGroupSelector != nil {
		in, out := &in.ClusterGroupSelector, &out.ClusterGroupSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleTarget.
func (in *BundleTarget) DeepCopy() *BundleTarget {
	if in == nil {
		return nil
	}
	out := new(BundleTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.Elemental.DeepCopyInto(&out.Elemental)
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = make(map[string]runtime.RawExtension, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImage) DeepCopyInto(out *ContainerImage) {
	*out = *in
	out.Metadata = in.Metadata
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImage.
func (in *ContainerImage) DeepCopy() *ContainerImage {
	if in == nil {
		return nil
	}
	out := new(ContainerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Elemental) DeepCopyInto(out *Elemental) {
	*out = *in
	in.Install.DeepCopyInto(&out.Install)
	out.Registration = in.Registration
	out.SystemAgent = in.SystemAgent
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Elemental.
func (in *Elemental) DeepCopy() *Elemental {
	if in == nil {
		return nil
	}
	out := new(Elemental)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorMessage) DeepCopyInto(out *ErrorMessage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorMessage.
func (in *ErrorMessage) DeepCopy() *ErrorMessage {
	if in == nil {
		return nil
	}
	out := new(ErrorMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISO) DeepCopyInto(out *ISO) {
	*out = *in
	out.Metadata = in.Metadata
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISO.
func (in *ISO) DeepCopy() *ISO {
	if in == nil {
		return nil
	}
	out := new(ISO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Install) DeepCopyInto(out *Install) {
	*out = *in
	if in.ConfigURLs != nil {
		in, out := &in.ConfigURLs, &out.ConfigURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Install.
func (in *Install) DeepCopy() *Install {
	if in == nil {
		return nil
	}
	out := new(Install)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventory) DeepCopyInto(out *MachineInventory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventory.
func (in *MachineInventory) DeepCopy() *MachineInventory {
	if in == nil {
		return nil
	}
	out := new(MachineInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventoryList) DeepCopyInto(out *MachineInventoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineInventory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventoryList.
func (in *MachineInventoryList) DeepCopy() *MachineInventoryList {
	if in == nil {
		return nil
	}
	out := new(MachineInventoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelector) DeepCopyInto(out *MachineInventorySelector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelector.
func (in *MachineInventorySelector) DeepCopy() *MachineInventorySelector {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventorySelector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorList) DeepCopyInto(out *MachineInventorySelectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineInventorySelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorList.
func (in *MachineInventorySelectorList) DeepCopy() *MachineInventorySelectorList {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventorySelectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorSpec) DeepCopyInto(out *MachineInventorySelectorSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorSpec.
func (in *MachineInventorySelectorSpec) DeepCopy() *MachineInventorySelectorSpec {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorStatus) DeepCopyInto(out *MachineInventorySelectorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make(apiv1beta1.MachineAddresses, len(*in))
		copy(*out, *in)
	}
	if in.MachineInventoryRef != nil {
		in, out := &in.MachineInventoryRef, &out.MachineInventoryRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorStatus.
func (in *MachineInventorySelectorStatus) DeepCopy() *MachineInventorySelectorStatus {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorTemplate) DeepCopyInto(out *MachineInventorySelectorTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorTemplate.
func (in *MachineInventorySelectorTemplate) DeepCopy() *MachineInventorySelectorTemplate {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventorySelectorTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorTemplateList) DeepCopyInto(out *MachineInventorySelectorTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineInventorySelectorTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorTemplateList.
func (in *MachineInventorySelectorTemplateList) DeepCopy() *MachineInventorySelectorTemplateList {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineInventorySelectorTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySelectorTemplateSpec) DeepCopyInto(out *MachineInventorySelectorTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySelectorTemplateSpec.
func (in *MachineInventorySelectorTemplateSpec) DeepCopy() *MachineInventorySelectorTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySelectorTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventorySpec) DeepCopyInto(out *MachineInventorySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventorySpec.
func (in *MachineInventorySpec) DeepCopy() *MachineInventorySpec {
	if in == nil {
		return nil
	}
	out := new(MachineInventorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineInventoryStatus) DeepCopyInto(out *MachineInventoryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Plan != nil {
		in, out := &in.Plan, &out.Plan
		*out = new(PlanStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineInventoryStatus.
func (in *MachineInventoryStatus) DeepCopy() *MachineInventoryStatus {
	if in == nil {
		return nil
	}
	out := new(MachineInventoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRegistration) DeepCopyInto(out *MachineRegistration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRegistration.
func (in *MachineRegistration) DeepCopy() *MachineRegistration {
	if in == nil {
		return nil
	}
	out := new(MachineRegistration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRegistration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRegistrationList) DeepCopyInto(out *MachineRegistrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MachineRegistration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRegistrationList.
func (in *MachineRegistrationList) DeepCopy() *MachineRegistrationList {
	if in == nil {
		return nil
	}
	out := new(MachineRegistrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineRegistrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRegistrationSpec) DeepCopyInto(out *MachineRegistrationSpec) {
	*out = *in
	if in.MachineInventoryLabels != nil {
		in, out := &in.MachineInventoryLabels, &out.MachineInventoryLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.MachineInventoryAnnotations != nil {
		in, out := &in.MachineInventoryAnnotations, &out.MachineInventoryAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(Config)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRegistrationSpec.
func (in *MachineRegistrationSpec) DeepCopy() *MachineRegistrationSpec {
	if in == nil {
		return nil
	}
	out := new(MachineRegistrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineRegistrationStatus) DeepCopyInto(out *MachineRegistrationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceAccountRef != nil {
		in, out := &in.ServiceAccountRef, &out.ServiceAccountRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineRegistrationStatus.
func (in *MachineRegistrationStatus) DeepCopy() *MachineRegistrationStatus {
	if in == nil {
		return nil
	}
	out := new(MachineRegistrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImage) DeepCopyInto(out *ManagedOSImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImage.
func (in *ManagedOSImage) DeepCopy() *ManagedOSImage {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageList) DeepCopyInto(out *ManagedOSImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedOSImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageList.
func (in *ManagedOSImageList) DeepCopy() *ManagedOSImageList {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageSpec) DeepCopyInto(out *ManagedOSImageSpec) {
	*out = *in
	if in.CloudConfig != nil {
		in, out := &in.CloudConfig, &out.CloudConfig
		*out = (*in).DeepCopy()
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Concurrency != nil {
		in, out := &in.Concurrency, &out.Concurrency
		*out = new(int64)
		**out = **in
	}
	if in.Prepare != nil {
		in, out := &in.Prepare, &out.Prepare
		*out = new(upgrade_cattle_iov1.ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Cordon != nil {
		in, out := &in.Cordon, &out.Cordon
		*out = new(bool)
		**out = **in
	}
	if in.Drain != nil {
		in, out := &in.Drain, &out.Drain
		*out = new(upgrade_cattle_iov1.DrainSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UpgradeContainer != nil {
		in, out := &in.UpgradeContainer, &out.UpgradeContainer
		*out = new(upgrade_cattle_iov1.ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterRolloutStrategy != nil {
		in, out := &in.ClusterRolloutStrategy, &out.ClusterRolloutStrategy
		*out = new(v1alpha1.RolloutStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]BundleTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageSpec.
func (in *ManagedOSImageSpec) DeepCopy() *ManagedOSImageSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSImageStatus) DeepCopyInto(out *ManagedOSImageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSImageStatus.
func (in *ManagedOSImageStatus) DeepCopy() *ManagedOSImageStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedOSImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersion) DeepCopyInto(out *ManagedOSVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersion.
func (in *ManagedOSVersion) DeepCopy() *ManagedOSVersion {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionChannel) DeepCopyInto(out *ManagedOSVersionChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionChannel.
func (in *ManagedOSVersionChannel) DeepCopy() *ManagedOSVersionChannel {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSVersionChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionChannelList) DeepCopyInto(out *ManagedOSVersionChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedOSVersionChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionChannelList.
func (in *ManagedOSVersionChannelList) DeepCopy() *ManagedOSVersionChannelList {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSVersionChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionChannelSpec) DeepCopyInto(out *ManagedOSVersionChannelSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]runtime.RawExtension, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.UpgradeContainer != nil {
		in, out := &in.UpgradeContainer, &out.UpgradeContainer
		*out = new(upgrade_cattle_iov1.ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionChannelSpec.
func (in *ManagedOSVersionChannelSpec) DeepCopy() *ManagedOSVersionChannelSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionChannelStatus) DeepCopyInto(out *ManagedOSVersionChannelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastSyncedTime != nil {
		in, out := &in.LastSyncedTime, &out.LastSyncedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionChannelStatus.
func (in *ManagedOSVersionChannelStatus) DeepCopy() *ManagedOSVersionChannelStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionList) DeepCopyInto(out *ManagedOSVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedOSVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionList.
func (in *ManagedOSVersionList) DeepCopy() *ManagedOSVersionList {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedOSVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionSpec) DeepCopyInto(out *ManagedOSVersionSpec) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]runtime.RawExtension, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.UpgradeContainer != nil {
		in, out := &in.UpgradeContainer, &out.UpgradeContainer
		*out = new(upgrade_cattle_iov1.ContainerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionSpec.
func (in *ManagedOSVersionSpec) DeepCopy() *ManagedOSVersionSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedOSVersionStatus) DeepCopyInto(out *ManagedOSVersionStatus) {
	*out = *in
	in.BundleStatus.DeepCopyInto(&out.BundleStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedOSVersionStatus.
func (in *ManagedOSVersionStatus) DeepCopy() *ManagedOSVersionStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedOSVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
	if in.PlanSecretRef != nil {
		in, out := &in.PlanSecretRef, &out.PlanSecretRef
		*out = new(corev1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registration) DeepCopyInto(out *Registration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registration.
func (in *Registration) DeepCopy() *Registration {
	if in == nil {
		return nil
	}
	out := new(Registration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemAgent) DeepCopyInto(out *SystemAgent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemAgent.
func (in *SystemAgent) DeepCopy() *SystemAgent {
	if in == nil {
		return nil
	}
	out := new(SystemAgent)
	in.DeepCopyInto(out)
	return out
}
