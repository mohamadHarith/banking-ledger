// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.2
// source: shared/proto/transaction_logger_proto/transaction_logger_service.proto

package transaction_logger_proto

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

// TransactionLoggerServiceClient is the client API for TransactionLoggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionLoggerServiceClient interface {
	GetTransactionLogs(ctx context.Context, in *GetTransactionLogsRequest, opts ...grpc.CallOption) (*GetTransactionLogsResponse, error)
}

type transactionLoggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionLoggerServiceClient(cc grpc.ClientConnInterface) TransactionLoggerServiceClient {
	return &transactionLoggerServiceClient{cc}
}

func (c *transactionLoggerServiceClient) GetTransactionLogs(ctx context.Context, in *GetTransactionLogsRequest, opts ...grpc.CallOption) (*GetTransactionLogsResponse, error) {
	out := new(GetTransactionLogsResponse)
	err := c.cc.Invoke(ctx, "/TransactionLoggerService/GetTransactionLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionLoggerServiceServer is the server API for TransactionLoggerService service.
// All implementations must embed UnimplementedTransactionLoggerServiceServer
// for forward compatibility
type TransactionLoggerServiceServer interface {
	GetTransactionLogs(context.Context, *GetTransactionLogsRequest) (*GetTransactionLogsResponse, error)
	mustEmbedUnimplementedTransactionLoggerServiceServer()
}

// UnimplementedTransactionLoggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionLoggerServiceServer struct {
}

func (UnimplementedTransactionLoggerServiceServer) GetTransactionLogs(context.Context, *GetTransactionLogsRequest) (*GetTransactionLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionLogs not implemented")
}
func (UnimplementedTransactionLoggerServiceServer) mustEmbedUnimplementedTransactionLoggerServiceServer() {
}

// UnsafeTransactionLoggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionLoggerServiceServer will
// result in compilation errors.
type UnsafeTransactionLoggerServiceServer interface {
	mustEmbedUnimplementedTransactionLoggerServiceServer()
}

func RegisterTransactionLoggerServiceServer(s grpc.ServiceRegistrar, srv TransactionLoggerServiceServer) {
	s.RegisterService(&TransactionLoggerService_ServiceDesc, srv)
}

func _TransactionLoggerService_GetTransactionLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionLoggerServiceServer).GetTransactionLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TransactionLoggerService/GetTransactionLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionLoggerServiceServer).GetTransactionLogs(ctx, req.(*GetTransactionLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionLoggerService_ServiceDesc is the grpc.ServiceDesc for TransactionLoggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionLoggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionLoggerService",
	HandlerType: (*TransactionLoggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionLogs",
			Handler:    _TransactionLoggerService_GetTransactionLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shared/proto/transaction_logger_proto/transaction_logger_service.proto",
}
