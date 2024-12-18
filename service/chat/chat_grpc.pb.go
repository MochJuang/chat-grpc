// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: chat.proto

package chat

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
	ChatService_AddMessageToConversation_FullMethodName = "/chat.ChatService/AddMessageToConversation"
	ChatService_GetConversationDetails_FullMethodName   = "/chat.ChatService/GetConversationDetails"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	AddMessageToConversation(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error)
	GetConversationDetails(ctx context.Context, in *ConversationRequest, opts ...grpc.CallOption) (*ConversationResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) AddMessageToConversation(ctx context.Context, in *AddMessageRequest, opts ...grpc.CallOption) (*AddMessageResponse, error) {
	out := new(AddMessageResponse)
	err := c.cc.Invoke(ctx, ChatService_AddMessageToConversation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetConversationDetails(ctx context.Context, in *ConversationRequest, opts ...grpc.CallOption) (*ConversationResponse, error) {
	out := new(ConversationResponse)
	err := c.cc.Invoke(ctx, ChatService_GetConversationDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	AddMessageToConversation(context.Context, *AddMessageRequest) (*AddMessageResponse, error)
	GetConversationDetails(context.Context, *ConversationRequest) (*ConversationResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) AddMessageToConversation(context.Context, *AddMessageRequest) (*AddMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMessageToConversation not implemented")
}
func (UnimplementedChatServiceServer) GetConversationDetails(context.Context, *ConversationRequest) (*ConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationDetails not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_AddMessageToConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).AddMessageToConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_AddMessageToConversation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).AddMessageToConversation(ctx, req.(*AddMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetConversationDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetConversationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChatService_GetConversationDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetConversationDetails(ctx, req.(*ConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMessageToConversation",
			Handler:    _ChatService_AddMessageToConversation_Handler,
		},
		{
			MethodName: "GetConversationDetails",
			Handler:    _ChatService_GetConversationDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
