// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/services.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ticker struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ticker) Reset()         { *m = Ticker{} }
func (m *Ticker) String() string { return proto.CompactTextString(m) }
func (*Ticker) ProtoMessage()    {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f1dc4323d296bf0, []int{0}
}

func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticker.Unmarshal(m, b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
}
func (m *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(m, src)
}
func (m *Ticker) XXX_Size() int {
	return xxx_messageInfo_Ticker.Size(m)
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Stock struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Eps                  string   `protobuf:"bytes,2,opt,name=eps,proto3" json:"eps,omitempty"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stock) Reset()         { *m = Stock{} }
func (m *Stock) String() string { return proto.CompactTextString(m) }
func (*Stock) ProtoMessage()    {}
func (*Stock) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f1dc4323d296bf0, []int{1}
}

func (m *Stock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stock.Unmarshal(m, b)
}
func (m *Stock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stock.Marshal(b, m, deterministic)
}
func (m *Stock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stock.Merge(m, src)
}
func (m *Stock) XXX_Size() int {
	return xxx_messageInfo_Stock.Size(m)
}
func (m *Stock) XXX_DiscardUnknown() {
	xxx_messageInfo_Stock.DiscardUnknown(m)
}

var xxx_messageInfo_Stock proto.InternalMessageInfo

func (m *Stock) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stock) GetEps() string {
	if m != nil {
		return m.Eps
	}
	return ""
}

func (m *Stock) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

type AllStocks struct {
	Data                 []*Stock `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllStocks) Reset()         { *m = AllStocks{} }
func (m *AllStocks) String() string { return proto.CompactTextString(m) }
func (*AllStocks) ProtoMessage()    {}
func (*AllStocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f1dc4323d296bf0, []int{2}
}

func (m *AllStocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllStocks.Unmarshal(m, b)
}
func (m *AllStocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllStocks.Marshal(b, m, deterministic)
}
func (m *AllStocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllStocks.Merge(m, src)
}
func (m *AllStocks) XXX_Size() int {
	return xxx_messageInfo_AllStocks.Size(m)
}
func (m *AllStocks) XXX_DiscardUnknown() {
	xxx_messageInfo_AllStocks.DiscardUnknown(m)
}

var xxx_messageInfo_AllStocks proto.InternalMessageInfo

func (m *AllStocks) GetData() []*Stock {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Ticker)(nil), "services.Ticker")
	proto.RegisterType((*Stock)(nil), "services.Stock")
	proto.RegisterType((*AllStocks)(nil), "services.allStocks")
}

func init() { proto.RegisterFile("protobuf/services.proto", fileDescriptor_8f1dc4323d296bf0) }

var fileDescriptor_8f1dc4323d296bf0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8f, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x13, 0xd3, 0x86, 0x76, 0x14, 0x2c, 0xab, 0x68, 0x88, 0x1e, 0xca, 0x7a, 0xe9, 0x69,
	0xa3, 0xf1, 0xe4, 0xc1, 0x83, 0x88, 0x77, 0xa9, 0xbe, 0xc0, 0x36, 0x3b, 0x86, 0xd0, 0xd4, 0x0d,
	0xbb, 0xab, 0xe0, 0x33, 0xf9, 0x92, 0x92, 0x59, 0xbb, 0x05, 0x89, 0xd0, 0xdb, 0xfc, 0xf9, 0xe6,
	0x9b, 0xdf, 0x07, 0xe7, 0x9d, 0xd1, 0x4e, 0xaf, 0x3e, 0xde, 0x0a, 0x8b, 0xe6, 0xb3, 0xa9, 0xd0,
	0x0a, 0x9a, 0xb0, 0xc9, 0xb6, 0xcf, 0x2f, 0x6a, 0xad, 0xeb, 0x16, 0x8b, 0xa0, 0xc4, 0x4d, 0xe7,
	0xbe, 0xbc, 0x8c, 0x5f, 0x42, 0xfa, 0xda, 0x54, 0x6b, 0x34, 0x8c, 0xc1, 0xe8, 0x5d, 0x6e, 0x30,
	0x8b, 0xe7, 0xf1, 0x62, 0xba, 0xa4, 0x9a, 0x3f, 0xc2, 0xf8, 0xc5, 0xe9, 0x6a, 0x3d, 0xb4, 0x64,
	0x33, 0x48, 0xb0, 0xb3, 0xd9, 0x01, 0x8d, 0xfa, 0x92, 0x9d, 0xc2, 0xb8, 0x33, 0x4d, 0x85, 0x59,
	0x42, 0x33, 0xdf, 0xf0, 0x6b, 0x98, 0xca, 0xb6, 0x25, 0x1f, 0xcb, 0xae, 0x60, 0xa4, 0xa4, 0x93,
	0x59, 0x3c, 0x4f, 0x16, 0x87, 0xe5, 0xb1, 0x08, 0xd4, 0xb4, 0x5f, 0xd2, 0xb2, 0xfc, 0x8e, 0x21,
	0xfd, 0xd5, 0x17, 0x30, 0xa9, 0xd1, 0x79, 0x88, 0xd9, 0x4e, 0xed, 0x99, 0xf3, 0xbf, 0xf7, 0x3c,
	0x62, 0xf7, 0x70, 0x54, 0xa3, 0x7b, 0x08, 0x0f, 0xcf, 0x84, 0x8f, 0x2f, 0xb6, 0xf1, 0xc5, 0x53,
	0x1f, 0x3f, 0x3f, 0xd9, 0x9d, 0x06, 0x3a, 0x1e, 0xb1, 0x1b, 0x48, 0xa4, 0x52, 0x03, 0xaf, 0xfe,
	0xf1, 0xe1, 0x51, 0x79, 0x07, 0xe9, 0xb3, 0x34, 0x16, 0x4d, 0x0f, 0x2b, 0x95, 0xa2, 0x66, 0x2f,
	0xd8, 0x55, 0x4a, 0x66, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xd4, 0xcd, 0x46, 0xc6,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StocksClient is the client API for Stocks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StocksClient interface {
	GetStock(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Stock, error)
	GetAllStocks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllStocks, error)
	Add(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*empty.Empty, error)
}

