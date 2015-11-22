// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/remote_api/remote_api.proto
// DO NOT EDIT!

/*
Package remote_api is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/remote_api/remote_api.proto

It has these top-level messages:
	Request
	ApplicationError
	RpcError
	Response
*/
package remote_api

import proto "github.com/eBay/fabio/_third_party/github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type RpcError_ErrorCode int32

const (
	RpcError_UNKNOWN             RpcError_ErrorCode = 0
	RpcError_CALL_NOT_FOUND      RpcError_ErrorCode = 1
	RpcError_PARSE_ERROR         RpcError_ErrorCode = 2
	RpcError_SECURITY_VIOLATION  RpcError_ErrorCode = 3
	RpcError_OVER_QUOTA          RpcError_ErrorCode = 4
	RpcError_REQUEST_TOO_LARGE   RpcError_ErrorCode = 5
	RpcError_CAPABILITY_DISABLED RpcError_ErrorCode = 6
	RpcError_FEATURE_DISABLED    RpcError_ErrorCode = 7
	RpcError_BAD_REQUEST         RpcError_ErrorCode = 8
	RpcError_RESPONSE_TOO_LARGE  RpcError_ErrorCode = 9
	RpcError_CANCELLED           RpcError_ErrorCode = 10
	RpcError_REPLAY_ERROR        RpcError_ErrorCode = 11
	RpcError_DEADLINE_EXCEEDED   RpcError_ErrorCode = 12
)

var RpcError_ErrorCode_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CALL_NOT_FOUND",
	2:  "PARSE_ERROR",
	3:  "SECURITY_VIOLATION",
	4:  "OVER_QUOTA",
	5:  "REQUEST_TOO_LARGE",
	6:  "CAPABILITY_DISABLED",
	7:  "FEATURE_DISABLED",
	8:  "BAD_REQUEST",
	9:  "RESPONSE_TOO_LARGE",
	10: "CANCELLED",
	11: "REPLAY_ERROR",
	12: "DEADLINE_EXCEEDED",
}
var RpcError_ErrorCode_value = map[string]int32{
	"UNKNOWN":             0,
	"CALL_NOT_FOUND":      1,
	"PARSE_ERROR":         2,
	"SECURITY_VIOLATION":  3,
	"OVER_QUOTA":          4,
	"REQUEST_TOO_LARGE":   5,
	"CAPABILITY_DISABLED": 6,
	"FEATURE_DISABLED":    7,
	"BAD_REQUEST":         8,
	"RESPONSE_TOO_LARGE":  9,
	"CANCELLED":           10,
	"REPLAY_ERROR":        11,
	"DEADLINE_EXCEEDED":   12,
}

func (x RpcError_ErrorCode) Enum() *RpcError_ErrorCode {
	p := new(RpcError_ErrorCode)
	*p = x
	return p
}
func (x RpcError_ErrorCode) String() string {
	return proto.EnumName(RpcError_ErrorCode_name, int32(x))
}
func (x *RpcError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcError_ErrorCode_value, data, "RpcError_ErrorCode")
	if err != nil {
		return err
	}
	*x = RpcError_ErrorCode(value)
	return nil
}

type Request struct {
	ServiceName      *string `protobuf:"bytes,2,req,name=service_name" json:"service_name,omitempty"`
	Method           *string `protobuf:"bytes,3,req,name=method" json:"method,omitempty"`
	Request          []byte  `protobuf:"bytes,4,req,name=request" json:"request,omitempty"`
	RequestId        *string `protobuf:"bytes,5,opt,name=request_id" json:"request_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetServiceName() string {
	if m != nil && m.ServiceName != nil {
		return *m.ServiceName
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *Request) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Request) GetRequestId() string {
	if m != nil && m.RequestId != nil {
		return *m.RequestId
	}
	return ""
}

type ApplicationError struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail           *string `protobuf:"bytes,2,req,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ApplicationError) Reset()         { *m = ApplicationError{} }
func (m *ApplicationError) String() string { return proto.CompactTextString(m) }
func (*ApplicationError) ProtoMessage()    {}

func (m *ApplicationError) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *ApplicationError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type RpcError struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	Detail           *string `protobuf:"bytes,2,opt,name=detail" json:"detail,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpcError) Reset()         { *m = RpcError{} }
func (m *RpcError) String() string { return proto.CompactTextString(m) }
func (*RpcError) ProtoMessage()    {}

func (m *RpcError) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *RpcError) GetDetail() string {
	if m != nil && m.Detail != nil {
		return *m.Detail
	}
	return ""
}

type Response struct {
	Response         []byte            `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	Exception        []byte            `protobuf:"bytes,2,opt,name=exception" json:"exception,omitempty"`
	ApplicationError *ApplicationError `protobuf:"bytes,3,opt,name=application_error" json:"application_error,omitempty"`
	JavaException    []byte            `protobuf:"bytes,4,opt,name=java_exception" json:"java_exception,omitempty"`
	RpcError         *RpcError         `protobuf:"bytes,5,opt,name=rpc_error" json:"rpc_error,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *Response) GetException() []byte {
	if m != nil {
		return m.Exception
	}
	return nil
}

func (m *Response) GetApplicationError() *ApplicationError {
	if m != nil {
		return m.ApplicationError
	}
	return nil
}

func (m *Response) GetJavaException() []byte {
	if m != nil {
		return m.JavaException
	}
	return nil
}

func (m *Response) GetRpcError() *RpcError {
	if m != nil {
		return m.RpcError
	}
	return nil
}

func init() {
}
