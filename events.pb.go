// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package main

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

type L9Event struct {
	ID                   string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp            int64             `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Component            string            `protobuf:"bytes,3,opt,name=Component,proto3" json:"Component,omitempty"`
	Host                 string            `protobuf:"bytes,4,opt,name=Host,proto3" json:"Host,omitempty"`
	Message              string            `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	Namespace            string            `protobuf:"bytes,6,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Reason               string            `protobuf:"bytes,7,opt,name=Reason,proto3" json:"Reason,omitempty"`
	ReferenceUID         string            `protobuf:"bytes,8,opt,name=ReferenceUID,proto3" json:"ReferenceUID,omitempty"`
	ReferenceNamespace   string            `protobuf:"bytes,9,opt,name=ReferenceNamespace,proto3" json:"ReferenceNamespace,omitempty"`
	ReferenceName        string            `protobuf:"bytes,10,opt,name=ReferenceName,proto3" json:"ReferenceName,omitempty"`
	ReferenceKind        string            `protobuf:"bytes,11,opt,name=ReferenceKind,proto3" json:"ReferenceKind,omitempty"`
	ReferenceVersion     string            `protobuf:"bytes,12,opt,name=ReferenceVersion,proto3" json:"ReferenceVersion,omitempty"`
	ObjectUid            string            `protobuf:"bytes,13,opt,name=ObjectUid,proto3" json:"ObjectUid,omitempty"`
	Labels               map[string]string `protobuf:"bytes,14,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,15,rep,name=Annotations,proto3" json:"Annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Address              []string          `protobuf:"bytes,16,rep,name=Address,proto3" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L9Event) Reset()         { *m = L9Event{} }
func (m *L9Event) String() string { return proto.CompactTextString(m) }
func (*L9Event) ProtoMessage()    {}
func (*L9Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *L9Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L9Event.Unmarshal(m, b)
}
func (m *L9Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L9Event.Marshal(b, m, deterministic)
}
func (m *L9Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L9Event.Merge(m, src)
}
func (m *L9Event) XXX_Size() int {
	return xxx_messageInfo_L9Event.Size(m)
}
func (m *L9Event) XXX_DiscardUnknown() {
	xxx_messageInfo_L9Event.DiscardUnknown(m)
}

var xxx_messageInfo_L9Event proto.InternalMessageInfo

func (m *L9Event) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *L9Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *L9Event) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *L9Event) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *L9Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *L9Event) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *L9Event) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *L9Event) GetReferenceUID() string {
	if m != nil {
		return m.ReferenceUID
	}
	return ""
}

func (m *L9Event) GetReferenceNamespace() string {
	if m != nil {
		return m.ReferenceNamespace
	}
	return ""
}

func (m *L9Event) GetReferenceName() string {
	if m != nil {
		return m.ReferenceName
	}
	return ""
}

func (m *L9Event) GetReferenceKind() string {
	if m != nil {
		return m.ReferenceKind
	}
	return ""
}

func (m *L9Event) GetReferenceVersion() string {
	if m != nil {
		return m.ReferenceVersion
	}
	return ""
}

func (m *L9Event) GetObjectUid() string {
	if m != nil {
		return m.ObjectUid
	}
	return ""
}

func (m *L9Event) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *L9Event) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *L9Event) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

type L9EventBatch struct {
	Events               []*L9Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *L9EventBatch) Reset()         { *m = L9EventBatch{} }
func (m *L9EventBatch) String() string { return proto.CompactTextString(m) }
func (*L9EventBatch) ProtoMessage()    {}
func (*L9EventBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *L9EventBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L9EventBatch.Unmarshal(m, b)
}
func (m *L9EventBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L9EventBatch.Marshal(b, m, deterministic)
}
func (m *L9EventBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L9EventBatch.Merge(m, src)
}
func (m *L9EventBatch) XXX_Size() int {
	return xxx_messageInfo_L9EventBatch.Size(m)
}
func (m *L9EventBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_L9EventBatch.DiscardUnknown(m)
}

