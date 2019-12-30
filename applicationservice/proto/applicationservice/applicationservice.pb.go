// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/applicationservice/applicationservice.proto

package go_micro_srv_applicationservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  string   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{1}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

type AllPersonResponse struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AllPersonResponse) Reset()         { *m = AllPersonResponse{} }
func (m *AllPersonResponse) String() string { return proto.CompactTextString(m) }
func (*AllPersonResponse) ProtoMessage()    {}
func (*AllPersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{2}
}
func (m *AllPersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllPersonResponse.Unmarshal(m, b)
}
func (m *AllPersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllPersonResponse.Marshal(b, m, deterministic)
}
func (dst *AllPersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllPersonResponse.Merge(dst, src)
}
func (m *AllPersonResponse) XXX_Size() int {
	return xxx_messageInfo_AllPersonResponse.Size(m)
}
func (m *AllPersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllPersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllPersonResponse proto.InternalMessageInfo

func (m *AllPersonResponse) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{3}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{5}
}
func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (dst *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(dst, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{6}
}
func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (dst *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(dst, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{7}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_applicationservice_3ab1d368a63cba6e, []int{8}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (dst *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(dst, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.applicationservice.Message")
	proto.RegisterType((*Person)(nil), "go.micro.srv.applicationservice.Person")
	proto.RegisterType((*AllPersonResponse)(nil), "go.micro.srv.applicationservice.AllPersonResponse")
	proto.RegisterType((*Request)(nil), "go.micro.srv.applicationservice.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.applicationservice.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.applicationservice.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.applicationservice.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.applicationservice.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.applicationservice.Pong")
}

func init() {
	proto.RegisterFile("proto/applicationservice/applicationservice.proto", fileDescriptor_applicationservice_3ab1d368a63cba6e)
}

var fileDescriptor_applicationservice_3ab1d368a63cba6e = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x4a, 0xeb, 0x40,
	0x10, 0xc6, 0x93, 0x93, 0x9e, 0xb4, 0x67, 0x0e, 0x42, 0xbb, 0x88, 0x94, 0xfa, 0xaf, 0x2c, 0x88,
	0xe9, 0x4d, 0x6c, 0xeb, 0x13, 0x14, 0x2f, 0xbc, 0x12, 0x4a, 0x04, 0xef, 0x14, 0xd6, 0x30, 0x2c,
	0xc1, 0x64, 0x37, 0x66, 0xd2, 0x82, 0x8f, 0xe6, 0xdb, 0x49, 0x36, 0xdb, 0x22, 0xa6, 0x92, 0xdc,
	0xed, 0xee, 0xfc, 0x66, 0x66, 0xbf, 0xf9, 0x76, 0x61, 0x91, 0x17, 0xba, 0xd4, 0x37, 0x22, 0xcf,
	0xd3, 0x24, 0x16, 0x65, 0xa2, 0x15, 0x61, 0xb1, 0x4d, 0x62, 0x3c, 0x70, 0x14, 0x1a, 0x96, 0x5d,
	0x4a, 0x1d, 0x66, 0x49, 0x5c, 0xe8, 0x90, 0x8a, 0x6d, 0xd8, 0xc4, 0xf8, 0x29, 0xf4, 0x1f, 0x90,
	0x48, 0x48, 0x64, 0x43, 0xf0, 0x48, 0x7c, 0x8c, 0xdd, 0xa9, 0x1b, 0xfc, 0x8b, 0xaa, 0x25, 0x0f,
	0xc1, 0x5f, 0x63, 0x41, 0x5a, 0x31, 0x06, 0x3d, 0x25, 0x32, 0xb4, 0x41, 0xb3, 0xae, 0x78, 0x21,
	0x71, 0xfc, 0xa7, 0xe6, 0x85, 0x44, 0xfe, 0x04, 0xa3, 0x55, 0x9a, 0xd6, 0x29, 0x11, 0x52, 0x5e,
	0xb5, 0x61, 0x2b, 0xe8, 0xe7, 0xe6, 0x84, 0xc6, 0xee, 0xd4, 0x0b, 0xfe, 0x2f, 0xaf, 0xc3, 0x96,
	0x4b, 0x85, 0xb6, 0xc2, 0x2e, 0x8f, 0x9f, 0x43, 0x3f, 0xc2, 0xf7, 0x0d, 0x52, 0x79, 0xe8, 0x22,
	0xfc, 0x0c, 0x06, 0xfb, 0x6e, 0x43, 0xf0, 0x32, 0x92, 0x3b, 0x11, 0x19, 0x49, 0x1e, 0xc0, 0xf0,
	0xb1, 0x2c, 0x50, 0x64, 0x89, 0x92, 0xbb, 0x2a, 0xc7, 0xf0, 0x37, 0xd6, 0x1b, 0x55, 0x1a, 0xce,
	0x8b, 0xea, 0x0d, 0x9f, 0xc1, 0xe8, 0x1b, 0x69, 0x0b, 0x1e, 0x46, 0x2f, 0xa0, 0xb7, 0x4e, 0x94,
	0x64, 0x27, 0xe0, 0x53, 0x59, 0xe8, 0x37, 0xb4, 0x61, 0xbb, 0x33, 0x71, 0xfd, 0x7b, 0x7c, 0xf9,
	0xe9, 0x01, 0x5b, 0x35, 0x84, 0xb3, 0x67, 0xe8, 0xdd, 0x89, 0x34, 0x65, 0x41, 0xeb, 0x88, 0xac,
	0x92, 0xc9, 0xac, 0x03, 0x59, 0x2b, 0xe1, 0x0e, 0xd3, 0x70, 0x74, 0x8f, 0xe5, 0xde, 0x22, 0xea,
	0xd0, 0xc7, 0x3e, 0x8e, 0xc9, 0xb2, 0x95, 0x6c, 0x38, 0xcf, 0x1d, 0x46, 0xe0, 0xd7, 0x13, 0x65,
	0x8b, 0xd6, 0xfc, 0x9f, 0x26, 0x75, 0x68, 0xd9, 0x70, 0x8b, 0x3b, 0x73, 0x97, 0xbd, 0xc0, 0xa0,
	0xf2, 0xc6, 0xcc, 0xff, 0xaa, 0xfd, 0xad, 0x25, 0x4a, 0x4e, 0x3a, 0x60, 0x5a, 0x49, 0xee, 0x04,
	0xee, 0xdc, 0x7d, 0xf5, 0xcd, 0xd7, 0xba, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb5, 0x87,
	0x63, 0x8f, 0x03, 0x00, 0x00,
}
