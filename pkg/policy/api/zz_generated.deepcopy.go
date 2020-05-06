// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package api

import (
	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSGroup) DeepCopyInto(out *AWSGroup) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroupsIds != nil {
		in, out := &in.SecurityGroupsIds, &out.SecurityGroupsIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupsNames != nil {
		in, out := &in.SecurityGroupsNames, &out.SecurityGroupsNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSGroup.
func (in *AWSGroup) DeepCopy() *AWSGroup {
	if in == nil {
		return nil
	}
	out := new(AWSGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIDRRule) DeepCopyInto(out *CIDRRule) {
	*out = *in
	if in.ExceptCIDRs != nil {
		in, out := &in.ExceptCIDRs, &out.ExceptCIDRs
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRRule.
func (in *CIDRRule) DeepCopy() *CIDRRule {
	if in == nil {
		return nil
	}
	out := new(CIDRRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CIDRRuleSlice) DeepCopyInto(out *CIDRRuleSlice) {
	{
		in := &in
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRRuleSlice.
func (in CIDRRuleSlice) DeepCopy() CIDRRuleSlice {
	if in == nil {
		return nil
	}
	out := new(CIDRRuleSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CIDRSlice) DeepCopyInto(out *CIDRSlice) {
	{
		in := &in
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRSlice.
func (in CIDRSlice) DeepCopy() CIDRSlice {
	if in == nil {
		return nil
	}
	out := new(CIDRSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	if in.ToEndpoints != nil {
		in, out := &in.ToEndpoints, &out.ToEndpoints
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToRequires != nil {
		in, out := &in.ToRequires, &out.ToRequires
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make([]PortRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToCIDR != nil {
		in, out := &in.ToCIDR, &out.ToCIDR
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
	}
	if in.ToCIDRSet != nil {
		in, out := &in.ToCIDRSet, &out.ToCIDRSet
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToEntities != nil {
		in, out := &in.ToEntities, &out.ToEntities
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
	}
	if in.ToServices != nil {
		in, out := &in.ToServices, &out.ToServices
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToFQDNs != nil {
		in, out := &in.ToFQDNs, &out.ToFQDNs
		*out = make(FQDNSelectorSlice, len(*in))
		copy(*out, *in)
	}
	if in.ToGroups != nil {
		in, out := &in.ToGroups, &out.ToGroups
		*out = make([]ToGroups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.aggregatedSelectors != nil {
		in, out := &in.aggregatedSelectors, &out.aggregatedSelectors
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSelector) DeepCopyInto(out *EndpointSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.requirements != nil {
		in, out := &in.requirements, &out.requirements
		*out = new(labels.Requirements)
		if **in != nil {
			in, out := *in, *out
			*out = make([]labels.Requirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSelector.
func (in *EndpointSelector) DeepCopy() *EndpointSelector {
	if in == nil {
		return nil
	}
	out := new(EndpointSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EndpointSelectorSlice) DeepCopyInto(out *EndpointSelectorSlice) {
	{
		in := &in
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSelectorSlice.
func (in EndpointSelectorSlice) DeepCopy() EndpointSelectorSlice {
	if in == nil {
		return nil
	}
	out := new(EndpointSelectorSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EntitySlice) DeepCopyInto(out *EntitySlice) {
	{
		in := &in
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntitySlice.
func (in EntitySlice) DeepCopy() EntitySlice {
	if in == nil {
		return nil
	}
	out := new(EntitySlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FQDNSelector) DeepCopyInto(out *FQDNSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FQDNSelector.
func (in *FQDNSelector) DeepCopy() *FQDNSelector {
	if in == nil {
		return nil
	}
	out := new(FQDNSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FQDNSelectorSlice) DeepCopyInto(out *FQDNSelectorSlice) {
	{
		in := &in
		*out = make(FQDNSelectorSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FQDNSelectorSlice.
func (in FQDNSelectorSlice) DeepCopy() FQDNSelectorSlice {
	if in == nil {
		return nil
	}
	out := new(FQDNSelectorSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(Secret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	if in.FromEndpoints != nil {
		in, out := &in.FromEndpoints, &out.FromEndpoints
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromRequires != nil {
		in, out := &in.FromRequires, &out.FromRequires
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make([]PortRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromCIDR != nil {
		in, out := &in.FromCIDR, &out.FromCIDR
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
	}
	if in.FromCIDRSet != nil {
		in, out := &in.FromCIDRSet, &out.FromCIDRSet
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromEntities != nil {
		in, out := &in.FromEntities, &out.FromEntities
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
	}
	if in.aggregatedSelectors != nil {
		in, out := &in.aggregatedSelectors, &out.aggregatedSelectors
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceNamespace) DeepCopyInto(out *K8sServiceNamespace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceNamespace.
func (in *K8sServiceNamespace) DeepCopy() *K8sServiceNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceSelectorNamespace) DeepCopyInto(out *K8sServiceSelectorNamespace) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceSelectorNamespace.
func (in *K8sServiceSelectorNamespace) DeepCopy() *K8sServiceSelectorNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceSelectorNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in KafkaRole) DeepCopyInto(out *KafkaRole) {
	{
		in := &in
		*out = make(KafkaRole, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaRole.
func (in KafkaRole) DeepCopy() KafkaRole {
	if in == nil {
		return nil
	}
	out := new(KafkaRole)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Rules) DeepCopyInto(out *L7Rules) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]PortRuleHTTP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = make([]PortRuleKafka, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]PortRuleDNS, len(*in))
		copy(*out, *in)
	}
	if in.L7 != nil {
		in, out := &in.L7, &out.L7
		*out = make([]PortRuleL7, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(PortRuleL7, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Rules.
func (in *L7Rules) DeepCopy() *L7Rules {
	if in == nil {
		return nil
	}
	out := new(L7Rules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortProtocol) DeepCopyInto(out *PortProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortProtocol.
func (in *PortProtocol) DeepCopy() *PortProtocol {
	if in == nil {
		return nil
	}
	out := new(PortProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRule) DeepCopyInto(out *PortRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortProtocol, len(*in))
		copy(*out, *in)
	}
	if in.TerminatingTLS != nil {
		in, out := &in.TerminatingTLS, &out.TerminatingTLS
		*out = new(TLSContext)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginatingTLS != nil {
		in, out := &in.OriginatingTLS, &out.OriginatingTLS
		*out = new(TLSContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new(L7Rules)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRule.
func (in *PortRule) DeepCopy() *PortRule {
	if in == nil {
		return nil
	}
	out := new(PortRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleDNS) DeepCopyInto(out *PortRuleDNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleDNS.
func (in *PortRuleDNS) DeepCopy() *PortRuleDNS {
	if in == nil {
		return nil
	}
	out := new(PortRuleDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleHTTP) DeepCopyInto(out *PortRuleHTTP) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HeaderMatches != nil {
		in, out := &in.HeaderMatches, &out.HeaderMatches
		*out = make([]*HeaderMatch, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HeaderMatch)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleHTTP.
func (in *PortRuleHTTP) DeepCopy() *PortRuleHTTP {
	if in == nil {
		return nil
	}
	out := new(PortRuleHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleKafka) DeepCopyInto(out *PortRuleKafka) {
	*out = *in
	if in.apiKeyInt != nil {
		in, out := &in.apiKeyInt, &out.apiKeyInt
		*out = make(KafkaRole, len(*in))
		copy(*out, *in)
	}
	if in.apiVersionInt != nil {
		in, out := &in.apiVersionInt, &out.apiVersionInt
		*out = new(int16)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleKafka.
func (in *PortRuleKafka) DeepCopy() *PortRuleKafka {
	if in == nil {
		return nil
	}
	out := new(PortRuleKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PortRuleL7) DeepCopyInto(out *PortRuleL7) {
	{
		in := &in
		*out = make(PortRuleL7, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleL7.
func (in PortRuleL7) DeepCopy() PortRuleL7 {
	if in == nil {
		return nil
	}
	out := new(PortRuleL7)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.EndpointSelector.DeepCopyInto(&out.EndpointSelector)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Labels = in.Labels.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Rules) DeepCopyInto(out *Rules) {
	{
		in := &in
		*out = make(Rules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Rule)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rules.
func (in Rules) DeepCopy() Rules {
	if in == nil {
		return nil
	}
	out := new(Rules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.K8sServiceSelector != nil {
		in, out := &in.K8sServiceSelector, &out.K8sServiceSelector
		*out = new(K8sServiceSelectorNamespace)
		(*in).DeepCopyInto(*out)
	}
	if in.K8sService != nil {
		in, out := &in.K8sService, &out.K8sService
		*out = new(K8sServiceNamespace)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.requirements != nil {
		in, out := &in.requirements, &out.requirements
		*out = new(labels.Requirements)
		if **in != nil {
			in, out := *in, *out
			*out = make([]labels.Requirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSContext) DeepCopyInto(out *TLSContext) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(Secret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSContext.
func (in *TLSContext) DeepCopy() *TLSContext {
	if in == nil {
		return nil
	}
	out := new(TLSContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToGroups) DeepCopyInto(out *ToGroups) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSGroup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToGroups.
func (in *ToGroups) DeepCopy() *ToGroups {
	if in == nil {
		return nil
	}
	out := new(ToGroups)
	in.DeepCopyInto(out)
	return out
}
