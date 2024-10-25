// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: log.proto

package eslog

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
	Log_Debug_FullMethodName = "/eslog.Log/Debug"
	Log_Info_FullMethodName  = "/eslog.Log/Info"
	Log_Warn_FullMethodName  = "/eslog.Log/Warn"
	Log_Error_FullMethodName = "/eslog.Log/Error"
	Log_Fatal_FullMethodName = "/eslog.Log/Fatal"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	Debug(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
	Info(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
	Warn(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
	Error(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
	Fatal(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Debug(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogReply)
	err := c.cc.Invoke(ctx, Log_Debug_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Info(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogReply)
	err := c.cc.Invoke(ctx, Log_Info_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Warn(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogReply)
	err := c.cc.Invoke(ctx, Log_Warn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Error(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogReply)
	err := c.cc.Invoke(ctx, Log_Error_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Fatal(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogReply)
	err := c.cc.Invoke(ctx, Log_Fatal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility.
type LogServer interface {
	Debug(context.Context, *LogRequest) (*LogReply, error)
	Info(context.Context, *LogRequest) (*LogReply, error)
	Warn(context.Context, *LogRequest) (*LogReply, error)
	Error(context.Context, *LogRequest) (*LogReply, error)
	Fatal(context.Context, *LogRequest) (*LogReply, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServer struct{}

func (UnimplementedLogServer) Debug(context.Context, *LogRequest) (*LogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Debug not implemented")
}
func (UnimplementedLogServer) Info(context.Context, *LogRequest) (*LogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedLogServer) Warn(context.Context, *LogRequest) (*LogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Warn not implemented")
}
func (UnimplementedLogServer) Error(context.Context, *LogRequest) (*LogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Error not implemented")
}
func (UnimplementedLogServer) Fatal(context.Context, *LogRequest) (*LogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fatal not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}
func (UnimplementedLogServer) testEmbeddedByValue()             {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	// If the following call pancis, it indicates UnimplementedLogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_Debug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Debug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Debug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Debug(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Info(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Warn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Warn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Warn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Warn(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Error_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Error(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Error_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Error(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Fatal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Fatal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Fatal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Fatal(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eslog.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Debug",
			Handler:    _Log_Debug_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Log_Info_Handler,
		},
		{
			MethodName: "Warn",
			Handler:    _Log_Warn_Handler,
		},
		{
			MethodName: "Error",
			Handler:    _Log_Error_Handler,
		},
		{
			MethodName: "Fatal",
			Handler:    _Log_Fatal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
