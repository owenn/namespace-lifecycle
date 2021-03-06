// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceLifecycle) DeepCopyInto(out *NamespaceLifecycle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceLifecycle.
func (in *NamespaceLifecycle) DeepCopy() *NamespaceLifecycle {
	if in == nil {
		return nil
	}
	out := new(NamespaceLifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceLifecycle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceLifecycleList) DeepCopyInto(out *NamespaceLifecycleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceLifecycle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceLifecycleList.
func (in *NamespaceLifecycleList) DeepCopy() *NamespaceLifecycleList {
	if in == nil {
		return nil
	}
	out := new(NamespaceLifecycleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceLifecycleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceLifecycleSpec) DeepCopyInto(out *NamespaceLifecycleSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceLifecycleSpec.
func (in *NamespaceLifecycleSpec) DeepCopy() *NamespaceLifecycleSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceLifecycleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceLifecycleStatus) DeepCopyInto(out *NamespaceLifecycleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceLifecycleStatus.
func (in *NamespaceLifecycleStatus) DeepCopy() *NamespaceLifecycleStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceLifecycleStatus)
	in.DeepCopyInto(out)
	return out
}
