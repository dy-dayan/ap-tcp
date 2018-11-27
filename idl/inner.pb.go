// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inner.proto

package access

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CODE int32

const (
	CODE_OK CODE = 0
	CODE_NO CODE = 1
)

var CODE_name = map[int32]string{
	0: "OK",
	1: "NO",
}

var CODE_value = map[string]int32{
	"OK": 0,
	"NO": 1,
}

func (x CODE) String() string {
	return proto.EnumName(CODE_name, int32(x))
}

func (CODE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{0}
}

type PushReq struct {
	Uid                  uint64   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Content              []byte   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Seq                  uint64   `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushReq) Reset()         { *m = PushReq{} }
func (m *PushReq) String() string { return proto.CompactTextString(m) }
func (*PushReq) ProtoMessage()    {}
func (*PushReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{0}
}

func (m *PushReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushReq.Unmarshal(m, b)
}
func (m *PushReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushReq.Marshal(b, m, deterministic)
}
func (m *PushReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushReq.Merge(m, src)
}
func (m *PushReq) XXX_Size() int {
	return xxx_messageInfo_PushReq.Size(m)
}
func (m *PushReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushReq proto.InternalMessageInfo

func (m *PushReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PushReq) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *PushReq) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PushReq) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PushReq) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type PushRsp struct {
	Code                 CODE     `protobuf:"varint,1,opt,name=code,proto3,enum=access.CODE" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRsp) Reset()         { *m = PushRsp{} }
func (m *PushRsp) String() string { return proto.CompactTextString(m) }
func (*PushRsp) ProtoMessage()    {}
func (*PushRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778376c45113e49, []int{1}
}

func (m *PushRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRsp.Unmarshal(m, b)
}
func (m *PushRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRsp.Marshal(b, m, deterministic)
}
func (m *PushRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRsp.Merge(m, src)
}
func (m *PushRsp) XXX_Size() int {
	return xxx_messageInfo_PushRsp.Size(m)
}
func (m *PushRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRsp.DiscardUnknown(m)
}

var xxx_messageInfo_PushRsp proto.InternalMessageInfo

func (m *PushRsp) GetCode() CODE {
	if m != nil {
		return m.Code
	}
	return CODE_OK
}

func init() {
	proto.RegisterEnum("access.CODE", CODE_name, CODE_value)
	proto.RegisterType((*PushReq)(nil), "access.PushReq")
	proto.RegisterType((*PushRsp)(nil), "access.PushRsp")
}

func init() { proto.RegisterFile("inner.proto", fileDescriptor_9778376c45113e49) }

var fileDescriptor_9778376c45113e49 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4f, 0xc4, 0x20,
	0x10, 0x85, 0x65, 0x17, 0xd9, 0x38, 0x6e, 0xb4, 0x99, 0xc3, 0x86, 0x78, 0x22, 0x7b, 0x22, 0x9a,
	0xf4, 0x50, 0x7f, 0x81, 0x51, 0x4f, 0x26, 0xd6, 0xf0, 0x0f, 0x94, 0x4e, 0xb2, 0x3d, 0x08, 0x6d,
	0x87, 0x35, 0xf1, 0xdf, 0x1b, 0x68, 0x7b, 0xf0, 0x04, 0x1f, 0x8f, 0xf7, 0x1e, 0x0c, 0x5c, 0xf7,
	0x21, 0xd0, 0x54, 0x0f, 0x53, 0x4c, 0x11, 0xd5, 0xa7, 0xf7, 0xc4, 0x7c, 0xfc, 0x85, 0xdd, 0xc7,
	0x99, 0x4f, 0x8e, 0x46, 0xac, 0x60, 0x7b, 0xee, 0x3b, 0x2d, 0x8c, 0xb0, 0xd2, 0xe5, 0x2d, 0x6a,
	0xd8, 0x31, 0x4d, 0x3f, 0xbd, 0x27, 0xbd, 0x31, 0xc2, 0x5e, 0xb9, 0x15, 0xf1, 0x00, 0xea, 0x9b,
	0xd2, 0x29, 0x76, 0x7a, 0x5b, 0x84, 0x85, 0xb2, 0xc3, 0xc7, 0x90, 0x28, 0x24, 0x2d, 0x8d, 0xb0,
	0x7b, 0xb7, 0x62, 0x4e, 0x67, 0x1a, 0xf5, 0xe5, 0x9c, 0xce, 0x34, 0x1e, 0x1f, 0x96, 0x6a, 0x1e,
	0xd0, 0x80, 0xf4, 0xb1, 0xa3, 0xd2, 0x7d, 0xd3, 0xec, 0xeb, 0xf9, 0x71, 0xf5, 0x73, 0xfb, 0xf2,
	0xea, 0x8a, 0x72, 0x7f, 0x00, 0x99, 0x09, 0x15, 0x6c, 0xda, 0xb7, 0xea, 0x22, 0xaf, 0xef, 0x6d,
	0x25, 0x9a, 0x06, 0xd4, 0x53, 0xb9, 0x8c, 0x16, 0x64, 0x8e, 0xc3, 0xdb, 0xd5, 0xbd, 0xfc, 0xeb,
	0xee, 0xff, 0x01, 0x0f, 0x5f, 0xaa, 0x8c, 0xe0, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x49,
	0x26, 0xb1, 0x11, 0x01, 0x00, 0x00,
}
