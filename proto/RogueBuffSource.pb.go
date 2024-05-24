// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueBuffSource.proto

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

type RogueBuffSource int32

const (
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_NONE                RogueBuffSource = 0
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_SELECT              RogueBuffSource = 1
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_ENHANCE             RogueBuffSource = 2
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_MIRACLE             RogueBuffSource = 3
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_DIALOGUE            RogueBuffSource = 4
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_BONUS               RogueBuffSource = 5
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_MAZE_SKILL          RogueBuffSource = 6
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_SHOP                RogueBuffSource = 7
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_LEVEL_MECHANISM     RogueBuffSource = 8
	RogueBuffSource_ROGUE_BUFF_SOURCE_TYPE_ENDLESS_LEVEL_START RogueBuffSource = 9
)

// Enum value maps for RogueBuffSource.
var (
	RogueBuffSource_name = map[int32]string{
		0: "ROGUE_BUFF_SOURCE_TYPE_NONE",
		1: "ROGUE_BUFF_SOURCE_TYPE_SELECT",
		2: "ROGUE_BUFF_SOURCE_TYPE_ENHANCE",
		3: "ROGUE_BUFF_SOURCE_TYPE_MIRACLE",
		4: "ROGUE_BUFF_SOURCE_TYPE_DIALOGUE",
		5: "ROGUE_BUFF_SOURCE_TYPE_BONUS",
		6: "ROGUE_BUFF_SOURCE_TYPE_MAZE_SKILL",
		7: "ROGUE_BUFF_SOURCE_TYPE_SHOP",
		8: "ROGUE_BUFF_SOURCE_TYPE_LEVEL_MECHANISM",
		9: "ROGUE_BUFF_SOURCE_TYPE_ENDLESS_LEVEL_START",
	}
	RogueBuffSource_value = map[string]int32{
		"ROGUE_BUFF_SOURCE_TYPE_NONE":                0,
		"ROGUE_BUFF_SOURCE_TYPE_SELECT":              1,
		"ROGUE_BUFF_SOURCE_TYPE_ENHANCE":             2,
		"ROGUE_BUFF_SOURCE_TYPE_MIRACLE":             3,
		"ROGUE_BUFF_SOURCE_TYPE_DIALOGUE":            4,
		"ROGUE_BUFF_SOURCE_TYPE_BONUS":               5,
		"ROGUE_BUFF_SOURCE_TYPE_MAZE_SKILL":          6,
		"ROGUE_BUFF_SOURCE_TYPE_SHOP":                7,
		"ROGUE_BUFF_SOURCE_TYPE_LEVEL_MECHANISM":     8,
		"ROGUE_BUFF_SOURCE_TYPE_ENDLESS_LEVEL_START": 9,
	}
)

func (x RogueBuffSource) Enum() *RogueBuffSource {
	p := new(RogueBuffSource)
	*p = x
	return p
}

func (x RogueBuffSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RogueBuffSource) Descriptor() protoreflect.EnumDescriptor {
	return file_RogueBuffSource_proto_enumTypes[0].Descriptor()
}

func (RogueBuffSource) Type() protoreflect.EnumType {
	return &file_RogueBuffSource_proto_enumTypes[0]
}

func (x RogueBuffSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RogueBuffSource.Descriptor instead.
func (RogueBuffSource) EnumDescriptor() ([]byte, []int) {
	return file_RogueBuffSource_proto_rawDescGZIP(), []int{0}
}

var File_RogueBuffSource_proto protoreflect.FileDescriptor

var file_RogueBuffSource_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x88, 0x03, 0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12,
	0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x48, 0x41, 0x4e, 0x43,
	0x45, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46,
	0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49,
	0x52, 0x41, 0x43, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47, 0x55, 0x45, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53, 0x10, 0x05, 0x12, 0x25,
	0x0a, 0x21, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x41, 0x5a, 0x45, 0x5f, 0x53, 0x4b,
	0x49, 0x4c, 0x4c, 0x10, 0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42,
	0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x48, 0x4f, 0x50, 0x10, 0x07, 0x12, 0x2a, 0x0a, 0x26, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f,
	0x42, 0x55, 0x46, 0x46, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x43, 0x48, 0x41, 0x4e, 0x49, 0x53, 0x4d,
	0x10, 0x08, 0x12, 0x2e, 0x0a, 0x2a, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x42, 0x55, 0x46, 0x46,
	0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x44,
	0x4c, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x09, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueBuffSource_proto_rawDescOnce sync.Once
	file_RogueBuffSource_proto_rawDescData = file_RogueBuffSource_proto_rawDesc
)

func file_RogueBuffSource_proto_rawDescGZIP() []byte {
	file_RogueBuffSource_proto_rawDescOnce.Do(func() {
		file_RogueBuffSource_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueBuffSource_proto_rawDescData)
	})
	return file_RogueBuffSource_proto_rawDescData
}

var file_RogueBuffSource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RogueBuffSource_proto_goTypes = []interface{}{
	(RogueBuffSource)(0), // 0: RogueBuffSource
}
var file_RogueBuffSource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueBuffSource_proto_init() }
func file_RogueBuffSource_proto_init() {
	if File_RogueBuffSource_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueBuffSource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueBuffSource_proto_goTypes,
		DependencyIndexes: file_RogueBuffSource_proto_depIdxs,
		EnumInfos:         file_RogueBuffSource_proto_enumTypes,
	}.Build()
	File_RogueBuffSource_proto = out.File
	file_RogueBuffSource_proto_rawDesc = nil
	file_RogueBuffSource_proto_goTypes = nil
	file_RogueBuffSource_proto_depIdxs = nil
}