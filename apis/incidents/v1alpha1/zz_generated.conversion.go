// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	incidents "go.searchlight.dev/icinga-operator/apis/incidents"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Acknowledgement)(nil), (*incidents.Acknowledgement)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Acknowledgement_To_incidents_Acknowledgement(a.(*Acknowledgement), b.(*incidents.Acknowledgement), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*incidents.Acknowledgement)(nil), (*Acknowledgement)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_incidents_Acknowledgement_To_v1alpha1_Acknowledgement(a.(*incidents.Acknowledgement), b.(*Acknowledgement), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AcknowledgementRequest)(nil), (*incidents.AcknowledgementRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest(a.(*AcknowledgementRequest), b.(*incidents.AcknowledgementRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*incidents.AcknowledgementRequest)(nil), (*AcknowledgementRequest)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest(a.(*incidents.AcknowledgementRequest), b.(*AcknowledgementRequest), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AcknowledgementResponse)(nil), (*incidents.AcknowledgementResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse(a.(*AcknowledgementResponse), b.(*incidents.AcknowledgementResponse), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*incidents.AcknowledgementResponse)(nil), (*AcknowledgementResponse)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse(a.(*incidents.AcknowledgementResponse), b.(*AcknowledgementResponse), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Acknowledgement_To_incidents_Acknowledgement(in *Acknowledgement, out *incidents.Acknowledgement, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest(&in.Request, &out.Request, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse(&in.Response, &out.Response, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Acknowledgement_To_incidents_Acknowledgement is an autogenerated conversion function.
func Convert_v1alpha1_Acknowledgement_To_incidents_Acknowledgement(in *Acknowledgement, out *incidents.Acknowledgement, s conversion.Scope) error {
	return autoConvert_v1alpha1_Acknowledgement_To_incidents_Acknowledgement(in, out, s)
}

func autoConvert_incidents_Acknowledgement_To_v1alpha1_Acknowledgement(in *incidents.Acknowledgement, out *Acknowledgement, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest(&in.Request, &out.Request, s); err != nil {
		return err
	}
	if err := Convert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse(&in.Response, &out.Response, s); err != nil {
		return err
	}
	return nil
}

// Convert_incidents_Acknowledgement_To_v1alpha1_Acknowledgement is an autogenerated conversion function.
func Convert_incidents_Acknowledgement_To_v1alpha1_Acknowledgement(in *incidents.Acknowledgement, out *Acknowledgement, s conversion.Scope) error {
	return autoConvert_incidents_Acknowledgement_To_v1alpha1_Acknowledgement(in, out, s)
}

func autoConvert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest(in *AcknowledgementRequest, out *incidents.AcknowledgementRequest, s conversion.Scope) error {
	out.Comment = in.Comment
	out.SkipNotify = in.SkipNotify
	return nil
}

// Convert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest is an autogenerated conversion function.
func Convert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest(in *AcknowledgementRequest, out *incidents.AcknowledgementRequest, s conversion.Scope) error {
	return autoConvert_v1alpha1_AcknowledgementRequest_To_incidents_AcknowledgementRequest(in, out, s)
}

func autoConvert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest(in *incidents.AcknowledgementRequest, out *AcknowledgementRequest, s conversion.Scope) error {
	out.Comment = in.Comment
	out.SkipNotify = in.SkipNotify
	return nil
}

// Convert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest is an autogenerated conversion function.
func Convert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest(in *incidents.AcknowledgementRequest, out *AcknowledgementRequest, s conversion.Scope) error {
	return autoConvert_incidents_AcknowledgementRequest_To_v1alpha1_AcknowledgementRequest(in, out, s)
}

func autoConvert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse(in *AcknowledgementResponse, out *incidents.AcknowledgementResponse, s conversion.Scope) error {
	out.Timestamp = in.Timestamp
	return nil
}

// Convert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse is an autogenerated conversion function.
func Convert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse(in *AcknowledgementResponse, out *incidents.AcknowledgementResponse, s conversion.Scope) error {
	return autoConvert_v1alpha1_AcknowledgementResponse_To_incidents_AcknowledgementResponse(in, out, s)
}

func autoConvert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse(in *incidents.AcknowledgementResponse, out *AcknowledgementResponse, s conversion.Scope) error {
	out.Timestamp = in.Timestamp
	return nil
}

// Convert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse is an autogenerated conversion function.
func Convert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse(in *incidents.AcknowledgementResponse, out *AcknowledgementResponse, s conversion.Scope) error {
	return autoConvert_incidents_AcknowledgementResponse_To_v1alpha1_AcknowledgementResponse(in, out, s)
}
