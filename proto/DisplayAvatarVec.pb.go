// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DisplayAvatarVec.proto

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

type DisplayAvatarVec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayAvatarList []*DisplayAvatarData `protobuf:"bytes,11,rep,name=display_avatar_list,json=displayAvatarList,proto3" json:"display_avatar_list,omitempty"`
	IsDisplay         bool                 `protobuf:"varint,12,opt,name=is_display,json=isDisplay,proto3" json:"is_display,omitempty"`
}

func (x *DisplayAvatarVec) Reset() {
	*x = DisplayAvatarVec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DisplayAvatarVec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisplayAvatarVec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisplayAvatarVec) ProtoMessage() {}

func (x *DisplayAvatarVec) ProtoReflect() protoreflect.Message {
	mi := &file_DisplayAvatarVec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisplayAvatarVec.ProtoReflect.Descriptor instead.
func (*DisplayAvatarVec) Descriptor() ([]byte, []int) {
	return file_DisplayAvatarVec_proto_rawDescGZIP(), []int{0}
}

func (x *DisplayAvatarVec) GetDisplayAvatarList() []*DisplayAvatarData {
	if x != nil {
		return x.DisplayAvatarList
	}
	return nil
}

func (x *DisplayAvatarVec) GetIsDisplay() bool {
	if x != nil {
		return x.IsDisplay
	}
	return false
}

var File_DisplayAvatarVec_proto protoreflect.FileDescriptor

var file_DisplayAvatarVec_proto_rawDesc = []byte{
	0x0a, 0x16, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x56,
	0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x75, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x56, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x11, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DisplayAvatarVec_proto_rawDescOnce sync.Once
	file_DisplayAvatarVec_proto_rawDescData = file_DisplayAvatarVec_proto_rawDesc
)

func file_DisplayAvatarVec_proto_rawDescGZIP() []byte {
	file_DisplayAvatarVec_proto_rawDescOnce.Do(func() {
		file_DisplayAvatarVec_proto_rawDescData = protoimpl.X.CompressGZIP(file_DisplayAvatarVec_proto_rawDescData)
	})
	return file_DisplayAvatarVec_proto_rawDescData
}

var file_DisplayAvatarVec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DisplayAvatarVec_proto_goTypes = []interface{}{
	(*DisplayAvatarVec)(nil),  // 0: DisplayAvatarVec
	(*DisplayAvatarData)(nil), // 1: DisplayAvatarData
}
var file_DisplayAvatarVec_proto_depIdxs = []int32{
	1, // 0: DisplayAvatarVec.display_avatar_list:type_name -> DisplayAvatarData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DisplayAvatarVec_proto_init() }
func file_DisplayAvatarVec_proto_init() {
	if File_DisplayAvatarVec_proto != nil {
		return
	}
	file_DisplayAvatarData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DisplayAvatarVec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisplayAvatarVec); i {
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
			RawDescriptor: file_DisplayAvatarVec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DisplayAvatarVec_proto_goTypes,
		DependencyIndexes: file_DisplayAvatarVec_proto_depIdxs,
		MessageInfos:      file_DisplayAvatarVec_proto_msgTypes,
	}.Build()
	File_DisplayAvatarVec_proto = out.File
	file_DisplayAvatarVec_proto_rawDesc = nil
	file_DisplayAvatarVec_proto_goTypes = nil
	file_DisplayAvatarVec_proto_depIdxs = nil
}
