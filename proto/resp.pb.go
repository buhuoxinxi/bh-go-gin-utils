// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resp.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	resp.proto

It has these top-level messages:
	ProtoResp
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ProtoResp proto resp with google/protobuf/any.proto
type ProtoResp struct {
	Code    int32                `protobuf:"varint,1,opt,name=Code" json:"Code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	TipInfo string               `protobuf:"bytes,3,opt,name=TipInfo" json:"TipInfo,omitempty"`
	Data    *google_protobuf.Any `protobuf:"bytes,4,opt,name=Data" json:"Data,omitempty"`
}

func (m *ProtoResp) Reset()                    { *m = ProtoResp{} }
func (m *ProtoResp) String() string            { return proto.CompactTextString(m) }
func (*ProtoResp) ProtoMessage()               {}
func (*ProtoResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ProtoResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ProtoResp) GetTipInfo() string {
	if m != nil {
		return m.TipInfo
	}
	return ""
}

func (m *ProtoResp) GetData() *google_protobuf.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtoResp)(nil), "protos.ProtoResp")
}

func init() { proto.RegisterFile("resp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0x2d, 0x2e,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0x92, 0xe9, 0xf9, 0xf9,
	0xe9, 0x39, 0xa9, 0xfa, 0x60, 0x6e, 0x52, 0x69, 0x9a, 0x7e, 0x62, 0x5e, 0x25, 0x44, 0x89, 0x52,
	0x2d, 0x17, 0x67, 0x00, 0x88, 0x11, 0x94, 0x5a, 0x5c, 0x20, 0x24, 0xc4, 0xc5, 0xe2, 0x9c, 0x9f,
	0x92, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x0b, 0x49, 0x70, 0xb1, 0xfb, 0xa6,
	0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x20, 0x99,
	0x90, 0xcc, 0x02, 0xcf, 0xbc, 0xb4, 0x7c, 0x09, 0x66, 0x88, 0x0c, 0x94, 0x2b, 0xa4, 0xc1, 0xc5,
	0xe2, 0x92, 0x58, 0x92, 0x28, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa2, 0x07, 0xb1, 0x5e,
	0x0f, 0x66, 0xbd, 0x9e, 0x63, 0x5e, 0x65, 0x10, 0x58, 0x85, 0x13, 0x47, 0x14, 0xd4, 0x8d, 0x49,
	0x10, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x10, 0xdc, 0x6f, 0xc0, 0x00, 0x00, 0x00,
}
