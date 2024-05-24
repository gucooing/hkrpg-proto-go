// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DevelopmentType.proto

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

type DevelopmentType int32

const (
	DevelopmentType_DEVELOPMENT_NONE             DevelopmentType = 0
	DevelopmentType_DEVELOPMENT_ROGUE_COSMOS     DevelopmentType = 1
	DevelopmentType_DEVELOPMENT_ROGUE_CHESS      DevelopmentType = 2
	DevelopmentType_DEVELOPMENT_ROGUE_CHESS_NOUS DevelopmentType = 3
	DevelopmentType_DEVELOPMENT_MEMORY_CHALLENGE DevelopmentType = 4
	DevelopmentType_DEVELOPMENT_STORY_CHALLENGE  DevelopmentType = 5
	DevelopmentType_DEVELOPMENT_UNLOCK_AVATAR    DevelopmentType = 6
	DevelopmentType_DEVELOPMENT_UNLOCK_EQUIPMENT DevelopmentType = 7
	DevelopmentType_DEVELOPMENT_ACTIVITY_START   DevelopmentType = 8
	DevelopmentType_DEVELOPMENT_ACTIVITY_END     DevelopmentType = 9
)

// Enum value maps for DevelopmentType.
var (
	DevelopmentType_name = map[int32]string{
		0: "DEVELOPMENT_NONE",
		1: "DEVELOPMENT_ROGUE_COSMOS",
		2: "DEVELOPMENT_ROGUE_CHESS",
		3: "DEVELOPMENT_ROGUE_CHESS_NOUS",
		4: "DEVELOPMENT_MEMORY_CHALLENGE",
		5: "DEVELOPMENT_STORY_CHALLENGE",
		6: "DEVELOPMENT_UNLOCK_AVATAR",
		7: "DEVELOPMENT_UNLOCK_EQUIPMENT",
		8: "DEVELOPMENT_ACTIVITY_START",
		9: "DEVELOPMENT_ACTIVITY_END",
	}
	DevelopmentType_value = map[string]int32{
		"DEVELOPMENT_NONE":             0,
		"DEVELOPMENT_ROGUE_COSMOS":     1,
		"DEVELOPMENT_ROGUE_CHESS":      2,
		"DEVELOPMENT_ROGUE_CHESS_NOUS": 3,
		"DEVELOPMENT_MEMORY_CHALLENGE": 4,
		"DEVELOPMENT_STORY_CHALLENGE":  5,
		"DEVELOPMENT_UNLOCK_AVATAR":    6,
		"DEVELOPMENT_UNLOCK_EQUIPMENT": 7,
		"DEVELOPMENT_ACTIVITY_START":   8,
		"DEVELOPMENT_ACTIVITY_END":     9,
	}
)

func (x DevelopmentType) Enum() *DevelopmentType {
	p := new(DevelopmentType)
	*p = x
	return p
}

func (x DevelopmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DevelopmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_DevelopmentType_proto_enumTypes[0].Descriptor()
}

func (DevelopmentType) Type() protoreflect.EnumType {
	return &file_DevelopmentType_proto_enumTypes[0]
}

func (x DevelopmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DevelopmentType.Descriptor instead.
func (DevelopmentType) EnumDescriptor() ([]byte, []int) {
	return file_DevelopmentType_proto_rawDescGZIP(), []int{0}
}

var File_DevelopmentType_proto protoreflect.FileDescriptor

var file_DevelopmentType_proto_rawDesc = []byte{
	0x0a, 0x15, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc6, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x44,
	0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x53, 0x4d, 0x4f, 0x53, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x4f, 0x47, 0x55,
	0x45, 0x5f, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x4e, 0x4f, 0x55, 0x53, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45,
	0x4d, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10, 0x04,
	0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4c, 0x4c, 0x45, 0x4e, 0x47, 0x45, 0x10,
	0x05, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x10, 0x06,
	0x12, 0x20, 0x0a, 0x1c, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x45, 0x51, 0x55, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x56, 0x45, 0x4c, 0x4f, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x09,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DevelopmentType_proto_rawDescOnce sync.Once
	file_DevelopmentType_proto_rawDescData = file_DevelopmentType_proto_rawDesc
)

func file_DevelopmentType_proto_rawDescGZIP() []byte {
	file_DevelopmentType_proto_rawDescOnce.Do(func() {
		file_DevelopmentType_proto_rawDescData = protoimpl.X.CompressGZIP(file_DevelopmentType_proto_rawDescData)
	})
	return file_DevelopmentType_proto_rawDescData
}

var file_DevelopmentType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DevelopmentType_proto_goTypes = []interface{}{
	(DevelopmentType)(0), // 0: DevelopmentType
}
var file_DevelopmentType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DevelopmentType_proto_init() }
func file_DevelopmentType_proto_init() {
	if File_DevelopmentType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DevelopmentType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DevelopmentType_proto_goTypes,
		DependencyIndexes: file_DevelopmentType_proto_depIdxs,
		EnumInfos:         file_DevelopmentType_proto_enumTypes,
	}.Build()
	File_DevelopmentType_proto = out.File
	file_DevelopmentType_proto_rawDesc = nil
	file_DevelopmentType_proto_goTypes = nil
	file_DevelopmentType_proto_depIdxs = nil
}
