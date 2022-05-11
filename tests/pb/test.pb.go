// Copyright 2022 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tests/pb/test.proto

package test

import (
	_ "github.com/alta/protopatch/patch/gopb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnaryEmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnaryEmptyRequest) Reset() {
	*x = UnaryEmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryEmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryEmptyRequest) ProtoMessage() {}

func (x *UnaryEmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryEmptyRequest.ProtoReflect.Descriptor instead.
func (*UnaryEmptyRequest) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{0}
}

type UnaryEmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnaryEmptyResponse) Reset() {
	*x = UnaryEmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryEmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryEmptyResponse) ProtoMessage() {}

func (x *UnaryEmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryEmptyResponse.ProtoReflect.Descriptor instead.
func (*UnaryEmptyResponse) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{1}
}

type UnaryRequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnaryRequestParams) Reset() {
	*x = UnaryRequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryRequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequestParams) ProtoMessage() {}

func (x *UnaryRequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequestParams.ProtoReflect.Descriptor instead.
func (*UnaryRequestParams) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{2}
}

func (x *UnaryRequestParams) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type UnaryResponseParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnaryResponseParams) Reset() {
	*x = UnaryResponseParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryResponseParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryResponseParams) ProtoMessage() {}

func (x *UnaryResponseParams) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryResponseParams.ProtoReflect.Descriptor instead.
func (*UnaryResponseParams) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{3}
}

func (x *UnaryResponseParams) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type UnaryRequestParamsAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Any     *anypb.Any `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	String_ String     `protobuf:"bytes,2,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *UnaryRequestParamsAny) Reset() {
	*x = UnaryRequestParamsAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryRequestParamsAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRequestParamsAny) ProtoMessage() {}

func (x *UnaryRequestParamsAny) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRequestParamsAny.ProtoReflect.Descriptor instead.
func (*UnaryRequestParamsAny) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{4}
}

func (x *UnaryRequestParamsAny) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *UnaryRequestParamsAny) GetString_() String {
	if x != nil {
		return x.String_
	}
	return ""
}

