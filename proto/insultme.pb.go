// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insultme.proto

/*
Package kickme is a generated protocol buffer package.

It is generated from these files:
	insultme.proto

It has these top-level messages:
	InsultRequest
	InsultResponse
*/
package kickme

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InsultRequest struct {
}

func (m *InsultRequest) Reset()                    { *m = InsultRequest{} }
func (m *InsultRequest) String() string            { return proto.CompactTextString(m) }
func (*InsultRequest) ProtoMessage()               {}
func (*InsultRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type InsultResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *InsultResponse) Reset()                    { *m = InsultResponse{} }
func (m *InsultResponse) String() string            { return proto.CompactTextString(m) }
func (*InsultResponse) ProtoMessage()               {}
func (*InsultResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InsultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InsultRequest)(nil), "kickme.InsultRequest")
	proto.RegisterType((*InsultResponse)(nil), "kickme.InsultResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Insult service

type InsultClient interface {
	// Sends a greeting
	GoForIt(ctx context.Context, in *InsultRequest, opts ...grpc.CallOption) (*InsultResponse, error)
}

type insultClient struct {
	cc *grpc.ClientConn
}

func NewInsultClient(cc *grpc.ClientConn) InsultClient {
	return &insultClient{cc}
}

func (c *insultClient) GoForIt(ctx context.Context, in *InsultRequest, opts ...grpc.CallOption) (*InsultResponse, error) {
	out := new(InsultResponse)
	err := grpc.Invoke(ctx, "/kickme.Insult/GoForIt", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Insult service

type InsultServer interface {
	// Sends a greeting
	GoForIt(context.Context, *InsultRequest) (*InsultResponse, error)
}

func RegisterInsultServer(s *grpc.Server, srv InsultServer) {
	s.RegisterService(&_Insult_serviceDesc, srv)
}

func _Insult_GoForIt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsultServer).GoForIt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kickme.Insult/GoForIt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsultServer).GoForIt(ctx, req.(*InsultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Insult_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kickme.Insult",
	HandlerType: (*InsultServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GoForIt",
			Handler:    _Insult_GoForIt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "insultme.proto",
}

func init() { proto.RegisterFile("insultme.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x2e,
	0xcd, 0x29, 0xc9, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xce, 0x4c, 0xce,
	0xce, 0x4d, 0x55, 0xe2, 0xe7, 0xe2, 0xf5, 0x04, 0xcb, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x28, 0x69, 0x71, 0xf1, 0xc1, 0x04, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c,
	0x23, 0x17, 0x2e, 0x36, 0x88, 0x5a, 0x21, 0x2b, 0x2e, 0x76, 0xf7, 0x7c, 0xb7, 0xfc, 0x22, 0xcf,
	0x12, 0x21, 0x51, 0x3d, 0x88, 0xd1, 0x7a, 0x28, 0xe6, 0x4a, 0x89, 0xa1, 0x0b, 0x43, 0x4c, 0x57,
	0x62, 0x48, 0x62, 0x03, 0xbb, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x56, 0x4a, 0xb3,
	0xa3, 0x00, 0x00, 0x00,
}