// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: protos/service.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MailVerifier_SyntaxVerification_FullMethodName = "/MailVerifier/SyntaxVerification"
	MailVerifier_SimpleVerification_FullMethodName = "/MailVerifier/SimpleVerification"
	MailVerifier_FullVerification_FullMethodName   = "/MailVerifier/FullVerification"
)

// MailVerifierClient is the client API for MailVerifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailVerifierClient interface {
	SyntaxVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error)
	SimpleVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error)
	FullVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error)
}

type mailVerifierClient struct {
	cc grpc.ClientConnInterface
}

func NewMailVerifierClient(cc grpc.ClientConnInterface) MailVerifierClient {
	return &mailVerifierClient{cc}
}

func (c *mailVerifierClient) SyntaxVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, MailVerifier_SyntaxVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailVerifierClient) SimpleVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, MailVerifier_SimpleVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailVerifierClient) FullVerification(ctx context.Context, in *VerificationRequest, opts ...grpc.CallOption) (*VerificationResponse, error) {
	out := new(VerificationResponse)
	err := c.cc.Invoke(ctx, MailVerifier_FullVerification_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailVerifierServer is the server API for MailVerifier service.
// All implementations must embed UnimplementedMailVerifierServer
// for forward compatibility
type MailVerifierServer interface {
	SyntaxVerification(context.Context, *VerificationRequest) (*VerificationResponse, error)
	SimpleVerification(context.Context, *VerificationRequest) (*VerificationResponse, error)
	FullVerification(context.Context, *VerificationRequest) (*VerificationResponse, error)
	mustEmbedUnimplementedMailVerifierServer()
}

// UnimplementedMailVerifierServer must be embedded to have forward compatible implementations.
type UnimplementedMailVerifierServer struct {
}

func (UnimplementedMailVerifierServer) SyntaxVerification(context.Context, *VerificationRequest) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyntaxVerification not implemented")
}
func (UnimplementedMailVerifierServer) SimpleVerification(context.Context, *VerificationRequest) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleVerification not implemented")
}
func (UnimplementedMailVerifierServer) FullVerification(context.Context, *VerificationRequest) (*VerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FullVerification not implemented")
}
func (UnimplementedMailVerifierServer) mustEmbedUnimplementedMailVerifierServer() {}

// UnsafeMailVerifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailVerifierServer will
// result in compilation errors.
type UnsafeMailVerifierServer interface {
	mustEmbedUnimplementedMailVerifierServer()
}

func RegisterMailVerifierServer(s grpc.ServiceRegistrar, srv MailVerifierServer) {
	s.RegisterService(&MailVerifier_ServiceDesc, srv)
}

func _MailVerifier_SyntaxVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailVerifierServer).SyntaxVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailVerifier_SyntaxVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailVerifierServer).SyntaxVerification(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailVerifier_SimpleVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailVerifierServer).SimpleVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailVerifier_SimpleVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailVerifierServer).SimpleVerification(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailVerifier_FullVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailVerifierServer).FullVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailVerifier_FullVerification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailVerifierServer).FullVerification(ctx, req.(*VerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailVerifier_ServiceDesc is the grpc.ServiceDesc for MailVerifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailVerifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MailVerifier",
	HandlerType: (*MailVerifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyntaxVerification",
			Handler:    _MailVerifier_SyntaxVerification_Handler,
		},
		{
			MethodName: "SimpleVerification",
			Handler:    _MailVerifier_SimpleVerification_Handler,
		},
		{
			MethodName: "FullVerification",
			Handler:    _MailVerifier_FullVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service.proto",
}
