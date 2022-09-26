// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package proto

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

type GreetRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0x4d,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x74, 0xb9, 0x78, 0xdc,
	0x41, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x59, 0x2e, 0xae, 0xb4, 0xcc, 0xa2,
	0xe2, 0x92, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0,
	0x88, 0x5f, 0x62, 0x6e, 0xaa, 0x92, 0x3a, 0x17, 0x2f, 0x54, 0x79, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x54, 0x2d, 0x94, 0x67, 0xf4,
	0x9b, 0x11, 0x6a, 0x70, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x11, 0x17, 0x2b, 0x98,
	0x2f, 0x24, 0xac, 0x07, 0x71, 0x06, 0xb2, 0xb5, 0x52, 0x22, 0xa8, 0x82, 0x50, 0xc3, 0x6d, 0xb9,
	0xf8, 0xc0, 0x02, 0xbe, 0x89, 0x79, 0x95, 0x21, 0x99, 0xb9, 0xa9, 0xc5, 0x24, 0x68, 0x36, 0x60,
	0x14, 0xb2, 0xe0, 0xe2, 0xf4, 0xc9, 0xcf, 0x4b, 0x27, 0xd5, 0x5a, 0x0d, 0x46, 0x21, 0x3b, 0xa8,
	0x37, 0x5d, 0xcb, 0x52, 0x8b, 0x2a, 0xf3, 0xf3, 0x52, 0x49, 0xd2, 0x6d, 0xc0, 0xe8, 0x64, 0x18,
	0xa5, 0x9f, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x9b, 0x9a, 0x03,
	0xd2, 0x10, 0x90, 0x98, 0x59, 0x96, 0xa8, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x9b, 0x9c, 0x5f, 0x5a,
	0x54, 0x9c, 0xaa, 0x0f, 0xd6, 0xae, 0x0f, 0x8e, 0x8f, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x41, 0xdd, 0x57, 0xec, 0xa5, 0x01, 0x00, 0x00,
}
