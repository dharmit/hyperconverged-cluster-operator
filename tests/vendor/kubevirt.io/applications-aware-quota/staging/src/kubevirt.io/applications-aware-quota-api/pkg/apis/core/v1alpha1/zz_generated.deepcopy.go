//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The AAQ Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQ) DeepCopyInto(out *AAQ) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQ.
func (in *AAQ) DeepCopy() *AAQ {
	if in == nil {
		return nil
	}
	out := new(AAQ)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AAQ) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQCertConfig) DeepCopyInto(out *AAQCertConfig) {
	*out = *in
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CertConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Server != nil {
		in, out := &in.Server, &out.Server
		*out = new(CertConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQCertConfig.
func (in *AAQCertConfig) DeepCopy() *AAQCertConfig {
	if in == nil {
		return nil
	}
	out := new(AAQCertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQConfiguration) DeepCopyInto(out *AAQConfiguration) {
	*out = *in
	out.VmiCalculatorConfiguration = in.VmiCalculatorConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQConfiguration.
func (in *AAQConfiguration) DeepCopy() *AAQConfiguration {
	if in == nil {
		return nil
	}
	out := new(AAQConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQJobQueueConfig) DeepCopyInto(out *AAQJobQueueConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQJobQueueConfig.
func (in *AAQJobQueueConfig) DeepCopy() *AAQJobQueueConfig {
	if in == nil {
		return nil
	}
	out := new(AAQJobQueueConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AAQJobQueueConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQJobQueueConfigList) DeepCopyInto(out *AAQJobQueueConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AAQJobQueueConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQJobQueueConfigList.
func (in *AAQJobQueueConfigList) DeepCopy() *AAQJobQueueConfigList {
	if in == nil {
		return nil
	}
	out := new(AAQJobQueueConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AAQJobQueueConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQJobQueueConfigSpec) DeepCopyInto(out *AAQJobQueueConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQJobQueueConfigSpec.
func (in *AAQJobQueueConfigSpec) DeepCopy() *AAQJobQueueConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AAQJobQueueConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQJobQueueConfigStatus) DeepCopyInto(out *AAQJobQueueConfigStatus) {
	*out = *in
	if in.PodsInJobQueue != nil {
		in, out := &in.PodsInJobQueue, &out.PodsInJobQueue
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ControllerLock != nil {
		in, out := &in.ControllerLock, &out.ControllerLock
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQJobQueueConfigStatus.
func (in *AAQJobQueueConfigStatus) DeepCopy() *AAQJobQueueConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AAQJobQueueConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQList) DeepCopyInto(out *AAQList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AAQ, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQList.
func (in *AAQList) DeepCopy() *AAQList {
	if in == nil {
		return nil
	}
	out := new(AAQList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AAQList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQSpec) DeepCopyInto(out *AAQSpec) {
	*out = *in
	in.Infra.DeepCopyInto(&out.Infra)
	in.Workloads.DeepCopyInto(&out.Workloads)
	if in.CertConfig != nil {
		in, out := &in.CertConfig, &out.CertConfig
		*out = new(AAQCertConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PriorityClass != nil {
		in, out := &in.PriorityClass, &out.PriorityClass
		*out = new(AAQPriorityClass)
		**out = **in
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Configuration = in.Configuration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQSpec.
func (in *AAQSpec) DeepCopy() *AAQSpec {
	if in == nil {
		return nil
	}
	out := new(AAQSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AAQStatus) DeepCopyInto(out *AAQStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AAQStatus.
func (in *AAQStatus) DeepCopy() *AAQStatus {
	if in == nil {
		return nil
	}
	out := new(AAQStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationsResourceQuota) DeepCopyInto(out *ApplicationsResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationsResourceQuota.
func (in *ApplicationsResourceQuota) DeepCopy() *ApplicationsResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ApplicationsResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationsResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationsResourceQuotaList) DeepCopyInto(out *ApplicationsResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationsResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationsResourceQuotaList.
func (in *ApplicationsResourceQuotaList) DeepCopy() *ApplicationsResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ApplicationsResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationsResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationsResourceQuotaSpec) DeepCopyInto(out *ApplicationsResourceQuotaSpec) {
	*out = *in
	in.ResourceQuotaSpec.DeepCopyInto(&out.ResourceQuotaSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationsResourceQuotaSpec.
func (in *ApplicationsResourceQuotaSpec) DeepCopy() *ApplicationsResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationsResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationsResourceQuotaStatus) DeepCopyInto(out *ApplicationsResourceQuotaStatus) {
	*out = *in
	in.ResourceQuotaStatus.DeepCopyInto(&out.ResourceQuotaStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationsResourceQuotaStatus.
func (in *ApplicationsResourceQuotaStatus) DeepCopy() *ApplicationsResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationsResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterAppsResourceQuota) DeepCopyInto(out *AppliedClusterAppsResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterAppsResourceQuota.
func (in *AppliedClusterAppsResourceQuota) DeepCopy() *AppliedClusterAppsResourceQuota {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterAppsResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterAppsResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppliedClusterAppsResourceQuotaList) DeepCopyInto(out *AppliedClusterAppsResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedClusterAppsResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppliedClusterAppsResourceQuotaList.
func (in *AppliedClusterAppsResourceQuotaList) DeepCopy() *AppliedClusterAppsResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(AppliedClusterAppsResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AppliedClusterAppsResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertConfig) DeepCopyInto(out *CertConfig) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RenewBefore != nil {
		in, out := &in.RenewBefore, &out.RenewBefore
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertConfig.
func (in *CertConfig) DeepCopy() *CertConfig {
	if in == nil {
		return nil
	}
	out := new(CertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppsResourceQuota) DeepCopyInto(out *ClusterAppsResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppsResourceQuota.
func (in *ClusterAppsResourceQuota) DeepCopy() *ClusterAppsResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ClusterAppsResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAppsResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppsResourceQuotaList) DeepCopyInto(out *ClusterAppsResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAppsResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppsResourceQuotaList.
func (in *ClusterAppsResourceQuotaList) DeepCopy() *ClusterAppsResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ClusterAppsResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAppsResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppsResourceQuotaSpec) DeepCopyInto(out *ClusterAppsResourceQuotaSpec) {
	*out = *in
	in.ClusterResourceQuotaSpec.DeepCopyInto(&out.ClusterResourceQuotaSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppsResourceQuotaSpec.
func (in *ClusterAppsResourceQuotaSpec) DeepCopy() *ClusterAppsResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAppsResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAppsResourceQuotaStatus) DeepCopyInto(out *ClusterAppsResourceQuotaStatus) {
	*out = *in
	in.ClusterResourceQuotaStatus.DeepCopyInto(&out.ClusterResourceQuotaStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAppsResourceQuotaStatus.
func (in *ClusterAppsResourceQuotaStatus) DeepCopy() *ClusterAppsResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAppsResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VmiCalculatorConfiguration) DeepCopyInto(out *VmiCalculatorConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmiCalculatorConfiguration.
func (in *VmiCalculatorConfiguration) DeepCopy() *VmiCalculatorConfiguration {
	if in == nil {
		return nil
	}
	out := new(VmiCalculatorConfiguration)
	in.DeepCopyInto(out)
	return out
}
