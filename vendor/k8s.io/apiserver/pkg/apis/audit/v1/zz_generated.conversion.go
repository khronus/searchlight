// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	authenticationv1 "k8s.io/api/authentication/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	audit "k8s.io/apiserver/pkg/apis/audit"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Event)(nil), (*audit.Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Event_To_audit_Event(a.(*Event), b.(*audit.Event), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.Event)(nil), (*Event)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_Event_To_v1_Event(a.(*audit.Event), b.(*Event), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*EventList)(nil), (*audit.EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_EventList_To_audit_EventList(a.(*EventList), b.(*audit.EventList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.EventList)(nil), (*EventList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_EventList_To_v1_EventList(a.(*audit.EventList), b.(*EventList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*GroupResources)(nil), (*audit.GroupResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_GroupResources_To_audit_GroupResources(a.(*GroupResources), b.(*audit.GroupResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.GroupResources)(nil), (*GroupResources)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_GroupResources_To_v1_GroupResources(a.(*audit.GroupResources), b.(*GroupResources), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectReference)(nil), (*audit.ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ObjectReference_To_audit_ObjectReference(a.(*ObjectReference), b.(*audit.ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.ObjectReference)(nil), (*ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_ObjectReference_To_v1_ObjectReference(a.(*audit.ObjectReference), b.(*ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Policy)(nil), (*audit.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Policy_To_audit_Policy(a.(*Policy), b.(*audit.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.Policy)(nil), (*Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_Policy_To_v1_Policy(a.(*audit.Policy), b.(*Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicyList)(nil), (*audit.PolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyList_To_audit_PolicyList(a.(*PolicyList), b.(*audit.PolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.PolicyList)(nil), (*PolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_PolicyList_To_v1_PolicyList(a.(*audit.PolicyList), b.(*PolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicyRule)(nil), (*audit.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyRule_To_audit_PolicyRule(a.(*PolicyRule), b.(*audit.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*audit.PolicyRule)(nil), (*PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_audit_PolicyRule_To_v1_PolicyRule(a.(*audit.PolicyRule), b.(*PolicyRule), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Event_To_audit_Event(in *Event, out *audit.Event, s conversion.Scope) error {
	out.Level = audit.Level(in.Level)
	out.AuditID = types.UID(in.AuditID)
	out.Stage = audit.Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.User, &out.User, 0); err != nil {
		return err
	}
	out.ImpersonatedUser = (*audit.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	out.UserAgent = in.UserAgent
	out.ObjectRef = (*audit.ObjectReference)(unsafe.Pointer(in.ObjectRef))
	out.ResponseStatus = (*metav1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*runtime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*runtime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_v1_Event_To_audit_Event is an autogenerated conversion function.
func Convert_v1_Event_To_audit_Event(in *Event, out *audit.Event, s conversion.Scope) error {
	return autoConvert_v1_Event_To_audit_Event(in, out, s)
}

func autoConvert_audit_Event_To_v1_Event(in *audit.Event, out *Event, s conversion.Scope) error {
	out.Level = Level(in.Level)
	out.AuditID = types.UID(in.AuditID)
	out.Stage = Stage(in.Stage)
	out.RequestURI = in.RequestURI
	out.Verb = in.Verb
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.User, &out.User, 0); err != nil {
		return err
	}
	out.ImpersonatedUser = (*authenticationv1.UserInfo)(unsafe.Pointer(in.ImpersonatedUser))
	out.SourceIPs = *(*[]string)(unsafe.Pointer(&in.SourceIPs))
	out.UserAgent = in.UserAgent
	out.ObjectRef = (*ObjectReference)(unsafe.Pointer(in.ObjectRef))
	out.ResponseStatus = (*metav1.Status)(unsafe.Pointer(in.ResponseStatus))
	out.RequestObject = (*runtime.Unknown)(unsafe.Pointer(in.RequestObject))
	out.ResponseObject = (*runtime.Unknown)(unsafe.Pointer(in.ResponseObject))
	out.RequestReceivedTimestamp = in.RequestReceivedTimestamp
	out.StageTimestamp = in.StageTimestamp
	out.Annotations = *(*map[string]string)(unsafe.Pointer(&in.Annotations))
	return nil
}

// Convert_audit_Event_To_v1_Event is an autogenerated conversion function.
func Convert_audit_Event_To_v1_Event(in *audit.Event, out *Event, s conversion.Scope) error {
	return autoConvert_audit_Event_To_v1_Event(in, out, s)
}

func autoConvert_v1_EventList_To_audit_EventList(in *EventList, out *audit.EventList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]audit.Event)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_EventList_To_audit_EventList is an autogenerated conversion function.
func Convert_v1_EventList_To_audit_EventList(in *EventList, out *audit.EventList, s conversion.Scope) error {
	return autoConvert_v1_EventList_To_audit_EventList(in, out, s)
}

func autoConvert_audit_EventList_To_v1_EventList(in *audit.EventList, out *EventList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Event)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_audit_EventList_To_v1_EventList is an autogenerated conversion function.
func Convert_audit_EventList_To_v1_EventList(in *audit.EventList, out *EventList, s conversion.Scope) error {
	return autoConvert_audit_EventList_To_v1_EventList(in, out, s)
}

