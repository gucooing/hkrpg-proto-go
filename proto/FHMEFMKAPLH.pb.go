// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FHMEFMKAPLH.proto

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

type FHMEFMKAPLH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FNFLAEFPODM uint32            `protobuf:"varint,2,opt,name=FNFLAEFPODM,proto3" json:"FNFLAEFPODM,omitempty"`
	GOMIABAKOJK uint32            `protobuf:"varint,9,opt,name=GOMIABAKOJK,proto3" json:"GOMIABAKOJK,omitempty"`
	OINDDHCMHHO map[uint32]uint32 `protobuf:"bytes,13,rep,name=OINDDHCMHHO,proto3" json:"OINDDHCMHHO,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *FHMEFMKAPLH) Reset() {
	*x = FHMEFMKAPLH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FHMEFMKAPLH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FHMEFMKAPLH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FHMEFMKAPLH) ProtoMessage() {}

func (x *FHMEFMKAPLH) ProtoReflect() protoreflect.Message {
	mi := &file_FHMEFMKAPLH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FHMEFMKAPLH.ProtoReflect.Descriptor instead.
func (*FHMEFMKAPLH) Descriptor() ([]byte, []int) {
	return file_FHMEFMKAPLH_proto_rawDescGZIP(), []int{0}
}

func (x *FHMEFMKAPLH) GetFNFLAEFPODM() uint32 {
	if x != nil {
		return x.FNFLAEFPODM
	}
	return 0
}

func (x *FHMEFMKAPLH) GetGOMIABAKOJK() uint32 {
	if x != nil {
		return x.GOMIABAKOJK
	}
	return 0
}

func (x *FHMEFMKAPLH) GetOINDDHCMHHO() map[uint32]uint32 {
	if x != nil {
		return x.OINDDHCMHHO
	}
	return nil
}

var File_FHMEFMKAPLH_proto protoreflect.FileDescriptor

var file_FHMEFMKAPLH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x48, 0x4d, 0x45, 0x46, 0x4d, 0x4b, 0x41, 0x50, 0x4c, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x46, 0x48, 0x4d, 0x45, 0x46, 0x4d, 0x4b, 0x41,
	0x50, 0x4c, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4e, 0x46, 0x4c, 0x41, 0x45, 0x46, 0x50, 0x4f,
	0x44, 0x4d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x4e, 0x46, 0x4c, 0x41, 0x45,
	0x46, 0x50, 0x4f, 0x44, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4f, 0x4d, 0x49, 0x41, 0x42, 0x41,
	0x4b, 0x4f, 0x4a, 0x4b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4f, 0x4d, 0x49,
	0x41, 0x42, 0x41, 0x4b, 0x4f, 0x4a, 0x4b, 0x12, 0x3f, 0x0a, 0x0b, 0x4f, 0x49, 0x4e, 0x44, 0x44,
	0x48, 0x43, 0x4d, 0x48, 0x48, 0x4f, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46,
	0x48, 0x4d, 0x45, 0x46, 0x4d, 0x4b, 0x41, 0x50, 0x4c, 0x48, 0x2e, 0x4f, 0x49, 0x4e, 0x44, 0x44,
	0x48, 0x43, 0x4d, 0x48, 0x48, 0x4f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4f, 0x49, 0x4e,
	0x44, 0x44, 0x48, 0x43, 0x4d, 0x48, 0x48, 0x4f, 0x1a, 0x3e, 0x0a, 0x10, 0x4f, 0x49, 0x4e, 0x44,
	0x44, 0x48, 0x43, 0x4d, 0x48, 0x48, 0x4f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FHMEFMKAPLH_proto_rawDescOnce sync.Once
	file_FHMEFMKAPLH_proto_rawDescData = file_FHMEFMKAPLH_proto_rawDesc
)

func file_FHMEFMKAPLH_proto_rawDescGZIP() []byte {
	file_FHMEFMKAPLH_proto_rawDescOnce.Do(func() {
		file_FHMEFMKAPLH_proto_rawDescData = protoimpl.X.CompressGZIP(file_FHMEFMKAPLH_proto_rawDescData)
	})
	return file_FHMEFMKAPLH_proto_rawDescData
}

var file_FHMEFMKAPLH_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FHMEFMKAPLH_proto_goTypes = []interface{}{
	(*FHMEFMKAPLH)(nil), // 0: FHMEFMKAPLH
	nil,                 // 1: FHMEFMKAPLH.OINDDHCMHHOEntry
}
var file_FHMEFMKAPLH_proto_depIdxs = []int32{
	1, // 0: FHMEFMKAPLH.OINDDHCMHHO:type_name -> FHMEFMKAPLH.OINDDHCMHHOEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FHMEFMKAPLH_proto_init() }
func file_FHMEFMKAPLH_proto_init() {
	if File_FHMEFMKAPLH_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FHMEFMKAPLH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FHMEFMKAPLH); i {
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
			RawDescriptor: file_FHMEFMKAPLH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FHMEFMKAPLH_proto_goTypes,
		DependencyIndexes: file_FHMEFMKAPLH_proto_depIdxs,
		MessageInfos:      file_FHMEFMKAPLH_proto_msgTypes,
	}.Build()
	File_FHMEFMKAPLH_proto = out.File
	file_FHMEFMKAPLH_proto_rawDesc = nil
	file_FHMEFMKAPLH_proto_goTypes = nil
	file_FHMEFMKAPLH_proto_depIdxs = nil
}