var xxx_messageInfo_L9EventBatch proto.InternalMessageInfo

func (m *L9EventBatch) GetEvents() []*L9Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type EventError struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventError) Reset()         { *m = EventError{} }
func (m *EventError) String() string { return proto.CompactTextString(m) }
func (*EventError) ProtoMessage()    {}
func (*EventError) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventError.Unmarshal(m, b)
}
func (m *EventError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventError.Marshal(b, m, deterministic)
}
func (m *EventError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventError.Merge(m, src)
}
func (m *EventError) XXX_Size() int {
	return xxx_messageInfo_EventError.Size(m)
}
func (m *EventError) XXX_DiscardUnknown() {
	xxx_messageInfo_EventError.DiscardUnknown(m)
}

var xxx_messageInfo_EventError proto.InternalMessageInfo

func (m *EventError) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EventError) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type Response struct {
	Ok                   bool          `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Batch                string        `protobuf:"bytes,2,opt,name=Batch,proto3" json:"Batch,omitempty"`
	Errors               []*EventError `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{3}
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

func (m *Response) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *Response) GetBatch() string {
	if m != nil {
		return m.Batch
	}
	return ""
}

func (m *Response) GetErrors() []*EventError {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*L9Event)(nil), "main.L9Event")
	proto.RegisterMapType((map[string]string)(nil), "main.L9Event.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "main.L9Event.LabelsEntry")
	proto.RegisterType((*L9EventBatch)(nil), "main.L9EventBatch")
	proto.RegisterType((*EventError)(nil), "main.EventError")
	proto.RegisterType((*Response)(nil), "main.Response")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x51, 0x8b, 0x13, 0x31,
	0x10, 0xa6, 0xbb, 0xed, 0xb6, 0x9d, 0xb6, 0xe7, 0x12, 0x44, 0xe2, 0x21, 0x52, 0x16, 0x85, 0xe2,
	0xc3, 0x82, 0x8a, 0xe0, 0xdd, 0x83, 0x78, 0xda, 0x82, 0xc5, 0xd3, 0x42, 0xf0, 0x7c, 0xf0, 0x2d,
	0xed, 0x8e, 0xba, 0xf6, 0x36, 0x59, 0x36, 0xf1, 0xf0, 0x7e, 0x94, 0xff, 0x51, 0x32, 0xc9, 0xb5,
	0xdd, 0xd3, 0x17, 0xdf, 0x32, 0xdf, 0x37, 0xdf, 0xcc, 0x97, 0xcc, 0x04, 0xc6, 0x78, 0x85, 0xca,
	0x9a, 0xbc, 0x6e, 0xb4, 0xd5, 0xac, 0x5b, 0xc9, 0x52, 0x65, 0xbf, 0x7b, 0xd0, 0x3f, 0x3f, 0x59,
	0x38, 0x82, 0x1d, 0x41, 0xb4, 0x9c, 0xf3, 0xce, 0xb4, 0x33, 0x1b, 0x8a, 0x68, 0x39, 0x67, 0x0f,
	0x60, 0xf8, 0xa9, 0xac, 0xd0, 0x58, 0x59, 0xd5, 0x3c, 0x9a, 0x76, 0x66, 0xb1, 0xd8, 0x03, 0x8e,
	0x7d, 0xab, 0xab, 0x5a, 0x2b, 0x54, 0x96, 0xc7, 0x24, 0xda, 0x03, 0x8c, 0x41, 0xf7, 0x9d, 0x36,
	0x96, 0x77, 0x89, 0xa0, 0x33, 0xe3, 0xd0, 0xff, 0x80, 0xc6, 0xc8, 0x6f, 0xc8, 0x7b, 0x04, 0xdf,
	0x84, 0xae, 0xd6, 0x47, 0x59, 0xa1, 0xa9, 0xe5, 0x06, 0x79, 0xe2, 0x6b, 0xed, 0x00, 0x76, 0x0f,
	0x12, 0x81, 0xd2, 0x68, 0xc5, 0xfb, 0x44, 0x85, 0x88, 0x65, 0x30, 0x16, 0xf8, 0x15, 0x1b, 0x54,
	0x1b, 0xbc, 0x58, 0xce, 0xf9, 0x80, 0xd8, 0x16, 0xc6, 0x72, 0x60, 0xbb, 0x78, 0xdf, 0x62, 0x48,
	0x99, 0xff, 0x60, 0xd8, 0x23, 0x98, 0xb4, 0x50, 0x0e, 0x94, 0xda, 0x06, 0x5b, 0x59, 0xef, 0x4b,
	0x55, 0xf0, 0xd1, 0xad, 0x2c, 0x07, 0xb2, 0x27, 0x90, 0xee, 0x80, 0xcf, 0xd8, 0x98, 0x52, 0x2b,
	0x3e, 0xa6, 0xc4, 0xbf, 0x70, 0xf7, 0x02, 0xab, 0xf5, 0x0f, 0xdc, 0xd8, 0x8b, 0xb2, 0xe0, 0x13,
	0xff, 0x02, 0x3b, 0x80, 0x3d, 0x85, 0xe4, 0x5c, 0xae, 0xf1, 0xd2, 0xf0, 0xa3, 0x69, 0x3c, 0x1b,
	0x3d, 0xbb, 0x9f, 0xbb, 0xe1, 0xe5, 0x61, 0x70, 0xb9, 0xe7, 0x16, 0xca, 0x36, 0xd7, 0x22, 0x24,
	0xb2, 0xd7, 0x30, 0x3a, 0x53, 0x4a, 0x5b, 0x69, 0x4b, 0xad, 0x0c, 0xbf, 0x43, 0xba, 0x87, 0x6d,
	0xdd, 0x41, 0x82, 0x17, 0x1f, 0x4a, 0xdc, 0xb8, 0xce, 0x8a, 0xa2, 0x41, 0x63, 0x78, 0x3a, 0x8d,
	0xdd, 0xb8, 0x42, 0x78, 0x7c, 0x02, 0xa3, 0x83, 0x96, 0x2c, 0x85, 0x78, 0x8b, 0xd7, 0x61, 0x71,
	0xdc, 0x91, 0xdd, 0x85, 0xde, 0x95, 0xbc, 0xfc, 0x89, 0xb4, 0x35, 0x43, 0xe1, 0x83, 0xd3, 0xe8,
	0x65, 0xe7, 0xf8, 0x15, 0xa4, 0xb7, 0xbb, 0xfe, 0x8f, 0x3e, 0x7b, 0x01, 0xe3, 0xe0, 0xfe, 0x8d,
	0xb4, 0x9b, 0xef, 0xec, 0x31, 0x24, 0x7e, 0xab, 0x79, 0x87, 0x6e, 0x38, 0x69, 0xdd, 0x50, 0x04,
	0x32, 0x3b, 0x05, 0x20, 0x60, 0xd1, 0x34, 0xba, 0x71, 0xe5, 0x4b, 0x55, 0xe0, 0x2f, 0x6a, 0xd9,
	0x13, 0x3e, 0x70, 0x6b, 0xd6, 0xf8, 0x35, 0xf3, 0x5d, 0x43, 0x94, 0x7d, 0x81, 0x81, 0x40, 0x53,
	0x6b, 0x65, 0xd0, 0x7d, 0x91, 0xd5, 0x96, 0x64, 0x03, 0x11, 0xad, 0xb6, 0xae, 0x12, 0xf9, 0xb8,
	0x31, 0xea, 0x4d, 0xcd, 0x20, 0x41, 0xd7, 0xc8, 0xf0, 0x98, 0x4c, 0xa5, 0xde, 0xd4, 0xde, 0x81,
	0x08, 0xfc, 0x3a, 0xa1, 0xbf, 0xf8, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x51, 0xae,
	0xd0, 0x9b, 0x03, 0x00, 0x00,
}
