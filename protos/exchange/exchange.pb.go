// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/exchange/exchange.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/bittorrent/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("protos/exchange/exchange.proto", fileDescriptor_b8e4cc6fcea5eff3) }

var fileDescriptor_b8e4cc6fcea5eff3 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x40, 0x2f, 0xa2, 0x25, 0x07, 0x95, 0x1c, 0xc4, 0xed, 0x41, 0x57, 0x0f, 0x1e, 0x2b, 0xe8,
	0x5d, 0x28, 0xab, 0x67, 0x3f, 0x51, 0x3c, 0x88, 0xc4, 0x38, 0xda, 0x80, 0x9b, 0x89, 0x99, 0xa9,
	0x1f, 0x3f, 0xd1, 0x7f, 0x25, 0xd4, 0x64, 0xbb, 0x66, 0xb3, 0x3d, 0xb5, 0xf0, 0x5e, 0x5e, 0x86,
	0x8c, 0xd8, 0x71, 0x1e, 0x19, 0xe9, 0x10, 0xbe, 0x74, 0xa3, 0xec, 0x2b, 0xcc, 0x7e, 0xaa, 0x0e,
	0xc8, 0xa2, 0xfb, 0x68, 0x7c, 0x2b, 0x0f, 0x96, 0x99, 0x8f, 0x53, 0x20, 0x52, 0xf1, 0xc4, 0xd1,
	0xcf, 0x8a, 0x28, 0xce, 0x02, 0x92, 0xb7, 0x62, 0xe3, 0xc2, 0x83, 0x53, 0x1e, 0xee, 0x0c, 0x37,
	0xcf, 0x5e, 0x7d, 0xca, 0x71, 0x15, 0x93, 0x55, 0x82, 0xae, 0xe0, 0xbd, 0x05, 0xe2, 0x72, 0x6f,
	0xc0, 0x20, 0x87, 0x96, 0x40, 0xd6, 0xa2, 0x98, 0x05, 0x47, 0xbd, 0x9e, 0x96, 0xca, 0x1c, 0x0a,
	0x89, 0x6b, 0xb1, 0x1e, 0xea, 0xa7, 0xe0, 0x90, 0x0c, 0xcb, 0xdd, 0x85, 0x7b, 0x03, 0x89, 0xb9,
	0xf1, 0x72, 0x21, 0x44, 0x4f, 0xc4, 0x5a, 0xac, 0x6d, 0xf7, 0x72, 0x92, 0x19, 0x65, 0x48, 0x3f,
	0xd4, 0x04, 0xed, 0x8b, 0xf1, 0xd3, 0xcc, 0x50, 0xff, 0x49, 0x66, 0xa8, 0x54, 0x08, 0xd1, 0x7b,
	0xb1, 0x79, 0xd9, 0x82, 0xff, 0xbe, 0xf1, 0xca, 0x92, 0xd2, 0x6c, 0xd0, 0xca, 0xb9, 0x37, 0x4e,
	0x59, 0x0c, 0xef, 0x0f, 0x29, 0x21, 0xfd, 0x20, 0xb6, 0x6a, 0xcd, 0xe6, 0x43, 0x31, 0xd4, 0x5a,
	0x63, 0x6b, 0xf9, 0xdc, 0x4e, 0x1a, 0x65, 0xec, 0xfc, 0x9a, 0x13, 0x23, 0xb3, 0xe6, 0x05, 0xe3,
	0x2f, 0xff, 0xb4, 0xda, 0x19, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xf4, 0x2e, 0x8c,
	0xa6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExchangeClient is the client API for Exchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExchangeClient interface {
	PrepareWithdraw(ctx context.Context, in *PrepareWithdrawRequest, opts ...grpc.CallOption) (*PrepareWithdrawResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	PrepareDeposit(ctx context.Context, in *PrepareDepositRequest, opts ...grpc.CallOption) (*PrepareDepositResponse, error)
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error)
	ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error)
	QueryTransaction(ctx context.Context, in *QueryTransactionRequest, opts ...grpc.CallOption) (*QueryTransactionResponse, error)
	ActivateAccountOnChain(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error)
}

type exchangeClient struct {
	cc *grpc.ClientConn
}

func NewExchangeClient(cc *grpc.ClientConn) ExchangeClient {
	return &exchangeClient{cc}
}

