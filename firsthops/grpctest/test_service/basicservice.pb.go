// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test_service/basicservice.proto

/*
Package test_service is a generated protocol buffer package.

It is generated from these files:
	test_service/basicservice.proto

It has these top-level messages:
	InfoRequest
	InfoList
*/
package test_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type InfoRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type InfoList struct {
	OsName    string `protobuf:"bytes,1,opt,name=osName" json:"osName,omitempty"`
	OsLocale  string `protobuf:"bytes,2,opt,name=osLocale" json:"osLocale,omitempty"`
	LocalTime string `protobuf:"bytes,3,opt,name=localTime" json:"localTime,omitempty"`
}

func (m *InfoList) Reset()                    { *m = InfoList{} }
func (m *InfoList) String() string            { return proto.CompactTextString(m) }
func (*InfoList) ProtoMessage()               {}
func (*InfoList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InfoList) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *InfoList) GetOsLocale() string {
	if m != nil {
		return m.OsLocale
	}
	return ""
}

func (m *InfoList) GetLocalTime() string {
	if m != nil {
		return m.LocalTime
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "test_service.InfoRequest")
	proto.RegisterType((*InfoList)(nil), "test_service.InfoList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ServerChecker service

type ServerCheckerClient interface {
	GetServerInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoList, error)
}

type serverCheckerClient struct {
	cc *grpc.ClientConn
}

func NewServerCheckerClient(cc *grpc.ClientConn) ServerCheckerClient {
	return &serverCheckerClient{cc}
}

func (c *serverCheckerClient) GetServerInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoList, error) {
	out := new(InfoList)
	err := grpc.Invoke(ctx, "/test_service.ServerChecker/GetServerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServerChecker service

type ServerCheckerServer interface {
	GetServerInfo(context.Context, *InfoRequest) (*InfoList, error)
}

func RegisterServerCheckerServer(s *grpc.Server, srv ServerCheckerServer) {
	s.RegisterService(&_ServerChecker_serviceDesc, srv)
}

func _ServerChecker_GetServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerCheckerServer).GetServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test_service.ServerChecker/GetServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerCheckerServer).GetServerInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test_service.ServerChecker",
	HandlerType: (*ServerCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServerInfo",
			Handler:    _ServerChecker_GetServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test_service/basicservice.proto",
}

func init() { proto.RegisterFile("test_service/basicservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x49, 0x2d, 0x2e,
	0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x4a, 0x2c, 0xce, 0x4c, 0x86, 0x72,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x90, 0x15, 0x28, 0x69, 0x72, 0x71, 0x7b, 0xe6,
	0xa5, 0xe5, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7,
	0x16, 0xf9, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4a, 0x31,
	0x5c, 0x1c, 0x20, 0xa5, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0xf9, 0xc5, 0x48, 0xaa,
	0xa0, 0x3c, 0x90, 0xfe, 0xfc, 0x62, 0x9f, 0xfc, 0xe4, 0xc4, 0x9c, 0x54, 0x09, 0x26, 0x88, 0x7e,
	0x18, 0x5f, 0x48, 0x86, 0x8b, 0x33, 0x07, 0xc4, 0x0a, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x06, 0x4b,
	0x22, 0x04, 0x8c, 0x42, 0xb9, 0x78, 0x83, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x9c, 0x33, 0x52, 0x93,
	0xb3, 0x53, 0x8b, 0x84, 0x5c, 0xb8, 0x78, 0xdd, 0x53, 0x4b, 0x20, 0x62, 0x20, 0x7b, 0x85, 0x24,
	0xf5, 0x90, 0x5d, 0xae, 0x87, 0xe4, 0x6c, 0x29, 0x31, 0x4c, 0x29, 0x90, 0x33, 0x95, 0x18, 0x92,
	0xd8, 0xc0, 0x9e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0xf4, 0xb3, 0x4b, 0x17, 0x01,
	0x00, 0x00,
}
