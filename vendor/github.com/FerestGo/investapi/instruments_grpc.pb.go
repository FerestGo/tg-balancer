// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// InstrumentsServiceClient is the client API for InstrumentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstrumentsServiceClient interface {
	//Метод получения расписания торгов торговых площадок.
	TradingSchedules(ctx context.Context, in *TradingSchedulesRequest, opts ...grpc.CallOption) (*TradingSchedulesResponse, error)
	//Метод получения облигации по её идентификатору.
	BondBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*BondResponse, error)
	//Метод получения списка облигаций.
	Bonds(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*BondsResponse, error)
	//Метод получения валюты по её идентификатору.
	CurrencyBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*CurrencyResponse, error)
	//Метод получения списка валют.
	Currencies(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*CurrenciesResponse, error)
	//Метод получения инвестиционного фонда по его идентификатору.
	EtfBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*EtfResponse, error)
	//Метод получения списка инвестиционных фондов.
	Etfs(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*EtfsResponse, error)
	//Метод получения фьючерса по его идентификатору.
	FutureBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*FutureResponse, error)
	//Метод получения списка фьючерсов.
	Futures(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*FuturesResponse, error)
	//Метод получения акции по её идентификатору.
	ShareBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*ShareResponse, error)
	//Метод получения списка акций.
	Shares(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*SharesResponse, error)
	//Метод получения накопленного купонного дохода по облигации.
	GetAccruedInterests(ctx context.Context, in *GetAccruedInterestsRequest, opts ...grpc.CallOption) (*GetAccruedInterestsResponse, error)
	//Метод получения размера гарантийного обеспечения по фьючерсам.
	GetFuturesMargin(ctx context.Context, in *GetFuturesMarginRequest, opts ...grpc.CallOption) (*GetFuturesMarginResponse, error)
	//Метод получения основной информации об инструменте.
	GetInstrumentBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*InstrumentResponse, error)
	//Метод для получения событий выплаты дивидендов по инструменту.
	GetDividends(ctx context.Context, in *GetDividendsRequest, opts ...grpc.CallOption) (*GetDividendsResponse, error)
}

type instrumentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstrumentsServiceClient(cc grpc.ClientConnInterface) InstrumentsServiceClient {
	return &instrumentsServiceClient{cc}
}

