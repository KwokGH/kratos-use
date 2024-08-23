// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.22.0
// source: mini/diary/v1/diary.proto

package v1

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
	Diary_ListDiary_FullMethodName   = "/api.mini.diary.v1.Diary/ListDiary"
	Diary_GetDiary_FullMethodName    = "/api.mini.diary.v1.Diary/GetDiary"
	Diary_CreateDiary_FullMethodName = "/api.mini.diary.v1.Diary/CreateDiary"
	Diary_UpdateDiary_FullMethodName = "/api.mini.diary.v1.Diary/UpdateDiary"
	Diary_DeleteDiary_FullMethodName = "/api.mini.diary.v1.Diary/DeleteDiary"
)

// DiaryClient is the client API for Diary service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiaryClient interface {
	// 日志列表
	ListDiary(ctx context.Context, in *ListDiaryReq, opts ...grpc.CallOption) (*ListDiaryResp, error)
	// 日志详情
	GetDiary(ctx context.Context, in *GetDiaryReq, opts ...grpc.CallOption) (*GetDiaryResp, error)
	// 日志创建
	CreateDiary(ctx context.Context, in *CreateDiaryReq, opts ...grpc.CallOption) (*CreateDiaryResp, error)
	// 日志修改
	UpdateDiary(ctx context.Context, in *UpdateDiaryReq, opts ...grpc.CallOption) (*UpdateDiaryResp, error)
	// 日志删除
	DeleteDiary(ctx context.Context, in *DeleteDiaryReq, opts ...grpc.CallOption) (*DeleteDiaryResp, error)
}

type diaryClient struct {
	cc grpc.ClientConnInterface
}

func NewDiaryClient(cc grpc.ClientConnInterface) DiaryClient {
	return &diaryClient{cc}
}

func (c *diaryClient) ListDiary(ctx context.Context, in *ListDiaryReq, opts ...grpc.CallOption) (*ListDiaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDiaryResp)
	err := c.cc.Invoke(ctx, Diary_ListDiary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryClient) GetDiary(ctx context.Context, in *GetDiaryReq, opts ...grpc.CallOption) (*GetDiaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDiaryResp)
	err := c.cc.Invoke(ctx, Diary_GetDiary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryClient) CreateDiary(ctx context.Context, in *CreateDiaryReq, opts ...grpc.CallOption) (*CreateDiaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDiaryResp)
	err := c.cc.Invoke(ctx, Diary_CreateDiary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryClient) UpdateDiary(ctx context.Context, in *UpdateDiaryReq, opts ...grpc.CallOption) (*UpdateDiaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDiaryResp)
	err := c.cc.Invoke(ctx, Diary_UpdateDiary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diaryClient) DeleteDiary(ctx context.Context, in *DeleteDiaryReq, opts ...grpc.CallOption) (*DeleteDiaryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDiaryResp)
	err := c.cc.Invoke(ctx, Diary_DeleteDiary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiaryServer is the server API for Diary service.
// All implementations must embed UnimplementedDiaryServer
// for forward compatibility.
type DiaryServer interface {
	// 日志列表
	ListDiary(context.Context, *ListDiaryReq) (*ListDiaryResp, error)
	// 日志详情
	GetDiary(context.Context, *GetDiaryReq) (*GetDiaryResp, error)
	// 日志创建
	CreateDiary(context.Context, *CreateDiaryReq) (*CreateDiaryResp, error)
	// 日志修改
	UpdateDiary(context.Context, *UpdateDiaryReq) (*UpdateDiaryResp, error)
	// 日志删除
	DeleteDiary(context.Context, *DeleteDiaryReq) (*DeleteDiaryResp, error)
	mustEmbedUnimplementedDiaryServer()
}

// UnimplementedDiaryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiaryServer struct{}

func (UnimplementedDiaryServer) ListDiary(context.Context, *ListDiaryReq) (*ListDiaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDiary not implemented")
}
func (UnimplementedDiaryServer) GetDiary(context.Context, *GetDiaryReq) (*GetDiaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDiary not implemented")
}
func (UnimplementedDiaryServer) CreateDiary(context.Context, *CreateDiaryReq) (*CreateDiaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiary not implemented")
}
func (UnimplementedDiaryServer) UpdateDiary(context.Context, *UpdateDiaryReq) (*UpdateDiaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDiary not implemented")
}
func (UnimplementedDiaryServer) DeleteDiary(context.Context, *DeleteDiaryReq) (*DeleteDiaryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDiary not implemented")
}
func (UnimplementedDiaryServer) mustEmbedUnimplementedDiaryServer() {}
func (UnimplementedDiaryServer) testEmbeddedByValue()               {}

// UnsafeDiaryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiaryServer will
// result in compilation errors.
type UnsafeDiaryServer interface {
	mustEmbedUnimplementedDiaryServer()
}

func RegisterDiaryServer(s grpc.ServiceRegistrar, srv DiaryServer) {
	// If the following call pancis, it indicates UnimplementedDiaryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Diary_ServiceDesc, srv)
}

func _Diary_ListDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServer).ListDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diary_ListDiary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServer).ListDiary(ctx, req.(*ListDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diary_GetDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServer).GetDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diary_GetDiary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServer).GetDiary(ctx, req.(*GetDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diary_CreateDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServer).CreateDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diary_CreateDiary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServer).CreateDiary(ctx, req.(*CreateDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diary_UpdateDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServer).UpdateDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diary_UpdateDiary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServer).UpdateDiary(ctx, req.(*UpdateDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diary_DeleteDiary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDiaryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiaryServer).DeleteDiary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diary_DeleteDiary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiaryServer).DeleteDiary(ctx, req.(*DeleteDiaryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Diary_ServiceDesc is the grpc.ServiceDesc for Diary service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diary_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mini.diary.v1.Diary",
	HandlerType: (*DiaryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDiary",
			Handler:    _Diary_ListDiary_Handler,
		},
		{
			MethodName: "GetDiary",
			Handler:    _Diary_GetDiary_Handler,
		},
		{
			MethodName: "CreateDiary",
			Handler:    _Diary_CreateDiary_Handler,
		},
		{
			MethodName: "UpdateDiary",
			Handler:    _Diary_UpdateDiary_Handler,
		},
		{
			MethodName: "DeleteDiary",
			Handler:    _Diary_DeleteDiary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mini/diary/v1/diary.proto",
}
