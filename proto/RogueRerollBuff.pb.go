// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueRerollBuff.proto

package proto

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

type RogueRerollBuff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BuffSelectInfo *RogueCommonBuffSelectInfo `protobuf:"bytes,15,opt,name=buff_select_info,json=buffSelectInfo,proto3" json:"buff_select_info,omitempty"`
}

func (x *RogueRerollBuff) Reset() {
	*x = RogueRerollBuff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueRerollBuff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueRerollBuff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueRerollBuff) ProtoMessage() {}

func (x *RogueRerollBuff) ProtoReflect() protoreflect.Message {
	mi := &file_RogueRerollBuff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueRerollBuff.ProtoReflect.Descriptor instead.
func (*RogueRerollBuff) Descriptor() ([]byte, []int) {
	return file_RogueRerollBuff_proto_rawDescGZIP(), []int{0}
}

func (x *RogueRerollBuff) GetBuffSelectInfo() *RogueCommonBuffSelectInfo {
	if x != nil {
		return x.BuffSelectInfo
	}
	return nil
}

var File_RogueRerollBuff_proto protoreflect.FileDescriptor

var file_RogueRerollBuff_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x42, 0x75, 0x66,
	0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x52, 0x65, 0x72, 0x6f, 0x6c, 0x6c, 0x42, 0x75, 0x66, 0x66, 0x12, 0x44, 0x0a, 0x10, 0x62,
	0x75, 0x66, 0x66, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueRerollBuff_proto_rawDescOnce sync.Once
	file_RogueRerollBuff_proto_rawDescData = file_RogueRerollBuff_proto_rawDesc
)

func file_RogueRerollBuff_proto_rawDescGZIP() []byte {
	file_RogueRerollBuff_proto_rawDescOnce.Do(func() {
		file_RogueRerollBuff_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueRerollBuff_proto_rawDescData)
	})
	return file_RogueRerollBuff_proto_rawDescData
}

var file_RogueRerollBuff_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueRerollBuff_proto_goTypes = []interface{}{
	(*RogueRerollBuff)(nil),           // 0: RogueRerollBuff
	(*RogueCommonBuffSelectInfo)(nil), // 1: RogueCommonBuffSelectInfo
}
var file_RogueRerollBuff_proto_depIdxs = []int32{
	1, // 0: RogueRerollBuff.buff_select_info:type_name -> RogueCommonBuffSelectInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueRerollBuff_proto_init() }
func file_RogueRerollBuff_proto_init() {
	if File_RogueRerollBuff_proto != nil {
		return
	}
	file_RogueCommonBuffSelectInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueRerollBuff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueRerollBuff); i {
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
			RawDescriptor: file_RogueRerollBuff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueRerollBuff_proto_goTypes,
		DependencyIndexes: file_RogueRerollBuff_proto_depIdxs,
		MessageInfos:      file_RogueRerollBuff_proto_msgTypes,
	}.Build()
	File_RogueRerollBuff_proto = out.File
	file_RogueRerollBuff_proto_rawDesc = nil
	file_RogueRerollBuff_proto_goTypes = nil
	file_RogueRerollBuff_proto_depIdxs = nil
}
