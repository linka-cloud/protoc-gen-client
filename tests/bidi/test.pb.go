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
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: tests/bidi/test.proto

package test

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_bidi_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tests_bidi_test_proto_msgTypes[0]
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
	return file_tests_bidi_test_proto_rawDescGZIP(), []int{0}
}

var File_tests_bidi_test_proto protoreflect.FileDescriptor

var file_tests_bidi_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x69, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x46, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x61, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_bidi_test_proto_rawDescOnce sync.Once
	file_tests_bidi_test_proto_rawDescData = file_tests_bidi_test_proto_rawDesc
)

func file_tests_bidi_test_proto_rawDescGZIP() []byte {
	file_tests_bidi_test_proto_rawDescOnce.Do(func() {
		file_tests_bidi_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_bidi_test_proto_rawDescData)
	})
	return file_tests_bidi_test_proto_rawDescData
}

var file_tests_bidi_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_bidi_test_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: go.client.test.Message
}
var file_tests_bidi_test_proto_depIdxs = []int32{
	0, // 0: go.client.test.Test.Stream:input_type -> go.client.test.Message
	0, // 1: go.client.test.Test.Stream:output_type -> go.client.test.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_bidi_test_proto_init() }
func file_tests_bidi_test_proto_init() {
	if File_tests_bidi_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_bidi_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_tests_bidi_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tests_bidi_test_proto_goTypes,
		DependencyIndexes: file_tests_bidi_test_proto_depIdxs,
		MessageInfos:      file_tests_bidi_test_proto_msgTypes,
	}.Build()
	File_tests_bidi_test_proto = out.File
	file_tests_bidi_test_proto_rawDesc = nil
	file_tests_bidi_test_proto_goTypes = nil
	file_tests_bidi_test_proto_depIdxs = nil
}
