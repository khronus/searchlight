// Code generated by protoc-gen-go.
// source: purchase.proto
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

type PurchaseBeginRequest struct {
	ProductSku string `protobuf:"bytes,1,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	Count      int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *PurchaseBeginRequest) Reset()                    { *m = PurchaseBeginRequest{} }
func (m *PurchaseBeginRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseBeginRequest) ProtoMessage()               {}
func (*PurchaseBeginRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PurchaseBeginRequest) GetProductSku() string {
	if m != nil {
		return m.ProductSku
	}
	return ""
}

func (m *PurchaseBeginRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type PurchaseBeginResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phids  []string                `protobuf:"bytes,2,rep,name=phids" json:"phids,omitempty"`
}

func (m *PurchaseBeginResponse) Reset()                    { *m = PurchaseBeginResponse{} }
func (m *PurchaseBeginResponse) String() string            { return proto.CompactTextString(m) }
func (*PurchaseBeginResponse) ProtoMessage()               {}
func (*PurchaseBeginResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *PurchaseBeginResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PurchaseBeginResponse) GetPhids() []string {
	if m != nil {
		return m.Phids
	}
	return nil
}

type PurchaseCompleteRequest struct {
	Phid       string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	ObjectPhid string            `protobuf:"bytes,2,opt,name=object_phid,json=objectPhid" json:"object_phid,omitempty"`
	Failed     bool              `protobuf:"varint,3,opt,name=failed" json:"failed,omitempty"`
	Metadata   map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PurchaseCompleteRequest) Reset()                    { *m = PurchaseCompleteRequest{} }
func (m *PurchaseCompleteRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseCompleteRequest) ProtoMessage()               {}
func (*PurchaseCompleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *PurchaseCompleteRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *PurchaseCompleteRequest) GetObjectPhid() string {
	if m != nil {
		return m.ObjectPhid
	}
	return ""
}

func (m *PurchaseCompleteRequest) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *PurchaseCompleteRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PurchaseCloseRequest struct {
	ObjectPhid string `protobuf:"bytes,1,opt,name=object_phid,json=objectPhid" json:"object_phid,omitempty"`
}

func (m *PurchaseCloseRequest) Reset()                    { *m = PurchaseCloseRequest{} }
func (m *PurchaseCloseRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseCloseRequest) ProtoMessage()               {}
func (*PurchaseCloseRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *PurchaseCloseRequest) GetObjectPhid() string {
	if m != nil {
		return m.ObjectPhid
	}
	return ""
}

func init() {
	proto.RegisterType((*PurchaseBeginRequest)(nil), "appscode.billing.v1beta1.PurchaseBeginRequest")
	proto.RegisterType((*PurchaseBeginResponse)(nil), "appscode.billing.v1beta1.PurchaseBeginResponse")
	proto.RegisterType((*PurchaseCompleteRequest)(nil), "appscode.billing.v1beta1.PurchaseCompleteRequest")
	proto.RegisterType((*PurchaseCloseRequest)(nil), "appscode.billing.v1beta1.PurchaseCloseRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Purchases service

type PurchasesClient interface {
	Begin(ctx context.Context, in *PurchaseBeginRequest, opts ...grpc.CallOption) (*PurchaseBeginResponse, error)
	Complete(ctx context.Context, in *PurchaseCompleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Close(ctx context.Context, in *PurchaseCloseRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type purchasesClient struct {
	cc *grpc.ClientConn
}

func NewPurchasesClient(cc *grpc.ClientConn) PurchasesClient {
	return &purchasesClient{cc}
}

func (c *purchasesClient) Begin(ctx context.Context, in *PurchaseBeginRequest, opts ...grpc.CallOption) (*PurchaseBeginResponse, error) {
	out := new(PurchaseBeginResponse)
	err := grpc.Invoke(ctx, "/appscode.billing.v1beta1.Purchases/Begin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasesClient) Complete(ctx context.Context, in *PurchaseCompleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.billing.v1beta1.Purchases/Complete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasesClient) Close(ctx context.Context, in *PurchaseCloseRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.billing.v1beta1.Purchases/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Purchases service

type PurchasesServer interface {
	Begin(context.Context, *PurchaseBeginRequest) (*PurchaseBeginResponse, error)
	Complete(context.Context, *PurchaseCompleteRequest) (*appscode_dtypes.VoidResponse, error)
	Close(context.Context, *PurchaseCloseRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterPurchasesServer(s *grpc.Server, srv PurchasesServer) {
	s.RegisterService(&_Purchases_serviceDesc, srv)
}

func _Purchases_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseBeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchasesServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.billing.v1beta1.Purchases/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchasesServer).Begin(ctx, req.(*PurchaseBeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchases_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchasesServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.billing.v1beta1.Purchases/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchasesServer).Complete(ctx, req.(*PurchaseCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchases_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseCloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchasesServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.billing.v1beta1.Purchases/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchasesServer).Close(ctx, req.(*PurchaseCloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Purchases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.billing.v1beta1.Purchases",
	HandlerType: (*PurchasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Begin",
			Handler:    _Purchases_Begin_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _Purchases_Complete_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Purchases_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "purchase.proto",
}

func init() { proto.RegisterFile("purchase.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x94, 0xef, 0x6b, 0xd3, 0x40,
	0x18, 0xc7, 0xb9, 0x64, 0xad, 0xed, 0x53, 0x26, 0x72, 0x4c, 0x17, 0x82, 0xb2, 0x90, 0x37, 0x86,
	0x21, 0x17, 0xda, 0x21, 0xd3, 0xa9, 0x08, 0x1d, 0xbe, 0x1c, 0x94, 0x0c, 0x7c, 0xa1, 0xe0, 0xb8,
	0x26, 0x67, 0x17, 0x9b, 0xe6, 0xce, 0xde, 0x65, 0x50, 0x64, 0x20, 0xfe, 0x0b, 0x22, 0xbe, 0xf1,
	0xbf, 0xf2, 0xad, 0x2f, 0xfd, 0x43, 0x24, 0x97, 0x4b, 0x97, 0x55, 0x46, 0xbb, 0x37, 0xe5, 0x9e,
	0x7b, 0x7e, 0xdc, 0xe7, 0x79, 0xbe, 0x4f, 0x03, 0x77, 0x45, 0x31, 0x8f, 0xcf, 0xa9, 0x64, 0x44,
	0xcc, 0xb9, 0xe2, 0xd8, 0xa1, 0x42, 0xc8, 0x98, 0x27, 0x8c, 0x8c, 0xd3, 0x2c, 0x4b, 0xf3, 0x09,
	0xb9, 0xe8, 0x8f, 0x99, 0xa2, 0x7d, 0xf7, 0xe1, 0x84, 0xf3, 0x49, 0xc6, 0x42, 0x2a, 0xd2, 0x90,
	0xe6, 0x39, 0x57, 0x54, 0xa5, 0x3c, 0x97, 0x55, 0x9e, 0xbb, 0x57, 0xe7, 0x69, 0x7f, 0xa2, 0x16,
	0x82, 0xc9, 0x50, 0xff, 0x56, 0x01, 0xfe, 0x09, 0xec, 0x8c, 0xcc, 0x53, 0x43, 0x36, 0x49, 0xf3,
	0x88, 0x7d, 0x2e, 0x98, 0x54, 0x78, 0x0f, 0x7a, 0x62, 0xce, 0x93, 0x22, 0x56, 0x67, 0x72, 0x5a,
	0x38, 0xc8, 0x43, 0x41, 0x37, 0x02, 0x73, 0x75, 0x3a, 0x2d, 0xf0, 0x0e, 0xb4, 0x62, 0x5e, 0xe4,
	0xca, 0xb1, 0x3c, 0x14, 0xb4, 0xa2, 0xca, 0xf0, 0x3f, 0xc0, 0xfd, 0x95, 0x72, 0x52, 0xf0, 0x5c,
	0x32, 0x1c, 0x42, 0x5b, 0x2a, 0xaa, 0x0a, 0xa9, 0x4b, 0xf5, 0x06, 0xbb, 0x64, 0xd9, 0x51, 0x45,
	0x45, 0x4e, 0xb5, 0x3b, 0x32, 0x61, 0x65, 0x7d, 0x71, 0x9e, 0x26, 0xd2, 0xb1, 0x3c, 0x3b, 0xe8,
	0x46, 0x95, 0xe1, 0x7f, 0xb5, 0x60, 0xb7, 0x7e, 0xe0, 0x98, 0xcf, 0x44, 0xc6, 0x14, 0xab, 0x91,
	0x31, 0x6c, 0x95, 0x41, 0x86, 0x55, 0x9f, 0xcb, 0x36, 0xf8, 0xf8, 0x13, 0x8b, 0xd5, 0x99, 0x76,
	0x59, 0x55, 0x1b, 0xd5, 0xd5, 0xa8, 0x0c, 0x78, 0x00, 0xed, 0x8f, 0x34, 0xcd, 0x58, 0xe2, 0xd8,
	0x1e, 0x0a, 0x3a, 0x91, 0xb1, 0xf0, 0x7b, 0xe8, 0xcc, 0x98, 0xa2, 0x09, 0x55, 0xd4, 0xd9, 0xf2,
	0xec, 0xa0, 0x37, 0x78, 0x4d, 0x6e, 0xd2, 0x80, 0xdc, 0x40, 0x44, 0x4e, 0x4c, 0x85, 0x37, 0xb9,
	0x9a, 0x2f, 0xa2, 0x65, 0x41, 0xf7, 0x05, 0x6c, 0x5f, 0x73, 0xe1, 0x7b, 0x60, 0x4f, 0xd9, 0xc2,
	0x90, 0x97, 0xc7, 0xb2, 0xfd, 0x0b, 0x9a, 0x15, 0xcc, 0x20, 0x57, 0xc6, 0x91, 0xf5, 0x0c, 0xf9,
	0x87, 0x57, 0x8a, 0x1d, 0x67, 0x5c, 0xb2, 0x86, 0x62, 0xcd, 0x56, 0xd1, 0x6a, 0xab, 0x83, 0x3f,
	0x36, 0x74, 0xeb, 0x4c, 0x89, 0x7f, 0x21, 0x68, 0x69, 0x89, 0x30, 0x59, 0xdf, 0x58, 0x73, 0x35,
	0xdc, 0x70, 0xe3, 0xf8, 0x4a, 0x7b, 0x7f, 0xf0, 0xed, 0xf7, 0xdf, 0xef, 0xd6, 0x13, 0xff, 0x71,
	0x78, 0x6d, 0x1b, 0x4d, 0x72, 0x68, 0x92, 0xc3, 0x7a, 0xe5, 0xe5, 0x11, 0xda, 0xc7, 0x3f, 0x11,
	0x74, 0xea, 0x71, 0xe2, 0xfe, 0xad, 0x47, 0xef, 0x3e, 0xfa, 0x6f, 0xbf, 0xde, 0xf2, 0x34, 0x59,
	0x22, 0x3d, 0xd7, 0x48, 0x07, 0x2e, 0xd9, 0x10, 0x29, 0xfc, 0x52, 0x0e, 0xf5, 0xb2, 0x24, 0xfb,
	0x81, 0xa0, 0xa5, 0x07, 0xbf, 0xc9, 0xe0, 0x9a, 0x0a, 0xad, 0x63, 0x7a, 0xa5, 0x99, 0x0e, 0xf7,
	0x9f, 0x6e, 0xcc, 0xd4, 0xd0, 0xfb, 0x72, 0xf8, 0x12, 0xbc, 0x98, 0xcf, 0xae, 0x9e, 0xa0, 0x22,
	0x5d, 0xc5, 0x1a, 0x6e, 0xd7, 0x5c, 0xa3, 0xf2, 0xcf, 0x3f, 0x42, 0xef, 0xee, 0x18, 0xcf, 0xb8,
	0xad, 0x3f, 0x07, 0x07, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x69, 0xaa, 0x78, 0x79, 0x04,
	0x00, 0x00,
}