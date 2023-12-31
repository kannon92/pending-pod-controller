//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingConditionsEntry) DeepCopyInto(out *PendingConditionsEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingConditionsEntry.
func (in *PendingConditionsEntry) DeepCopy() *PendingConditionsEntry {
	if in == nil {
		return nil
	}
	out := new(PendingConditionsEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingContainerStatusEntry) DeepCopyInto(out *PendingContainerStatusEntry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingContainerStatusEntry.
func (in *PendingContainerStatusEntry) DeepCopy() *PendingContainerStatusEntry {
	if in == nil {
		return nil
	}
	out := new(PendingContainerStatusEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingHandler) DeepCopyInto(out *PendingHandler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingHandler.
func (in *PendingHandler) DeepCopy() *PendingHandler {
	if in == nil {
		return nil
	}
	out := new(PendingHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PendingHandler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingHandlerList) DeepCopyInto(out *PendingHandlerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PendingHandler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingHandlerList.
func (in *PendingHandlerList) DeepCopy() *PendingHandlerList {
	if in == nil {
		return nil
	}
	out := new(PendingHandlerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PendingHandlerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingHandlerSpec) DeepCopyInto(out *PendingHandlerSpec) {
	*out = *in
	if in.PendingConditions != nil {
		in, out := &in.PendingConditions, &out.PendingConditions
		*out = make([]PendingConditionsEntry, len(*in))
		copy(*out, *in)
	}
	if in.PendingContainerStatuses != nil {
		in, out := &in.PendingContainerStatuses, &out.PendingContainerStatuses
		*out = make([]PendingContainerStatusEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingHandlerSpec.
func (in *PendingHandlerSpec) DeepCopy() *PendingHandlerSpec {
	if in == nil {
		return nil
	}
	out := new(PendingHandlerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PendingHandlerStatus) DeepCopyInto(out *PendingHandlerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PendingHandlerStatus.
func (in *PendingHandlerStatus) DeepCopy() *PendingHandlerStatus {
	if in == nil {
		return nil
	}
	out := new(PendingHandlerStatus)
	in.DeepCopyInto(out)
	return out
}
