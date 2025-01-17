//go:build !ignore_autogenerated

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityZone) DeepCopyInto(out *AvailabilityZone) {
	*out = *in
	if in.LoadBalancerAddresses != nil {
		in, out := &in.LoadBalancerAddresses, &out.LoadBalancerAddresses
		*out = make([]*LoadBalancerAddress, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(LoadBalancerAddress)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.OutpostID != nil {
		in, out := &in.OutpostID, &out.OutpostID
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.ZoneName != nil {
		in, out := &in.ZoneName, &out.ZoneName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityZone.
func (in *AvailabilityZone) DeepCopy() *AvailabilityZone {
	if in == nil {
		return nil
	}
	out := new(AvailabilityZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	if in.LoadBalancerARN != nil {
		in, out := &in.LoadBalancerARN, &out.LoadBalancerARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer) DeepCopyInto(out *LoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer.
func (in *LoadBalancer) DeepCopy() *LoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerAddress) DeepCopyInto(out *LoadBalancerAddress) {
	*out = *in
	if in.AllocationID != nil {
		in, out := &in.AllocationID, &out.AllocationID
		*out = new(string)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPv4Address != nil {
		in, out := &in.PrivateIPv4Address, &out.PrivateIPv4Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerAddress.
func (in *LoadBalancerAddress) DeepCopy() *LoadBalancerAddress {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerList) DeepCopyInto(out *LoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerList.
func (in *LoadBalancerList) DeepCopy() *LoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.CustomerOwnedIPv4Pool != nil {
		in, out := &in.CustomerOwnedIPv4Pool, &out.CustomerOwnedIPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetMappings != nil {
		in, out := &in.SubnetMappings, &out.SubnetMappings
		*out = make([]*SubnetMapping, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SubnetMapping)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerState) DeepCopyInto(out *LoadBalancerState) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerState.
func (in *LoadBalancerState) DeepCopy() *LoadBalancerState {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerStatus) DeepCopyInto(out *LoadBalancerStatus) {
	*out = *in
	if in.ACKResourceMetadata != nil {
		in, out := &in.ACKResourceMetadata, &out.ACKResourceMetadata
		*out = new(corev1alpha1.ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]*corev1alpha1.Condition, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(corev1alpha1.Condition)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*AvailabilityZone, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AvailabilityZone)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CanonicalHostedZoneID != nil {
		in, out := &in.CanonicalHostedZoneID, &out.CanonicalHostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = (*in).DeepCopy()
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic != nil {
		in, out := &in.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic, &out.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerState)
		(*in).DeepCopyInto(*out)
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerStatus.
func (in *LoadBalancerStatus) DeepCopy() *LoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancer_SDK) DeepCopyInto(out *LoadBalancer_SDK) {
	*out = *in
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]*AvailabilityZone, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(AvailabilityZone)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.CanonicalHostedZoneID != nil {
		in, out := &in.CanonicalHostedZoneID, &out.CanonicalHostedZoneID
		*out = new(string)
		**out = **in
	}
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = (*in).DeepCopy()
	}
	if in.CustomerOwnedIPv4Pool != nil {
		in, out := &in.CustomerOwnedIPv4Pool, &out.CustomerOwnedIPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic != nil {
		in, out := &in.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic, &out.EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic
		*out = new(string)
		**out = **in
	}
	if in.IPAddressType != nil {
		in, out := &in.IPAddressType, &out.IPAddressType
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerARN != nil {
		in, out := &in.LoadBalancerARN, &out.LoadBalancerARN
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerName != nil {
		in, out := &in.LoadBalancerName, &out.LoadBalancerName
		*out = new(string)
		**out = **in
	}
	if in.Scheme != nil {
		in, out := &in.Scheme, &out.Scheme
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerState)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancer_SDK.
func (in *LoadBalancer_SDK) DeepCopy() *LoadBalancer_SDK {
	if in == nil {
		return nil
	}
	out := new(LoadBalancer_SDK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetMapping) DeepCopyInto(out *SubnetMapping) {
	*out = *in
	if in.AllocationID != nil {
		in, out := &in.AllocationID, &out.AllocationID
		*out = new(string)
		**out = **in
	}
	if in.IPv6Address != nil {
		in, out := &in.IPv6Address, &out.IPv6Address
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPv4Address != nil {
		in, out := &in.PrivateIPv4Address, &out.PrivateIPv4Address
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetMapping.
func (in *SubnetMapping) DeepCopy() *SubnetMapping {
	if in == nil {
		return nil
	}
	out := new(SubnetMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagDescription) DeepCopyInto(out *TagDescription) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagDescription.
func (in *TagDescription) DeepCopy() *TagDescription {
	if in == nil {
		return nil
	}
	out := new(TagDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetDescription) DeepCopyInto(out *TargetDescription) {
	*out = *in
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetDescription.
func (in *TargetDescription) DeepCopy() *TargetDescription {
	if in == nil {
		return nil
	}
	out := new(TargetDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroup) DeepCopyInto(out *TargetGroup) {
	*out = *in
	if in.LoadBalancerARNs != nil {
		in, out := &in.LoadBalancerARNs, &out.LoadBalancerARNs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroup.
func (in *TargetGroup) DeepCopy() *TargetGroup {
	if in == nil {
		return nil
	}
	out := new(TargetGroup)
	in.DeepCopyInto(out)
	return out
}
