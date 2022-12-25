// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/services/post/savepost.proto

package post

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

type SavePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *SavePost `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
}

func (x *SavePostRequest) Reset() {
	*x = SavePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_post_savepost_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePostRequest) ProtoMessage() {}

func (x *SavePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_post_savepost_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePostRequest.ProtoReflect.Descriptor instead.
func (*SavePostRequest) Descriptor() ([]byte, []int) {
	return file_proto_services_post_savepost_proto_rawDescGZIP(), []int{0}
}

func (x *SavePostRequest) GetPost() *SavePost {
	if x != nil {
		return x.Post
	}
	return nil
}

type SavePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string   `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Tags         []string `protobuf:"bytes,2,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Draft        string   `protobuf:"bytes,3,opt,name=Draft,proto3" json:"Draft,omitempty"`
	Summary      string   `protobuf:"bytes,4,opt,name=Summary,proto3" json:"Summary,omitempty"`
	Overview     string   `protobuf:"bytes,5,opt,name=Overview,proto3" json:"Overview,omitempty"`
	Bibliography string   `protobuf:"bytes,6,opt,name=Bibliography,proto3" json:"Bibliography,omitempty"`
	Slug         string   `protobuf:"bytes,7,opt,name=Slug,proto3" json:"Slug,omitempty"`
}

func (x *SavePost) Reset() {
	*x = SavePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_services_post_savepost_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePost) ProtoMessage() {}

func (x *SavePost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_post_savepost_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePost.ProtoReflect.Descriptor instead.
func (*SavePost) Descriptor() ([]byte, []int) {
	return file_proto_services_post_savepost_proto_rawDescGZIP(), []int{1}
}

func (x *SavePost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SavePost) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SavePost) GetDraft() string {
	if x != nil {
		return x.Draft
	}
	return ""
}

func (x *SavePost) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *SavePost) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *SavePost) GetBibliography() string {
	if x != nil {
		return x.Bibliography
	}
	return ""
}

func (x *SavePost) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_proto_services_post_savepost_proto protoreflect.FileDescriptor

var file_proto_services_post_savepost_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x66,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x44, 0x72, 0x61, 0x66, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x69, 0x62, 0x6c, 0x69, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x42, 0x69, 0x62, 0x6c,
	0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6c, 0x75, 0x67,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x42, 0x2b, 0x5a, 0x29,
	0x66, 0x61, 0x6c, 0x74, 0x61, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_services_post_savepost_proto_rawDescOnce sync.Once
	file_proto_services_post_savepost_proto_rawDescData = file_proto_services_post_savepost_proto_rawDesc
)

func file_proto_services_post_savepost_proto_rawDescGZIP() []byte {
	file_proto_services_post_savepost_proto_rawDescOnce.Do(func() {
		file_proto_services_post_savepost_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_post_savepost_proto_rawDescData)
	})
	return file_proto_services_post_savepost_proto_rawDescData
}

var file_proto_services_post_savepost_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_services_post_savepost_proto_goTypes = []interface{}{
	(*SavePostRequest)(nil), // 0: post.SavePostRequest
	(*SavePost)(nil),        // 1: post.SavePost
}
var file_proto_services_post_savepost_proto_depIdxs = []int32{
	1, // 0: post.SavePostRequest.Post:type_name -> post.SavePost
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_services_post_savepost_proto_init() }
func file_proto_services_post_savepost_proto_init() {
	if File_proto_services_post_savepost_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_services_post_savepost_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePostRequest); i {
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
		file_proto_services_post_savepost_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePost); i {
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
			RawDescriptor: file_proto_services_post_savepost_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_services_post_savepost_proto_goTypes,
		DependencyIndexes: file_proto_services_post_savepost_proto_depIdxs,
		MessageInfos:      file_proto_services_post_savepost_proto_msgTypes,
	}.Build()
	File_proto_services_post_savepost_proto = out.File
	file_proto_services_post_savepost_proto_rawDesc = nil
	file_proto_services_post_savepost_proto_goTypes = nil
	file_proto_services_post_savepost_proto_depIdxs = nil
}