func autoConvert_v1_GroupResources_To_audit_GroupResources(in *GroupResources, out *audit.GroupResources, s conversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_v1_GroupResources_To_audit_GroupResources is an autogenerated conversion function.
func Convert_v1_GroupResources_To_audit_GroupResources(in *GroupResources, out *audit.GroupResources, s conversion.Scope) error {
	return autoConvert_v1_GroupResources_To_audit_GroupResources(in, out, s)
}

func autoConvert_audit_GroupResources_To_v1_GroupResources(in *audit.GroupResources, out *GroupResources, s conversion.Scope) error {
	out.Group = in.Group
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	return nil
}

// Convert_audit_GroupResources_To_v1_GroupResources is an autogenerated conversion function.
func Convert_audit_GroupResources_To_v1_GroupResources(in *audit.GroupResources, out *GroupResources, s conversion.Scope) error {
	return autoConvert_audit_GroupResources_To_v1_GroupResources(in, out, s)
}

func autoConvert_v1_ObjectReference_To_audit_ObjectReference(in *ObjectReference, out *audit.ObjectReference, s conversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.APIGroup = in.APIGroup
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

// Convert_v1_ObjectReference_To_audit_ObjectReference is an autogenerated conversion function.
func Convert_v1_ObjectReference_To_audit_ObjectReference(in *ObjectReference, out *audit.ObjectReference, s conversion.Scope) error {
	return autoConvert_v1_ObjectReference_To_audit_ObjectReference(in, out, s)
}

func autoConvert_audit_ObjectReference_To_v1_ObjectReference(in *audit.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	out.Resource = in.Resource
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.UID = types.UID(in.UID)
	out.APIGroup = in.APIGroup
	out.APIVersion = in.APIVersion
	out.ResourceVersion = in.ResourceVersion
	out.Subresource = in.Subresource
	return nil
}

// Convert_audit_ObjectReference_To_v1_ObjectReference is an autogenerated conversion function.
func Convert_audit_ObjectReference_To_v1_ObjectReference(in *audit.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	return autoConvert_audit_ObjectReference_To_v1_ObjectReference(in, out, s)
}

func autoConvert_v1_Policy_To_audit_Policy(in *Policy, out *audit.Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]audit.PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]audit.Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_v1_Policy_To_audit_Policy is an autogenerated conversion function.
func Convert_v1_Policy_To_audit_Policy(in *Policy, out *audit.Policy, s conversion.Scope) error {
	return autoConvert_v1_Policy_To_audit_Policy(in, out, s)
}

func autoConvert_audit_Policy_To_v1_Policy(in *audit.Policy, out *Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Rules = *(*[]PolicyRule)(unsafe.Pointer(&in.Rules))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_audit_Policy_To_v1_Policy is an autogenerated conversion function.
func Convert_audit_Policy_To_v1_Policy(in *audit.Policy, out *Policy, s conversion.Scope) error {
	return autoConvert_audit_Policy_To_v1_Policy(in, out, s)
}

func autoConvert_v1_PolicyList_To_audit_PolicyList(in *PolicyList, out *audit.PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]audit.Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_PolicyList_To_audit_PolicyList is an autogenerated conversion function.
func Convert_v1_PolicyList_To_audit_PolicyList(in *PolicyList, out *audit.PolicyList, s conversion.Scope) error {
	return autoConvert_v1_PolicyList_To_audit_PolicyList(in, out, s)
}

func autoConvert_audit_PolicyList_To_v1_PolicyList(in *audit.PolicyList, out *PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_audit_PolicyList_To_v1_PolicyList is an autogenerated conversion function.
func Convert_audit_PolicyList_To_v1_PolicyList(in *audit.PolicyList, out *PolicyList, s conversion.Scope) error {
	return autoConvert_audit_PolicyList_To_v1_PolicyList(in, out, s)
}

func autoConvert_v1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *audit.PolicyRule, s conversion.Scope) error {
	out.Level = audit.Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]audit.GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]audit.Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_v1_PolicyRule_To_audit_PolicyRule is an autogenerated conversion function.
func Convert_v1_PolicyRule_To_audit_PolicyRule(in *PolicyRule, out *audit.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1_PolicyRule_To_audit_PolicyRule(in, out, s)
}

func autoConvert_audit_PolicyRule_To_v1_PolicyRule(in *audit.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	out.Level = Level(in.Level)
	out.Users = *(*[]string)(unsafe.Pointer(&in.Users))
	out.UserGroups = *(*[]string)(unsafe.Pointer(&in.UserGroups))
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.Resources = *(*[]GroupResources)(unsafe.Pointer(&in.Resources))
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	out.OmitStages = *(*[]Stage)(unsafe.Pointer(&in.OmitStages))
	return nil
}

// Convert_audit_PolicyRule_To_v1_PolicyRule is an autogenerated conversion function.
func Convert_audit_PolicyRule_To_v1_PolicyRule(in *audit.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	return autoConvert_audit_PolicyRule_To_v1_PolicyRule(in, out, s)
}
