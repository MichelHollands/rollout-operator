//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneIngesterAutoScaler) DeepCopyInto(out *MultiZoneIngesterAutoScaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneIngesterAutoScaler.
func (in *MultiZoneIngesterAutoScaler) DeepCopy() *MultiZoneIngesterAutoScaler {
	if in == nil {
		return nil
	}
	out := new(MultiZoneIngesterAutoScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiZoneIngesterAutoScaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneIngesterAutoScalerList) DeepCopyInto(out *MultiZoneIngesterAutoScalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MultiZoneIngesterAutoScaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneIngesterAutoScalerList.
func (in *MultiZoneIngesterAutoScalerList) DeepCopy() *MultiZoneIngesterAutoScalerList {
	if in == nil {
		return nil
	}
	out := new(MultiZoneIngesterAutoScalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MultiZoneIngesterAutoScalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneIngesterAutoScalerSpec) DeepCopyInto(out *MultiZoneIngesterAutoScalerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneIngesterAutoScalerSpec.
func (in *MultiZoneIngesterAutoScalerSpec) DeepCopy() *MultiZoneIngesterAutoScalerSpec {
	if in == nil {
		return nil
	}
	out := new(MultiZoneIngesterAutoScalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiZoneIngesterAutoScalerStatus) DeepCopyInto(out *MultiZoneIngesterAutoScalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiZoneIngesterAutoScalerStatus.
func (in *MultiZoneIngesterAutoScalerStatus) DeepCopy() *MultiZoneIngesterAutoScalerStatus {
	if in == nil {
		return nil
	}
	out := new(MultiZoneIngesterAutoScalerStatus)
	in.DeepCopyInto(out)
	return out
}
