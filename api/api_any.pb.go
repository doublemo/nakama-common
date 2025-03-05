// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.27.3
// source: api_any.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 微服务请求结构
type AnyQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *AnyQuery) Reset() {
	*x = AnyQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyQuery) ProtoMessage() {}

func (x *AnyQuery) ProtoReflect() protoreflect.Message {
	mi := &file_api_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyQuery.ProtoReflect.Descriptor instead.
func (*AnyQuery) Descriptor() ([]byte, []int) {
	return file_api_any_proto_rawDescGZIP(), []int{0}
}

func (x *AnyQuery) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type AnyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid     string               `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	Name    string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Header  map[string]string    `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Query   map[string]*AnyQuery `protobuf:"bytes,4,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Context map[string]string    `protobuf:"bytes,5,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HttpKey string               `protobuf:"bytes,6,opt,name=http_key,json=httpKey,proto3" json:"http_key,omitempty"`
	// Types that are assignable to Body:
	//
	//	*AnyRequest_BytesContent
	//	*AnyRequest_StringContent
	Body isAnyRequest_Body `protobuf_oneof:"body"`
}

func (x *AnyRequest) Reset() {
	*x = AnyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_any_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyRequest) ProtoMessage() {}

func (x *AnyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_any_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyRequest.ProtoReflect.Descriptor instead.
func (*AnyRequest) Descriptor() ([]byte, []int) {
	return file_api_any_proto_rawDescGZIP(), []int{1}
}

func (x *AnyRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *AnyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnyRequest) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *AnyRequest) GetQuery() map[string]*AnyQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *AnyRequest) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *AnyRequest) GetHttpKey() string {
	if x != nil {
		return x.HttpKey
	}
	return ""
}

func (m *AnyRequest) GetBody() isAnyRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *AnyRequest) GetBytesContent() []byte {
	if x, ok := x.GetBody().(*AnyRequest_BytesContent); ok {
		return x.BytesContent
	}
	return nil
}

func (x *AnyRequest) GetStringContent() string {
	if x, ok := x.GetBody().(*AnyRequest_StringContent); ok {
		return x.StringContent
	}
	return ""
}

type isAnyRequest_Body interface {
	isAnyRequest_Body()
}

type AnyRequest_BytesContent struct {
	BytesContent []byte `protobuf:"bytes,8,opt,name=bytesContent,proto3,oneof"`
}

type AnyRequest_StringContent struct {
	StringContent string `protobuf:"bytes,9,opt,name=stringContent,proto3,oneof"`
}

func (*AnyRequest_BytesContent) isAnyRequest_Body() {}

func (*AnyRequest_StringContent) isAnyRequest_Body() {}

type AnyResponseWriter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to Body:
	//
	//	*AnyResponseWriter_BytesContent
	//	*AnyResponseWriter_StringContent
	Body isAnyResponseWriter_Body `protobuf_oneof:"body"`
}

func (x *AnyResponseWriter) Reset() {
	*x = AnyResponseWriter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_any_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyResponseWriter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyResponseWriter) ProtoMessage() {}

func (x *AnyResponseWriter) ProtoReflect() protoreflect.Message {
	mi := &file_api_any_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyResponseWriter.ProtoReflect.Descriptor instead.
func (*AnyResponseWriter) Descriptor() ([]byte, []int) {
	return file_api_any_proto_rawDescGZIP(), []int{2}
}

func (x *AnyResponseWriter) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (m *AnyResponseWriter) GetBody() isAnyResponseWriter_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *AnyResponseWriter) GetBytesContent() []byte {
	if x, ok := x.GetBody().(*AnyResponseWriter_BytesContent); ok {
		return x.BytesContent
	}
	return nil
}

func (x *AnyResponseWriter) GetStringContent() string {
	if x, ok := x.GetBody().(*AnyResponseWriter_StringContent); ok {
		return x.StringContent
	}
	return ""
}

type isAnyResponseWriter_Body interface {
	isAnyResponseWriter_Body()
}

type AnyResponseWriter_BytesContent struct {
	BytesContent []byte `protobuf:"bytes,2,opt,name=bytesContent,proto3,oneof"`
}

type AnyResponseWriter_StringContent struct {
	StringContent string `protobuf:"bytes,3,opt,name=stringContent,proto3,oneof"`
}

func (*AnyResponseWriter_BytesContent) isAnyResponseWriter_Body() {}

func (*AnyResponseWriter_StringContent) isAnyResponseWriter_Body() {}

var File_api_any_proto protoreflect.FileDescriptor

var file_api_any_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x20, 0x0a, 0x08, 0x41,
	0x6e, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9e, 0x04,
	0x0a, 0x0a, 0x41, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x37,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x74, 0x74, 0x70, 0x4b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a,
	0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4e, 0x0a, 0x0a, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x61, 0x6b, 0x61,
	0x6d, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6e, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0xe7,
	0x01, 0x0a, 0x11, 0x41, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0c, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0c, 0x62, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x61, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x72, 0x6f, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x6e, 0x61, 0x6b, 0x61, 0x6d,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x09, 0x4e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x41, 0x70, 0x69,
	0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x6d, 0x6f, 0x2f, 0x6e, 0x61, 0x6b, 0x61, 0x6d, 0x61, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0f, 0x4e, 0x61, 0x6b, 0x61,
	0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_any_proto_rawDescOnce sync.Once
	file_api_any_proto_rawDescData = file_api_any_proto_rawDesc
)

func file_api_any_proto_rawDescGZIP() []byte {
	file_api_any_proto_rawDescOnce.Do(func() {
		file_api_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_any_proto_rawDescData)
	})
	return file_api_any_proto_rawDescData
}

var file_api_any_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_any_proto_goTypes = []interface{}{
	(*AnyQuery)(nil),          // 0: nakama.api.AnyQuery
	(*AnyRequest)(nil),        // 1: nakama.api.AnyRequest
	(*AnyResponseWriter)(nil), // 2: nakama.api.AnyResponseWriter
	nil,                       // 3: nakama.api.AnyRequest.HeaderEntry
	nil,                       // 4: nakama.api.AnyRequest.QueryEntry
	nil,                       // 5: nakama.api.AnyRequest.ContextEntry
	nil,                       // 6: nakama.api.AnyResponseWriter.HeaderEntry
}
var file_api_any_proto_depIdxs = []int32{
	3, // 0: nakama.api.AnyRequest.header:type_name -> nakama.api.AnyRequest.HeaderEntry
	4, // 1: nakama.api.AnyRequest.query:type_name -> nakama.api.AnyRequest.QueryEntry
	5, // 2: nakama.api.AnyRequest.context:type_name -> nakama.api.AnyRequest.ContextEntry
	6, // 3: nakama.api.AnyResponseWriter.header:type_name -> nakama.api.AnyResponseWriter.HeaderEntry
	0, // 4: nakama.api.AnyRequest.QueryEntry.value:type_name -> nakama.api.AnyQuery
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_any_proto_init() }
func file_api_any_proto_init() {
	if File_api_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_any_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_any_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyResponseWriter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_any_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AnyRequest_BytesContent)(nil),
		(*AnyRequest_StringContent)(nil),
	}
	file_api_any_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AnyResponseWriter_BytesContent)(nil),
		(*AnyResponseWriter_StringContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_any_proto_goTypes,
		DependencyIndexes: file_api_any_proto_depIdxs,
		MessageInfos:      file_api_any_proto_msgTypes,
	}.Build()
	File_api_any_proto = out.File
	file_api_any_proto_rawDesc = nil
	file_api_any_proto_goTypes = nil
	file_api_any_proto_depIdxs = nil
}
