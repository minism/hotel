// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package hotel_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Status int32

const (
	// Common status codes.
	Status_OK              Status = 0
	Status_UNKNOWN_ERROR   Status = 1
	Status_INVALID_REQUEST Status = 2
	// Semantic errors.
	Status_MAX_GAME_SERVERS_REACHED Status = 3
)

var Status_name = map[int32]string{
	0: "OK",
	1: "UNKNOWN_ERROR",
	2: "INVALID_REQUEST",
	3: "MAX_GAME_SERVERS_REACHED",
}

var Status_value = map[string]int32{
	"OK":                       0,
	"UNKNOWN_ERROR":            1,
	"INVALID_REQUEST":          2,
	"MAX_GAME_SERVERS_REACHED": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

type GameServer struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameServer) Reset()         { *m = GameServer{} }
func (m *GameServer) String() string { return proto.CompactTextString(m) }
func (*GameServer) ProtoMessage()    {}
func (*GameServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *GameServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameServer.Unmarshal(m, b)
}
func (m *GameServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameServer.Marshal(b, m, deterministic)
}
func (m *GameServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameServer.Merge(m, src)
}
func (m *GameServer) XXX_Size() int {
	return xxx_messageInfo_GameServer.Size(m)
}
func (m *GameServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GameServer.DiscardUnknown(m)
}

var xxx_messageInfo_GameServer proto.InternalMessageInfo

func (m *GameServer) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GameServer) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("hotel_pb.Status", Status_name, Status_value)
	proto.RegisterType((*GameServer)(nil), "hotel_pb.GameServer")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0x73, 0x84, 0x38, 0x32, 0xf2, 0x4b,
	0x52, 0x73, 0xe2, 0x0b, 0x92, 0x94, 0x4c, 0xb8, 0xb8, 0xdc, 0x13, 0x73, 0x53, 0x83, 0x53, 0x8b,
	0xca, 0x52, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0x32, 0xf2, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x41, 0x7e, 0x51, 0x89, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6f,
	0x10, 0x98, 0xad, 0x15, 0xc6, 0xc5, 0x16, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc4, 0xc6, 0xc5,
	0xe4, 0xef, 0x2d, 0xc0, 0x20, 0x24, 0xc8, 0xc5, 0x1b, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x17,
	0xef, 0x1a, 0x14, 0xe4, 0x1f, 0x24, 0xc0, 0x28, 0x24, 0xcc, 0xc5, 0xef, 0xe9, 0x17, 0xe6, 0xe8,
	0xe3, 0xe9, 0x12, 0x1f, 0xe4, 0x1a, 0x18, 0xea, 0x1a, 0x1c, 0x22, 0xc0, 0x24, 0x24, 0xc3, 0x25,
	0xe1, 0xeb, 0x18, 0x11, 0xef, 0xee, 0xe8, 0xeb, 0x1a, 0x1f, 0xec, 0x1a, 0x14, 0xe6, 0x1a, 0x14,
	0x1c, 0x1f, 0xe4, 0xea, 0xe8, 0xec, 0xe1, 0xea, 0x22, 0xc0, 0x9c, 0xc4, 0x06, 0x76, 0x9e, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0xa1, 0x05, 0x4c, 0xb4, 0x00, 0x00, 0x00,
}