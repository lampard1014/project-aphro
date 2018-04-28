// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room.proto

/*
Package Aphro_Room_pb is a generated protocol buffer package.

It is generated from these files:
	room.proto

It has these top-level messages:
	RSDeleteRequest
	RSDeleteResponse
	RSResult
	RSQueryRequest
	RSQueryResponse
	RSCreateRequest
	RSCreateResponse
	RSUpdateRequest
	RSUpdateResponse
	RSLocation
	RSTerminalBindRequest
	RSTerminalBindResponse
	RSTerminalUnbindRequest
	RSTerminalUnindResponse
*/
package Aphro_Room_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import Aphro_RoomChargeRule "."

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

type RSDeleteRequest struct {
	RoomID uint32 `protobuf:"varint,1,opt,name=roomID" json:"roomID,omitempty"`
}

func (m *RSDeleteRequest) Reset()                    { *m = RSDeleteRequest{} }
func (m *RSDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*RSDeleteRequest) ProtoMessage()               {}
func (*RSDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RSDeleteRequest) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type RSDeleteResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *RSDeleteResponse) Reset()                    { *m = RSDeleteResponse{} }
func (m *RSDeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*RSDeleteResponse) ProtoMessage()               {}
func (*RSDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RSDeleteResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type RSResult struct {
	RoomID       uint32                            `protobuf:"varint,1,opt,name=roomID" json:"roomID,omitempty"`
	MerchantID   uint32                            `protobuf:"varint,2,opt,name=merchantID" json:"merchantID,omitempty"`
	TerminalCode string                            `protobuf:"bytes,3,opt,name=terminalCode" json:"terminalCode,omitempty"`
	Location     *RSLocation                       `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Status       uint32                            `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	RoomName     string                            `protobuf:"bytes,6,opt,name=roomName" json:"roomName,omitempty"`
	ChargeRules  []*Aphro_RoomChargeRule.RCRResult `protobuf:"bytes,7,rep,name=chargeRules" json:"chargeRules,omitempty"`
}

func (m *RSResult) Reset()                    { *m = RSResult{} }
func (m *RSResult) String() string            { return proto.CompactTextString(m) }
func (*RSResult) ProtoMessage()               {}
func (*RSResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RSResult) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RSResult) GetMerchantID() uint32 {
	if m != nil {
		return m.MerchantID
	}
	return 0
}

func (m *RSResult) GetTerminalCode() string {
	if m != nil {
		return m.TerminalCode
	}
	return ""
}

func (m *RSResult) GetLocation() *RSLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RSResult) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RSResult) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *RSResult) GetChargeRules() []*Aphro_RoomChargeRule.RCRResult {
	if m != nil {
		return m.ChargeRules
	}
	return nil
}

type RSQueryRequest struct {
	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken" json:"sessionToken,omitempty"`
	RoomID       uint32 `protobuf:"varint,2,opt,name=roomID" json:"roomID,omitempty"`
	MerchantID   uint32 `protobuf:"varint,3,opt,name=merchantID" json:"merchantID,omitempty"`
}

func (m *RSQueryRequest) Reset()                    { *m = RSQueryRequest{} }
func (m *RSQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*RSQueryRequest) ProtoMessage()               {}
func (*RSQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RSQueryRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *RSQueryRequest) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RSQueryRequest) GetMerchantID() uint32 {
	if m != nil {
		return m.MerchantID
	}
	return 0
}

type RSQueryResponse struct {
	Successed bool        `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
	Results   []*RSResult `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
}

