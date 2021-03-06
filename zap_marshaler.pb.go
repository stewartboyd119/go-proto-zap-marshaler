// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zap_marshaler.proto

package zap_marshaler // import "github.com/kazegusuri/go-proto-zap-marshaler"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ZapMarshalerRule struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZapMarshalerRule) Reset()         { *m = ZapMarshalerRule{} }
func (m *ZapMarshalerRule) String() string { return proto.CompactTextString(m) }
func (*ZapMarshalerRule) ProtoMessage()    {}
func (*ZapMarshalerRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_zap_marshaler_e0b1c9cd6a84cae1, []int{0}
}
func (m *ZapMarshalerRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZapMarshalerRule.Unmarshal(m, b)
}
func (m *ZapMarshalerRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZapMarshalerRule.Marshal(b, m, deterministic)
}
func (dst *ZapMarshalerRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZapMarshalerRule.Merge(dst, src)
}
func (m *ZapMarshalerRule) XXX_Size() int {
	return xxx_messageInfo_ZapMarshalerRule.Size(m)
}
func (m *ZapMarshalerRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ZapMarshalerRule.DiscardUnknown(m)
}

var xxx_messageInfo_ZapMarshalerRule proto.InternalMessageInfo

func (m *ZapMarshalerRule) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ZapMarshalerRule)(nil),
	Field:         64553,
	Name:          "kazegusuri.zap_mashaler.field",
	Tag:           "bytes,64553,opt,name=field",
	Filename:      "zap_marshaler.proto",
}

func init() {
	proto.RegisterType((*ZapMarshalerRule)(nil), "kazegusuri.zap_mashaler.ZapMarshalerRule")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("zap_marshaler.proto", fileDescriptor_zap_marshaler_e0b1c9cd6a84cae1) }

var fileDescriptor_zap_marshaler_e0b1c9cd6a84cae1 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xae, 0x4a, 0x2c, 0x88,
	0xcf, 0x4d, 0x2c, 0x2a, 0xce, 0x48, 0xcc, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0xcf, 0x4e, 0xac, 0x4a, 0x4d, 0x2f, 0x2d, 0x2e, 0x2d, 0xca, 0xd4, 0x83, 0xc8, 0x43, 0xa4,
	0xa5, 0x14, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xca, 0x92, 0x4a, 0xd3, 0xf4, 0x53,
	0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0xf2, 0xa1, 0x5a, 0x95, 0x74, 0xb8, 0x04, 0xa2, 0x12,
	0x0b, 0x7c, 0x61, 0x06, 0x06, 0x95, 0xe6, 0xa4, 0x0a, 0x49, 0x70, 0xb1, 0xa7, 0xe6, 0x25, 0x26,
	0xe5, 0xa4, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0xc1, 0xb8, 0x56, 0x09, 0x5c, 0xac,
	0x69, 0x99, 0xa9, 0x39, 0x29, 0x42, 0xb2, 0x7a, 0x10, 0x93, 0xf5, 0x60, 0x26, 0xeb, 0xb9, 0x81,
	0xc4, 0xfd, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0x25, 0x56, 0xfe, 0x60, 0x56, 0x60, 0xd4, 0xe0,
	0x36, 0xd2, 0xd4, 0xc3, 0xe1, 0x32, 0x3d, 0x74, 0x4b, 0x83, 0x20, 0x06, 0x3b, 0xd9, 0x44, 0x59,
	0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x23, 0x74, 0xeb, 0xa7, 0xe7,
	0xeb, 0x82, 0xed, 0xd2, 0xad, 0x4a, 0x2c, 0xd0, 0x85, 0x07, 0x80, 0x35, 0x4a, 0x70, 0x24, 0xb1,
	0x81, 0x95, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0xe9, 0x41, 0xb1, 0x26, 0x01, 0x00,
	0x00,
}
