// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasources/sql/sql.proto

package sql

import (
	fmt "fmt"
	execution "github.com/cube2222/octosql/execution"
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

type QueueElement struct {
	// Types that are valid to be assigned to Type:
	//	*QueueElement_Record
	//	*QueueElement_EndOfStream
	//	*QueueElement_Error
	Type                 isQueueElement_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueueElement) Reset()         { *m = QueueElement{} }
func (m *QueueElement) String() string { return proto.CompactTextString(m) }
func (*QueueElement) ProtoMessage()    {}
func (*QueueElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebfa9ffc66e09d2b, []int{0}
}

func (m *QueueElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueElement.Unmarshal(m, b)
}
func (m *QueueElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueElement.Marshal(b, m, deterministic)
}
func (m *QueueElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueElement.Merge(m, src)
}
func (m *QueueElement) XXX_Size() int {
	return xxx_messageInfo_QueueElement.Size(m)
}
func (m *QueueElement) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueElement.DiscardUnknown(m)
}

var xxx_messageInfo_QueueElement proto.InternalMessageInfo

type isQueueElement_Type interface {
	isQueueElement_Type()
}

type QueueElement_Record struct {
	Record *execution.Record `protobuf:"bytes,1,opt,name=record,proto3,oneof"`
}

type QueueElement_EndOfStream struct {
	EndOfStream bool `protobuf:"varint,2,opt,name=endOfStream,proto3,oneof"`
}

type QueueElement_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*QueueElement_Record) isQueueElement_Type() {}

func (*QueueElement_EndOfStream) isQueueElement_Type() {}

func (*QueueElement_Error) isQueueElement_Type() {}

func (m *QueueElement) GetType() isQueueElement_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *QueueElement) GetRecord() *execution.Record {
	if x, ok := m.GetType().(*QueueElement_Record); ok {
		return x.Record
	}
	return nil
}

func (m *QueueElement) GetEndOfStream() bool {
	if x, ok := m.GetType().(*QueueElement_EndOfStream); ok {
		return x.EndOfStream
	}
	return false
}

func (m *QueueElement) GetError() string {
	if x, ok := m.GetType().(*QueueElement_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QueueElement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QueueElement_Record)(nil),
		(*QueueElement_EndOfStream)(nil),
		(*QueueElement_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*QueueElement)(nil), "sql.QueueElement")
}

func init() { proto.RegisterFile("datasources/sql/sql.proto", fileDescriptor_ebfa9ffc66e09d2b) }

var fileDescriptor_ebfa9ffc66e09d2b = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0x2c, 0x49,
	0x2c, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0xd6, 0x2f, 0x2e, 0xcc, 0x01, 0x61, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0xe6, 0xe2, 0xc2, 0x1c, 0x29, 0xb1, 0xd4, 0x8a, 0xd4, 0xe4, 0xd2, 0x92,
	0xcc, 0xfc, 0x3c, 0xfd, 0xa2, 0xd4, 0xe4, 0xfc, 0xa2, 0x14, 0x88, 0xa4, 0x52, 0x3d, 0x17, 0x4f,
	0x60, 0x69, 0x6a, 0x69, 0xaa, 0x6b, 0x4e, 0x6a, 0x6e, 0x6a, 0x5e, 0x89, 0x90, 0x36, 0x17, 0x1b,
	0x44, 0x5e, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x50, 0x0f, 0xae, 0x51, 0x2f, 0x08, 0x2c,
	0xe1, 0xc1, 0x10, 0x04, 0x55, 0x22, 0xa4, 0xc4, 0xc5, 0x9d, 0x9a, 0x97, 0xe2, 0x9f, 0x16, 0x5c,
	0x52, 0x94, 0x9a, 0x98, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe1, 0xc1, 0x10, 0x84, 0x2c, 0x28,
	0x24, 0xc6, 0xc5, 0x9a, 0x5a, 0x54, 0x94, 0x5f, 0x24, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0xe9, 0xc1,
	0x10, 0x04, 0xe1, 0x3a, 0xb1, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x3a, 0xe9, 0x46, 0x69, 0xa7,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x97, 0x26, 0xa5, 0x1a, 0x19,
	0x19, 0x19, 0xe9, 0xe7, 0x27, 0x97, 0xe4, 0x83, 0xbc, 0x81, 0xe6, 0xad, 0x24, 0x36, 0xb0, 0xb3,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x0c, 0x3c, 0xed, 0xf0, 0x00, 0x00, 0x00,
}
