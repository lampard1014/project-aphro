// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Commodity.proto

package Aphro_Commodity_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CommodityQueryRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	MerchantID           string   `protobuf:"bytes,2,opt,name=merchantID" json:"merchantID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommodityQueryRequest) Reset()         { *m = CommodityQueryRequest{} }
func (m *CommodityQueryRequest) String() string { return proto.CompactTextString(m) }
func (*CommodityQueryRequest) ProtoMessage()    {}
func (*CommodityQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{0}
}
func (m *CommodityQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityQueryRequest.Unmarshal(m, b)
}
func (m *CommodityQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityQueryRequest.Marshal(b, m, deterministic)
}
func (dst *CommodityQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityQueryRequest.Merge(dst, src)
}
func (m *CommodityQueryRequest) XXX_Size() int {
	return xxx_messageInfo_CommodityQueryRequest.Size(m)
}
func (m *CommodityQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityQueryRequest proto.InternalMessageInfo

func (m *CommodityQueryRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommodityQueryRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

type CommodityQueryResponse struct {
	Goods                []*InnerComodityInfo `protobuf:"bytes,1,rep,name=goods" json:"goods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommodityQueryResponse) Reset()         { *m = CommodityQueryResponse{} }
func (m *CommodityQueryResponse) String() string { return proto.CompactTextString(m) }
func (*CommodityQueryResponse) ProtoMessage()    {}
func (*CommodityQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{1}
}
func (m *CommodityQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityQueryResponse.Unmarshal(m, b)
}
func (m *CommodityQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityQueryResponse.Marshal(b, m, deterministic)
}
func (dst *CommodityQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityQueryResponse.Merge(dst, src)
}
func (m *CommodityQueryResponse) XXX_Size() int {
	return xxx_messageInfo_CommodityQueryResponse.Size(m)
}
func (m *CommodityQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityQueryResponse proto.InternalMessageInfo

func (m *CommodityQueryResponse) GetGoods() []*InnerComodityInfo {
	if m != nil {
		return m.Goods
	}
	return nil
}

type InnerComodityInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count" json:"count,omitempty"`
	Id                   uint64   `protobuf:"varint,4,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InnerComodityInfo) Reset()         { *m = InnerComodityInfo{} }
func (m *InnerComodityInfo) String() string { return proto.CompactTextString(m) }
func (*InnerComodityInfo) ProtoMessage()    {}
func (*InnerComodityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{2}
}
func (m *InnerComodityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InnerComodityInfo.Unmarshal(m, b)
}
func (m *InnerComodityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InnerComodityInfo.Marshal(b, m, deterministic)
}
func (dst *InnerComodityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InnerComodityInfo.Merge(dst, src)
}
func (m *InnerComodityInfo) XXX_Size() int {
	return xxx_messageInfo_InnerComodityInfo.Size(m)
}
func (m *InnerComodityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InnerComodityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InnerComodityInfo proto.InternalMessageInfo

func (m *InnerComodityInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InnerComodityInfo) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *InnerComodityInfo) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *InnerComodityInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CommodityCreateRequest struct {
	Token                string             `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	MerchantID           string             `protobuf:"bytes,2,opt,name=merchantID" json:"merchantID,omitempty"`
	Good                 *InnerComodityInfo `protobuf:"bytes,3,opt,name=good" json:"good,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommodityCreateRequest) Reset()         { *m = CommodityCreateRequest{} }
func (m *CommodityCreateRequest) String() string { return proto.CompactTextString(m) }
func (*CommodityCreateRequest) ProtoMessage()    {}
func (*CommodityCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{3}
}
func (m *CommodityCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityCreateRequest.Unmarshal(m, b)
}
func (m *CommodityCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityCreateRequest.Marshal(b, m, deterministic)
}
func (dst *CommodityCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityCreateRequest.Merge(dst, src)
}
func (m *CommodityCreateRequest) XXX_Size() int {
	return xxx_messageInfo_CommodityCreateRequest.Size(m)
}
func (m *CommodityCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityCreateRequest proto.InternalMessageInfo

func (m *CommodityCreateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommodityCreateRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *CommodityCreateRequest) GetGood() *InnerComodityInfo {
	if m != nil {
		return m.Good
	}
	return nil
}

type CommodityCreateResponse struct {
	Successed            bool     `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommodityCreateResponse) Reset()         { *m = CommodityCreateResponse{} }
