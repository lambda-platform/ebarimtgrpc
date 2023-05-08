// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: posapi.proto

package posapi

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

// PosApiClient is the client API for PosApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosApiClient interface {
	GetInformation(ctx context.Context, in *GetInformationRequest, opts ...grpc.CallOption) (*InformationOutput, error)
	CheckApi(ctx context.Context, in *CheckApiRequest, opts ...grpc.CallOption) (*CheckOutput, error)
	CallFunction(ctx context.Context, in *CallFunctionRequest, opts ...grpc.CallOption) (*CallFunctionResponse, error)
	Put(ctx context.Context, in *PutInput, opts ...grpc.CallOption) (*PutOutput, error)
	PutBatch(ctx context.Context, in *PutInputBatch, opts ...grpc.CallOption) (*PutOutput, error)
	ReturnBill(ctx context.Context, in *BillInput, opts ...grpc.CallOption) (*BillOutput, error)
	SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*DataOutput, error)
}

type posApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPosApiClient(cc grpc.ClientConnInterface) PosApiClient {
	return &posApiClient{cc}
}

func (c *posApiClient) GetInformation(ctx context.Context, in *GetInformationRequest, opts ...grpc.CallOption) (*InformationOutput, error) {
	out := new(InformationOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/GetInformation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) CheckApi(ctx context.Context, in *CheckApiRequest, opts ...grpc.CallOption) (*CheckOutput, error) {
	out := new(CheckOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/CheckApi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) CallFunction(ctx context.Context, in *CallFunctionRequest, opts ...grpc.CallOption) (*CallFunctionResponse, error) {
	out := new(CallFunctionResponse)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/CallFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) Put(ctx context.Context, in *PutInput, opts ...grpc.CallOption) (*PutOutput, error) {
	out := new(PutOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) PutBatch(ctx context.Context, in *PutInputBatch, opts ...grpc.CallOption) (*PutOutput, error) {
	out := new(PutOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/PutBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) ReturnBill(ctx context.Context, in *BillInput, opts ...grpc.CallOption) (*BillOutput, error) {
	out := new(BillOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/ReturnBill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posApiClient) SendData(ctx context.Context, in *SendDataRequest, opts ...grpc.CallOption) (*DataOutput, error) {
	out := new(DataOutput)
	err := c.cc.Invoke(ctx, "/posapi.PosApi/SendData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosApiServer is the server API for PosApi service.
// All implementations should embed UnimplementedPosApiServer
// for forward compatibility
type PosApiServer interface {
	GetInformation(context.Context, *GetInformationRequest) (*InformationOutput, error)
	CheckApi(context.Context, *CheckApiRequest) (*CheckOutput, error)
	CallFunction(context.Context, *CallFunctionRequest) (*CallFunctionResponse, error)
	Put(context.Context, *PutInput) (*PutOutput, error)
	PutBatch(context.Context, *PutInputBatch) (*PutOutput, error)
	ReturnBill(context.Context, *BillInput) (*BillOutput, error)
	SendData(context.Context, *SendDataRequest) (*DataOutput, error)
}

// UnimplementedPosApiServer should be embedded to have forward compatible implementations.
type UnimplementedPosApiServer struct {
}

func (UnimplementedPosApiServer) GetInformation(context.Context, *GetInformationRequest) (*InformationOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInformation not implemented")
}
func (UnimplementedPosApiServer) CheckApi(context.Context, *CheckApiRequest) (*CheckOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckApi not implemented")
}
func (UnimplementedPosApiServer) CallFunction(context.Context, *CallFunctionRequest) (*CallFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallFunction not implemented")
}
func (UnimplementedPosApiServer) Put(context.Context, *PutInput) (*PutOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedPosApiServer) PutBatch(context.Context, *PutInputBatch) (*PutOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBatch not implemented")
}
func (UnimplementedPosApiServer) ReturnBill(context.Context, *BillInput) (*BillOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReturnBill not implemented")
}
func (UnimplementedPosApiServer) SendData(context.Context, *SendDataRequest) (*DataOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendData not implemented")
}

// UnsafePosApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosApiServer will
// result in compilation errors.
type UnsafePosApiServer interface {
	mustEmbedUnimplementedPosApiServer()
}

func RegisterPosApiServer(s grpc.ServiceRegistrar, srv PosApiServer) {
	s.RegisterService(&PosApi_ServiceDesc, srv)
}

func _PosApi_GetInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).GetInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/GetInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).GetInformation(ctx, req.(*GetInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_CheckApi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckApiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).CheckApi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/CheckApi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).CheckApi(ctx, req.(*CheckApiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_CallFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).CallFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/CallFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).CallFunction(ctx, req.(*CallFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).Put(ctx, req.(*PutInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_PutBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutInputBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).PutBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/PutBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).PutBatch(ctx, req.(*PutInputBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_ReturnBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).ReturnBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/ReturnBill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).ReturnBill(ctx, req.(*BillInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosApi_SendData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosApiServer).SendData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posapi.PosApi/SendData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosApiServer).SendData(ctx, req.(*SendDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosApi_ServiceDesc is the grpc.ServiceDesc for PosApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posapi.PosApi",
	HandlerType: (*PosApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInformation",
			Handler:    _PosApi_GetInformation_Handler,
		},
		{
			MethodName: "CheckApi",
			Handler:    _PosApi_CheckApi_Handler,
		},
		{
			MethodName: "CallFunction",
			Handler:    _PosApi_CallFunction_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _PosApi_Put_Handler,
		},
		{
			MethodName: "PutBatch",
			Handler:    _PosApi_PutBatch_Handler,
		},
		{
			MethodName: "ReturnBill",
			Handler:    _PosApi_ReturnBill_Handler,
		},
		{
			MethodName: "SendData",
			Handler:    _PosApi_SendData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posapi.proto",
}