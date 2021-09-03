// Code generated by protoc-gen-go.
// source: cs_common.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	cs_common.proto
	ss_base.proto

It has these top-level messages:
	C2SClientSessionPing
	S2CClientSessionPong
	S2SServerSessionVeriryReq
	S2SServerSessionVeriryAck
	S2SServerSessionPing
	S2SServerSessionPong
	S2SClientSessionPing
	S2SClientSessionPong
	SdConnAttr
	ServiceDiscovery
	ServiceDiscoveryReq
	ServiceDiscoveryAck
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type C2SClientSessionPing struct {
}

func (m *C2SClientSessionPing) Reset()                    { *m = C2SClientSessionPing{} }
func (m *C2SClientSessionPing) String() string            { return proto.CompactTextString(m) }
func (*C2SClientSessionPing) ProtoMessage()               {}
func (*C2SClientSessionPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type S2CClientSessionPong struct {
}

func (m *S2CClientSessionPong) Reset()                    { *m = S2CClientSessionPong{} }
func (m *S2CClientSessionPong) String() string            { return proto.CompactTextString(m) }
func (*S2CClientSessionPong) ProtoMessage()               {}
func (*S2CClientSessionPong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*C2SClientSessionPing)(nil), "pb.c2s_client_session_ping")
	proto.RegisterType((*S2CClientSessionPong)(nil), "pb.s2c_client_session_pong")
}

var fileDescriptor0 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2e, 0x8e, 0x4f,
	0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x92, 0xe4, 0x12, 0x4f, 0x36, 0x02, 0x8a, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0xc4, 0x17, 0xa7, 0x16,
	0x17, 0x67, 0xe6, 0xe7, 0xc5, 0x17, 0x64, 0xe6, 0xa5, 0x83, 0xa4, 0x8a, 0x8d, 0x92, 0x31, 0xa4,
	0xf2, 0xf3, 0xd2, 0x93, 0xd8, 0xc0, 0x06, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0x10,
	0x0e, 0x4e, 0x53, 0x00, 0x00, 0x00,
}