func (c *instrumentsServiceClient) TradingSchedules(ctx context.Context, in *TradingSchedulesRequest, opts ...grpc.CallOption) (*TradingSchedulesResponse, error) {
	out := new(TradingSchedulesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/TradingSchedules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) BondBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*BondResponse, error) {
	out := new(BondResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/BondBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) Bonds(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*BondsResponse, error) {
	out := new(BondsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Bonds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) CurrencyBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*CurrencyResponse, error) {
	out := new(CurrencyResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CurrencyBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) Currencies(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*CurrenciesResponse, error) {
	out := new(CurrenciesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Currencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) EtfBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*EtfResponse, error) {
	out := new(EtfResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EtfBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) Etfs(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*EtfsResponse, error) {
	out := new(EtfsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Etfs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) FutureBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*FutureResponse, error) {
	out := new(FutureResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FutureBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) Futures(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*FuturesResponse, error) {
	out := new(FuturesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Futures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) ShareBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*ShareResponse, error) {
	out := new(ShareResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) Shares(ctx context.Context, in *InstrumentsRequest, opts ...grpc.CallOption) (*SharesResponse, error) {
	out := new(SharesResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Shares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) GetAccruedInterests(ctx context.Context, in *GetAccruedInterestsRequest, opts ...grpc.CallOption) (*GetAccruedInterestsResponse, error) {
	out := new(GetAccruedInterestsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAccruedInterests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) GetFuturesMargin(ctx context.Context, in *GetFuturesMarginRequest, opts ...grpc.CallOption) (*GetFuturesMarginResponse, error) {
	out := new(GetFuturesMarginResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFuturesMargin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) GetInstrumentBy(ctx context.Context, in *InstrumentRequest, opts ...grpc.CallOption) (*InstrumentResponse, error) {
	out := new(InstrumentResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInstrumentBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instrumentsServiceClient) GetDividends(ctx context.Context, in *GetDividendsRequest, opts ...grpc.CallOption) (*GetDividendsResponse, error) {
	out := new(GetDividendsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetDividends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstrumentsServiceServer is the server API for InstrumentsService service.
// All implementations must embed UnimplementedInstrumentsServiceServer
// for forward compatibility
type InstrumentsServiceServer interface {
	//Метод получения расписания торгов торговых площадок.
	TradingSchedules(context.Context, *TradingSchedulesRequest) (*TradingSchedulesResponse, error)
	//Метод получения облигации по её идентификатору.
	BondBy(context.Context, *InstrumentRequest) (*BondResponse, error)
	//Метод получения списка облигаций.
	Bonds(context.Context, *InstrumentsRequest) (*BondsResponse, error)
	//Метод получения валюты по её идентификатору.
	CurrencyBy(context.Context, *InstrumentRequest) (*CurrencyResponse, error)
	//Метод получения списка валют.
	Currencies(context.Context, *InstrumentsRequest) (*CurrenciesResponse, error)
	//Метод получения инвестиционного фонда по его идентификатору.
	EtfBy(context.Context, *InstrumentRequest) (*EtfResponse, error)
	//Метод получения списка инвестиционных фондов.
	Etfs(context.Context, *InstrumentsRequest) (*EtfsResponse, error)
	//Метод получения фьючерса по его идентификатору.
	FutureBy(context.Context, *InstrumentRequest) (*FutureResponse, error)
	//Метод получения списка фьючерсов.
	Futures(context.Context, *InstrumentsRequest) (*FuturesResponse, error)
	//Метод получения акции по её идентификатору.
	ShareBy(context.Context, *InstrumentRequest) (*ShareResponse, error)
	//Метод получения списка акций.
	Shares(context.Context, *InstrumentsRequest) (*SharesResponse, error)
	//Метод получения накопленного купонного дохода по облигации.
	GetAccruedInterests(context.Context, *GetAccruedInterestsRequest) (*GetAccruedInterestsResponse, error)
	//Метод получения размера гарантийного обеспечения по фьючерсам.
	GetFuturesMargin(context.Context, *GetFuturesMarginRequest) (*GetFuturesMarginResponse, error)
	//Метод получения основной информации об инструменте.
	GetInstrumentBy(context.Context, *InstrumentRequest) (*InstrumentResponse, error)
	//Метод для получения событий выплаты дивидендов по инструменту.
	GetDividends(context.Context, *GetDividendsRequest) (*GetDividendsResponse, error)
	mustEmbedUnimplementedInstrumentsServiceServer()
}

// UnimplementedInstrumentsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInstrumentsServiceServer struct {
}

func (UnimplementedInstrumentsServiceServer) TradingSchedules(context.Context, *TradingSchedulesRequest) (*TradingSchedulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradingSchedules not implemented")
}
func (UnimplementedInstrumentsServiceServer) BondBy(context.Context, *InstrumentRequest) (*BondResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BondBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) Bonds(context.Context, *InstrumentsRequest) (*BondsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bonds not implemented")
}
func (UnimplementedInstrumentsServiceServer) CurrencyBy(context.Context, *InstrumentRequest) (*CurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrencyBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) Currencies(context.Context, *InstrumentsRequest) (*CurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currencies not implemented")
}
func (UnimplementedInstrumentsServiceServer) EtfBy(context.Context, *InstrumentRequest) (*EtfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EtfBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) Etfs(context.Context, *InstrumentsRequest) (*EtfsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Etfs not implemented")
}
func (UnimplementedInstrumentsServiceServer) FutureBy(context.Context, *InstrumentRequest) (*FutureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FutureBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) Futures(context.Context, *InstrumentsRequest) (*FuturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Futures not implemented")
}
func (UnimplementedInstrumentsServiceServer) ShareBy(context.Context, *InstrumentRequest) (*ShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) Shares(context.Context, *InstrumentsRequest) (*SharesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shares not implemented")
}
func (UnimplementedInstrumentsServiceServer) GetAccruedInterests(context.Context, *GetAccruedInterestsRequest) (*GetAccruedInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccruedInterests not implemented")
}
func (UnimplementedInstrumentsServiceServer) GetFuturesMargin(context.Context, *GetFuturesMarginRequest) (*GetFuturesMarginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuturesMargin not implemented")
}
func (UnimplementedInstrumentsServiceServer) GetInstrumentBy(context.Context, *InstrumentRequest) (*InstrumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentBy not implemented")
}
func (UnimplementedInstrumentsServiceServer) GetDividends(context.Context, *GetDividendsRequest) (*GetDividendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDividends not implemented")
}
func (UnimplementedInstrumentsServiceServer) mustEmbedUnimplementedInstrumentsServiceServer() {}

// UnsafeInstrumentsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstrumentsServiceServer will
// result in compilation errors.
type UnsafeInstrumentsServiceServer interface {
	mustEmbedUnimplementedInstrumentsServiceServer()
}

func RegisterInstrumentsServiceServer(s grpc.ServiceRegistrar, srv InstrumentsServiceServer) {
	s.RegisterService(&InstrumentsService_ServiceDesc, srv)
}

func _InstrumentsService_TradingSchedules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradingSchedulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).TradingSchedules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/TradingSchedules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).TradingSchedules(ctx, req.(*TradingSchedulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_BondBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).BondBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/BondBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).BondBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_Bonds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).Bonds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Bonds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).Bonds(ctx, req.(*InstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_CurrencyBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).CurrencyBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/CurrencyBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).CurrencyBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_Currencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).Currencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Currencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).Currencies(ctx, req.(*InstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_EtfBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).EtfBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/EtfBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).EtfBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_Etfs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).Etfs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Etfs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).Etfs(ctx, req.(*InstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_FutureBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).FutureBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/FutureBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).FutureBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_Futures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).Futures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Futures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).Futures(ctx, req.(*InstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_ShareBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).ShareBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).ShareBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_Shares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).Shares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/Shares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).Shares(ctx, req.(*InstrumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_GetAccruedInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccruedInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).GetAccruedInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetAccruedInterests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).GetAccruedInterests(ctx, req.(*GetAccruedInterestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_GetFuturesMargin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFuturesMarginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).GetFuturesMargin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetFuturesMargin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).GetFuturesMargin(ctx, req.(*GetFuturesMarginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_GetInstrumentBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstrumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).GetInstrumentBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetInstrumentBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).GetInstrumentBy(ctx, req.(*InstrumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstrumentsService_GetDividends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDividendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstrumentsServiceServer).GetDividends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetDividends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstrumentsServiceServer).GetDividends(ctx, req.(*GetDividendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstrumentsService_ServiceDesc is the grpc.ServiceDesc for InstrumentsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstrumentsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.InstrumentsService",
	HandlerType: (*InstrumentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TradingSchedules",
			Handler:    _InstrumentsService_TradingSchedules_Handler,
		},
		{
			MethodName: "BondBy",
			Handler:    _InstrumentsService_BondBy_Handler,
		},
		{
			MethodName: "Bonds",
			Handler:    _InstrumentsService_Bonds_Handler,
		},
		{
			MethodName: "CurrencyBy",
			Handler:    _InstrumentsService_CurrencyBy_Handler,
		},
		{
			MethodName: "Currencies",
			Handler:    _InstrumentsService_Currencies_Handler,
		},
		{
			MethodName: "EtfBy",
			Handler:    _InstrumentsService_EtfBy_Handler,
		},
		{
			MethodName: "Etfs",
			Handler:    _InstrumentsService_Etfs_Handler,
		},
		{
			MethodName: "FutureBy",
			Handler:    _InstrumentsService_FutureBy_Handler,
		},
		{
			MethodName: "Futures",
			Handler:    _InstrumentsService_Futures_Handler,
		},
		{
			MethodName: "ShareBy",
			Handler:    _InstrumentsService_ShareBy_Handler,
		},
		{
			MethodName: "Shares",
			Handler:    _InstrumentsService_Shares_Handler,
		},
		{
			MethodName: "GetAccruedInterests",
			Handler:    _InstrumentsService_GetAccruedInterests_Handler,
		},
		{
			MethodName: "GetFuturesMargin",
			Handler:    _InstrumentsService_GetFuturesMargin_Handler,
		},
		{
			MethodName: "GetInstrumentBy",
			Handler:    _InstrumentsService_GetInstrumentBy_Handler,
		},
		{
			MethodName: "GetDividends",
			Handler:    _InstrumentsService_GetDividends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "instruments.proto",
}