type stocksClient struct {
	cc *grpc.ClientConn
}

func NewStocksClient(cc *grpc.ClientConn) StocksClient {
	return &stocksClient{cc}
}

func (c *stocksClient) GetStock(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Stock, error) {
	out := new(Stock)
	err := c.cc.Invoke(ctx, "/services.Stocks/getStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stocksClient) GetAllStocks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*AllStocks, error) {
	out := new(AllStocks)
	err := c.cc.Invoke(ctx, "/services.Stocks/getAllStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stocksClient) Add(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/services.Stocks/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StocksServer is the server API for Stocks service.
type StocksServer interface {
	GetStock(context.Context, *Ticker) (*Stock, error)
	GetAllStocks(context.Context, *empty.Empty) (*AllStocks, error)
	Add(context.Context, *Ticker) (*empty.Empty, error)
}

// UnimplementedStocksServer can be embedded to have forward compatible implementations.
type UnimplementedStocksServer struct {
}

func (*UnimplementedStocksServer) GetStock(ctx context.Context, req *Ticker) (*Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStock not implemented")
}
func (*UnimplementedStocksServer) GetAllStocks(ctx context.Context, req *empty.Empty) (*AllStocks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStocks not implemented")
}
func (*UnimplementedStocksServer) Add(ctx context.Context, req *Ticker) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterStocksServer(s *grpc.Server, srv StocksServer) {
	s.RegisterService(&_Stocks_serviceDesc, srv)
}

func _Stocks_GetStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).GetStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Stocks/GetStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).GetStock(ctx, req.(*Ticker))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stocks_GetAllStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).GetAllStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Stocks/GetAllStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).GetAllStocks(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stocks_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StocksServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Stocks/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StocksServer).Add(ctx, req.(*Ticker))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stocks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.Stocks",
	HandlerType: (*StocksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStock",
			Handler:    _Stocks_GetStock_Handler,
		},
		{
			MethodName: "getAllStocks",
			Handler:    _Stocks_GetAllStocks_Handler,
		},
		{
			MethodName: "add",
			Handler:    _Stocks_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/services.proto",
}

// ParserClient is the client API for Parser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ParserClient interface {
	AddParse(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Stock, error)
}

type parserClient struct {
	cc *grpc.ClientConn
}

func NewParserClient(cc *grpc.ClientConn) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) AddParse(ctx context.Context, in *Ticker, opts ...grpc.CallOption) (*Stock, error) {
	out := new(Stock)
	err := c.cc.Invoke(ctx, "/services.Parser/addParse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServer is the server API for Parser service.
type ParserServer interface {
	AddParse(context.Context, *Ticker) (*Stock, error)
}

// UnimplementedParserServer can be embedded to have forward compatible implementations.
type UnimplementedParserServer struct {
}

func (*UnimplementedParserServer) AddParse(ctx context.Context, req *Ticker) (*Stock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddParse not implemented")
}

func RegisterParserServer(s *grpc.Server, srv ParserServer) {
	s.RegisterService(&_Parser_serviceDesc, srv)
}

func _Parser_AddParse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ticker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServer).AddParse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Parser/AddParse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServer).AddParse(ctx, req.(*Ticker))
	}
	return interceptor(ctx, in, info, handler)
}

var _Parser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.Parser",
	HandlerType: (*ParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addParse",
			Handler:    _Parser_AddParse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/services.proto",
}
