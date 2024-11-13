//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserDefinedNetwork) DeepCopyInto(out *ClusterUserDefinedNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserDefinedNetwork.
func (in *ClusterUserDefinedNetwork) DeepCopy() *ClusterUserDefinedNetwork {
	if in == nil {
		return nil
	}
	out := new(ClusterUserDefinedNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUserDefinedNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserDefinedNetworkList) DeepCopyInto(out *ClusterUserDefinedNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterUserDefinedNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserDefinedNetworkList.
func (in *ClusterUserDefinedNetworkList) DeepCopy() *ClusterUserDefinedNetworkList {
	if in == nil {
		return nil
	}
	out := new(ClusterUserDefinedNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUserDefinedNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserDefinedNetworkSpec) DeepCopyInto(out *ClusterUserDefinedNetworkSpec) {
	*out = *in
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.Network.DeepCopyInto(&out.Network)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserDefinedNetworkSpec.
func (in *ClusterUserDefinedNetworkSpec) DeepCopy() *ClusterUserDefinedNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterUserDefinedNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserDefinedNetworkStatus) DeepCopyInto(out *ClusterUserDefinedNetworkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserDefinedNetworkStatus.
func (in *ClusterUserDefinedNetworkStatus) DeepCopy() *ClusterUserDefinedNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterUserDefinedNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DualStackCIDRs) DeepCopyInto(out *DualStackCIDRs) {
	{
		in := &in
		*out = make(DualStackCIDRs, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DualStackCIDRs.
func (in DualStackCIDRs) DeepCopy() DualStackCIDRs {
	if in == nil {
		return nil
	}
	out := new(DualStackCIDRs)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layer2Config) DeepCopyInto(out *Layer2Config) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make(DualStackCIDRs, len(*in))
		copy(*out, *in)
	}
	if in.JoinSubnets != nil {
		in, out := &in.JoinSubnets, &out.JoinSubnets
		*out = make(DualStackCIDRs, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer2Config.
func (in *Layer2Config) DeepCopy() *Layer2Config {
	if in == nil {
		return nil
	}
	out := new(Layer2Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layer3Config) DeepCopyInto(out *Layer3Config) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]Layer3Subnet, len(*in))
		copy(*out, *in)
	}
	if in.JoinSubnets != nil {
		in, out := &in.JoinSubnets, &out.JoinSubnets
		*out = make(DualStackCIDRs, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer3Config.
func (in *Layer3Config) DeepCopy() *Layer3Config {
	if in == nil {
		return nil
	}
	out := new(Layer3Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Layer3Subnet) DeepCopyInto(out *Layer3Subnet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Layer3Subnet.
func (in *Layer3Subnet) DeepCopy() *Layer3Subnet {
	if in == nil {
		return nil
	}
	out := new(Layer3Subnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Layer3 != nil {
		in, out := &in.Layer3, &out.Layer3
		*out = new(Layer3Config)
		(*in).DeepCopyInto(*out)
	}
	if in.Layer2 != nil {
		in, out := &in.Layer2, &out.Layer2
		*out = new(Layer2Config)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedNetwork) DeepCopyInto(out *UserDefinedNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedNetwork.
func (in *UserDefinedNetwork) DeepCopy() *UserDefinedNetwork {
	if in == nil {
		return nil
	}
	out := new(UserDefinedNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserDefinedNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedNetworkList) DeepCopyInto(out *UserDefinedNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserDefinedNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedNetworkList.
func (in *UserDefinedNetworkList) DeepCopy() *UserDefinedNetworkList {
	if in == nil {
		return nil
	}
	out := new(UserDefinedNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserDefinedNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedNetworkSpec) DeepCopyInto(out *UserDefinedNetworkSpec) {
	*out = *in
	if in.Layer3 != nil {
		in, out := &in.Layer3, &out.Layer3
		*out = new(Layer3Config)
		(*in).DeepCopyInto(*out)
	}
	if in.Layer2 != nil {
		in, out := &in.Layer2, &out.Layer2
		*out = new(Layer2Config)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedNetworkSpec.
func (in *UserDefinedNetworkSpec) DeepCopy() *UserDefinedNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(UserDefinedNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedNetworkStatus) DeepCopyInto(out *UserDefinedNetworkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedNetworkStatus.
func (in *UserDefinedNetworkStatus) DeepCopy() *UserDefinedNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(UserDefinedNetworkStatus)
	in.DeepCopyInto(out)
	return out
}
