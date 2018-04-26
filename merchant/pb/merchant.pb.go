// Code generated by protoc-gen-go. DO NOT EDIT.
// source: merchant.proto

/*
Package Aphro_Merchant_pb is a generated protocol buffer package.

It is generated from these files:
	merchant.proto

It has these top-level messages:
	MerchantOpenRequest
	MerchantOpenResponse
	MerchantRegisterRequest
	MerchantRegisterResponse
	MerchantChangePswRequest
	MerchantChangePswResponse
	MerchantLoginRequest
	MerchantLoginResponse
	MerchantInfoRequest
	MerchantInfoResponse
*/
package Aphro_Merchant_pb

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

type MerchantOpenRequest struct {
	Name       string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cellphone  string `protobuf:"bytes,2,opt,name=cellphone" json:"cellphone,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	PaymentBit uint32 `protobuf:"varint,4,opt,name=paymentBit" json:"paymentBit,omitempty"`
}

func (m *MerchantOpenRequest) Reset()                    { *m = MerchantOpenRequest{} }
func (m *MerchantOpenRequest) String() string            { return proto.CompactTextString(m) }
func (*MerchantOpenRequest) ProtoMessage()               {}
func (*MerchantOpenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MerchantOpenRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MerchantOpenRequest) GetCellphone() string {
	if m != nil {
		return m.Cellphone
	}
	return ""
}

func (m *MerchantOpenRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MerchantOpenRequest) GetPaymentBit() uint32 {
	if m != nil {
		return m.PaymentBit
	}
	return 0
}

type MerchantOpenResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *MerchantOpenResponse) Reset()                    { *m = MerchantOpenResponse{} }
func (m *MerchantOpenResponse) String() string            { return proto.CompactTextString(m) }
func (*MerchantOpenResponse) ProtoMessage()               {}
func (*MerchantOpenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MerchantOpenResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type MerchantRegisterRequest struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	VerifyCode string `protobuf:"bytes,3,opt,name=verifyCode" json:"verifyCode,omitempty"`
	Role       uint32 `protobuf:"varint,5,opt,name=role" json:"role,omitempty"`
	MerchantID uint32 `protobuf:"varint,6,opt,name=merchantID" json:"merchantID,omitempty"`
}

func (m *MerchantRegisterRequest) Reset()                    { *m = MerchantRegisterRequest{} }
func (m *MerchantRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*MerchantRegisterRequest) ProtoMessage()               {}
func (*MerchantRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MerchantRegisterRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MerchantRegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MerchantRegisterRequest) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

func (m *MerchantRegisterRequest) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *MerchantRegisterRequest) GetMerchantID() uint32 {
	if m != nil {
		return m.MerchantID
	}
	return 0
}

type MerchantRegisterResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *MerchantRegisterResponse) Reset()                    { *m = MerchantRegisterResponse{} }
func (m *MerchantRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*MerchantRegisterResponse) ProtoMessage()               {}
func (*MerchantRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MerchantRegisterResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type MerchantChangePswRequest struct {
	NewPsw       string `protobuf:"bytes,1,opt,name=newPsw" json:"newPsw,omitempty"`
	SessionToken string `protobuf:"bytes,2,opt,name=sessionToken" json:"sessionToken,omitempty"`
	VerifyCode   string `protobuf:"bytes,3,opt,name=verifyCode" json:"verifyCode,omitempty"`
	Cellphone    string `protobuf:"bytes,4,opt,name=cellphone" json:"cellphone,omitempty"`
}

func (m *MerchantChangePswRequest) Reset()                    { *m = MerchantChangePswRequest{} }
func (m *MerchantChangePswRequest) String() string            { return proto.CompactTextString(m) }
func (*MerchantChangePswRequest) ProtoMessage()               {}
func (*MerchantChangePswRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MerchantChangePswRequest) GetNewPsw() string {
	if m != nil {
		return m.NewPsw
	}
	return ""
}

func (m *MerchantChangePswRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *MerchantChangePswRequest) GetVerifyCode() string {
	if m != nil {
		return m.VerifyCode
	}
	return ""
}

func (m *MerchantChangePswRequest) GetCellphone() string {
	if m != nil {
		return m.Cellphone
	}
	return ""
}

type MerchantChangePswResponse struct {
	Successed bool `protobuf:"varint,1,opt,name=successed" json:"successed,omitempty"`
}

func (m *MerchantChangePswResponse) Reset()                    { *m = MerchantChangePswResponse{} }
func (m *MerchantChangePswResponse) String() string            { return proto.CompactTextString(m) }
func (*MerchantChangePswResponse) ProtoMessage()               {}
func (*MerchantChangePswResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MerchantChangePswResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

type MerchantLoginRequest struct {
	Key          string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	TokenRequest string `protobuf:"bytes,2,opt,name=tokenRequest" json:"tokenRequest,omitempty"`
	MerchantID   uint32 `protobuf:"varint,3,opt,name=merchantID" json:"merchantID,omitempty"`
}

func (m *MerchantLoginRequest) Reset()                    { *m = MerchantLoginRequest{} }
func (m *MerchantLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*MerchantLoginRequest) ProtoMessage()               {}
func (*MerchantLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MerchantLoginRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MerchantLoginRequest) GetTokenRequest() string {
	if m != nil {
		return m.TokenRequest
	}
	return ""
}

func (m *MerchantLoginRequest) GetMerchantID() uint32 {
	if m != nil {
		return m.MerchantID
	}
	return 0
}

type MerchantLoginResponse struct {
	SessionToken string `protobuf:"bytes,1,opt,name=sessionToken" json:"sessionToken,omitempty"`
}

func (m *MerchantLoginResponse) Reset()                    { *m = MerchantLoginResponse{} }
func (m *MerchantLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*MerchantLoginResponse) ProtoMessage()               {}
func (*MerchantLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MerchantLoginResponse) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type MerchantInfoRequest struct {
	Token      string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	MerchantID string `protobuf:"bytes,2,opt,name=merchantID" json:"merchantID,omitempty"`
}

func (m *MerchantInfoRequest) Reset()                    { *m = MerchantInfoRequest{} }
func (m *MerchantInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MerchantInfoRequest) ProtoMessage()               {}
func (*MerchantInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MerchantInfoRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MerchantInfoRequest) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

type MerchantInfoResponse struct {
	MerchantName    string                                    `protobuf:"bytes,1,opt,name=merchantName" json:"merchantName,omitempty"`
	MerchantAccount *MerchantInfoResponseInnerMerchantAccount `protobuf:"bytes,2,opt,name=merchantAccount" json:"merchantAccount,omitempty"`
}

func (m *MerchantInfoResponse) Reset()                    { *m = MerchantInfoResponse{} }
func (m *MerchantInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*MerchantInfoResponse) ProtoMessage()               {}
func (*MerchantInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MerchantInfoResponse) GetMerchantName() string {
	if m != nil {
		return m.MerchantName
	}
	return ""
}

func (m *MerchantInfoResponse) GetMerchantAccount() *MerchantInfoResponseInnerMerchantAccount {
	if m != nil {
		return m.MerchantAccount
	}
	return nil
}

type MerchantInfoResponseInnerMerchantAccount struct {
	Role uint32 `protobuf:"varint,1,opt,name=role" json:"role,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *MerchantInfoResponseInnerMerchantAccount) Reset() {
	*m = MerchantInfoResponseInnerMerchantAccount{}
}
func (m *MerchantInfoResponseInnerMerchantAccount) String() string { return proto.CompactTextString(m) }
func (*MerchantInfoResponseInnerMerchantAccount) ProtoMessage()    {}
func (*MerchantInfoResponseInnerMerchantAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{9, 0}
}

