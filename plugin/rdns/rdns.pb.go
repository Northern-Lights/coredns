// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdns.proto

package rdns

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RdnsRequest struct {
	IpAddr               string   `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RdnsRequest) Reset()         { *m = RdnsRequest{} }
func (m *RdnsRequest) String() string { return proto.CompactTextString(m) }
func (*RdnsRequest) ProtoMessage()    {}
func (*RdnsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_810f08610c23646e, []int{0}
}

func (m *RdnsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RdnsRequest.Unmarshal(m, b)
}
func (m *RdnsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RdnsRequest.Marshal(b, m, deterministic)
}
func (m *RdnsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RdnsRequest.Merge(m, src)
}
func (m *RdnsRequest) XXX_Size() int {
	return xxx_messageInfo_RdnsRequest.Size(m)
}
func (m *RdnsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RdnsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RdnsRequest proto.InternalMessageInfo

func (m *RdnsRequest) GetIpAddr() string {
	if m != nil {
		return m.IpAddr
	}
	return ""
}

type RdnsResponse struct {
	Domains              []string `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RdnsResponse) Reset()         { *m = RdnsResponse{} }
func (m *RdnsResponse) String() string { return proto.CompactTextString(m) }
func (*RdnsResponse) ProtoMessage()    {}
func (*RdnsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_810f08610c23646e, []int{1}
}

func (m *RdnsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RdnsResponse.Unmarshal(m, b)
}
func (m *RdnsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RdnsResponse.Marshal(b, m, deterministic)
}
func (m *RdnsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RdnsResponse.Merge(m, src)
}
func (m *RdnsResponse) XXX_Size() int {
	return xxx_messageInfo_RdnsResponse.Size(m)
}
func (m *RdnsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RdnsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RdnsResponse proto.InternalMessageInfo

func (m *RdnsResponse) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func init() {
	proto.RegisterType((*RdnsRequest)(nil), "RdnsRequest")
	proto.RegisterType((*RdnsResponse)(nil), "RdnsResponse")
}

func init() { proto.RegisterFile("rdns.proto", fileDescriptor_810f08610c23646e) }

var fileDescriptor_810f08610c23646e = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0xc9, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe3, 0xe2, 0x0e, 0x4a, 0xc9, 0x2b, 0x0e, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7, 0x62, 0xcf, 0x2c, 0x88, 0x4f, 0x4c, 0x49, 0x29,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xcb, 0x2c, 0x70, 0x4c, 0x49, 0x29, 0x52, 0xd2,
	0xe0, 0xe2, 0x81, 0xa8, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe0, 0x62, 0x4f, 0xc9,
	0xcf, 0x4d, 0xcc, 0xcc, 0x2b, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x71, 0x8d, 0x0c,
	0xb8, 0x38, 0x41, 0x2a, 0x9d, 0x13, 0x93, 0x33, 0x52, 0x85, 0x94, 0xb9, 0x58, 0x40, 0x1c, 0x21,
	0x1e, 0x3d, 0x24, 0x5b, 0xa4, 0x78, 0xf5, 0x90, 0xcd, 0x4a, 0x62, 0x03, 0x3b, 0xc5, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xa8, 0x5f, 0xe0, 0x4b, 0x98, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RdnsCacheClient is the client API for RdnsCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RdnsCacheClient interface {
	Rdns(ctx context.Context, in *RdnsRequest, opts ...grpc.CallOption) (*RdnsResponse, error)
}

type rdnsCacheClient struct {
	cc *grpc.ClientConn
}

func NewRdnsCacheClient(cc *grpc.ClientConn) RdnsCacheClient {
	return &rdnsCacheClient{cc}
}

func (c *rdnsCacheClient) Rdns(ctx context.Context, in *RdnsRequest, opts ...grpc.CallOption) (*RdnsResponse, error) {
	out := new(RdnsResponse)
	err := c.cc.Invoke(ctx, "/RdnsCache/Rdns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RdnsCacheServer is the server API for RdnsCache service.
type RdnsCacheServer interface {
	Rdns(context.Context, *RdnsRequest) (*RdnsResponse, error)
}

// UnimplementedRdnsCacheServer can be embedded to have forward compatible implementations.
type UnimplementedRdnsCacheServer struct {
}

func (*UnimplementedRdnsCacheServer) Rdns(ctx context.Context, req *RdnsRequest) (*RdnsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rdns not implemented")
}

func RegisterRdnsCacheServer(s *grpc.Server, srv RdnsCacheServer) {
	s.RegisterService(&_RdnsCache_serviceDesc, srv)
}

func _RdnsCache_Rdns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdnsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdnsCacheServer).Rdns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RdnsCache/Rdns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdnsCacheServer).Rdns(ctx, req.(*RdnsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RdnsCache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RdnsCache",
	HandlerType: (*RdnsCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rdns",
			Handler:    _RdnsCache_Rdns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rdns.proto",
}