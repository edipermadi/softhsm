// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: hsm/hsm.proto

package hsm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SessionManagement_OpenSession_FullMethodName      = "/hsm.SessionManagement/OpenSession"
	SessionManagement_CloseSession_FullMethodName     = "/hsm.SessionManagement/CloseSession"
	SessionManagement_CloseAllSessions_FullMethodName = "/hsm.SessionManagement/CloseAllSessions"
)

// SessionManagementClient is the client API for SessionManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionManagementClient interface {
	OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionResponse, error)
	CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error)
	CloseAllSessions(ctx context.Context, in *CloseAllSessionsRequest, opts ...grpc.CallOption) (*CloseAllSessionsResponse, error)
}

type sessionManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionManagementClient(cc grpc.ClientConnInterface) SessionManagementClient {
	return &sessionManagementClient{cc}
}

func (c *sessionManagementClient) OpenSession(ctx context.Context, in *OpenSessionRequest, opts ...grpc.CallOption) (*OpenSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OpenSessionResponse)
	err := c.cc.Invoke(ctx, SessionManagement_OpenSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagementClient) CloseSession(ctx context.Context, in *CloseSessionRequest, opts ...grpc.CallOption) (*CloseSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseSessionResponse)
	err := c.cc.Invoke(ctx, SessionManagement_CloseSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagementClient) CloseAllSessions(ctx context.Context, in *CloseAllSessionsRequest, opts ...grpc.CallOption) (*CloseAllSessionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CloseAllSessionsResponse)
	err := c.cc.Invoke(ctx, SessionManagement_CloseAllSessions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionManagementServer is the server API for SessionManagement service.
// All implementations must embed UnimplementedSessionManagementServer
// for forward compatibility.
type SessionManagementServer interface {
	OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionResponse, error)
	CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error)
	CloseAllSessions(context.Context, *CloseAllSessionsRequest) (*CloseAllSessionsResponse, error)
	mustEmbedUnimplementedSessionManagementServer()
}

// UnimplementedSessionManagementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSessionManagementServer struct{}

func (UnimplementedSessionManagementServer) OpenSession(context.Context, *OpenSessionRequest) (*OpenSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenSession not implemented")
}
func (UnimplementedSessionManagementServer) CloseSession(context.Context, *CloseSessionRequest) (*CloseSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (UnimplementedSessionManagementServer) CloseAllSessions(context.Context, *CloseAllSessionsRequest) (*CloseAllSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseAllSessions not implemented")
}
func (UnimplementedSessionManagementServer) mustEmbedUnimplementedSessionManagementServer() {}
func (UnimplementedSessionManagementServer) testEmbeddedByValue()                           {}

// UnsafeSessionManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionManagementServer will
// result in compilation errors.
type UnsafeSessionManagementServer interface {
	mustEmbedUnimplementedSessionManagementServer()
}

func RegisterSessionManagementServer(s grpc.ServiceRegistrar, srv SessionManagementServer) {
	// If the following call pancis, it indicates UnimplementedSessionManagementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SessionManagement_ServiceDesc, srv)
}

func _SessionManagement_OpenSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagementServer).OpenSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionManagement_OpenSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagementServer).OpenSession(ctx, req.(*OpenSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManagement_CloseSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagementServer).CloseSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionManagement_CloseSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagementServer).CloseSession(ctx, req.(*CloseSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManagement_CloseAllSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseAllSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagementServer).CloseAllSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SessionManagement_CloseAllSessions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagementServer).CloseAllSessions(ctx, req.(*CloseAllSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionManagement_ServiceDesc is the grpc.ServiceDesc for SessionManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hsm.SessionManagement",
	HandlerType: (*SessionManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenSession",
			Handler:    _SessionManagement_OpenSession_Handler,
		},
		{
			MethodName: "CloseSession",
			Handler:    _SessionManagement_CloseSession_Handler,
		},
		{
			MethodName: "CloseAllSessions",
			Handler:    _SessionManagement_CloseAllSessions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hsm/hsm.proto",
}
