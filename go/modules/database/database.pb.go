// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/modules/database/database.proto

package database

import (
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

type SendData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data *anypb.Any `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *SendData) Reset() {
	*x = SendData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_modules_database_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendData) ProtoMessage() {}

func (x *SendData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_modules_database_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendData.ProtoReflect.Descriptor instead.
func (*SendData) Descriptor() ([]byte, []int) {
	return file_proto_modules_database_database_proto_rawDescGZIP(), []int{0}
}

func (x *SendData) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SendData) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_proto_modules_database_database_proto protoreflect.FileDescriptor

var file_proto_modules_database_database_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x08,
	0x53, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x32, 0x5a, 0x30, 0x66, 0x61, 0x6c, 0x74, 0x61, 0x72, 0x2d, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x3b, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_modules_database_database_proto_rawDescOnce sync.Once
	file_proto_modules_database_database_proto_rawDescData = file_proto_modules_database_database_proto_rawDesc
)

func file_proto_modules_database_database_proto_rawDescGZIP() []byte {
	file_proto_modules_database_database_proto_rawDescOnce.Do(func() {
		file_proto_modules_database_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_modules_database_database_proto_rawDescData)
	})
	return file_proto_modules_database_database_proto_rawDescData
}

var file_proto_modules_database_database_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_modules_database_database_proto_goTypes = []interface{}{
	(*SendData)(nil),  // 0: database.SendData
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_proto_modules_database_database_proto_depIdxs = []int32{
	1, // 0: database.SendData.Data:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_modules_database_database_proto_init() }
func file_proto_modules_database_database_proto_init() {
	if File_proto_modules_database_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_modules_database_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendData); i {
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
			RawDescriptor: file_proto_modules_database_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_modules_database_database_proto_goTypes,
		DependencyIndexes: file_proto_modules_database_database_proto_depIdxs,
		MessageInfos:      file_proto_modules_database_database_proto_msgTypes,
	}.Build()
	File_proto_modules_database_database_proto = out.File
	file_proto_modules_database_database_proto_rawDesc = nil
	file_proto_modules_database_database_proto_goTypes = nil
	file_proto_modules_database_database_proto_depIdxs = nil
}
