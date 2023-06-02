// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protos/protocol/core/TronInventoryItems.proto

package protocol_core

import (
	fmt "fmt"
	proto "github.com/bittorrent/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type InventoryItems struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty" pg:"type"`
	Items                [][]byte `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty" pg:"items"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *InventoryItems) Reset()         { *m = InventoryItems{} }
func (m *InventoryItems) String() string { return proto.CompactTextString(m) }
func (*InventoryItems) ProtoMessage()    {}
func (*InventoryItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_25d52db6a1f7f6e0, []int{0}
}
func (m *InventoryItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryItems.Unmarshal(m, b)
}
func (m *InventoryItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryItems.Marshal(b, m, deterministic)
}
func (m *InventoryItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryItems.Merge(m, src)
}
func (m *InventoryItems) XXX_Size() int {
	return xxx_messageInfo_InventoryItems.Size(m)
}
func (m *InventoryItems) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryItems.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryItems proto.InternalMessageInfo

func (m *InventoryItems) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *InventoryItems) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*InventoryItems)(nil), "protocol.InventoryItems")
}

func init() {
	proto.RegisterFile("protos/protocol/core/TronInventoryItems.proto", fileDescriptor_25d52db6a1f7f6e0)
}

var fileDescriptor_25d52db6a1f7f6e0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x07, 0x53, 0xc9, 0xf9, 0x39, 0xfa, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x21, 0x45,
	0xf9, 0x79, 0x9e, 0x79, 0x65, 0xa9, 0x79, 0x25, 0xf9, 0x45, 0x95, 0x9e, 0x25, 0xa9, 0xb9, 0xc5,
	0x7a, 0x60, 0x05, 0x42, 0x1c, 0x30, 0x75, 0x4a, 0x56, 0x5c, 0x7c, 0xa8, 0x2a, 0x84, 0x84, 0xb8,
	0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11,
	0x2e, 0xd6, 0x4c, 0x90, 0xa4, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x4f, 0x10, 0x84, 0xe3, 0xc4, 0x1f,
	0xc5, 0x0b, 0x33, 0x47, 0x0f, 0x64, 0x5f, 0x12, 0x1b, 0x98, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x94, 0x46, 0xeb, 0xd3, 0x8e, 0x00, 0x00, 0x00,
}
