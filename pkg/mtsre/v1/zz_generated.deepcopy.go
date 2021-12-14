//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleParameters) DeepCopyInto(out *BundleParameters) {
	*out = *in
	if in.UseClusterStorage != nil {
		in, out := &in.UseClusterStorage, &out.UseClusterStorage
		*out = new(string)
		**out = **in
	}
	if in.AlertingEmailAddress != nil {
		in, out := &in.AlertingEmailAddress, &out.AlertingEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.BuAlertingEmailAddress != nil {
		in, out := &in.BuAlertingEmailAddress, &out.BuAlertingEmailAddress
		*out = new(string)
		**out = **in
	}
	if in.AlertSMTPFrom != nil {
		in, out := &in.AlertSMTPFrom, &out.AlertSMTPFrom
		*out = new(string)
		**out = **in
	}
	if in.AddonParamsSecretName != nil {
		in, out := &in.AddonParamsSecretName, &out.AddonParamsSecretName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleParameters.
func (in *BundleParameters) DeepCopy() *BundleParameters {
	if in == nil {
		return nil
	}
	out := new(BundleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeadmansSnitch) DeepCopyInto(out *DeadmansSnitch) {
	*out = *in
	if in.ClusterDeploymentSelector != nil {
		in, out := &in.ClusterDeploymentSelector, &out.ClusterDeploymentSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SnitchNamePostFix != nil {
		in, out := &in.SnitchNamePostFix, &out.SnitchNamePostFix
		*out = new(string)
		**out = **in
	}
	if in.TargetSecretRef != nil {
		in, out := &in.TargetSecretRef, &out.TargetSecretRef
		*out = new(TargetSecretRef)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeadmansSnitch.
func (in *DeadmansSnitch) DeepCopy() *DeadmansSnitch {
	if in == nil {
		return nil
	}
	out := new(DeadmansSnitch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvItem) DeepCopyInto(out *EnvItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvItem.
func (in *EnvItem) DeepCopy() *EnvItem {
	if in == nil {
		return nil
	}
	out := new(EnvItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Monitoring) DeepCopyInto(out *Monitoring) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Monitoring.
func (in *Monitoring) DeepCopy() *Monitoring {
	if in == nil {
		return nil
	}
	out := new(Monitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PagerDuty) DeepCopyInto(out *PagerDuty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PagerDuty.
func (in *PagerDuty) DeepCopy() *PagerDuty {
	if in == nil {
		return nil
	}
	out := new(PagerDuty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionConfig) DeepCopyInto(out *SubscriptionConfig) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = new([]EnvItem)
		if **in != nil {
			in, out := *in, *out
			*out = make([]EnvItem, len(*in))
			copy(*out, *in)
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionConfig.
func (in *SubscriptionConfig) DeepCopy() *SubscriptionConfig {
	if in == nil {
		return nil
	}
	out := new(SubscriptionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetSecretRef) DeepCopyInto(out *TargetSecretRef) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetSecretRef.
func (in *TargetSecretRef) DeepCopy() *TargetSecretRef {
	if in == nil {
		return nil
	}
	out := new(TargetSecretRef)
	in.DeepCopyInto(out)
	return out
}