type UnaryResponseParamsAny struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Any     *anypb.Any `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	String_ String     `protobuf:"bytes,2,opt,name=string,proto3" json:"string,omitempty"`
}

func (x *UnaryResponseParamsAny) Reset() {
	*x = UnaryResponseParamsAny{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryResponseParamsAny) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryResponseParamsAny) ProtoMessage() {}

func (x *UnaryResponseParamsAny) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryResponseParamsAny.ProtoReflect.Descriptor instead.
func (*UnaryResponseParamsAny) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{5}
}

func (x *UnaryResponseParamsAny) GetAny() *anypb.Any {
	if x != nil {
		return x.Any
	}
	return nil
}

func (x *UnaryResponseParamsAny) GetString_() String {
	if x != nil {
		return x.String_
	}
	return ""
}

type ClientStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientStreamRequest) Reset() {
	*x = ClientStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRequest) ProtoMessage() {}

func (x *ClientStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRequest.ProtoReflect.Descriptor instead.
func (*ClientStreamRequest) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{6}
}

type ClientStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientStreamResponse) Reset() {
	*x = ClientStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamResponse) ProtoMessage() {}

func (x *ClientStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamResponse.ProtoReflect.Descriptor instead.
func (*ClientStreamResponse) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{7}
}

type ServerStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerStreamRequest) Reset() {
	*x = ServerStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRequest) ProtoMessage() {}

func (x *ServerStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamRequest.ProtoReflect.Descriptor instead.
func (*ServerStreamRequest) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{8}
}

type ServerStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerStreamResponse) Reset() {
	*x = ServerStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamResponse) ProtoMessage() {}

func (x *ServerStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamResponse.ProtoReflect.Descriptor instead.
func (*ServerStreamResponse) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{9}
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{10}
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{11}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_pb_test_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tests_pb_test_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_tests_pb_test_proto_rawDescGZIP(), []int{12}
}

var File_tests_pb_test_proto protoreflect.FileDescriptor

var file_tests_pb_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x13, 0x0a, 0x11, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x12, 0x55,
	0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x29, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x40, 0x0a, 0x13,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x65,
	0x0a, 0x15, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x41, 0x6e, 0x79, 0x12, 0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12,
	0x24, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x66, 0x0a, 0x16, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x41, 0x6e, 0x79, 0x12,
	0x26, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x15, 0x0a,
	0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x0a, 0x13,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x09,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd5, 0x05, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x53, 0x0a, 0x0a, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x22, 0x2e,
	0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55,
	0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x58, 0x0a, 0x0e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x56, 0x0a, 0x0b, 0x55,
	0x6e, 0x61, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x23,
	0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x5f, 0x0a, 0x0e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x41, 0x6e, 0x79, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x41, 0x6e, 0x79, 0x1a, 0x26, 0x2e, 0x67,
	0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x41, 0x6e, 0x79, 0x12, 0x59, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x28, 0x01, 0x12,
	0x59, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_pb_test_proto_rawDescOnce sync.Once
	file_tests_pb_test_proto_rawDescData = file_tests_pb_test_proto_rawDesc
)

func file_tests_pb_test_proto_rawDescGZIP() []byte {
	file_tests_pb_test_proto_rawDescOnce.Do(func() {
		file_tests_pb_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_pb_test_proto_rawDescData)
	})
	return file_tests_pb_test_proto_rawDescData
}

var file_tests_pb_test_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tests_pb_test_proto_goTypes = []interface{}{
	(*UnaryEmptyRequest)(nil),      // 0: go.client.test.UnaryEmptyRequest
	(*UnaryEmptyResponse)(nil),     // 1: go.client.test.UnaryEmptyResponse
	(*UnaryRequestParams)(nil),     // 2: go.client.test.UnaryRequestParams
	(*UnaryResponseParams)(nil),    // 3: go.client.test.UnaryResponseParams
	(*UnaryRequestParamsAny)(nil),  // 4: go.client.test.UnaryRequestParamsAny
	(*UnaryResponseParamsAny)(nil), // 5: go.client.test.UnaryResponseParamsAny
	(*ClientStreamRequest)(nil),    // 6: go.client.test.ClientStreamRequest
	(*ClientStreamResponse)(nil),   // 7: go.client.test.ClientStreamResponse
	(*ServerStreamRequest)(nil),    // 8: go.client.test.ServerStreamRequest
	(*ServerStreamResponse)(nil),   // 9: go.client.test.ServerStreamResponse
	(*StreamRequest)(nil),          // 10: go.client.test.StreamRequest
	(*StreamResponse)(nil),         // 11: go.client.test.StreamResponse
	(*Message)(nil),                // 12: go.client.test.Message
	(*anypb.Any)(nil),              // 13: google.protobuf.Any
}
var file_tests_pb_test_proto_depIdxs = []int32{
	12, // 0: go.client.test.UnaryRequestParams.msg:type_name -> go.client.test.Message
	12, // 1: go.client.test.UnaryResponseParams.msg:type_name -> go.client.test.Message
	13, // 2: go.client.test.UnaryRequestParamsAny.any:type_name -> google.protobuf.Any
	13, // 3: go.client.test.UnaryResponseParamsAny.any:type_name -> google.protobuf.Any
	0,  // 4: go.client.test.Test.UnaryEmpty:input_type -> go.client.test.UnaryEmptyRequest
	2,  // 5: go.client.test.Test.UnaryReqParams:input_type -> go.client.test.UnaryRequestParams
	0,  // 6: go.client.test.Test.UnaryResParams:input_type -> go.client.test.UnaryEmptyRequest
	2,  // 7: go.client.test.Test.UnaryParams:input_type -> go.client.test.UnaryRequestParams
	4,  // 8: go.client.test.Test.UnaryParamsAny:input_type -> go.client.test.UnaryRequestParamsAny
	2,  // 9: go.client.test.Test.ClientStream:input_type -> go.client.test.UnaryRequestParams
	2,  // 10: go.client.test.Test.ServerStream:input_type -> go.client.test.UnaryRequestParams
	2,  // 11: go.client.test.Test.Stream:input_type -> go.client.test.UnaryRequestParams
	1,  // 12: go.client.test.Test.UnaryEmpty:output_type -> go.client.test.UnaryEmptyResponse
	1,  // 13: go.client.test.Test.UnaryReqParams:output_type -> go.client.test.UnaryEmptyResponse
	3,  // 14: go.client.test.Test.UnaryResParams:output_type -> go.client.test.UnaryResponseParams
	3,  // 15: go.client.test.Test.UnaryParams:output_type -> go.client.test.UnaryResponseParams
	5,  // 16: go.client.test.Test.UnaryParamsAny:output_type -> go.client.test.UnaryResponseParamsAny
	3,  // 17: go.client.test.Test.ClientStream:output_type -> go.client.test.UnaryResponseParams
	3,  // 18: go.client.test.Test.ServerStream:output_type -> go.client.test.UnaryResponseParams
	3,  // 19: go.client.test.Test.Stream:output_type -> go.client.test.UnaryResponseParams
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_tests_pb_test_proto_init() }
func file_tests_pb_test_proto_init() {
	if File_tests_pb_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_pb_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryEmptyRequest); i {
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
		file_tests_pb_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryEmptyResponse); i {
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
		file_tests_pb_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryRequestParams); i {
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
		file_tests_pb_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryResponseParams); i {
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
		file_tests_pb_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryRequestParamsAny); i {
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
		file_tests_pb_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryResponseParamsAny); i {
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
		file_tests_pb_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamRequest); i {
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
		file_tests_pb_test_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamResponse); i {
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
		file_tests_pb_test_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamRequest); i {
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
		file_tests_pb_test_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamResponse); i {
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
		file_tests_pb_test_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_tests_pb_test_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_tests_pb_test_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_pb_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tests_pb_test_proto_goTypes,
		DependencyIndexes: file_tests_pb_test_proto_depIdxs,
		MessageInfos:      file_tests_pb_test_proto_msgTypes,
	}.Build()
	File_tests_pb_test_proto = out.File
	file_tests_pb_test_proto_rawDesc = nil
	file_tests_pb_test_proto_goTypes = nil
	file_tests_pb_test_proto_depIdxs = nil
}