func (m *MerchantInfoResponseInnerMerchantAccount) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *MerchantInfoResponseInnerMerchantAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*MerchantOpenRequest)(nil), "Aphro.Merchant.pb.merchantOpenRequest")
	proto.RegisterType((*MerchantOpenResponse)(nil), "Aphro.Merchant.pb.merchantOpenResponse")
	proto.RegisterType((*MerchantRegisterRequest)(nil), "Aphro.Merchant.pb.merchantRegisterRequest")
	proto.RegisterType((*MerchantRegisterResponse)(nil), "Aphro.Merchant.pb.merchantRegisterResponse")
	proto.RegisterType((*MerchantChangePswRequest)(nil), "Aphro.Merchant.pb.merchantChangePswRequest")
	proto.RegisterType((*MerchantChangePswResponse)(nil), "Aphro.Merchant.pb.merchantChangePswResponse")
	proto.RegisterType((*MerchantLoginRequest)(nil), "Aphro.Merchant.pb.merchantLoginRequest")
	proto.RegisterType((*MerchantLoginResponse)(nil), "Aphro.Merchant.pb.merchantLoginResponse")
	proto.RegisterType((*MerchantInfoRequest)(nil), "Aphro.Merchant.pb.merchantInfoRequest")
	proto.RegisterType((*MerchantInfoResponse)(nil), "Aphro.Merchant.pb.merchantInfoResponse")
	proto.RegisterType((*MerchantInfoResponseInnerMerchantAccount)(nil), "Aphro.Merchant.pb.merchantInfoResponse.innerMerchantAccount")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MerchantService service

