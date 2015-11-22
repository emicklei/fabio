// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/base/api_base.proto
// DO NOT EDIT!

/*
Package base is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/base/api_base.proto

It has these top-level messages:
	StringProto
	Integer32Proto
	Integer64Proto
	BoolProto
	DoubleProto
	BytesProto
	VoidProto
*/
package base

import proto "github.com/eBay/fabio/_third_party/github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type StringProto struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StringProto) Reset()         { *m = StringProto{} }
func (m *StringProto) String() string { return proto.CompactTextString(m) }
func (*StringProto) ProtoMessage()    {}

func (m *StringProto) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type Integer32Proto struct {
	Value            *int32 `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Integer32Proto) Reset()         { *m = Integer32Proto{} }
func (m *Integer32Proto) String() string { return proto.CompactTextString(m) }
func (*Integer32Proto) ProtoMessage()    {}

func (m *Integer32Proto) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type Integer64Proto struct {
	Value            *int64 `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Integer64Proto) Reset()         { *m = Integer64Proto{} }
func (m *Integer64Proto) String() string { return proto.CompactTextString(m) }
func (*Integer64Proto) ProtoMessage()    {}

func (m *Integer64Proto) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type BoolProto struct {
	Value            *bool  `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BoolProto) Reset()         { *m = BoolProto{} }
func (m *BoolProto) String() string { return proto.CompactTextString(m) }
func (*BoolProto) ProtoMessage()    {}

func (m *BoolProto) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

type DoubleProto struct {
	Value            *float64 `protobuf:"fixed64,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *DoubleProto) Reset()         { *m = DoubleProto{} }
func (m *DoubleProto) String() string { return proto.CompactTextString(m) }
func (*DoubleProto) ProtoMessage()    {}

func (m *DoubleProto) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type BytesProto struct {
	Value            []byte `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BytesProto) Reset()         { *m = BytesProto{} }
func (m *BytesProto) String() string { return proto.CompactTextString(m) }
func (*BytesProto) ProtoMessage()    {}

func (m *BytesProto) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VoidProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *VoidProto) Reset()         { *m = VoidProto{} }
func (m *VoidProto) String() string { return proto.CompactTextString(m) }
func (*VoidProto) ProtoMessage()    {}

func init() {
}
