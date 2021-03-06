// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/auth/v2alpha/external_auth.proto

package v2alpha

import (
	context "context"
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("envoy/service/auth/v2alpha/external_auth.proto", fileDescriptor_878c0ddb0c43de8d)
}

var fileDescriptor_878c0ddb0c43de8d = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f,
	0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x89,
	0x07, 0x89, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x81, 0xd5, 0xeb, 0x41, 0xd5, 0xeb,
	0x81, 0x65, 0xa0, 0xea, 0xa5, 0x34, 0xb1, 0x9a, 0x85, 0xcd, 0x18, 0xa3, 0x64, 0x2e, 0x5e, 0xc7,
	0xd2, 0x92, 0x8c, 0xfc, 0xa2, 0xcc, 0xaa, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0xa1, 0x20, 0x2e, 0x56,
	0xe7, 0x8c, 0xd4, 0xe4, 0x6c, 0x21, 0x65, 0x3d, 0xac, 0x36, 0xe8, 0x81, 0x65, 0x83, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0xa4, 0x54, 0xf0, 0x2b, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x75, 0x0a,
	0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xb9, 0x34, 0x32,
	0xf3, 0x21, 0xba, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x70, 0xfb, 0xc3, 0x89, 0xcb, 0x39, 0xb5,
	0xa8, 0xa4, 0x38, 0x00, 0xe4, 0xd0, 0x00, 0xc6, 0x28, 0x76, 0xa8, 0x70, 0x07, 0x23, 0x63, 0x12,
	0x1b, 0xd8, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xd9, 0x7c, 0x50, 0x36, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *v2.CheckRequest, opts ...grpc.CallOption) (*v2.CheckResponse, error) {
	out := new(v2.CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v2alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	// Performs authorization check based on the attributes associated with the
	// incoming request, and returns status `OK` or not `OK`.
	Check(context.Context, *v2.CheckRequest) (*v2.CheckResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v2alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*v2.CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v2alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v2alpha/external_auth.proto",
}
