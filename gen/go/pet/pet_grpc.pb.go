// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pet/pet.proto

package petrosh_pet_v1_petv1

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

// GoodsHandClient is the client API for GoodsHand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsHandClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsResponse, error)
	HistoryGoods(ctx context.Context, in *HistoryGoodsRequest, opts ...grpc.CallOption) (*HistoryGoodsResponse, error)
}

type goodsHandClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsHandClient(cc grpc.ClientConnInterface) GoodsHandClient {
	return &goodsHandClient{cc}
}

func (c *goodsHandClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsHandClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsHandClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsHandClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/GetGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsHandClient) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...grpc.CallOption) (*ListGoodsResponse, error) {
	out := new(ListGoodsResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/ListGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsHandClient) HistoryGoods(ctx context.Context, in *HistoryGoodsRequest, opts ...grpc.CallOption) (*HistoryGoodsResponse, error) {
	out := new(HistoryGoodsResponse)
	err := c.cc.Invoke(ctx, "/goodsHand.GoodsHand/HistoryGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsHandServer is the server API for GoodsHand service.
// All implementations must embed UnimplementedGoodsHandServer
// for forward compatibility
type GoodsHandServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsResponse, error)
	HistoryGoods(context.Context, *HistoryGoodsRequest) (*HistoryGoodsResponse, error)
	mustEmbedUnimplementedGoodsHandServer()
}

// UnimplementedGoodsHandServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsHandServer struct {
}

func (UnimplementedGoodsHandServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGoodsHandServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGoodsHandServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGoodsHandServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodsHandServer) ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGoods not implemented")
}
func (UnimplementedGoodsHandServer) HistoryGoods(context.Context, *HistoryGoodsRequest) (*HistoryGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HistoryGoods not implemented")
}
func (UnimplementedGoodsHandServer) mustEmbedUnimplementedGoodsHandServer() {}

// UnsafeGoodsHandServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsHandServer will
// result in compilation errors.
type UnsafeGoodsHandServer interface {
	mustEmbedUnimplementedGoodsHandServer()
}

func RegisterGoodsHandServer(s grpc.ServiceRegistrar, srv GoodsHandServer) {
	s.RegisterService(&GoodsHand_ServiceDesc, srv)
}

func _GoodsHand_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsHand_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsHand_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsHand_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/GetGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsHand_ListGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).ListGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/ListGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).ListGoods(ctx, req.(*ListGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoodsHand_HistoryGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsHandServer).HistoryGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goodsHand.GoodsHand/HistoryGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsHandServer).HistoryGoods(ctx, req.(*HistoryGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoodsHand_ServiceDesc is the grpc.ServiceDesc for GoodsHand service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoodsHand_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goodsHand.GoodsHand",
	HandlerType: (*GoodsHandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GoodsHand_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GoodsHand_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GoodsHand_Delete_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _GoodsHand_GetGoods_Handler,
		},
		{
			MethodName: "ListGoods",
			Handler:    _GoodsHand_ListGoods_Handler,
		},
		{
			MethodName: "HistoryGoods",
			Handler:    _GoodsHand_HistoryGoods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet/pet.proto",
}
