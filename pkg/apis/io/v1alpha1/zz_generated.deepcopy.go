// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosMonkey) DeepCopyInto(out *ChaosMonkey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosMonkey.
func (in *ChaosMonkey) DeepCopy() *ChaosMonkey {
	if in == nil {
		return nil
	}
	out := new(ChaosMonkey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosMonkey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosMonkeyList) DeepCopyInto(out *ChaosMonkeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosMonkey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosMonkeyList.
func (in *ChaosMonkeyList) DeepCopy() *ChaosMonkeyList {
	if in == nil {
		return nil
	}
	out := new(ChaosMonkeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosMonkeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosMonkeySpec) DeepCopyInto(out *ChaosMonkeySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosMonkeySpec.
func (in *ChaosMonkeySpec) DeepCopy() *ChaosMonkeySpec {
	if in == nil {
		return nil
	}
	out := new(ChaosMonkeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosMonkeyStatus) DeepCopyInto(out *ChaosMonkeyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosMonkeyStatus.
func (in *ChaosMonkeyStatus) DeepCopy() *ChaosMonkeyStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosMonkeyStatus)
	in.DeepCopyInto(out)
	return out
}