func (m *CommodityCreateResponse) String() string { return proto.CompactTextString(m) }
func (*CommodityCreateResponse) ProtoMessage()    {}
func (*CommodityCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{4}
}
func (m *CommodityCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityCreateResponse.Unmarshal(m, b)
}
func (m *CommodityCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityCreateResponse.Marshal(b, m, deterministic)
}
func (dst *CommodityCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityCreateResponse.Merge(dst, src)
}
func (m *CommodityCreateResponse) XXX_Size() int {
	return xxx_messageInfo_CommodityCreateResponse.Size(m)
}
func (m *CommodityCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityCreateResponse proto.InternalMessageInfo

func (m *CommodityCreateResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type CommodityDeleteRequest struct {
	Token                string             `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Merchant             string             `protobuf:"bytes,2,opt,name=merchant" json:"merchant,omitempty"`
	Good                 *InnerComodityInfo `protobuf:"bytes,3,opt,name=good" json:"good,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommodityDeleteRequest) Reset()         { *m = CommodityDeleteRequest{} }
func (m *CommodityDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*CommodityDeleteRequest) ProtoMessage()    {}
func (*CommodityDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{5}
}
func (m *CommodityDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityDeleteRequest.Unmarshal(m, b)
}
func (m *CommodityDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *CommodityDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityDeleteRequest.Merge(dst, src)
}
func (m *CommodityDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_CommodityDeleteRequest.Size(m)
}
func (m *CommodityDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityDeleteRequest proto.InternalMessageInfo

func (m *CommodityDeleteRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommodityDeleteRequest) GetMerchant() string {
	if m != nil {
		return m.Merchant
	}
	return ""
}

func (m *CommodityDeleteRequest) GetGood() *InnerComodityInfo {
	if m != nil {
		return m.Good
	}
	return nil
}

type CommodityDeleteResponse struct {
	Successed            bool     `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommodityDeleteResponse) Reset()         { *m = CommodityDeleteResponse{} }
func (m *CommodityDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*CommodityDeleteResponse) ProtoMessage()    {}
func (*CommodityDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{6}
}
func (m *CommodityDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityDeleteResponse.Unmarshal(m, b)
}
func (m *CommodityDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityDeleteResponse.Marshal(b, m, deterministic)
}
func (dst *CommodityDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityDeleteResponse.Merge(dst, src)
}
func (m *CommodityDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_CommodityDeleteResponse.Size(m)
}
func (m *CommodityDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityDeleteResponse proto.InternalMessageInfo

func (m *CommodityDeleteResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type CommodityUpdateRequest struct {
	Token                string             `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	MerchantID           string             `protobuf:"bytes,2,opt,name=merchantID" json:"merchantID,omitempty"`
	Good                 *InnerComodityInfo `protobuf:"bytes,3,opt,name=good" json:"good,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CommodityUpdateRequest) Reset()         { *m = CommodityUpdateRequest{} }
func (m *CommodityUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*CommodityUpdateRequest) ProtoMessage()    {}
func (*CommodityUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{7}
}
func (m *CommodityUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityUpdateRequest.Unmarshal(m, b)
}
func (m *CommodityUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *CommodityUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityUpdateRequest.Merge(dst, src)
}
func (m *CommodityUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_CommodityUpdateRequest.Size(m)
}
func (m *CommodityUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityUpdateRequest proto.InternalMessageInfo

func (m *CommodityUpdateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommodityUpdateRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *CommodityUpdateRequest) GetGood() *InnerComodityInfo {
	if m != nil {
		return m.Good
	}
	return nil
}

type CommodityUpdateResponse struct {
	Successed            bool     `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommodityUpdateResponse) Reset()         { *m = CommodityUpdateResponse{} }
func (m *CommodityUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*CommodityUpdateResponse) ProtoMessage()    {}
func (*CommodityUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_Commodity_0cfce62f7c0560c4, []int{8}
}
func (m *CommodityUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommodityUpdateResponse.Unmarshal(m, b)
}
func (m *CommodityUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommodityUpdateResponse.Marshal(b, m, deterministic)
}
func (dst *CommodityUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommodityUpdateResponse.Merge(dst, src)
}
func (m *CommodityUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_CommodityUpdateResponse.Size(m)
}
func (m *CommodityUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommodityUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommodityUpdateResponse proto.InternalMessageInfo

func (m *CommodityUpdateResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func init() {
	proto.RegisterType((*CommodityQueryRequest)(nil), "Aphro.Commodity.pb.CommodityQueryRequest")
	proto.RegisterType((*CommodityQueryResponse)(nil), "Aphro.Commodity.pb.CommodityQueryResponse")
	proto.RegisterType((*InnerComodityInfo)(nil), "Aphro.Commodity.pb.InnerComodityInfo")
	proto.RegisterType((*CommodityCreateRequest)(nil), "Aphro.Commodity.pb.CommodityCreateRequest")
	proto.RegisterType((*CommodityCreateResponse)(nil), "Aphro.Commodity.pb.CommodityCreateResponse")
	proto.RegisterType((*CommodityDeleteRequest)(nil), "Aphro.Commodity.pb.CommodityDeleteRequest")
	proto.RegisterType((*CommodityDeleteResponse)(nil), "Aphro.Commodity.pb.CommodityDeleteResponse")
	proto.RegisterType((*CommodityUpdateRequest)(nil), "Aphro.Commodity.pb.CommodityUpdateRequest")
	proto.RegisterType((*CommodityUpdateResponse)(nil), "Aphro.Commodity.pb.CommodityUpdateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommodityService service

type CommodityServiceClient interface {
	// 查看商品
	CommodityQuery(ctx context.Context, in *CommodityQueryRequest, opts ...grpc.CallOption) (*CommodityQueryResponse, error)
	// 增加商品
	CommodityCreate(ctx context.Context, in *CommodityCreateRequest, opts ...grpc.CallOption) (*CommodityCreateResponse, error)
	// 删除商品
	CommodityDelete(ctx context.Context, in *CommodityDeleteRequest, opts ...grpc.CallOption) (*CommodityDeleteResponse, error)
	// 更新商品
	CommodityUpdate(ctx context.Context, in *CommodityUpdateRequest, opts ...grpc.CallOption) (*CommodityUpdateResponse, error)
}

type commodityServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommodityServiceClient(cc *grpc.ClientConn) CommodityServiceClient {
	return &commodityServiceClient{cc}
}

func (c *commodityServiceClient) CommodityQuery(ctx context.Context, in *CommodityQueryRequest, opts ...grpc.CallOption) (*CommodityQueryResponse, error) {
	out := new(CommodityQueryResponse)
	err := grpc.Invoke(ctx, "/Aphro.Commodity.pb.CommodityService/CommodityQuery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commodityServiceClient) CommodityCreate(ctx context.Context, in *CommodityCreateRequest, opts ...grpc.CallOption) (*CommodityCreateResponse, error) {
	out := new(CommodityCreateResponse)
	err := grpc.Invoke(ctx, "/Aphro.Commodity.pb.CommodityService/CommodityCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commodityServiceClient) CommodityDelete(ctx context.Context, in *CommodityDeleteRequest, opts ...grpc.CallOption) (*CommodityDeleteResponse, error) {
	out := new(CommodityDeleteResponse)
	err := grpc.Invoke(ctx, "/Aphro.Commodity.pb.CommodityService/CommodityDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commodityServiceClient) CommodityUpdate(ctx context.Context, in *CommodityUpdateRequest, opts ...grpc.CallOption) (*CommodityUpdateResponse, error) {
	out := new(CommodityUpdateResponse)
	err := grpc.Invoke(ctx, "/Aphro.Commodity.pb.CommodityService/CommodityUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommodityService service

type CommodityServiceServer interface {
	// 查看商品
	CommodityQuery(context.Context, *CommodityQueryRequest) (*CommodityQueryResponse, error)
	// 增加商品
	CommodityCreate(context.Context, *CommodityCreateRequest) (*CommodityCreateResponse, error)
	// 删除商品
	CommodityDelete(context.Context, *CommodityDeleteRequest) (*CommodityDeleteResponse, error)
	// 更新商品
	CommodityUpdate(context.Context, *CommodityUpdateRequest) (*CommodityUpdateResponse, error)
}

func RegisterCommodityServiceServer(s *grpc.Server, srv CommodityServiceServer) {
	s.RegisterService(&_CommodityService_serviceDesc, srv)
}

func _CommodityService_CommodityQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommodityQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServiceServer).CommodityQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Commodity.pb.CommodityService/CommodityQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServiceServer).CommodityQuery(ctx, req.(*CommodityQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommodityService_CommodityCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommodityCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServiceServer).CommodityCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Commodity.pb.CommodityService/CommodityCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServiceServer).CommodityCreate(ctx, req.(*CommodityCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommodityService_CommodityDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommodityDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServiceServer).CommodityDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Commodity.pb.CommodityService/CommodityDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServiceServer).CommodityDelete(ctx, req.(*CommodityDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommodityService_CommodityUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommodityUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommodityServiceServer).CommodityUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Commodity.pb.CommodityService/CommodityUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommodityServiceServer).CommodityUpdate(ctx, req.(*CommodityUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommodityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aphro.Commodity.pb.CommodityService",
	HandlerType: (*CommodityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommodityQuery",
			Handler:    _CommodityService_CommodityQuery_Handler,
		},
		{
			MethodName: "CommodityCreate",
			Handler:    _CommodityService_CommodityCreate_Handler,
		},
		{
			MethodName: "CommodityDelete",
			Handler:    _CommodityService_CommodityDelete_Handler,
		},
		{
			MethodName: "CommodityUpdate",
			Handler:    _CommodityService_CommodityUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commodity/pb/Commodity.proto",
}

func init() {
	proto.RegisterFile("commodity/pb/Commodity.proto", fileDescriptor_Commodity_0cfce62f7c0560c4)
}

var fileDescriptor_Commodity_0cfce62f7c0560c4 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb5, 0xa9, 0x8b, 0xda, 0x41, 0x14, 0x58, 0xf1, 0xc7, 0x58, 0x51, 0x89, 0x56, 0x54,
	0x0a, 0x41, 0x8a, 0xa5, 0x72, 0x40, 0x94, 0x13, 0x4a, 0x2f, 0x39, 0x70, 0xc0, 0xa8, 0x0f, 0xe0,
	0xd8, 0x43, 0x6a, 0xd1, 0xec, 0xb8, 0xde, 0x35, 0x52, 0xaf, 0x48, 0x1c, 0x38, 0x70, 0x40, 0x3c,
	0x1a, 0x77, 0x4e, 0x3c, 0x08, 0xf2, 0x6e, 0x9c, 0xda, 0x71, 0x14, 0x5b, 0x82, 0x43, 0x6f, 0x99,
	0xf1, 0x37, 0xd9, 0xdf, 0xb7, 0xfb, 0xed, 0x42, 0x3f, 0xa2, 0xc5, 0x82, 0xe2, 0x44, 0x5f, 0xf9,
	0xe9, 0xcc, 0x9f, 0x94, 0xc5, 0x38, 0xcd, 0x48, 0x13, 0xe7, 0x6f, 0xd3, 0xf3, 0x8c, 0xc6, 0x95,
	0xf6, 0xcc, 0xeb, 0xcf, 0x89, 0xe6, 0x17, 0xe8, 0x87, 0x69, 0xe2, 0x87, 0x52, 0x92, 0x0e, 0x75,
	0x42, 0x52, 0xd9, 0x09, 0xf1, 0x0e, 0x1e, 0xae, 0xd4, 0xef, 0x73, 0xcc, 0xae, 0x02, 0xbc, 0xcc,
	0x51, 0x69, 0xfe, 0x00, 0x76, 0x35, 0x7d, 0x42, 0xe9, 0xb2, 0x01, 0x1b, 0xee, 0x07, 0xb6, 0xe0,
	0x87, 0x00, 0x0b, 0xcc, 0xa2, 0xf3, 0x50, 0xea, 0xe9, 0xa9, 0xdb, 0x33, 0x9f, 0x2a, 0x1d, 0x71,
	0x06, 0x8f, 0xd6, 0xff, 0x4e, 0xa5, 0x24, 0x15, 0xf2, 0x37, 0xb0, 0x3b, 0x27, 0x8a, 0x95, 0xcb,
	0x06, 0x3b, 0xc3, 0xdb, 0xc7, 0x47, 0xe3, 0x26, 0xea, 0x78, 0x2a, 0x25, 0x66, 0x13, 0xb2, 0x8d,
	0xa9, 0xfc, 0x48, 0x81, 0x9d, 0x11, 0x11, 0xdc, 0x6f, 0x7c, 0xe3, 0x1c, 0x1c, 0x19, 0x2e, 0x70,
	0x09, 0x68, 0x7e, 0x17, 0xd4, 0x69, 0x96, 0x44, 0x68, 0xd0, 0x58, 0x60, 0x8b, 0xa2, 0x1b, 0x51,
	0x2e, 0xb5, 0xbb, 0x33, 0x60, 0xc3, 0x3b, 0x81, 0x2d, 0xf8, 0x01, 0xf4, 0x92, 0xd8, 0x75, 0x06,
	0x6c, 0xe8, 0x04, 0xbd, 0x24, 0x16, 0xdf, 0x58, 0x05, 0x7e, 0x92, 0x61, 0xa8, 0xf1, 0x9f, 0x36,
	0x83, 0xbf, 0x06, 0xa7, 0xc0, 0x37, 0xab, 0x76, 0x76, 0x6c, 0x46, 0xc4, 0x2b, 0x78, 0xdc, 0x40,
	0x59, 0x6e, 0x64, 0x1f, 0xf6, 0x55, 0x1e, 0x45, 0xa8, 0x14, 0xc6, 0x86, 0x67, 0x2f, 0xb8, 0x6e,
	0x88, 0xaf, 0x55, 0x13, 0xa7, 0x78, 0x81, 0x6d, 0x26, 0x3c, 0xd8, 0x2b, 0x91, 0x97, 0x16, 0x56,
	0xf5, 0xff, 0x32, 0x50, 0x62, 0x74, 0x32, 0x50, 0x3b, 0x85, 0xb3, 0x34, 0xbe, 0x29, 0xa7, 0x50,
	0xa2, 0x74, 0x31, 0x71, 0xfc, 0xdb, 0x81, 0x7b, 0xab, 0xc9, 0x0f, 0x98, 0x7d, 0x2e, 0x52, 0xf8,
	0x9d, 0xc1, 0x41, 0xfd, 0x72, 0xf0, 0xe7, 0x9b, 0x68, 0x36, 0xde, 0x47, 0x6f, 0xd4, 0x45, 0x6a,
	0xe1, 0xc4, 0xb3, 0x2f, 0xbf, 0xfe, 0xfc, 0xec, 0x1d, 0x8a, 0x27, 0x7e, 0xb9, 0x0f, 0xfe, 0xf5,
	0xb3, 0x71, 0x59, 0x48, 0x4f, 0xd8, 0x88, 0xff, 0x60, 0x70, 0x77, 0x2d, 0x64, 0x7c, 0xfb, 0x2a,
	0xb5, 0x4b, 0xe1, 0xbd, 0xe8, 0xa4, 0x5d, 0x22, 0x1d, 0x19, 0xa4, 0xa7, 0xc2, 0xdb, 0x84, 0x14,
	0x19, 0x6d, 0x83, 0xc9, 0xe6, 0xa6, 0x85, 0xa9, 0x96, 0xf1, 0x16, 0xa6, 0x7a, 0x10, 0xb7, 0x33,
	0xc5, 0x46, 0xdb, 0x60, 0xb2, 0x31, 0x68, 0x61, 0xaa, 0xc5, 0xb6, 0x85, 0xa9, 0x9e, 0xab, 0xed,
	0x4c, 0xb9, 0xd1, 0x9e, 0xb0, 0xd1, 0xec, 0x96, 0x79, 0xbd, 0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xf8, 0x4a, 0x29, 0x9e, 0x0f, 0x06, 0x00, 0x00,
}
