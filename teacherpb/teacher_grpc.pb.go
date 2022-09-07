// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: teacherpb/teacher.proto

package teacherpb

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

// TeacherServiceClient is the client API for TeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherServiceClient interface {
	RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*Teacher, error)
	CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*Subject, error)
	GetTeacher(ctx context.Context, in *GetTeacherRequest, opts ...grpc.CallOption) (*Teacher, error)
	GetSubject(ctx context.Context, in *GetSubjectRequest, opts ...grpc.CallOption) (*Subject, error)
	DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTeachers(ctx context.Context, in *ListTeachersRequest, opts ...grpc.CallOption) (*ListTeachersResponse, error)
	ListSubjects(ctx context.Context, in *ListSubjectsRequest, opts ...grpc.CallOption) (*ListSubjectsResponse, error)
}

type teacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherServiceClient(cc grpc.ClientConnInterface) TeacherServiceClient {
	return &teacherServiceClient{cc}
}

func (c *teacherServiceClient) RegisterTeacher(ctx context.Context, in *RegisterTeacherRequest, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/RegisterTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) CreateSubject(ctx context.Context, in *CreateSubjectRequest, opts ...grpc.CallOption) (*Subject, error) {
	out := new(Subject)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/CreateSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetTeacher(ctx context.Context, in *GetTeacherRequest, opts ...grpc.CallOption) (*Teacher, error) {
	out := new(Teacher)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/GetTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetSubject(ctx context.Context, in *GetSubjectRequest, opts ...grpc.CallOption) (*Subject, error) {
	out := new(Subject)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/GetSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) DeleteTeacher(ctx context.Context, in *DeleteTeacherRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/DeleteTeacher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) DeleteSubject(ctx context.Context, in *DeleteSubjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/DeleteSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) ListTeachers(ctx context.Context, in *ListTeachersRequest, opts ...grpc.CallOption) (*ListTeachersResponse, error) {
	out := new(ListTeachersResponse)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/ListTeachers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) ListSubjects(ctx context.Context, in *ListSubjectsRequest, opts ...grpc.CallOption) (*ListSubjectsResponse, error) {
	out := new(ListSubjectsResponse)
	err := c.cc.Invoke(ctx, "/teacherpb.TeacherService/ListSubjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServiceServer is the server API for TeacherService service.
// All implementations must embed UnimplementedTeacherServiceServer
// for forward compatibility
type TeacherServiceServer interface {
	RegisterTeacher(context.Context, *RegisterTeacherRequest) (*Teacher, error)
	CreateSubject(context.Context, *CreateSubjectRequest) (*Subject, error)
	GetTeacher(context.Context, *GetTeacherRequest) (*Teacher, error)
	GetSubject(context.Context, *GetSubjectRequest) (*Subject, error)
	DeleteTeacher(context.Context, *DeleteTeacherRequest) (*emptypb.Empty, error)
	DeleteSubject(context.Context, *DeleteSubjectRequest) (*emptypb.Empty, error)
	ListTeachers(context.Context, *ListTeachersRequest) (*ListTeachersResponse, error)
	ListSubjects(context.Context, *ListSubjectsRequest) (*ListSubjectsResponse, error)
	mustEmbedUnimplementedTeacherServiceServer()
}

// UnimplementedTeacherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeacherServiceServer struct {
}

func (UnimplementedTeacherServiceServer) RegisterTeacher(context.Context, *RegisterTeacherRequest) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) CreateSubject(context.Context, *CreateSubjectRequest) (*Subject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubject not implemented")
}
func (UnimplementedTeacherServiceServer) GetTeacher(context.Context, *GetTeacherRequest) (*Teacher, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) GetSubject(context.Context, *GetSubjectRequest) (*Subject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubject not implemented")
}
func (UnimplementedTeacherServiceServer) DeleteTeacher(context.Context, *DeleteTeacherRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeacher not implemented")
}
func (UnimplementedTeacherServiceServer) DeleteSubject(context.Context, *DeleteSubjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubject not implemented")
}
func (UnimplementedTeacherServiceServer) ListTeachers(context.Context, *ListTeachersRequest) (*ListTeachersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeachers not implemented")
}
func (UnimplementedTeacherServiceServer) ListSubjects(context.Context, *ListSubjectsRequest) (*ListSubjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubjects not implemented")
}
func (UnimplementedTeacherServiceServer) mustEmbedUnimplementedTeacherServiceServer() {}

// UnsafeTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServiceServer will
// result in compilation errors.
type UnsafeTeacherServiceServer interface {
	mustEmbedUnimplementedTeacherServiceServer()
}

func RegisterTeacherServiceServer(s grpc.ServiceRegistrar, srv TeacherServiceServer) {
	s.RegisterService(&TeacherService_ServiceDesc, srv)
}

func _TeacherService_RegisterTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).RegisterTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/RegisterTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).RegisterTeacher(ctx, req.(*RegisterTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_CreateSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).CreateSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/CreateSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).CreateSubject(ctx, req.(*CreateSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/GetTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetTeacher(ctx, req.(*GetTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/GetSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetSubject(ctx, req.(*GetSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_DeleteTeacher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTeacherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).DeleteTeacher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/DeleteTeacher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).DeleteTeacher(ctx, req.(*DeleteTeacherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_DeleteSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).DeleteSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/DeleteSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).DeleteSubject(ctx, req.(*DeleteSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_ListTeachers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTeachersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).ListTeachers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/ListTeachers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).ListTeachers(ctx, req.(*ListTeachersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_ListSubjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSubjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).ListSubjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teacherpb.TeacherService/ListSubjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).ListSubjects(ctx, req.(*ListSubjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeacherService_ServiceDesc is the grpc.ServiceDesc for TeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teacherpb.TeacherService",
	HandlerType: (*TeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterTeacher",
			Handler:    _TeacherService_RegisterTeacher_Handler,
		},
		{
			MethodName: "CreateSubject",
			Handler:    _TeacherService_CreateSubject_Handler,
		},
		{
			MethodName: "GetTeacher",
			Handler:    _TeacherService_GetTeacher_Handler,
		},
		{
			MethodName: "GetSubject",
			Handler:    _TeacherService_GetSubject_Handler,
		},
		{
			MethodName: "DeleteTeacher",
			Handler:    _TeacherService_DeleteTeacher_Handler,
		},
		{
			MethodName: "DeleteSubject",
			Handler:    _TeacherService_DeleteSubject_Handler,
		},
		{
			MethodName: "ListTeachers",
			Handler:    _TeacherService_ListTeachers_Handler,
		},
		{
			MethodName: "ListSubjects",
			Handler:    _TeacherService_ListSubjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teacherpb/teacher.proto",
}
