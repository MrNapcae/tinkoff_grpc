// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: marketdata.proto

package investapi

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

// MarketDataServiceClient is the client API for MarketDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataServiceClient interface {
	//Метод запроса исторических свечей по инструменту.
	GetCandles(ctx context.Context, in *GetCandlesRequest, opts ...grpc.CallOption) (*GetCandlesResponse, error)
	//Метод запроса последних цен по инструментам.
	GetLastPrices(ctx context.Context, in *GetLastPricesRequest, opts ...grpc.CallOption) (*GetLastPricesResponse, error)
	//Метод получения стакана по инструменту.
	GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error)
	//Метод запроса статуса торгов по инструментам.
	GetTradingStatus(ctx context.Context, in *GetTradingStatusRequest, opts ...grpc.CallOption) (*GetTradingStatusResponse, error)
}

type marketDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataServiceClient(cc grpc.ClientConnInterface) MarketDataServiceClient {
	return &marketDataServiceClient{cc}
}

func (c *marketDataServiceClient) GetCandles(ctx context.Context, in *GetCandlesRequest, opts ...grpc.CallOption) (*GetCandlesResponse, error) {
	out := new(GetCandlesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetLastPrices(ctx context.Context, in *GetLastPricesRequest, opts ...grpc.CallOption) (*GetLastPricesResponse, error) {
	out := new(GetLastPricesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error) {
	out := new(GetOrderBookResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketDataServiceClient) GetTradingStatus(ctx context.Context, in *GetTradingStatusRequest, opts ...grpc.CallOption) (*GetTradingStatusResponse, error) {
	out := new(GetTradingStatusResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketDataServiceServer is the server API for MarketDataService service.
// All implementations must embed UnimplementedMarketDataServiceServer
// for forward compatibility
type MarketDataServiceServer interface {
	//Метод запроса исторических свечей по инструменту.
	GetCandles(context.Context, *GetCandlesRequest) (*GetCandlesResponse, error)
	//Метод запроса последних цен по инструментам.
	GetLastPrices(context.Context, *GetLastPricesRequest) (*GetLastPricesResponse, error)
	//Метод получения стакана по инструменту.
	GetOrderBook(context.Context, *GetOrderBookRequest) (*GetOrderBookResponse, error)
	//Метод запроса статуса торгов по инструментам.
	GetTradingStatus(context.Context, *GetTradingStatusRequest) (*GetTradingStatusResponse, error)
	mustEmbedUnimplementedMarketDataServiceServer()
}

// UnimplementedMarketDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataServiceServer struct {
}

func (UnimplementedMarketDataServiceServer) GetCandles(context.Context, *GetCandlesRequest) (*GetCandlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCandles not implemented")
}
func (UnimplementedMarketDataServiceServer) GetLastPrices(context.Context, *GetLastPricesRequest) (*GetLastPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastPrices not implemented")
}
func (UnimplementedMarketDataServiceServer) GetOrderBook(context.Context, *GetOrderBookRequest) (*GetOrderBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderBook not implemented")
}
func (UnimplementedMarketDataServiceServer) GetTradingStatus(context.Context, *GetTradingStatusRequest) (*GetTradingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTradingStatus not implemented")
}
func (UnimplementedMarketDataServiceServer) mustEmbedUnimplementedMarketDataServiceServer() {}

// UnsafeMarketDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataServiceServer will
// result in compilation errors.
type UnsafeMarketDataServiceServer interface {
	mustEmbedUnimplementedMarketDataServiceServer()
}

func RegisterMarketDataServiceServer(s grpc.ServiceRegistrar, srv MarketDataServiceServer) {
	s.RegisterService(&MarketDataService_ServiceDesc, srv)
}

func _MarketDataService_GetCandles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCandlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetCandles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetCandles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetCandles(ctx, req.(*GetCandlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetLastPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetLastPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetLastPrices(ctx, req.(*GetLastPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetOrderBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetOrderBook(ctx, req.(*GetOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketDataService_GetTradingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketDataServiceServer).GetTradingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.MarketDataService/GetTradingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketDataServiceServer).GetTradingStatus(ctx, req.(*GetTradingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketDataService_ServiceDesc is the grpc.ServiceDesc for MarketDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.MarketDataService",
	HandlerType: (*MarketDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCandles",
			Handler:    _MarketDataService_GetCandles_Handler,
		},
		{
			MethodName: "GetLastPrices",
			Handler:    _MarketDataService_GetLastPrices_Handler,
		},
		{
			MethodName: "GetOrderBook",
			Handler:    _MarketDataService_GetOrderBook_Handler,
		},
		{
			MethodName: "GetTradingStatus",
			Handler:    _MarketDataService_GetTradingStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketdata.proto",
}

// MarketDataStreamServiceClient is the client API for MarketDataStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketDataStreamServiceClient interface {
	//Bi-directional стрим предоставления биржевой информации.
	MarketDataStream(ctx context.Context, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataStreamClient, error)
}

type marketDataStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketDataStreamServiceClient(cc grpc.ClientConnInterface) MarketDataStreamServiceClient {
	return &marketDataStreamServiceClient{cc}
}

func (c *marketDataStreamServiceClient) MarketDataStream(ctx context.Context, opts ...grpc.CallOption) (MarketDataStreamService_MarketDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &MarketDataStreamService_ServiceDesc.Streams[0], "/tinkoff.public.invest.api.contract.v1.MarketDataStreamService/MarketDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &marketDataStreamServiceMarketDataStreamClient{stream}
	return x, nil
}

type MarketDataStreamService_MarketDataStreamClient interface {
	Send(*MarketDataRequest) error
	Recv() (*MarketDataResponse, error)
	grpc.ClientStream
}

type marketDataStreamServiceMarketDataStreamClient struct {
	grpc.ClientStream
}

func (x *marketDataStreamServiceMarketDataStreamClient) Send(m *MarketDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *marketDataStreamServiceMarketDataStreamClient) Recv() (*MarketDataResponse, error) {
	m := new(MarketDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketDataStreamServiceServer is the server API for MarketDataStreamService service.
// All implementations must embed UnimplementedMarketDataStreamServiceServer
// for forward compatibility
type MarketDataStreamServiceServer interface {
	//Bi-directional стрим предоставления биржевой информации.
	MarketDataStream(MarketDataStreamService_MarketDataStreamServer) error
	mustEmbedUnimplementedMarketDataStreamServiceServer()
}

// UnimplementedMarketDataStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketDataStreamServiceServer struct {
}

func (UnimplementedMarketDataStreamServiceServer) MarketDataStream(MarketDataStreamService_MarketDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MarketDataStream not implemented")
}
func (UnimplementedMarketDataStreamServiceServer) mustEmbedUnimplementedMarketDataStreamServiceServer() {
}

// UnsafeMarketDataStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketDataStreamServiceServer will
// result in compilation errors.
type UnsafeMarketDataStreamServiceServer interface {
	mustEmbedUnimplementedMarketDataStreamServiceServer()
}

func RegisterMarketDataStreamServiceServer(s grpc.ServiceRegistrar, srv MarketDataStreamServiceServer) {
	s.RegisterService(&MarketDataStreamService_ServiceDesc, srv)
}

func _MarketDataStreamService_MarketDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MarketDataStreamServiceServer).MarketDataStream(&marketDataStreamServiceMarketDataStreamServer{stream})
}

type MarketDataStreamService_MarketDataStreamServer interface {
	Send(*MarketDataResponse) error
	Recv() (*MarketDataRequest, error)
	grpc.ServerStream
}

type marketDataStreamServiceMarketDataStreamServer struct {
	grpc.ServerStream
}

func (x *marketDataStreamServiceMarketDataStreamServer) Send(m *MarketDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *marketDataStreamServiceMarketDataStreamServer) Recv() (*MarketDataRequest, error) {
	m := new(MarketDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MarketDataStreamService_ServiceDesc is the grpc.ServiceDesc for MarketDataStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketDataStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.MarketDataStreamService",
	HandlerType: (*MarketDataStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MarketDataStream",
			Handler:       _MarketDataStreamService_MarketDataStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "marketdata.proto",
}
