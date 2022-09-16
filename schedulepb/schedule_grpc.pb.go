// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: schedulepb/schedule.proto

package schedulepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScheduleServiceClient is the client API for ScheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScheduleServiceClient interface {
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*Schedule, error)
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*Schedule, error)
	UpdateSchedule(ctx context.Context, in *Schedule, opts ...grpc.CallOption) (*Schedule, error)
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFullScheduleForGroup(ctx context.Context, in *GetFullScheduleForGroupRequest, opts ...grpc.CallOption) (*ScheduleList, error)
	GetSpecificDateScheduleForGroup(ctx context.Context, in *GetSpecificDateScheduleForGroupRequest, opts ...grpc.CallOption) (*ScheduleList, error)
	GetFullScheduleForTeacher(ctx context.Context, in *GetFullScheduleForTeacherRequest, opts ...grpc.CallOption) (*ScheduleList, error)
	GetSpecificDateScheduleForTeacher(ctx context.Context, in *GetSpecificDateScheduleForTeacherRequest, opts ...grpc.CallOption) (*ScheduleList, error)
}

type scheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScheduleServiceClient(cc grpc.ClientConnInterface) ScheduleServiceClient {
	return &scheduleServiceClient{cc}
}

func (c *scheduleServiceClient) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/CreateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/GetSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) UpdateSchedule(ctx context.Context, in *Schedule, opts ...grpc.CallOption) (*Schedule, error) {
	out := new(Schedule)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/UpdateSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/DeleteSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetFullScheduleForGroup(ctx context.Context, in *GetFullScheduleForGroupRequest, opts ...grpc.CallOption) (*ScheduleList, error) {
	out := new(ScheduleList)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/GetFullScheduleForGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetSpecificDateScheduleForGroup(ctx context.Context, in *GetSpecificDateScheduleForGroupRequest, opts ...grpc.CallOption) (*ScheduleList, error) {
	out := new(ScheduleList)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/GetSpecificDateScheduleForGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetFullScheduleForTeacher(ctx context.Context, in *GetFullScheduleForTeacherRequest, opts ...grpc.CallOption) (*ScheduleList, error) {
	out := new(ScheduleList)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/GetFullScheduleForTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleServiceClient) GetSpecificDateScheduleForTeacher(ctx context.Context, in *GetSpecificDateScheduleForTeacherRequest, opts ...grpc.CallOption) (*ScheduleList, error) {
	out := new(ScheduleList)
	err := c.cc.Invoke(ctx, "/schedulepb.ScheduleService/GetSpecificDateScheduleForTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScheduleServiceServer is the server API for ScheduleService service.
// All implementations must embed UnimplementedScheduleServiceServer
// for forward compatibility
type ScheduleServiceServer interface {
	CreateSchedule(context.Context, *CreateScheduleRequest) (*Schedule, error)
	GetSchedule(context.Context, *GetScheduleRequest) (*Schedule, error)
	UpdateSchedule(context.Context, *Schedule) (*Schedule, error)
	DeleteSchedule(context.Context, *DeleteScheduleRequest) (*emptypb.Empty, error)
	GetFullScheduleForGroup(context.Context, *GetFullScheduleForGroupRequest) (*ScheduleList, error)
	GetSpecificDateScheduleForGroup(context.Context, *GetSpecificDateScheduleForGroupRequest) (*ScheduleList, error)
	GetFullScheduleForTeacher(context.Context, *GetFullScheduleForTeacherRequest) (*ScheduleList, error)
	GetSpecificDateScheduleForTeacher(context.Context, *GetSpecificDateScheduleForTeacherRequest) (*ScheduleList, error)
	mustEmbedUnimplementedScheduleServiceServer()
}

// UnimplementedScheduleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScheduleServiceServer struct {
}

func (UnimplementedScheduleServiceServer) CreateSchedule(context.Context, *CreateScheduleRequest) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) GetSchedule(context.Context, *GetScheduleRequest) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) UpdateSchedule(context.Context, *Schedule) (*Schedule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) DeleteSchedule(context.Context, *DeleteScheduleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchedule not implemented")
}
func (UnimplementedScheduleServiceServer) GetFullScheduleForGroup(context.Context, *GetFullScheduleForGroupRequest) (*ScheduleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullScheduleForGroup not implemented")
}
func (UnimplementedScheduleServiceServer) GetSpecificDateScheduleForGroup(context.Context, *GetSpecificDateScheduleForGroupRequest) (*ScheduleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecificDateScheduleForGroup not implemented")
}
func (UnimplementedScheduleServiceServer) GetFullScheduleForTeacher(context.Context, *GetFullScheduleForTeacherRequest) (*ScheduleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullScheduleForTeacher not implemented")
}
func (UnimplementedScheduleServiceServer) GetSpecificDateScheduleForTeacher(context.Context, *GetSpecificDateScheduleForTeacherRequest) (*ScheduleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecificDateScheduleForTeacher not implemented")
}
func (UnimplementedScheduleServiceServer) mustEmbedUnimplementedScheduleServiceServer() {}

// UnsafeScheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScheduleServiceServer will
// result in compilation errors.
type UnsafeScheduleServiceServer interface {
	mustEmbedUnimplementedScheduleServiceServer()
}

func RegisterScheduleServiceServer(s grpc.ServiceRegistrar, srv ScheduleServiceServer) {
	s.RegisterService(&ScheduleService_ServiceDesc, srv)
}

func _ScheduleService_CreateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/CreateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).CreateSchedule(ctx, req.(*CreateScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/GetSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSchedule(ctx, req.(*GetScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_UpdateSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Schedule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).UpdateSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/UpdateSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).UpdateSchedule(ctx, req.(*Schedule))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_DeleteSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).DeleteSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/DeleteSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).DeleteSchedule(ctx, req.(*DeleteScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetFullScheduleForGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullScheduleForGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetFullScheduleForGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/GetFullScheduleForGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetFullScheduleForGroup(ctx, req.(*GetFullScheduleForGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetSpecificDateScheduleForGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificDateScheduleForGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSpecificDateScheduleForGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/GetSpecificDateScheduleForGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSpecificDateScheduleForGroup(ctx, req.(*GetSpecificDateScheduleForGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetFullScheduleForTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFullScheduleForTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetFullScheduleForTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/GetFullScheduleForTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetFullScheduleForTeacher(ctx, req.(*GetFullScheduleForTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScheduleService_GetSpecificDateScheduleForTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecificDateScheduleForTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScheduleServiceServer).GetSpecificDateScheduleForTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/schedulepb.ScheduleService/GetSpecificDateScheduleForTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScheduleServiceServer).GetSpecificDateScheduleForTeacher(ctx, req.(*GetSpecificDateScheduleForTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScheduleService_ServiceDesc is the grpc.ServiceDesc for ScheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schedulepb.ScheduleService",
	HandlerType: (*ScheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchedule",
			Handler:    _ScheduleService_CreateSchedule_Handler,
		},
		{
			MethodName: "GetSchedule",
			Handler:    _ScheduleService_GetSchedule_Handler,
		},
		{
			MethodName: "UpdateSchedule",
			Handler:    _ScheduleService_UpdateSchedule_Handler,
		},
		{
			MethodName: "DeleteSchedule",
			Handler:    _ScheduleService_DeleteSchedule_Handler,
		},
		{
			MethodName: "GetFullScheduleForGroup",
			Handler:    _ScheduleService_GetFullScheduleForGroup_Handler,
		},
		{
			MethodName: "GetSpecificDateScheduleForGroup",
			Handler:    _ScheduleService_GetSpecificDateScheduleForGroup_Handler,
		},
		{
			MethodName: "GetFullScheduleForTeacher",
			Handler:    _ScheduleService_GetFullScheduleForTeacher_Handler,
		},
		{
			MethodName: "GetSpecificDateScheduleForTeacher",
			Handler:    _ScheduleService_GetSpecificDateScheduleForTeacher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schedulepb/schedule.proto",
}
