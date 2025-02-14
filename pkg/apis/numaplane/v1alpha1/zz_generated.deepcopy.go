//go:build !ignore_autogenerated

/*
Copyright 2024.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Controller) DeepCopyInto(out *Controller) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Controller.
func (in *Controller) DeepCopy() *Controller {
	if in == nil {
		return nil
	}
	out := new(Controller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerDefinitions) DeepCopyInto(out *ControllerDefinitions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerDefinitions.
func (in *ControllerDefinitions) DeepCopy() *ControllerDefinitions {
	if in == nil {
		return nil
	}
	out := new(ControllerDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISBServiceProgressiveStatus) DeepCopyInto(out *ISBServiceProgressiveStatus) {
	*out = *in
	if in.UpgradingISBServiceStatus != nil {
		in, out := &in.UpgradingISBServiceStatus, &out.UpgradingISBServiceStatus
		*out = new(UpgradingISBServiceStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PromotedISBServiceStatus != nil {
		in, out := &in.PromotedISBServiceStatus, &out.PromotedISBServiceStatus
		*out = new(PromotedISBServiceStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISBServiceProgressiveStatus.
func (in *ISBServiceProgressiveStatus) DeepCopy() *ISBServiceProgressiveStatus {
	if in == nil {
		return nil
	}
	out := new(ISBServiceProgressiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISBServiceRollout) DeepCopyInto(out *ISBServiceRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISBServiceRollout.
func (in *ISBServiceRollout) DeepCopy() *ISBServiceRollout {
	if in == nil {
		return nil
	}
	out := new(ISBServiceRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISBServiceRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISBServiceRolloutList) DeepCopyInto(out *ISBServiceRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ISBServiceRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISBServiceRolloutList.
func (in *ISBServiceRolloutList) DeepCopy() *ISBServiceRolloutList {
	if in == nil {
		return nil
	}
	out := new(ISBServiceRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ISBServiceRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISBServiceRolloutSpec) DeepCopyInto(out *ISBServiceRolloutSpec) {
	*out = *in
	in.InterStepBufferService.DeepCopyInto(&out.InterStepBufferService)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISBServiceRolloutSpec.
func (in *ISBServiceRolloutSpec) DeepCopy() *ISBServiceRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(ISBServiceRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ISBServiceRolloutStatus) DeepCopyInto(out *ISBServiceRolloutStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.PauseRequestStatus.DeepCopyInto(&out.PauseRequestStatus)
	if in.NameCount != nil {
		in, out := &in.NameCount, &out.NameCount
		*out = new(int32)
		**out = **in
	}
	in.ProgressiveStatus.DeepCopyInto(&out.ProgressiveStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ISBServiceRolloutStatus.
func (in *ISBServiceRolloutStatus) DeepCopy() *ISBServiceRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(ISBServiceRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterStepBufferService) DeepCopyInto(out *InterStepBufferService) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterStepBufferService.
func (in *InterStepBufferService) DeepCopy() *InterStepBufferService {
	if in == nil {
		return nil
	}
	out := new(InterStepBufferService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
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
func (in *MonoVertex) DeepCopyInto(out *MonoVertex) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertex.
func (in *MonoVertex) DeepCopy() *MonoVertex {
	if in == nil {
		return nil
	}
	out := new(MonoVertex)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonoVertexProgressiveStatus) DeepCopyInto(out *MonoVertexProgressiveStatus) {
	*out = *in
	if in.UpgradingMonoVertexStatus != nil {
		in, out := &in.UpgradingMonoVertexStatus, &out.UpgradingMonoVertexStatus
		*out = new(UpgradingMonoVertexStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PromotedMonoVertexStatus != nil {
		in, out := &in.PromotedMonoVertexStatus, &out.PromotedMonoVertexStatus
		*out = new(PromotedMonoVertexStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertexProgressiveStatus.
func (in *MonoVertexProgressiveStatus) DeepCopy() *MonoVertexProgressiveStatus {
	if in == nil {
		return nil
	}
	out := new(MonoVertexProgressiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonoVertexRollout) DeepCopyInto(out *MonoVertexRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertexRollout.
func (in *MonoVertexRollout) DeepCopy() *MonoVertexRollout {
	if in == nil {
		return nil
	}
	out := new(MonoVertexRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonoVertexRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonoVertexRolloutList) DeepCopyInto(out *MonoVertexRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MonoVertexRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertexRolloutList.
func (in *MonoVertexRolloutList) DeepCopy() *MonoVertexRolloutList {
	if in == nil {
		return nil
	}
	out := new(MonoVertexRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MonoVertexRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonoVertexRolloutSpec) DeepCopyInto(out *MonoVertexRolloutSpec) {
	*out = *in
	in.MonoVertex.DeepCopyInto(&out.MonoVertex)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertexRolloutSpec.
func (in *MonoVertexRolloutSpec) DeepCopy() *MonoVertexRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(MonoVertexRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonoVertexRolloutStatus) DeepCopyInto(out *MonoVertexRolloutStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.NameCount != nil {
		in, out := &in.NameCount, &out.NameCount
		*out = new(int32)
		**out = **in
	}
	in.ProgressiveStatus.DeepCopyInto(&out.ProgressiveStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonoVertexRolloutStatus.
func (in *MonoVertexRolloutStatus) DeepCopy() *MonoVertexRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(MonoVertexRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowController) DeepCopyInto(out *NumaflowController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowController.
func (in *NumaflowController) DeepCopy() *NumaflowController {
	if in == nil {
		return nil
	}
	out := new(NumaflowController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NumaflowController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerList) DeepCopyInto(out *NumaflowControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NumaflowController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerList.
func (in *NumaflowControllerList) DeepCopy() *NumaflowControllerList {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NumaflowControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerRollout) DeepCopyInto(out *NumaflowControllerRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerRollout.
func (in *NumaflowControllerRollout) DeepCopy() *NumaflowControllerRollout {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NumaflowControllerRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerRolloutList) DeepCopyInto(out *NumaflowControllerRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NumaflowControllerRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerRolloutList.
func (in *NumaflowControllerRolloutList) DeepCopy() *NumaflowControllerRolloutList {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NumaflowControllerRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerRolloutSpec) DeepCopyInto(out *NumaflowControllerRolloutSpec) {
	*out = *in
	out.Controller = in.Controller
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerRolloutSpec.
func (in *NumaflowControllerRolloutSpec) DeepCopy() *NumaflowControllerRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerRolloutStatus) DeepCopyInto(out *NumaflowControllerRolloutStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.PauseRequestStatus.DeepCopyInto(&out.PauseRequestStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerRolloutStatus.
func (in *NumaflowControllerRolloutStatus) DeepCopy() *NumaflowControllerRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerSpec) DeepCopyInto(out *NumaflowControllerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerSpec.
func (in *NumaflowControllerSpec) DeepCopy() *NumaflowControllerSpec {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NumaflowControllerStatus) DeepCopyInto(out *NumaflowControllerStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NumaflowControllerStatus.
func (in *NumaflowControllerStatus) DeepCopy() *NumaflowControllerStatus {
	if in == nil {
		return nil
	}
	out := new(NumaflowControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PauseStatus) DeepCopyInto(out *PauseStatus) {
	*out = *in
	in.LastPauseBeginTime.DeepCopyInto(&out.LastPauseBeginTime)
	in.LastPauseTransitionTime.DeepCopyInto(&out.LastPauseTransitionTime)
	in.LastPauseEndTime.DeepCopyInto(&out.LastPauseEndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PauseStatus.
func (in *PauseStatus) DeepCopy() *PauseStatus {
	if in == nil {
		return nil
	}
	out := new(PauseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineProgressiveStatus) DeepCopyInto(out *PipelineProgressiveStatus) {
	*out = *in
	if in.UpgradingPipelineStatus != nil {
		in, out := &in.UpgradingPipelineStatus, &out.UpgradingPipelineStatus
		*out = new(UpgradingPipelineStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.PromotedPipelineStatus != nil {
		in, out := &in.PromotedPipelineStatus, &out.PromotedPipelineStatus
		*out = new(PromotedPipelineStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineProgressiveStatus.
func (in *PipelineProgressiveStatus) DeepCopy() *PipelineProgressiveStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineProgressiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRollout) DeepCopyInto(out *PipelineRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRollout.
func (in *PipelineRollout) DeepCopy() *PipelineRollout {
	if in == nil {
		return nil
	}
	out := new(PipelineRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRolloutList) DeepCopyInto(out *PipelineRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRolloutList.
func (in *PipelineRolloutList) DeepCopy() *PipelineRolloutList {
	if in == nil {
		return nil
	}
	out := new(PipelineRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRolloutSpec) DeepCopyInto(out *PipelineRolloutSpec) {
	*out = *in
	in.Pipeline.DeepCopyInto(&out.Pipeline)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRolloutSpec.
func (in *PipelineRolloutSpec) DeepCopy() *PipelineRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineRolloutStatus) DeepCopyInto(out *PipelineRolloutStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.PauseStatus.DeepCopyInto(&out.PauseStatus)
	if in.NameCount != nil {
		in, out := &in.NameCount, &out.NameCount
		*out = new(int32)
		**out = **in
	}
	in.ProgressiveStatus.DeepCopyInto(&out.ProgressiveStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineRolloutStatus.
func (in *PipelineRolloutStatus) DeepCopy() *PipelineRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotedChildStatus) DeepCopyInto(out *PromotedChildStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotedChildStatus.
func (in *PromotedChildStatus) DeepCopy() *PromotedChildStatus {
	if in == nil {
		return nil
	}
	out := new(PromotedChildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotedISBServiceStatus) DeepCopyInto(out *PromotedISBServiceStatus) {
	*out = *in
	out.PromotedChildStatus = in.PromotedChildStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotedISBServiceStatus.
func (in *PromotedISBServiceStatus) DeepCopy() *PromotedISBServiceStatus {
	if in == nil {
		return nil
	}
	out := new(PromotedISBServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotedMonoVertexStatus) DeepCopyInto(out *PromotedMonoVertexStatus) {
	*out = *in
	in.PromotedPipelineTypeStatus.DeepCopyInto(&out.PromotedPipelineTypeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotedMonoVertexStatus.
func (in *PromotedMonoVertexStatus) DeepCopy() *PromotedMonoVertexStatus {
	if in == nil {
		return nil
	}
	out := new(PromotedMonoVertexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotedPipelineStatus) DeepCopyInto(out *PromotedPipelineStatus) {
	*out = *in
	in.PromotedPipelineTypeStatus.DeepCopyInto(&out.PromotedPipelineTypeStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotedPipelineStatus.
func (in *PromotedPipelineStatus) DeepCopy() *PromotedPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PromotedPipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotedPipelineTypeStatus) DeepCopyInto(out *PromotedPipelineTypeStatus) {
	*out = *in
	out.PromotedChildStatus = in.PromotedChildStatus
	if in.ScaleValues != nil {
		in, out := &in.ScaleValues, &out.ScaleValues
		*out = make(map[string]ScaleValues, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotedPipelineTypeStatus.
func (in *PromotedPipelineTypeStatus) DeepCopy() *PromotedPipelineTypeStatus {
	if in == nil {
		return nil
	}
	out := new(PromotedPipelineTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleValues) DeepCopyInto(out *ScaleValues) {
	*out = *in
	if in.DesiredMin != nil {
		in, out := &in.DesiredMin, &out.DesiredMin
		*out = new(int64)
		**out = **in
	}
	if in.DesiredMax != nil {
		in, out := &in.DesiredMax, &out.DesiredMax
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleValues.
func (in *ScaleValues) DeepCopy() *ScaleValues {
	if in == nil {
		return nil
	}
	out := new(ScaleValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastFailureTime.DeepCopyInto(&out.LastFailureTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradingChildStatus) DeepCopyInto(out *UpgradingChildStatus) {
	*out = *in
	if in.NextAssessmentTime != nil {
		in, out := &in.NextAssessmentTime, &out.NextAssessmentTime
		*out = (*in).DeepCopy()
	}
	if in.AssessUntil != nil {
		in, out := &in.AssessUntil, &out.AssessUntil
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradingChildStatus.
func (in *UpgradingChildStatus) DeepCopy() *UpgradingChildStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradingChildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradingISBServiceStatus) DeepCopyInto(out *UpgradingISBServiceStatus) {
	*out = *in
	in.UpgradingChildStatus.DeepCopyInto(&out.UpgradingChildStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradingISBServiceStatus.
func (in *UpgradingISBServiceStatus) DeepCopy() *UpgradingISBServiceStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradingISBServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradingMonoVertexStatus) DeepCopyInto(out *UpgradingMonoVertexStatus) {
	*out = *in
	in.UpgradingChildStatus.DeepCopyInto(&out.UpgradingChildStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradingMonoVertexStatus.
func (in *UpgradingMonoVertexStatus) DeepCopy() *UpgradingMonoVertexStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradingMonoVertexStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradingPipelineStatus) DeepCopyInto(out *UpgradingPipelineStatus) {
	*out = *in
	in.UpgradingChildStatus.DeepCopyInto(&out.UpgradingChildStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradingPipelineStatus.
func (in *UpgradingPipelineStatus) DeepCopy() *UpgradingPipelineStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradingPipelineStatus)
	in.DeepCopyInto(out)
	return out
}
