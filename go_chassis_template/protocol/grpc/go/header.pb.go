// Code generated by protoc-gen-go. DO NOT EDIT.
// source: header.proto

package template_proto

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

type ReqHeader struct {
	RequestId            *string  `protobuf:"bytes,1,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	Token                *string  `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	Timestamp            *uint32  `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	Nonce                *uint32  `protobuf:"varint,4,req,name=nonce" json:"nonce,omitempty"`
	Sign                 *uint32  `protobuf:"varint,5,req,name=sign" json:"sign,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqHeader) Reset()         { *m = ReqHeader{} }
func (m *ReqHeader) String() string { return proto.CompactTextString(m) }
func (*ReqHeader) ProtoMessage()    {}
func (*ReqHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{0}
}

func (m *ReqHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqHeader.Unmarshal(m, b)
}
func (m *ReqHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqHeader.Marshal(b, m, deterministic)
}
func (m *ReqHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqHeader.Merge(m, src)
}
func (m *ReqHeader) XXX_Size() int {
	return xxx_messageInfo_ReqHeader.Size(m)
}
func (m *ReqHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ReqHeader proto.InternalMessageInfo

func (m *ReqHeader) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *ReqHeader) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *ReqHeader) GetTimestamp() uint32 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *ReqHeader) GetNonce() uint32 {
	if m != nil && m.Nonce != nil {
		return *m.Nonce
	}
	return 0
}

func (m *ReqHeader) GetSign() uint32 {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return 0
}

type RespHeader struct {
	RequestId            *string  `protobuf:"bytes,1,req,name=request_id,json=requestId" json:"request_id,omitempty"`
	Retcode              *int32   `protobuf:"varint,2,req,name=retcode" json:"retcode,omitempty"`
	Message              *string  `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespHeader) Reset()         { *m = RespHeader{} }
func (m *RespHeader) String() string { return proto.CompactTextString(m) }
func (*RespHeader) ProtoMessage()    {}
func (*RespHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6398613e36d6c2ce, []int{1}
}

func (m *RespHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespHeader.Unmarshal(m, b)
}
func (m *RespHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespHeader.Marshal(b, m, deterministic)
}
func (m *RespHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespHeader.Merge(m, src)
}
func (m *RespHeader) XXX_Size() int {
	return xxx_messageInfo_RespHeader.Size(m)
}
func (m *RespHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RespHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RespHeader proto.InternalMessageInfo

func (m *RespHeader) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

func (m *RespHeader) GetRetcode() int32 {
	if m != nil && m.Retcode != nil {
		return *m.Retcode
	}
	return 0
}

func (m *RespHeader) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqHeader)(nil), "template_proto.ReqHeader")
	proto.RegisterType((*RespHeader)(nil), "template_proto.RespHeader")
}

func init() {
	proto.RegisterFile("header.proto", fileDescriptor_6398613e36d6c2ce)
}

var fileDescriptor_6398613e36d6c2ce = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0xcd, 0xaa, 0xc2, 0x30,
	0x10, 0x46, 0x69, 0x6f, 0xcb, 0x25, 0xc3, 0xbd, 0x2e, 0x82, 0x8b, 0x2c, 0x14, 0x4a, 0x57, 0x5d,
	0xf9, 0x1c, 0xba, 0xcd, 0x0b, 0x94, 0xd0, 0x0c, 0xb5, 0x68, 0x7e, 0x9a, 0x8c, 0xef, 0xe0, 0x63,
	0x4b, 0xa7, 0x16, 0xb7, 0xee, 0x72, 0xbe, 0x93, 0x81, 0x03, 0x7f, 0x57, 0x34, 0x16, 0xd3, 0x29,
	0xa6, 0x40, 0x41, 0xee, 0x08, 0x5d, 0xbc, 0x1b, 0xc2, 0x9e, 0xb9, 0x7d, 0x16, 0x20, 0x34, 0xce,
	0x67, 0xfe, 0x23, 0x8f, 0x00, 0x09, 0xe7, 0x07, 0x66, 0xea, 0x27, 0xab, 0x8a, 0xa6, 0xec, 0x84,
	0x16, 0xef, 0xe5, 0x62, 0xe5, 0x1e, 0x6a, 0x0a, 0x37, 0xf4, 0xaa, 0x64, 0xb3, 0x82, 0x3c, 0x80,
	0xa0, 0xc9, 0x61, 0x26, 0xe3, 0xa2, 0xfa, 0x69, 0xca, 0xee, 0x5f, 0x7f, 0x86, 0xe5, 0xc6, 0x07,
	0x3f, 0xa0, 0xaa, 0xd8, 0xac, 0x20, 0x25, 0x54, 0x79, 0x1a, 0xbd, 0xaa, 0x79, 0xe4, 0x77, 0xdb,
	0x03, 0x68, 0xcc, 0xf1, 0xbb, 0x14, 0x05, 0xbf, 0x09, 0x69, 0x08, 0x16, 0x39, 0xa6, 0xd6, 0x1b,
	0x2e, 0xc6, 0x61, 0xce, 0x66, 0x44, 0x8e, 0x11, 0x7a, 0xc3, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x70, 0xfb, 0xca, 0x3d, 0x0a, 0x01, 0x00, 0x00,
}
