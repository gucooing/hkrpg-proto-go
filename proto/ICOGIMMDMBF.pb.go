// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ICOGIMMDMBF.proto

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

type ICOGIMMDMBF int32

const (
	ICOGIMMDMBF_PLAYER_RETURN_NONE       ICOGIMMDMBF = 0
	ICOGIMMDMBF_PLAYER_RETURN_PROCESSING ICOGIMMDMBF = 1
	ICOGIMMDMBF_PLAYER_RETURN_FINISH     ICOGIMMDMBF = 2
)

// Enum value maps for ICOGIMMDMBF.
var (
	ICOGIMMDMBF_name = map[int32]string{
		0: "PLAYER_RETURN_NONE",
		1: "PLAYER_RETURN_PROCESSING",
		2: "PLAYER_RETURN_FINISH",
	}
	ICOGIMMDMBF_value = map[string]int32{
		"PLAYER_RETURN_NONE":       0,
		"PLAYER_RETURN_PROCESSING": 1,
		"PLAYER_RETURN_FINISH":     2,
	}
)

func (x ICOGIMMDMBF) Enum() *ICOGIMMDMBF {
	p := new(ICOGIMMDMBF)
	*p = x
	return p
}

func (x ICOGIMMDMBF) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ICOGIMMDMBF) Descriptor() protoreflect.EnumDescriptor {
	return file_ICOGIMMDMBF_proto_enumTypes[0].Descriptor()
}

func (ICOGIMMDMBF) Type() protoreflect.EnumType {
	return &file_ICOGIMMDMBF_proto_enumTypes[0]
}

func (x ICOGIMMDMBF) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ICOGIMMDMBF.Descriptor instead.
func (ICOGIMMDMBF) EnumDescriptor() ([]byte, []int) {
	return file_ICOGIMMDMBF_proto_rawDescGZIP(), []int{0}
}

var File_ICOGIMMDMBF_proto protoreflect.FileDescriptor

var file_ICOGIMMDMBF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x43, 0x4f, 0x47, 0x49, 0x4d, 0x4d, 0x44, 0x4d, 0x42, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x5d, 0x0a, 0x0b, 0x49, 0x43, 0x4f, 0x47, 0x49, 0x4d, 0x4d, 0x44, 0x4d,
	0x42, 0x46, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x54,
	0x55, 0x52, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c,
	0x41, 0x59, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43,
	0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x10, 0x02, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ICOGIMMDMBF_proto_rawDescOnce sync.Once
	file_ICOGIMMDMBF_proto_rawDescData = file_ICOGIMMDMBF_proto_rawDesc
)

func file_ICOGIMMDMBF_proto_rawDescGZIP() []byte {
	file_ICOGIMMDMBF_proto_rawDescOnce.Do(func() {
		file_ICOGIMMDMBF_proto_rawDescData = protoimpl.X.CompressGZIP(file_ICOGIMMDMBF_proto_rawDescData)
	})
	return file_ICOGIMMDMBF_proto_rawDescData
}

var file_ICOGIMMDMBF_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ICOGIMMDMBF_proto_goTypes = []interface{}{
	(ICOGIMMDMBF)(0), // 0: ICOGIMMDMBF
}
var file_ICOGIMMDMBF_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ICOGIMMDMBF_proto_init() }
func file_ICOGIMMDMBF_proto_init() {
	if File_ICOGIMMDMBF_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ICOGIMMDMBF_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ICOGIMMDMBF_proto_goTypes,
		DependencyIndexes: file_ICOGIMMDMBF_proto_depIdxs,
		EnumInfos:         file_ICOGIMMDMBF_proto_enumTypes,
	}.Build()
	File_ICOGIMMDMBF_proto = out.File
	file_ICOGIMMDMBF_proto_rawDesc = nil
	file_ICOGIMMDMBF_proto_goTypes = nil
	file_ICOGIMMDMBF_proto_depIdxs = nil
}
