// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartChallengeStoryBuffInfo.proto

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

type StartChallengeStoryBuffInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoryBuffOne uint32 `protobuf:"varint,10,opt,name=story_buff_one,json=storyBuffOne,proto3" json:"story_buff_one,omitempty"`
	StoryBuffTwo uint32 `protobuf:"varint,7,opt,name=story_buff_two,json=storyBuffTwo,proto3" json:"story_buff_two,omitempty"`
}

func (x *StartChallengeStoryBuffInfo) Reset() {
	*x = StartChallengeStoryBuffInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartChallengeStoryBuffInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartChallengeStoryBuffInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartChallengeStoryBuffInfo) ProtoMessage() {}

func (x *StartChallengeStoryBuffInfo) ProtoReflect() protoreflect.Message {
	mi := &file_StartChallengeStoryBuffInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartChallengeStoryBuffInfo.ProtoReflect.Descriptor instead.
func (*StartChallengeStoryBuffInfo) Descriptor() ([]byte, []int) {
	return file_StartChallengeStoryBuffInfo_proto_rawDescGZIP(), []int{0}
}

func (x *StartChallengeStoryBuffInfo) GetStoryBuffOne() uint32 {
	if x != nil {
		return x.StoryBuffOne
	}
	return 0
}

func (x *StartChallengeStoryBuffInfo) GetStoryBuffTwo() uint32 {
	if x != nil {
		return x.StoryBuffTwo
	}
	return 0
}

var File_StartChallengeStoryBuffInfo_proto protoreflect.FileDescriptor

var file_StartChallengeStoryBuffInfo_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x1b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x75, 0x66, 0x66,
	0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x42, 0x75, 0x66, 0x66, 0x4f, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x75, 0x66, 0x66, 0x54, 0x77, 0x6f, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartChallengeStoryBuffInfo_proto_rawDescOnce sync.Once
	file_StartChallengeStoryBuffInfo_proto_rawDescData = file_StartChallengeStoryBuffInfo_proto_rawDesc
)

func file_StartChallengeStoryBuffInfo_proto_rawDescGZIP() []byte {
	file_StartChallengeStoryBuffInfo_proto_rawDescOnce.Do(func() {
		file_StartChallengeStoryBuffInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartChallengeStoryBuffInfo_proto_rawDescData)
	})
	return file_StartChallengeStoryBuffInfo_proto_rawDescData
}

var file_StartChallengeStoryBuffInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartChallengeStoryBuffInfo_proto_goTypes = []interface{}{
	(*StartChallengeStoryBuffInfo)(nil), // 0: StartChallengeStoryBuffInfo
}
var file_StartChallengeStoryBuffInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_StartChallengeStoryBuffInfo_proto_init() }
func file_StartChallengeStoryBuffInfo_proto_init() {
	if File_StartChallengeStoryBuffInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_StartChallengeStoryBuffInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartChallengeStoryBuffInfo); i {
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
			RawDescriptor: file_StartChallengeStoryBuffInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartChallengeStoryBuffInfo_proto_goTypes,
		DependencyIndexes: file_StartChallengeStoryBuffInfo_proto_depIdxs,
		MessageInfos:      file_StartChallengeStoryBuffInfo_proto_msgTypes,
	}.Build()
	File_StartChallengeStoryBuffInfo_proto = out.File
	file_StartChallengeStoryBuffInfo_proto_rawDesc = nil
	file_StartChallengeStoryBuffInfo_proto_goTypes = nil
	file_StartChallengeStoryBuffInfo_proto_depIdxs = nil
}
