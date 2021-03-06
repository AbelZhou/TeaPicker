// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Tree.Service.proto

package tree

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	JsonData             string   `protobuf:"bytes,3,opt,name=jsonData,proto3" json:"jsonData,omitempty"`
	Chinese              string   `protobuf:"bytes,4,opt,name=chinese,proto3" json:"chinese,omitempty"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{0}
}

func (m *AddRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *AddRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AddRequest) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *AddRequest) GetChinese() string {
	if m != nil {
		return m.Chinese
	}
	return ""
}

//正确与否回应
type BooleanResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *BooleanResponse) Reset()         { *m = BooleanResponse{} }
func (m *BooleanResponse) String() string { return proto.CompactTextString(m) }
func (*BooleanResponse) ProtoMessage()    {}
func (*BooleanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{1}
}

func (m *BooleanResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type DelRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Chinese              string   `protobuf:"bytes,3,opt,name=chinese,proto3" json:"chinese,omitempty"`
}

func (m *DelRequest) Reset()         { *m = DelRequest{} }
func (m *DelRequest) String() string { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()    {}
func (*DelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{2}
}

func (m *DelRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DelRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelRequest) GetChinese() string {
	if m != nil {
		return m.Chinese
	}
	return ""
}

type SearchRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Context              string   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{3}
}

func (m *SearchRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SearchRequest) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

type SearchLenovosRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Context              string   `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
}

func (m *SearchLenovosRequest) Reset()         { *m = SearchLenovosRequest{} }
func (m *SearchLenovosRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLenovosRequest) ProtoMessage()    {}
func (*SearchLenovosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{4}
}

func (m *SearchLenovosRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SearchLenovosRequest) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