func (m *RSQueryResponse) Reset()                    { *m = RSQueryResponse{} }
func (m *RSQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*RSQueryResponse) ProtoMessage()               {}
func (*RSQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RSQueryResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func (m *RSQueryResponse) GetResults() []*RSResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type RSCreateRequest struct {
	TerminalCode string      `protobuf:"bytes,1,opt,name=terminalCode" json:"terminalCode,omitempty"`
	SessionToken string      `protobuf:"bytes,2,opt,name=sessionToken" json:"sessionToken,omitempty"`
	Location     *RSLocation `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	RoomName     string      `protobuf:"bytes,4,opt,name=roomName" json:"roomName,omitempty"`
}

func (m *RSCreateRequest) Reset()                    { *m = RSCreateRequest{} }
func (m *RSCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*RSCreateRequest) ProtoMessage()               {}
func (*RSCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RSCreateRequest) GetTerminalCode() string {
	if m != nil {
		return m.TerminalCode
	}
	return ""
}

func (m *RSCreateRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *RSCreateRequest) GetLocation() *RSLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RSCreateRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

type RSCreateResponse struct {
	RoomID    uint32 `protobuf:"varint,1,opt,name=roomID" json:"roomID,omitempty"`
	Successed bool   `protobuf:"varint,2,opt,name=successed" json:"successed,omitempty"`
}

func (m *RSCreateResponse) Reset()                    { *m = RSCreateResponse{} }
func (m *RSCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*RSCreateResponse) ProtoMessage()               {}
func (*RSCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RSCreateResponse) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RSCreateResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type RSUpdateRequest struct {
	SessionToken string                                   `protobuf:"bytes,1,opt,name=sessionToken" json:"sessionToken,omitempty"`
	RoomID       uint32                                   `protobuf:"varint,2,opt,name=roomID" json:"roomID,omitempty"`
	TerminalCode string                                   `protobuf:"bytes,3,opt,name=terminalCode" json:"terminalCode,omitempty"`
	Location     *RSLocation                              `protobuf:"bytes,4,opt,name=location" json:"location,omitempty"`
	Status       uint32                                   `protobuf:"varint,5,opt,name=status" json:"status,omitempty"`
	RoomName     string                                   `protobuf:"bytes,7,opt,name=roomName" json:"roomName,omitempty"`
	ChargeRules  []*Aphro_RoomChargeRule.RCRCreateRequest `protobuf:"bytes,6,rep,name=chargeRules" json:"chargeRules,omitempty"`
}

func (m *RSUpdateRequest) Reset()                    { *m = RSUpdateRequest{} }
func (m *RSUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*RSUpdateRequest) ProtoMessage()               {}
func (*RSUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RSUpdateRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *RSUpdateRequest) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RSUpdateRequest) GetTerminalCode() string {
	if m != nil {
		return m.TerminalCode
	}
	return ""
}

func (m *RSUpdateRequest) GetLocation() *RSLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RSUpdateRequest) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RSUpdateRequest) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *RSUpdateRequest) GetChargeRules() []*Aphro_RoomChargeRule.RCRCreateRequest {
	if m != nil {
		return m.ChargeRules
	}
	return nil
}

type RSUpdateResponse struct {
	Successed bool `protobuf:"varint,2,opt,name=successed" json:"successed,omitempty"`
}

func (m *RSUpdateResponse) Reset()                    { *m = RSUpdateResponse{} }
func (m *RSUpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*RSUpdateResponse) ProtoMessage()               {}
func (*RSUpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RSUpdateResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type RSLocation struct {
	Longitude string `protobuf:"bytes,1,opt,name=longitude" json:"longitude,omitempty"`
	Latitude  string `protobuf:"bytes,2,opt,name=latitude" json:"latitude,omitempty"`
}

func (m *RSLocation) Reset()                    { *m = RSLocation{} }
func (m *RSLocation) String() string            { return proto.CompactTextString(m) }
func (*RSLocation) ProtoMessage()               {}
func (*RSLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RSLocation) GetLongitude() string {
	if m != nil {
		return m.Longitude
	}
	return ""
}

func (m *RSLocation) GetLatitude() string {
	if m != nil {
		return m.Latitude
	}
	return ""
}

type RSTerminalBindRequest struct {
	SessionToken string      `protobuf:"bytes,1,opt,name=sessionToken" json:"sessionToken,omitempty"`
	TerminalCode string      `protobuf:"bytes,2,opt,name=terminalCode" json:"terminalCode,omitempty"`
	Location     *RSLocation `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	RoomID       uint32      `protobuf:"varint,4,opt,name=roomID" json:"roomID,omitempty"`
}

func (m *RSTerminalBindRequest) Reset()                    { *m = RSTerminalBindRequest{} }
func (m *RSTerminalBindRequest) String() string            { return proto.CompactTextString(m) }
func (*RSTerminalBindRequest) ProtoMessage()               {}
func (*RSTerminalBindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RSTerminalBindRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *RSTerminalBindRequest) GetTerminalCode() string {
	if m != nil {
		return m.TerminalCode
	}
	return ""
}

func (m *RSTerminalBindRequest) GetLocation() *RSLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RSTerminalBindRequest) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

type RSTerminalBindResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *RSTerminalBindResponse) Reset()                    { *m = RSTerminalBindResponse{} }
func (m *RSTerminalBindResponse) String() string            { return proto.CompactTextString(m) }
func (*RSTerminalBindResponse) ProtoMessage()               {}
func (*RSTerminalBindResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RSTerminalBindResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type RSTerminalUnbindRequest struct {
	RoomID       uint32 `protobuf:"varint,1,opt,name=roomID" json:"roomID,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken" json:"sessionToken,omitempty"`
}

func (m *RSTerminalUnbindRequest) Reset()                    { *m = RSTerminalUnbindRequest{} }
func (m *RSTerminalUnbindRequest) String() string            { return proto.CompactTextString(m) }
func (*RSTerminalUnbindRequest) ProtoMessage()               {}
func (*RSTerminalUnbindRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RSTerminalUnbindRequest) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *RSTerminalUnbindRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type RSTerminalUnindResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *RSTerminalUnindResponse) Reset()                    { *m = RSTerminalUnindResponse{} }
func (m *RSTerminalUnindResponse) String() string            { return proto.CompactTextString(m) }
func (*RSTerminalUnindResponse) ProtoMessage()               {}
func (*RSTerminalUnindResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RSTerminalUnindResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func init() {
	proto.RegisterType((*RSDeleteRequest)(nil), "Aphro.Room.pb.RSDeleteRequest")
	proto.RegisterType((*RSDeleteResponse)(nil), "Aphro.Room.pb.RSDeleteResponse")
	proto.RegisterType((*RSResult)(nil), "Aphro.Room.pb.RSResult")
	proto.RegisterType((*RSQueryRequest)(nil), "Aphro.Room.pb.RSQueryRequest")
	proto.RegisterType((*RSQueryResponse)(nil), "Aphro.Room.pb.RSQueryResponse")
	proto.RegisterType((*RSCreateRequest)(nil), "Aphro.Room.pb.RSCreateRequest")
	proto.RegisterType((*RSCreateResponse)(nil), "Aphro.Room.pb.RSCreateResponse")
	proto.RegisterType((*RSUpdateRequest)(nil), "Aphro.Room.pb.RSUpdateRequest")
	proto.RegisterType((*RSUpdateResponse)(nil), "Aphro.Room.pb.RSUpdateResponse")
	proto.RegisterType((*RSLocation)(nil), "Aphro.Room.pb.RSLocation")
	proto.RegisterType((*RSTerminalBindRequest)(nil), "Aphro.Room.pb.RSTerminalBindRequest")
	proto.RegisterType((*RSTerminalBindResponse)(nil), "Aphro.Room.pb.RSTerminalBindResponse")
	proto.RegisterType((*RSTerminalUnbindRequest)(nil), "Aphro.Room.pb.RSTerminalUnbindRequest")
	proto.RegisterType((*RSTerminalUnindResponse)(nil), "Aphro.Room.pb.RSTerminalUnindResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RoomService service

type RoomServiceClient interface {
	// 绑定二维码终端信息
	TerminalBind(ctx context.Context, in *RSTerminalBindRequest, opts ...grpc.CallOption) (*RSTerminalBindResponse, error)
	// 解绑二维码终端信息
	TerminalUnbind(ctx context.Context, in *RSTerminalUnbindRequest, opts ...grpc.CallOption) (*RSTerminalUnindResponse, error)
	// 直接创建房间 同时绑定二维码
	Create(ctx context.Context, in *RSCreateRequest, opts ...grpc.CallOption) (*RSCreateResponse, error)
	// 更新房间的信息 同时更新计费方式
	Update(ctx context.Context, in *RSUpdateRequest, opts ...grpc.CallOption) (*RSUpdateResponse, error)
	// 删除房间
	Delete(ctx context.Context, in *RSDeleteRequest, opts ...grpc.CallOption) (*RSDeleteResponse, error)
	// 查询房间
	Query(ctx context.Context, in *RSQueryRequest, opts ...grpc.CallOption) (*RSQueryResponse, error)
}

type roomServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoomServiceClient(cc *grpc.ClientConn) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) TerminalBind(ctx context.Context, in *RSTerminalBindRequest, opts ...grpc.CallOption) (*RSTerminalBindResponse, error) {
	out := new(RSTerminalBindResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/terminalBind", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) TerminalUnbind(ctx context.Context, in *RSTerminalUnbindRequest, opts ...grpc.CallOption) (*RSTerminalUnindResponse, error) {
	out := new(RSTerminalUnindResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/terminalUnbind", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Create(ctx context.Context, in *RSCreateRequest, opts ...grpc.CallOption) (*RSCreateResponse, error) {
	out := new(RSCreateResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Update(ctx context.Context, in *RSUpdateRequest, opts ...grpc.CallOption) (*RSUpdateResponse, error) {
	out := new(RSUpdateResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Delete(ctx context.Context, in *RSDeleteRequest, opts ...grpc.CallOption) (*RSDeleteResponse, error) {
	out := new(RSDeleteResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Query(ctx context.Context, in *RSQueryRequest, opts ...grpc.CallOption) (*RSQueryResponse, error) {
	out := new(RSQueryResponse)
	err := grpc.Invoke(ctx, "/Aphro.Room.pb.RoomService/query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomService service

type RoomServiceServer interface {
	// 绑定二维码终端信息
	TerminalBind(context.Context, *RSTerminalBindRequest) (*RSTerminalBindResponse, error)
	// 解绑二维码终端信息
	TerminalUnbind(context.Context, *RSTerminalUnbindRequest) (*RSTerminalUnindResponse, error)
	// 直接创建房间 同时绑定二维码
	Create(context.Context, *RSCreateRequest) (*RSCreateResponse, error)
	// 更新房间的信息 同时更新计费方式
	Update(context.Context, *RSUpdateRequest) (*RSUpdateResponse, error)
	// 删除房间
	Delete(context.Context, *RSDeleteRequest) (*RSDeleteResponse, error)
	// 查询房间
	Query(context.Context, *RSQueryRequest) (*RSQueryResponse, error)
}

func RegisterRoomServiceServer(s *grpc.Server, srv RoomServiceServer) {
	s.RegisterService(&_RoomService_serviceDesc, srv)
}

func _RoomService_TerminalBind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSTerminalBindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).TerminalBind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/TerminalBind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).TerminalBind(ctx, req.(*RSTerminalBindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_TerminalUnbind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSTerminalUnbindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).TerminalUnbind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/TerminalUnbind",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).TerminalUnbind(ctx, req.(*RSTerminalUnbindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Create(ctx, req.(*RSCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Update(ctx, req.(*RSUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Delete(ctx, req.(*RSDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RSQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Room.pb.RoomService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Query(ctx, req.(*RSQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aphro.Room.pb.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "terminalBind",
			Handler:    _RoomService_TerminalBind_Handler,
		},
		{
			MethodName: "terminalUnbind",
			Handler:    _RoomService_TerminalUnbind_Handler,
		},
		{
			MethodName: "create",
			Handler:    _RoomService_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _RoomService_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _RoomService_Delete_Handler,
		},
		{
			MethodName: "query",
			Handler:    _RoomService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "room.proto",
}

func init() { proto.RegisterFile("room.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xad, 0xeb, 0x5e, 0xb7, 0x01, 0x9e, 0xb6, 0x85, 0x68, 0x74, 0x95, 0x05, 0x53,
	0xe1, 0xd0, 0xc2, 0x10, 0x20, 0x71, 0x1b, 0x9d, 0xd0, 0x26, 0x21, 0x24, 0xdc, 0xed, 0xc4, 0x61,
	0x4a, 0x5b, 0xab, 0x8b, 0x96, 0xc6, 0x5d, 0x9c, 0x20, 0xa1, 0x89, 0xcb, 0xfe, 0x02, 0xfc, 0x01,
	0x2e, 0x9c, 0x10, 0x7f, 0x86, 0xbf, 0xc0, 0x0f, 0x41, 0xb6, 0x93, 0x26, 0x4e, 0x93, 0xad, 0xe3,
	0xc2, 0xd1, 0xcf, 0x2f, 0xef, 0x7b, 0xef, 0xfb, 0x3e, 0x3b, 0x06, 0x08, 0x18, 0x1b, 0xb7, 0x27,
	0x01, 0x0b, 0x19, 0x5a, 0xdd, 0x9f, 0x9c, 0x05, 0xac, 0x4d, 0x64, 0xa4, 0x6f, 0x6f, 0x8f, 0x18,
	0x1b, 0x79, 0xb4, 0xe3, 0x4c, 0xdc, 0x8e, 0xe3, 0xfb, 0x2c, 0x74, 0x42, 0x97, 0xf9, 0x5c, 0x25,
	0xdb, 0x9b, 0xe2, 0xc3, 0xd3, 0xc1, 0x99, 0x13, 0x8c, 0xe8, 0x69, 0x10, 0x79, 0x54, 0xc5, 0xf1,
	0x63, 0xb8, 0x43, 0x7a, 0x07, 0xd4, 0xa3, 0x21, 0x25, 0xf4, 0x22, 0xa2, 0x3c, 0x44, 0x9b, 0x50,
	0x15, 0xc9, 0x47, 0x07, 0x96, 0xd1, 0x34, 0x5a, 0xab, 0x24, 0x5e, 0xe1, 0xa7, 0x70, 0x37, 0x4d,
	0xe5, 0x13, 0xe6, 0x73, 0x8a, 0xb6, 0x61, 0x99, 0x47, 0x83, 0x01, 0xe5, 0x9c, 0x0e, 0x65, 0x7a,
	0x8d, 0xa4, 0x01, 0xfc, 0xcd, 0x84, 0x1a, 0xe9, 0x11, 0xca, 0x23, 0xaf, 0xb4, 0x2c, 0x6a, 0x00,
	0x8c, 0x69, 0x30, 0x38, 0x73, 0xfc, 0xf0, 0xe8, 0xc0, 0x32, 0xe5, 0x5e, 0x26, 0x82, 0x30, 0xac,
	0x84, 0x34, 0x18, 0xbb, 0xbe, 0xe3, 0x75, 0xd9, 0x90, 0x5a, 0x95, 0xa6, 0xd1, 0x5a, 0x26, 0x5a,
	0x0c, 0xbd, 0x80, 0x9a, 0xc7, 0x06, 0x72, 0x60, 0x6b, 0xa1, 0x69, 0xb4, 0xea, 0x7b, 0xf7, 0xdb,
	0x1a, 0x3b, 0x6d, 0xd2, 0x7b, 0x17, 0x27, 0x90, 0x69, 0xaa, 0x68, 0x89, 0x87, 0x4e, 0x18, 0x71,
	0x6b, 0x51, 0xb5, 0xa4, 0x56, 0xc8, 0x86, 0x9a, 0x68, 0xee, 0xbd, 0x33, 0xa6, 0x56, 0x55, 0xc2,
	0x4d, 0xd7, 0x68, 0x1f, 0xea, 0x8a, 0x45, 0x12, 0x79, 0x94, 0x5b, 0x4b, 0xcd, 0x4a, 0xab, 0xbe,
	0xb7, 0x93, 0x41, 0xeb, 0x4e, 0x77, 0xdb, 0xa4, 0x4b, 0xd4, 0xf0, 0x24, 0xfb, 0x0d, 0xf6, 0x60,
	0x8d, 0xf4, 0x3e, 0x44, 0x34, 0xf8, 0x9c, 0x50, 0x8e, 0x61, 0x85, 0x53, 0xce, 0x5d, 0xe6, 0x1f,
	0xb3, 0x73, 0xea, 0x4b, 0x86, 0x96, 0x89, 0x16, 0xcb, 0xf0, 0x67, 0x5e, 0xc3, 0x5f, 0x25, 0xcf,
	0x1f, 0xee, 0x0b, 0x85, 0x63, 0xb4, 0x79, 0x54, 0x43, 0xcf, 0x60, 0x29, 0x90, 0x5d, 0x73, 0xcb,
	0x94, 0xd3, 0x6d, 0xcd, 0x70, 0x19, 0x4f, 0x95, 0xe4, 0xe1, 0x9f, 0x86, 0x00, 0xe9, 0x06, 0xd4,
	0x49, 0x6d, 0x94, 0xd7, 0xcd, 0x28, 0xd0, 0x2d, 0x3f, 0xb7, 0x59, 0x30, 0x77, 0x56, 0xdb, 0xca,
	0xfc, 0xda, 0x66, 0x35, 0x5c, 0xd0, 0x35, 0xc4, 0x87, 0xc2, 0xc9, 0x49, 0xb7, 0x31, 0x27, 0x65,
	0xf6, 0xd4, 0xb8, 0x32, 0xf3, 0x0e, 0xff, 0x61, 0x8a, 0xc1, 0x4f, 0x26, 0x43, 0x7d, 0xf0, 0x7f,
	0x16, 0xf3, 0x3f, 0x9b, 0x7d, 0x29, 0x67, 0xf6, 0x43, 0xdd, 0xec, 0x55, 0x69, 0x87, 0xdd, 0x52,
	0xb3, 0x6b, 0x06, 0xd0, 0x3d, 0x2f, 0x2f, 0x8f, 0x84, 0xa7, 0x22, 0x1b, 0xce, 0x50, 0xfb, 0x16,
	0x20, 0x9d, 0x43, 0xe4, 0x7a, 0xcc, 0x1f, 0xb9, 0x61, 0x34, 0xb5, 0x52, 0x1a, 0x10, 0x33, 0x78,
	0x4e, 0xa8, 0x36, 0x95, 0x87, 0xa6, 0x6b, 0xfc, 0xcb, 0x80, 0x0d, 0xd2, 0x3b, 0x8e, 0x19, 0x7c,
	0xe3, 0xfa, 0xc3, 0xdb, 0x08, 0x95, 0x17, 0xc4, 0xbc, 0x41, 0x90, 0xca, 0xad, 0x04, 0x89, 0x3d,
	0xb0, 0xa0, 0xdd, 0xb3, 0x2f, 0x61, 0x33, 0xdf, 0xef, 0x5c, 0xb7, 0xed, 0x09, 0x6c, 0xa5, 0xdf,
	0x9d, 0xf8, 0xfd, 0xcc, 0xa4, 0x65, 0xe6, 0x9e, 0xe3, 0xfc, 0xe1, 0x57, 0x7a, 0xd9, 0xb9, 0xfb,
	0xd9, 0xfb, 0xbe, 0x08, 0x75, 0x41, 0x40, 0x8f, 0x06, 0x9f, 0xdc, 0x01, 0x45, 0x97, 0x29, 0x95,
	0x62, 0x2a, 0xf4, 0x70, 0x86, 0xa4, 0x02, 0x91, 0xec, 0x47, 0x37, 0x64, 0xa9, 0x56, 0x70, 0xe3,
	0xea, 0xf7, 0x9f, 0xaf, 0xa6, 0x85, 0xd7, 0x3b, 0x62, 0xb4, 0x4e, 0x02, 0xd4, 0x11, 0x2c, 0xbc,
	0x36, 0x9e, 0xa0, 0x2b, 0x03, 0xd6, 0x42, 0x8d, 0x1b, 0xb4, 0x5b, 0x5a, 0x59, 0x23, 0xcf, 0xbe,
	0x2e, 0x2f, 0xdb, 0x42, 0x53, 0xb6, 0x60, 0xe3, 0x8d, 0x5c, 0x0b, 0x91, 0x9f, 0x34, 0xd1, 0x87,
	0xea, 0x40, 0x1e, 0x11, 0xd4, 0x98, 0xa9, 0xa9, 0x9d, 0x1d, 0x7b, 0xa7, 0x74, 0x3f, 0x06, 0xdb,
	0x92, 0x60, 0xf7, 0xf0, 0x8a, 0x02, 0x53, 0x65, 0x63, 0x8c, 0x48, 0x1e, 0xb3, 0x02, 0x0c, 0xed,
	0x9e, 0x2a, 0xc0, 0xd0, 0xcf, 0x67, 0x1e, 0x43, 0x95, 0x15, 0x18, 0xe7, 0x50, 0x1d, 0xca, 0x77,
	0x40, 0x01, 0x86, 0xf6, 0x96, 0x28, 0xc0, 0xd0, 0x1f, 0x10, 0x09, 0x69, 0x76, 0x4c, 0x9a, 0x2a,
	0xdb, 0xb9, 0x54, 0xfe, 0xfc, 0x22, 0xc0, 0x3e, 0xc2, 0xe2, 0x85, 0xf8, 0x7b, 0xa1, 0x07, 0x33,
	0xb5, 0xb2, 0xff, 0x50, 0xbb, 0x51, 0xb6, 0x1d, 0x23, 0xad, 0x4b, 0xa4, 0x55, 0x54, 0x57, 0x48,
	0xb2, 0x66, 0xbf, 0x2a, 0x5f, 0x41, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x73, 0xb1, 0x8a,
	0x89, 0x58, 0x09, 0x00, 0x00,
}
