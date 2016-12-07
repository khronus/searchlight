// Code generated by protoc-gen-go.
// source: master.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MasterCreateRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	VolumeId    string `protobuf:"bytes,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	Namespace   string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *MasterCreateRequest) Reset()                    { *m = MasterCreateRequest{} }
func (m *MasterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*MasterCreateRequest) ProtoMessage()               {}
func (*MasterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *MasterCreateRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MasterCreateRequest) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *MasterCreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type MasterDeleteRequest struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *MasterDeleteRequest) Reset()                    { *m = MasterDeleteRequest{} }
func (m *MasterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*MasterDeleteRequest) ProtoMessage()               {}
func (*MasterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *MasterDeleteRequest) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MasterDeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*MasterCreateRequest)(nil), "appscode.ci.v1beta1.MasterCreateRequest")
	proto.RegisterType((*MasterDeleteRequest)(nil), "appscode.ci.v1beta1.MasterDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Master service

type MasterClient interface {
	Create(ctx context.Context, in *MasterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *MasterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Create(ctx context.Context, in *MasterCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Master/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Delete(ctx context.Context, in *MasterDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.LongRunningResponse, error) {
	out := new(appscode_dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Master/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Create(context.Context, *MasterCreateRequest) (*appscode_dtypes.LongRunningResponse, error)
	Delete(context.Context, *MasterDeleteRequest) (*appscode_dtypes.LongRunningResponse, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Master/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Create(ctx, req.(*MasterCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Master/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Delete(ctx, req.(*MasterDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Master_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Master_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}

func init() { proto.RegisterFile("master.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x25, 0x2b, 0xac, 0x36, 0xed, 0x29, 0xbd, 0x94, 0xb6, 0x60, 0x5d, 0x14, 0x4a, 0x85, 0x84,
	0x2a, 0x5e, 0x3c, 0x56, 0x2f, 0x82, 0x4a, 0xd9, 0x83, 0x07, 0x2f, 0x25, 0xcd, 0x0e, 0x4b, 0x60,
	0x37, 0x89, 0x4d, 0xb6, 0xe0, 0xb5, 0xe0, 0x2f, 0x10, 0x7f, 0x99, 0x7f, 0xc1, 0x1f, 0x22, 0x4d,
	0x5a, 0x6a, 0xa5, 0xf8, 0x71, 0xc9, 0xe1, 0xcd, 0xcc, 0x9b, 0xf7, 0x5e, 0x06, 0x37, 0x4a, 0x6e,
	0x1d, 0xcc, 0xa8, 0x99, 0x69, 0xa7, 0x49, 0x93, 0x1b, 0x63, 0x85, 0xce, 0x80, 0x0a, 0x49, 0xe7,
	0xc3, 0x29, 0x38, 0x3e, 0x6c, 0x77, 0x73, 0xad, 0xf3, 0x02, 0x18, 0x37, 0x92, 0x71, 0xa5, 0xb4,
	0xe3, 0x4e, 0x6a, 0x65, 0xc3, 0x48, 0xfb, 0x70, 0x3d, 0xe2, 0xeb, 0x99, 0x7b, 0x36, 0x60, 0x99,
	0x7f, 0x43, 0x43, 0x62, 0x71, 0xf3, 0xce, 0xef, 0xb8, 0x9a, 0x01, 0x77, 0x90, 0xc2, 0x53, 0x05,
	0xd6, 0x91, 0x23, 0xdc, 0x10, 0x45, 0xb5, 0xc4, 0x27, 0x8a, 0x97, 0xd0, 0x42, 0x3d, 0xd4, 0xaf,
	0xa5, 0xf5, 0x15, 0x76, 0xcf, 0x4b, 0x20, 0x1d, 0x5c, 0x9b, 0xeb, 0xa2, 0x2a, 0x61, 0x22, 0xb3,
	0x56, 0xe4, 0xeb, 0x07, 0x01, 0xb8, 0xc9, 0x48, 0x17, 0xd7, 0x96, 0x73, 0xd6, 0x70, 0x01, 0xad,
	0x3d, 0x5f, 0xdc, 0x00, 0xc9, 0xc3, 0x7a, 0xe9, 0x35, 0x14, 0xf0, 0xaf, 0xa5, 0x5b, 0xbc, 0xd1,
	0x37, 0xde, 0xb3, 0xb7, 0x08, 0xc7, 0x81, 0x98, 0xbc, 0x20, 0x1c, 0x07, 0x4b, 0xa4, 0x4f, 0x77,
	0xe4, 0x46, 0x77, 0xb8, 0x6e, 0x1f, 0x6f, 0x3a, 0x43, 0x54, 0xf4, 0x56, 0xab, 0x3c, 0xad, 0x94,
	0x92, 0x2a, 0x4f, 0xc1, 0x1a, 0xad, 0x2c, 0x24, 0xa7, 0x8b, 0xf7, 0x8f, 0xd7, 0xe8, 0x24, 0xe9,
	0xb1, 0xad, 0x70, 0x85, 0x64, 0x2b, 0x6e, 0x16, 0x7e, 0xcd, 0x5e, 0xa2, 0x01, 0x59, 0x20, 0x1c,
	0x07, 0x97, 0x3f, 0xea, 0xd8, 0x0a, 0xe2, 0x8f, 0x3a, 0xfa, 0x5e, 0x47, 0x32, 0xf8, 0x55, 0xc7,
	0xe8, 0x02, 0x77, 0x84, 0x2e, 0x37, 0xa4, 0xdc, 0xc8, 0x2f, 0x12, 0x46, 0xf5, 0xa0, 0x61, 0xbc,
	0x3c, 0x88, 0x31, 0x7a, 0xdc, 0x5f, 0xe1, 0xd3, 0xd8, 0x9f, 0xc8, 0xf9, 0x67, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x5b, 0x1d, 0x02, 0x86, 0x02, 0x00, 0x00,
}