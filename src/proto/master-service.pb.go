// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/master-service.proto

package hotel_pb

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

type RegisterSpawnerRequest struct {
	// The hostname that the spawner should be accessed by.
	// Optional, if this is set to the emptry string, the host is inferred automatically by the request.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The port that this spawner should be accessed by.
	// Required.
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSpawnerRequest) Reset()         { *m = RegisterSpawnerRequest{} }
func (m *RegisterSpawnerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterSpawnerRequest) ProtoMessage()    {}
func (*RegisterSpawnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3cfa9e19c71843, []int{0}
}

func (m *RegisterSpawnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSpawnerRequest.Unmarshal(m, b)
}
func (m *RegisterSpawnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSpawnerRequest.Marshal(b, m, deterministic)
}
func (m *RegisterSpawnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSpawnerRequest.Merge(m, src)
}
func (m *RegisterSpawnerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterSpawnerRequest.Size(m)
}
func (m *RegisterSpawnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSpawnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSpawnerRequest proto.InternalMessageInfo

func (m *RegisterSpawnerRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *RegisterSpawnerRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type RegisterSpawnerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterSpawnerResponse) Reset()         { *m = RegisterSpawnerResponse{} }
func (m *RegisterSpawnerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterSpawnerResponse) ProtoMessage()    {}
func (*RegisterSpawnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3cfa9e19c71843, []int{1}
}

func (m *RegisterSpawnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterSpawnerResponse.Unmarshal(m, b)
}
func (m *RegisterSpawnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterSpawnerResponse.Marshal(b, m, deterministic)
}
func (m *RegisterSpawnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterSpawnerResponse.Merge(m, src)
}
func (m *RegisterSpawnerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterSpawnerResponse.Size(m)
}
func (m *RegisterSpawnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterSpawnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterSpawnerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterSpawnerRequest)(nil), "hotel_pb.RegisterSpawnerRequest")
	proto.RegisterType((*RegisterSpawnerResponse)(nil), "hotel_pb.RegisterSpawnerResponse")
}

func init() { proto.RegisterFile("proto/master-service.proto", fileDescriptor_0a3cfa9e19c71843) }

var fileDescriptor_0a3cfa9e19c71843 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0x2e, 0x49, 0x2d, 0xd2, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0xd5, 0x03, 0x0b, 0x0a, 0x71, 0x64, 0xe4, 0x97, 0xa4, 0xe6, 0xc4, 0x17, 0x24, 0x49, 0x09, 0x41,
	0x54, 0x25, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x41, 0x64, 0x95, 0x1c, 0xb8, 0xc4, 0x82, 0x52, 0xd3,
	0x33, 0x41, 0xfa, 0x82, 0x0b, 0x12, 0xcb, 0xf3, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x84, 0xb8, 0x58, 0x32, 0xf2, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x90, 0x58, 0x41, 0x7e, 0x51, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x98,
	0xad, 0x24, 0xc9, 0x25, 0x8e, 0x61, 0x42, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x51, 0x26, 0x17,
	0xaf, 0x2f, 0xd8, 0x49, 0xc1, 0x10, 0x17, 0x09, 0x45, 0x70, 0xf1, 0xa3, 0xa9, 0x15, 0x52, 0xd0,
	0x83, 0xb9, 0x4f, 0x0f, 0xbb, 0x43, 0xa4, 0x14, 0xf1, 0xa8, 0x80, 0x58, 0xa4, 0xc4, 0x90, 0xc4,
	0x06, 0xf6, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x66, 0xb0, 0x7c, 0x0a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterServiceClient interface {
	// Notify the master service that a spawner is available. This is used
	// instead of some service discovery mechanism by the master.
	RegisterSpawner(ctx context.Context, in *RegisterSpawnerRequest, opts ...grpc.CallOption) (*RegisterSpawnerResponse, error)
}

type masterServiceClient struct {
	cc *grpc.ClientConn
}

func NewMasterServiceClient(cc *grpc.ClientConn) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) RegisterSpawner(ctx context.Context, in *RegisterSpawnerRequest, opts ...grpc.CallOption) (*RegisterSpawnerResponse, error) {
	out := new(RegisterSpawnerResponse)
	err := c.cc.Invoke(ctx, "/hotel_pb.MasterService/RegisterSpawner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
type MasterServiceServer interface {
	// Notify the master service that a spawner is available. This is used
	// instead of some service discovery mechanism by the master.
	RegisterSpawner(context.Context, *RegisterSpawnerRequest) (*RegisterSpawnerResponse, error)
}

// UnimplementedMasterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (*UnimplementedMasterServiceServer) RegisterSpawner(ctx context.Context, req *RegisterSpawnerRequest) (*RegisterSpawnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSpawner not implemented")
}

func RegisterMasterServiceServer(s *grpc.Server, srv MasterServiceServer) {
	s.RegisterService(&_MasterService_serviceDesc, srv)
}

func _MasterService_RegisterSpawner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSpawnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).RegisterSpawner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hotel_pb.MasterService/RegisterSpawner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).RegisterSpawner(ctx, req.(*RegisterSpawnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MasterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hotel_pb.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSpawner",
			Handler:    _MasterService_RegisterSpawner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/master-service.proto",
}