func (c *exchangeClient) PrepareWithdraw(ctx context.Context, in *PrepareWithdrawRequest, opts ...grpc.CallOption) (*PrepareWithdrawResponse, error) {
	out := new(PrepareWithdrawResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/PrepareWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) PrepareDeposit(ctx context.Context, in *PrepareDepositRequest, opts ...grpc.CallOption) (*PrepareDepositResponse, error) {
	out := new(PrepareDepositResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/PrepareDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositResponse, error) {
	out := new(DepositResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error) {
	out := new(ConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/ConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) QueryTransaction(ctx context.Context, in *QueryTransactionRequest, opts ...grpc.CallOption) (*QueryTransactionResponse, error) {
	out := new(QueryTransactionResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/QueryTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exchangeClient) ActivateAccountOnChain(ctx context.Context, in *ActivateAccountRequest, opts ...grpc.CallOption) (*ActivateAccountResponse, error) {
	out := new(ActivateAccountResponse)
	err := c.cc.Invoke(ctx, "/protocol.Exchange/ActivateAccountOnChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExchangeServer is the server API for Exchange service.
type ExchangeServer interface {
	PrepareWithdraw(context.Context, *PrepareWithdrawRequest) (*PrepareWithdrawResponse, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	PrepareDeposit(context.Context, *PrepareDepositRequest) (*PrepareDepositResponse, error)
	Deposit(context.Context, *DepositRequest) (*DepositResponse, error)
	ConfirmDeposit(context.Context, *ConfirmDepositRequest) (*ConfirmDepositResponse, error)
	QueryTransaction(context.Context, *QueryTransactionRequest) (*QueryTransactionResponse, error)
	ActivateAccountOnChain(context.Context, *ActivateAccountRequest) (*ActivateAccountResponse, error)
}

// UnimplementedExchangeServer can be embedded to have forward compatible implementations.
type UnimplementedExchangeServer struct {
}

func (*UnimplementedExchangeServer) PrepareWithdraw(ctx context.Context, req *PrepareWithdrawRequest) (*PrepareWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareWithdraw not implemented")
}
func (*UnimplementedExchangeServer) Withdraw(ctx context.Context, req *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (*UnimplementedExchangeServer) PrepareDeposit(ctx context.Context, req *PrepareDepositRequest) (*PrepareDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDeposit not implemented")
}
func (*UnimplementedExchangeServer) Deposit(ctx context.Context, req *DepositRequest) (*DepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (*UnimplementedExchangeServer) ConfirmDeposit(ctx context.Context, req *ConfirmDepositRequest) (*ConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeposit not implemented")
}
func (*UnimplementedExchangeServer) QueryTransaction(ctx context.Context, req *QueryTransactionRequest) (*QueryTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryTransaction not implemented")
}
func (*UnimplementedExchangeServer) ActivateAccountOnChain(ctx context.Context, req *ActivateAccountRequest) (*ActivateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateAccountOnChain not implemented")
}

func RegisterExchangeServer(s *grpc.Server, srv ExchangeServer) {
	s.RegisterService(&_Exchange_serviceDesc, srv)
}

func _Exchange_PrepareWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).PrepareWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/PrepareWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).PrepareWithdraw(ctx, req.(*PrepareWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_PrepareDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).PrepareDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/PrepareDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).PrepareDeposit(ctx, req.(*PrepareDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_ConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).ConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/ConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).ConfirmDeposit(ctx, req.(*ConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_QueryTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).QueryTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/QueryTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).QueryTransaction(ctx, req.(*QueryTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Exchange_ActivateAccountOnChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExchangeServer).ActivateAccountOnChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Exchange/ActivateAccountOnChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExchangeServer).ActivateAccountOnChain(ctx, req.(*ActivateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exchange_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Exchange",
	HandlerType: (*ExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareWithdraw",
			Handler:    _Exchange_PrepareWithdraw_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Exchange_Withdraw_Handler,
		},
		{
			MethodName: "PrepareDeposit",
			Handler:    _Exchange_PrepareDeposit_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Exchange_Deposit_Handler,
		},
		{
			MethodName: "ConfirmDeposit",
			Handler:    _Exchange_ConfirmDeposit_Handler,
		},
		{
			MethodName: "QueryTransaction",
			Handler:    _Exchange_QueryTransaction_Handler,
		},
		{
			MethodName: "ActivateAccountOnChain",
			Handler:    _Exchange_ActivateAccountOnChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/exchange/exchange.proto",
}
