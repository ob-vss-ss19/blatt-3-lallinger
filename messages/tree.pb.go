// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tree.proto

package messages

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

type Usage int32

const (
	Usage_CREATE   Usage = 0
	Usage_ADD      Usage = 1
	Usage_FIND     Usage = 2
	Usage_REMOVE   Usage = 3
	Usage_TRAVERSE Usage = 4
	Usage_DELETE   Usage = 5
)

var Usage_name = map[int32]string{
	0: "CREATE",
	1: "ADD",
	2: "FIND",
	3: "REMOVE",
	4: "TRAVERSE",
	5: "DELETE",
}

var Usage_value = map[string]int32{
	"CREATE":   0,
	"ADD":      1,
	"FIND":     2,
	"REMOVE":   3,
	"TRAVERSE": 4,
	"DELETE":   5,
}

func (x Usage) String() string {
	return proto.EnumName(Usage_name, int32(x))
}

func (Usage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{0}
}

type Request struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Usage    `protobuf:"varint,3,opt,name=type,proto3,enum=messages.Usage" json:"type,omitempty"`
	Key                  int32    `protobuf:"varint,4,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Request) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request) GetType() Usage {
	if m != nil {
		return m.Type
	}
	return Usage_CREATE
}

func (m *Request) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	Key                  int32    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Traverse struct {
	Values               []*Response `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Traverse) Reset()         { *m = Traverse{} }
func (m *Traverse) String() string { return proto.CompactTextString(m) }
func (*Traverse) ProtoMessage()    {}
func (*Traverse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{2}
}

func (m *Traverse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Traverse.Unmarshal(m, b)
}
func (m *Traverse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Traverse.Marshal(b, m, deterministic)
}
func (m *Traverse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Traverse.Merge(m, src)
}
func (m *Traverse) XXX_Size() int {
	return xxx_messageInfo_Traverse.Size(m)
}
func (m *Traverse) XXX_DiscardUnknown() {
	xxx_messageInfo_Traverse.DiscardUnknown(m)
}

var xxx_messageInfo_Traverse proto.InternalMessageInfo

func (m *Traverse) GetValues() []*Response {
	if m != nil {
		return m.Values
	}
	return nil
}

type Error struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3889276909882a, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("messages.Usage", Usage_name, Usage_value)
	proto.RegisterType((*Request)(nil), "messages.Request")
	proto.RegisterType((*Response)(nil), "messages.Response")
	proto.RegisterType((*Traverse)(nil), "messages.Traverse")
	proto.RegisterType((*Error)(nil), "messages.Error")
}

func init() { proto.RegisterFile("tree.proto", fileDescriptor_cb3889276909882a) }

var fileDescriptor_cb3889276909882a = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x4d, 0xdb, 0xb4, 0xf5, 0x29, 0x33, 0x3c, 0x3c, 0xe4, 0x58, 0xea, 0xa5, 0xec, 0xd0,
	0x43, 0x05, 0xef, 0xc5, 0x46, 0x10, 0xa6, 0x42, 0xac, 0xbb, 0x57, 0xf6, 0x90, 0x31, 0x5d, 0x6b,
	0xd2, 0x0d, 0x06, 0xfe, 0xf1, 0x92, 0xcc, 0xa9, 0x87, 0x9d, 0x92, 0x97, 0x2f, 0xbf, 0x1f, 0x1f,
	0x0f, 0x60, 0x34, 0x44, 0xe5, 0x60, 0xfa, 0xb1, 0xc7, 0xf4, 0x83, 0xac, 0xed, 0xde, 0xc8, 0xe6,
	0x5f, 0x90, 0x68, 0xfa, 0xdc, 0x90, 0x1d, 0xf1, 0x12, 0xf8, 0xd8, 0xaf, 0x68, 0x2d, 0x59, 0xc6,
	0x8a, 0x53, 0xbd, 0x1f, 0x70, 0x02, 0xc1, 0x72, 0x21, 0x83, 0x8c, 0x15, 0x5c, 0x07, 0xcb, 0x05,
	0x5e, 0x41, 0x34, 0xee, 0x06, 0x92, 0x61, 0xc6, 0x8a, 0x49, 0x75, 0x51, 0x1e, 0x4c, 0xe5, 0x8b,
	0x3b, 0xb4, 0x0f, 0x51, 0x40, 0xb8, 0xa2, 0x9d, 0x8c, 0x3c, 0xe5, 0xae, 0x4e, 0xbe, 0xed, 0xde,
	0x37, 0x24, 0xf9, 0x5e, 0xee, 0x87, 0xbc, 0x82, 0x54, 0x93, 0x1d, 0xfa, 0xb5, 0xfd, 0x65, 0xd8,
	0x11, 0x26, 0xf8, 0xcf, 0xdc, 0x40, 0xda, 0x9a, 0x6e, 0x4b, 0xc6, 0x12, 0x4e, 0x21, 0xf6, 0x8f,
	0x56, 0xb2, 0x2c, 0x2c, 0xce, 0x2a, 0xfc, 0xab, 0x73, 0xf0, 0xea, 0x9f, 0x1f, 0x79, 0x02, 0x5c,
	0x19, 0xd3, 0x9b, 0xe9, 0x0c, 0xb8, 0xef, 0x8a, 0x00, 0xf1, 0xad, 0x56, 0x75, 0xab, 0xc4, 0x09,
	0x26, 0x10, 0xd6, 0x4d, 0x23, 0x18, 0xa6, 0x10, 0xdd, 0xdd, 0x3f, 0x36, 0x22, 0x70, 0xb1, 0x56,
	0x0f, 0x4f, 0x73, 0x25, 0x42, 0x3c, 0x87, 0xb4, 0xd5, 0xf5, 0x5c, 0xe9, 0x67, 0x25, 0x22, 0x97,
	0x34, 0x6a, 0xa6, 0x5a, 0x25, 0xf8, 0x6b, 0xec, 0x37, 0x7a, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x59, 0x39, 0xe8, 0xce, 0x5f, 0x01, 0x00, 0x00,
}