//搜索响应体
type SearchResponse struct {
	Ret                  []*SearchResult `protobuf:"bytes,1,rep,name=ret,proto3" json:"ret,omitempty"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{5}
}

func (m *SearchResponse) GetRet() []*SearchResult {
	if m != nil {
		return m.Ret
	}
	return nil
}

type SearchResult struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}
func (*SearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5056a8ab65a04467, []int{6}
}

func (m *SearchResult) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *SearchResult) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "tree.AddRequest")
	proto.RegisterType((*BooleanResponse)(nil), "tree.BooleanResponse")
	proto.RegisterType((*DelRequest)(nil), "tree.DelRequest")
	proto.RegisterType((*SearchRequest)(nil), "tree.SearchRequest")
	proto.RegisterType((*SearchLenovosRequest)(nil), "tree.SearchLenovosRequest")
	proto.RegisterType((*SearchResponse)(nil), "tree.SearchResponse")
	proto.RegisterType((*SearchResult)(nil), "tree.SearchResult")
}

func init() { proto.RegisterFile("Tree.Service.proto", fileDescriptor_5056a8ab65a04467) }

var fileDescriptor_5056a8ab65a04467 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x4f, 0xea, 0x40,
	0x10, 0x0e, 0xaf, 0x3c, 0x1e, 0x6f, 0xfc, 0x45, 0x46, 0x34, 0x4d, 0x4f, 0xa6, 0xf1, 0x80, 0x97,
	0x1e, 0xc0, 0x78, 0x32, 0x31, 0x35, 0x8d, 0x27, 0xe3, 0xa1, 0x78, 0xf1, 0x58, 0xd9, 0x49, 0x00,
	0x9b, 0x5d, 0xdc, 0x5d, 0x08, 0xfe, 0xf7, 0x66, 0x7f, 0x49, 0x21, 0xc1, 0x44, 0x6f, 0x33, 0xdf,
	0xee, 0x7c, 0xdf, 0x7c, 0x3b, 0xb3, 0x80, 0xcf, 0x92, 0x28, 0x1b, 0x93, 0x5c, 0xcd, 0x26, 0x94,
	0x2d, 0xa4, 0xd0, 0x02, 0xdb, 0x5a, 0x12, 0xa5, 0x73, 0x80, 0x9c, 0xb1, 0x92, 0xde, 0x97, 0xa4,
	0x34, 0xf6, 0xe1, 0xef, 0x8c, 0x33, 0x5a, 0xc7, 0xad, 0x8b, 0xd6, 0xe0, 0x7f, 0xe9, 0x12, 0xec,
	0x41, 0xf4, 0x46, 0x1f, 0xf1, 0x1f, 0x8b, 0x99, 0x10, 0x13, 0xe8, 0xce, 0x95, 0xe0, 0x45, 0xa5,
	0xab, 0x38, 0xb2, 0xf0, 0x57, 0x8e, 0x31, 0xfc, 0x9b, 0x4c, 0x67, 0x9c, 0x14, 0xc5, 0x6d, 0x7b,
	0x14, 0xd2, 0xf4, 0x0a, 0x4e, 0xee, 0x85, 0xa8, 0xa9, 0xe2, 0x25, 0xa9, 0x85, 0xe0, 0x8a, 0xf0,
	0x1c, 0x3a, 0x92, 0xd4, 0xb2, 0xd6, 0x56, 0xb1, 0x5b, 0xfa, 0x2c, 0x7d, 0x02, 0x28, 0xa8, 0xfe,
	0x69, 0x5b, 0x0d, 0xe9, 0x68, 0x5b, 0xfa, 0x0e, 0x8e, 0xc6, 0x54, 0xc9, 0xc9, 0xf4, 0x7b, 0x4a,
	0x43, 0x20, 0xb8, 0xa6, 0xb5, 0xf6, 0xb4, 0x21, 0x4d, 0x1f, 0xa0, 0xef, 0x08, 0x1e, 0x89, 0x8b,
	0x95, 0x50, 0xbf, 0xe5, 0xb9, 0x81, 0xe3, 0xd0, 0x88, 0x7f, 0x82, 0x4b, 0x88, 0x24, 0x19, 0xff,
	0xd1, 0xe0, 0x60, 0x88, 0x99, 0x76, 0x93, 0xf2, 0x57, 0x96, 0xb5, 0x2e, 0xcd, 0x71, 0x7a, 0x0d,
	0x87, 0x4d, 0xd0, 0x98, 0x57, 0x5a, 0x7a, 0x55, 0x13, 0x22, 0x42, 0x9b, 0x99, 0x79, 0x38, 0x41,
	0x1b, 0x0f, 0x8b, 0x60, 0xdb, 0x8f, 0x1e, 0x47, 0xd0, 0x71, 0x00, 0x9e, 0x6e, 0x2b, 0x59, 0x37,
	0x49, 0x7f, 0x47, 0xde, 0x76, 0x38, 0x7c, 0xd9, 0xf1, 0x1e, 0xc8, 0xf2, 0xc0, 0xee, 0x71, 0x4c,
	0x9a, 0xe5, 0xdb, 0x0f, 0xb5, 0x87, 0xfa, 0xd6, 0xce, 0x39, 0x10, 0x66, 0x10, 0x15, 0x54, 0x63,
	0xcf, 0x5d, 0xdd, 0x2c, 0x40, 0x72, 0xe6, 0x90, 0x9d, 0xed, 0x31, 0xd5, 0x39, 0x63, 0x8d, 0xea,
	0x9c, 0xb1, 0x50, 0xbd, 0xd9, 0xea, 0x3d, 0xd5, 0xaf, 0x1d, 0xfb, 0x0f, 0x46, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x0d, 0x19, 0x92, 0x2a, 0x1d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/tree.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tree.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tree.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Tree.Service.proto",
}

// SearchLenovosServiceClient is the client API for SearchLenovosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchLenovosServiceClient interface {
	SearchLenovos(ctx context.Context, in *SearchLenovosRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type searchLenovosServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchLenovosServiceClient(cc *grpc.ClientConn) SearchLenovosServiceClient {
	return &searchLenovosServiceClient{cc}
}

func (c *searchLenovosServiceClient) SearchLenovos(ctx context.Context, in *SearchLenovosRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/tree.SearchLenovosService/SearchLenovos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchLenovosServiceServer is the server API for SearchLenovosService service.
type SearchLenovosServiceServer interface {
	SearchLenovos(context.Context, *SearchLenovosRequest) (*SearchResponse, error)
}

func RegisterSearchLenovosServiceServer(s *grpc.Server, srv SearchLenovosServiceServer) {
	s.RegisterService(&_SearchLenovosService_serviceDesc, srv)
}

func _SearchLenovosService_SearchLenovos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLenovosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchLenovosServiceServer).SearchLenovos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tree.SearchLenovosService/SearchLenovos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchLenovosServiceServer).SearchLenovos(ctx, req.(*SearchLenovosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchLenovosService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tree.SearchLenovosService",
	HandlerType: (*SearchLenovosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchLenovos",
			Handler:    _SearchLenovosService_SearchLenovos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Tree.Service.proto",
}

// DelServiceClient is the client API for DelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DelServiceClient interface {
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
}

type delServiceClient struct {
	cc *grpc.ClientConn
}

func NewDelServiceClient(cc *grpc.ClientConn) DelServiceClient {
	return &delServiceClient{cc}
}

func (c *delServiceClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/tree.DelService/Del", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DelServiceServer is the server API for DelService service.
type DelServiceServer interface {
	Del(context.Context, *DelRequest) (*BooleanResponse, error)
}

func RegisterDelServiceServer(s *grpc.Server, srv DelServiceServer) {
	s.RegisterService(&_DelService_serviceDesc, srv)
}

func _DelService_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DelServiceServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tree.DelService/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DelServiceServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tree.DelService",
	HandlerType: (*DelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Del",
			Handler:    _DelService_Del_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Tree.Service.proto",
}

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*BooleanResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*BooleanResponse, error) {
	out := new(BooleanResponse)
	err := c.cc.Invoke(ctx, "/tree.AddService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	Add(context.Context, *AddRequest) (*BooleanResponse, error)
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tree.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tree.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Tree.Service.proto",
}