type MerchantServiceClient interface {
	// 商户开户
	MerchantOpen(ctx context.Context, in *MerchantOpenRequest, opts ...grpc.CallOption) (*MerchantOpenResponse, error)
	// 商户用户注册
	MerchantRegister(ctx context.Context, in *MerchantRegisterRequest, opts ...grpc.CallOption) (*MerchantRegisterResponse, error)
	// 商户用户修改密码
	MerchantChangePsw(ctx context.Context, in *MerchantChangePswRequest, opts ...grpc.CallOption) (*MerchantChangePswResponse, error)
	// 商户app登陆
	MerchantLogin(ctx context.Context, in *MerchantLoginRequest, opts ...grpc.CallOption) (*MerchantLoginResponse, error)
	// 商户app 信息
	MerchantInfo(ctx context.Context, in *MerchantInfoRequest, opts ...grpc.CallOption) (*MerchantInfoResponse, error)
}

type merchantServiceClient struct {
	cc *grpc.ClientConn
}

func NewMerchantServiceClient(cc *grpc.ClientConn) MerchantServiceClient {
	return &merchantServiceClient{cc}
}

func (c *merchantServiceClient) MerchantOpen(ctx context.Context, in *MerchantOpenRequest, opts ...grpc.CallOption) (*MerchantOpenResponse, error) {
	out := new(MerchantOpenResponse)
	err := grpc.Invoke(ctx, "/Aphro.Merchant.pb.MerchantService/merchantOpen", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) MerchantRegister(ctx context.Context, in *MerchantRegisterRequest, opts ...grpc.CallOption) (*MerchantRegisterResponse, error) {
	out := new(MerchantRegisterResponse)
	err := grpc.Invoke(ctx, "/Aphro.Merchant.pb.MerchantService/merchantRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) MerchantChangePsw(ctx context.Context, in *MerchantChangePswRequest, opts ...grpc.CallOption) (*MerchantChangePswResponse, error) {
	out := new(MerchantChangePswResponse)
	err := grpc.Invoke(ctx, "/Aphro.Merchant.pb.MerchantService/merchantChangePsw", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) MerchantLogin(ctx context.Context, in *MerchantLoginRequest, opts ...grpc.CallOption) (*MerchantLoginResponse, error) {
	out := new(MerchantLoginResponse)
	err := grpc.Invoke(ctx, "/Aphro.Merchant.pb.MerchantService/merchantLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServiceClient) MerchantInfo(ctx context.Context, in *MerchantInfoRequest, opts ...grpc.CallOption) (*MerchantInfoResponse, error) {
	out := new(MerchantInfoResponse)
	err := grpc.Invoke(ctx, "/Aphro.Merchant.pb.MerchantService/merchantInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MerchantService service

type MerchantServiceServer interface {
	// 商户开户
	MerchantOpen(context.Context, *MerchantOpenRequest) (*MerchantOpenResponse, error)
	// 商户用户注册
	MerchantRegister(context.Context, *MerchantRegisterRequest) (*MerchantRegisterResponse, error)
	// 商户用户修改密码
	MerchantChangePsw(context.Context, *MerchantChangePswRequest) (*MerchantChangePswResponse, error)
	// 商户app登陆
	MerchantLogin(context.Context, *MerchantLoginRequest) (*MerchantLoginResponse, error)
	// 商户app 信息
	MerchantInfo(context.Context, *MerchantInfoRequest) (*MerchantInfoResponse, error)
}

func RegisterMerchantServiceServer(s *grpc.Server, srv MerchantServiceServer) {
	s.RegisterService(&_MerchantService_serviceDesc, srv)
}

func _MerchantService_MerchantOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantOpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Merchant.pb.MerchantService/MerchantOpen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantOpen(ctx, req.(*MerchantOpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_MerchantRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Merchant.pb.MerchantService/MerchantRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantRegister(ctx, req.(*MerchantRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_MerchantChangePsw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantChangePswRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantChangePsw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Merchant.pb.MerchantService/MerchantChangePsw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantChangePsw(ctx, req.(*MerchantChangePswRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_MerchantLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Merchant.pb.MerchantService/MerchantLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantLogin(ctx, req.(*MerchantLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantService_MerchantInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServiceServer).MerchantInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aphro.Merchant.pb.MerchantService/MerchantInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServiceServer).MerchantInfo(ctx, req.(*MerchantInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MerchantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aphro.Merchant.pb.MerchantService",
	HandlerType: (*MerchantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "merchantOpen",
			Handler:    _MerchantService_MerchantOpen_Handler,
		},
		{
			MethodName: "merchantRegister",
			Handler:    _MerchantService_MerchantRegister_Handler,
		},
		{
			MethodName: "merchantChangePsw",
			Handler:    _MerchantService_MerchantChangePsw_Handler,
		},
		{
			MethodName: "merchantLogin",
			Handler:    _MerchantService_MerchantLogin_Handler,
		},
		{
			MethodName: "merchantInfo",
			Handler:    _MerchantService_MerchantInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merchant.proto",
}

func init() { proto.RegisterFile("merchant.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x7c, 0x29, 0x47, 0x10, 0x18, 0x11, 0x6a, 0x21, 0x88, 0x93, 0x88, 0x04, 0xcd,
	0x6e, 0x82, 0x5e, 0x28, 0x1a, 0x13, 0xc4, 0x1b, 0xa2, 0x28, 0xa9, 0xbe, 0x40, 0xe9, 0x9e, 0xdd,
	0x6d, 0x28, 0x33, 0xb5, 0xd3, 0x65, 0xb3, 0x97, 0x72, 0xe1, 0x03, 0x88, 0xf1, 0xc1, 0xf4, 0x0d,
	0x8c, 0x0f, 0x62, 0xda, 0xce, 0x6c, 0xa7, 0xed, 0xda, 0xdd, 0xbb, 0x99, 0xd3, 0x73, 0xe6, 0xfc,
	0xfe, 0x73, 0xce, 0x99, 0xc2, 0xed, 0x0b, 0x8c, 0xbc, 0xae, 0xcb, 0xe3, 0x46, 0x18, 0x89, 0x58,
	0xd0, 0x95, 0xc3, 0xb0, 0x1b, 0x89, 0xc6, 0xc9, 0xd0, 0x7a, 0x66, 0x6f, 0x76, 0x84, 0xe8, 0x04,
	0xd8, 0x74, 0x43, 0xbf, 0xe9, 0x72, 0x2e, 0x62, 0x37, 0xf6, 0x05, 0x97, 0x59, 0x00, 0xfb, 0x4a,
	0xe0, 0x8e, 0x3e, 0xe3, 0x63, 0x88, 0xdc, 0xc1, 0x2f, 0x3d, 0x94, 0x31, 0xa5, 0x30, 0xc3, 0xdd,
	0x0b, 0xb4, 0xc8, 0x36, 0xd9, 0x9d, 0x77, 0xd2, 0x35, 0xdd, 0x84, 0x79, 0x0f, 0x83, 0x20, 0xec,
	0x0a, 0x8e, 0xd6, 0x54, 0xfa, 0x21, 0x37, 0x50, 0x0b, 0x6e, 0xb8, 0xad, 0x56, 0x84, 0x52, 0x5a,
	0xd3, 0xe9, 0x37, 0xbd, 0xa5, 0x5b, 0x00, 0xa1, 0x3b, 0xb8, 0x40, 0x1e, 0xbf, 0xf1, 0x63, 0x6b,
	0x66, 0x9b, 0xec, 0x2e, 0x3a, 0x86, 0x85, 0x3d, 0x83, 0xd5, 0x22, 0x82, 0x0c, 0x05, 0x97, 0x69,
	0x3e, 0xd9, 0xf3, 0x3c, 0x94, 0x12, 0x5b, 0x29, 0xc8, 0x4d, 0x27, 0x37, 0xb0, 0x6b, 0x02, 0xeb,
	0x3a, 0xcc, 0xc1, 0x8e, 0x2f, 0x63, 0x8c, 0x34, 0xfd, 0x32, 0x4c, 0x9f, 0xe3, 0x40, 0xc1, 0x27,
	0xcb, 0xa1, 0x9e, 0x29, 0x43, 0xcf, 0x16, 0xc0, 0x25, 0x46, 0x7e, 0x7b, 0x70, 0x24, 0x5a, 0xa8,
	0xa0, 0x0d, 0x4b, 0x12, 0x13, 0x89, 0x00, 0xad, 0xd9, 0x94, 0x38, 0x5d, 0x27, 0x31, 0x3a, 0xe9,
	0xf1, 0x5b, 0x6b, 0x2e, 0xd3, 0x92, 0x5b, 0xd8, 0x73, 0xb0, 0xaa, 0x50, 0x13, 0xe9, 0xf9, 0x41,
	0xf2, 0xd0, 0xa3, 0xae, 0xcb, 0x3b, 0x78, 0x2a, 0xfb, 0x5a, 0xd0, 0x1a, 0xcc, 0x71, 0xec, 0x9f,
	0xca, 0xbe, 0xd2, 0xa4, 0x76, 0x94, 0xc1, 0x82, 0x44, 0x29, 0x7d, 0xc1, 0x3f, 0x8b, 0x73, 0xe4,
	0x4a, 0x5e, 0xc1, 0x36, 0x56, 0x66, 0xa1, 0xac, 0x33, 0xa5, 0xb2, 0xb2, 0x17, 0x70, 0x6f, 0x04,
	0xd5, 0x44, 0x8a, 0x82, 0xbc, 0xae, 0xef, 0x45, 0xc7, 0xe7, 0xff, 0xaf, 0x0e, 0x83, 0x85, 0x38,
	0x61, 0x55, 0x1e, 0x5a, 0x86, 0x69, 0x2b, 0xdd, 0xfc, 0x74, 0xe5, 0xe6, 0x5f, 0xc2, 0xdd, 0x52,
	0x36, 0x05, 0x59, 0xbe, 0x23, 0x52, 0xbd, 0x23, 0xf6, 0x2e, 0x9f, 0x82, 0x63, 0xde, 0x16, 0x3a,
	0xe7, 0x2a, 0xcc, 0xc6, 0x46, 0x4c, 0xb6, 0x29, 0x91, 0x64, 0xac, 0x26, 0xc9, 0x1f, 0x92, 0x0b,
	0xcf, 0x4e, 0xcb, 0x49, 0xb4, 0xfd, 0x43, 0x3e, 0x5c, 0x05, 0x1b, 0x6d, 0xc3, 0x92, 0xde, 0x1f,
	0x7a, 0x9e, 0xe8, 0xf1, 0xec, 0x36, 0x6e, 0xed, 0xbf, 0x6a, 0x54, 0x66, 0xbb, 0x31, 0x2a, 0x4b,
	0xc3, 0xe7, 0x1c, 0xa3, 0x93, 0xe2, 0x19, 0x4e, 0xf9, 0x50, 0xfb, 0x35, 0xac, 0x8e, 0x72, 0x1c,
	0x36, 0x3d, 0x31, 0x9a, 0x7e, 0xc4, 0xf0, 0xec, 0xff, 0x9a, 0x85, 0x25, 0x1d, 0xfb, 0x09, 0xa3,
	0x4b, 0xdf, 0x43, 0x7a, 0x45, 0x72, 0x81, 0xc9, 0x24, 0xd3, 0x9d, 0x1a, 0x66, 0xe3, 0xb5, 0xb1,
	0x1f, 0x8d, 0xf5, 0xcb, 0xb4, 0xb1, 0x07, 0x57, 0xbf, 0xff, 0x5e, 0x4f, 0x6d, 0xb0, 0xb5, 0xa6,
	0x2a, 0x5f, 0x53, 0xbb, 0x35, 0x45, 0x88, 0xfc, 0x80, 0xec, 0xd1, 0xef, 0x04, 0x96, 0xcb, 0x23,
	0x48, 0xf7, 0x6a, 0x12, 0x94, 0x1e, 0x0f, 0xfb, 0xf1, 0x44, 0xbe, 0x0a, 0xe8, 0x61, 0x0a, 0x74,
	0x9f, 0xd9, 0x55, 0xa0, 0x48, 0xf9, 0x26, 0x50, 0x3f, 0x09, 0xac, 0x54, 0xc6, 0x88, 0xd6, 0x65,
	0x2a, 0x3f, 0x01, 0xf6, 0x93, 0xc9, 0x9c, 0x15, 0xd7, 0x4e, 0xca, 0xb5, 0xcd, 0x36, 0xaa, 0x5c,
	0x9e, 0x76, 0x4e, 0xc0, 0xbe, 0x11, 0x58, 0x2c, 0x8c, 0x0d, 0xad, 0xab, 0x85, 0x39, 0xc6, 0xf6,
	0xee, 0x78, 0x47, 0x05, 0xc3, 0x52, 0x98, 0x4d, 0xb6, 0x5e, 0x85, 0x09, 0x12, 0xc7, 0x04, 0xc4,
	0xec, 0x9d, 0xa4, 0x9d, 0x6b, 0x7b, 0xc7, 0x98, 0xd1, 0xda, 0xde, 0x31, 0xe7, 0xa2, 0xae, 0x77,
	0x7c, 0xde, 0x16, 0x07, 0x64, 0xef, 0x6c, 0x2e, 0xfd, 0x29, 0x3e, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0xe0, 0x94, 0x38, 0x57, 0x07, 0x00, 0x00,
}
