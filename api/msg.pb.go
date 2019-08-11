// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package api

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

type Msg struct {
	Cmd                  int32    `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	File                 string   `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Md5                  string   `protobuf:"bytes,4,opt,name=md5,proto3" json:"md5,omitempty"`
	Msg                  []byte   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *Msg) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Msg) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Msg) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *Msg) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*Msg)(nil), "api.msg")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x8a, 0xe7, 0x62, 0xce, 0x2d,
	0x4e, 0x17, 0x12, 0xe0, 0x62, 0x4e, 0xce, 0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02,
	0x31, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x0b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32,
	0x0b, 0x84, 0x84, 0xb8, 0x58, 0xd2, 0x32, 0x73, 0x52, 0x25, 0x98, 0xc1, 0x22, 0x60, 0x36, 0x48,
	0x57, 0x6e, 0x8a, 0xa9, 0x04, 0x0b, 0x58, 0x08, 0xc4, 0x04, 0x8b, 0x14, 0xa7, 0x4b, 0xb0, 0x2a,
	0x30, 0x6a, 0xf0, 0x04, 0x81, 0x98, 0x49, 0x6c, 0x60, 0xcb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x13, 0x25, 0xbe, 0x79, 0x00, 0x00, 0x00,
}
