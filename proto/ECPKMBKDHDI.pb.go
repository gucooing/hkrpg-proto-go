// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ECPKMBKDHDI.proto

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

type ECPKMBKDHDI int32

const (
	ECPKMBKDHDI_OFFERING_STATE_NONE ECPKMBKDHDI = 0
	ECPKMBKDHDI_OFFERING_STATE_LOCK ECPKMBKDHDI = 1
	ECPKMBKDHDI_OFFERING_STATE_OPEN ECPKMBKDHDI = 2
)

// Enum value maps for ECPKMBKDHDI.
var (
	ECPKMBKDHDI_name = map[int32]string{
		0: "OFFERING_STATE_NONE",
		1: "OFFERING_STATE_LOCK",
		2: "OFFERING_STATE_OPEN",
	}
	ECPKMBKDHDI_value = map[string]int32{
		"OFFERING_STATE_NONE": 0,
		"OFFERING_STATE_LOCK": 1,
		"OFFERING_STATE_OPEN": 2,
	}
)

func (x ECPKMBKDHDI) Enum() *ECPKMBKDHDI {
	p := new(ECPKMBKDHDI)
	*p = x
	return p
}

func (x ECPKMBKDHDI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECPKMBKDHDI) Descriptor() protoreflect.EnumDescriptor {
	return file_ECPKMBKDHDI_proto_enumTypes[0].Descriptor()
}

func (ECPKMBKDHDI) Type() protoreflect.EnumType {
	return &file_ECPKMBKDHDI_proto_enumTypes[0]
}

func (x ECPKMBKDHDI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ECPKMBKDHDI.Descriptor instead.
func (ECPKMBKDHDI) EnumDescriptor() ([]byte, []int) {
	return file_ECPKMBKDHDI_proto_rawDescGZIP(), []int{0}
}

var File_ECPKMBKDHDI_proto protoreflect.FileDescriptor

var file_ECPKMBKDHDI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x43, 0x50, 0x4b, 0x4d, 0x42, 0x4b, 0x44, 0x48, 0x44, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x58, 0x0a, 0x0b, 0x45, 0x43, 0x50, 0x4b, 0x4d, 0x42, 0x4b, 0x44, 0x48,
	0x44, 0x49, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f,
	0x46, 0x46, 0x45, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x46, 0x46, 0x45, 0x52, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x02, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ECPKMBKDHDI_proto_rawDescOnce sync.Once
	file_ECPKMBKDHDI_proto_rawDescData = file_ECPKMBKDHDI_proto_rawDesc
)

func file_ECPKMBKDHDI_proto_rawDescGZIP() []byte {
	file_ECPKMBKDHDI_proto_rawDescOnce.Do(func() {
		file_ECPKMBKDHDI_proto_rawDescData = protoimpl.X.CompressGZIP(file_ECPKMBKDHDI_proto_rawDescData)
	})
	return file_ECPKMBKDHDI_proto_rawDescData
}

var file_ECPKMBKDHDI_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ECPKMBKDHDI_proto_goTypes = []interface{}{
	(ECPKMBKDHDI)(0), // 0: ECPKMBKDHDI
}
var file_ECPKMBKDHDI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ECPKMBKDHDI_proto_init() }
func file_ECPKMBKDHDI_proto_init() {
	if File_ECPKMBKDHDI_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ECPKMBKDHDI_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ECPKMBKDHDI_proto_goTypes,
		DependencyIndexes: file_ECPKMBKDHDI_proto_depIdxs,
		EnumInfos:         file_ECPKMBKDHDI_proto_enumTypes,
	}.Build()
	File_ECPKMBKDHDI_proto = out.File
	file_ECPKMBKDHDI_proto_rawDesc = nil
	file_ECPKMBKDHDI_proto_goTypes = nil
	file_ECPKMBKDHDI_proto_depIdxs = nil
}
