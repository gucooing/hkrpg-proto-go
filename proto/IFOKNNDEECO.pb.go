// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: IFOKNNDEECO.proto

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

type IFOKNNDEECO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GKEAKOLIHIL uint32         `protobuf:"varint,6,opt,name=GKEAKOLIHIL,proto3" json:"GKEAKOLIHIL,omitempty"`
	JODDIDDEKEM bool           `protobuf:"varint,14,opt,name=JODDIDDEKEM,proto3" json:"JODDIDDEKEM,omitempty"`
	KKHCJICIKCM uint32         `protobuf:"varint,1,opt,name=KKHCJICIKCM,proto3" json:"KKHCJICIKCM,omitempty"`
	FBJPEDFOADG uint32         `protobuf:"varint,8,opt,name=FBJPEDFOADG,proto3" json:"FBJPEDFOADG,omitempty"`
	ANJMKKMGELA uint32         `protobuf:"varint,15,opt,name=ANJMKKMGELA,proto3" json:"ANJMKKMGELA,omitempty"`
	ELFJKDBPEJJ bool           `protobuf:"varint,3,opt,name=ELFJKDBPEJJ,proto3" json:"ELFJKDBPEJJ,omitempty"`
	HOFALKJFBEI []*BIDFPLCIODP `protobuf:"bytes,9,rep,name=HOFALKJFBEI,proto3" json:"HOFALKJFBEI,omitempty"`
	NNHDHBMLHEG uint32         `protobuf:"varint,5,opt,name=NNHDHBMLHEG,proto3" json:"NNHDHBMLHEG,omitempty"`
}

func (x *IFOKNNDEECO) Reset() {
	*x = IFOKNNDEECO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IFOKNNDEECO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IFOKNNDEECO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IFOKNNDEECO) ProtoMessage() {}

func (x *IFOKNNDEECO) ProtoReflect() protoreflect.Message {
	mi := &file_IFOKNNDEECO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IFOKNNDEECO.ProtoReflect.Descriptor instead.
func (*IFOKNNDEECO) Descriptor() ([]byte, []int) {
	return file_IFOKNNDEECO_proto_rawDescGZIP(), []int{0}
}

func (x *IFOKNNDEECO) GetGKEAKOLIHIL() uint32 {
	if x != nil {
		return x.GKEAKOLIHIL
	}
	return 0
}

func (x *IFOKNNDEECO) GetJODDIDDEKEM() bool {
	if x != nil {
		return x.JODDIDDEKEM
	}
	return false
}

func (x *IFOKNNDEECO) GetKKHCJICIKCM() uint32 {
	if x != nil {
		return x.KKHCJICIKCM
	}
	return 0
}

func (x *IFOKNNDEECO) GetFBJPEDFOADG() uint32 {
	if x != nil {
		return x.FBJPEDFOADG
	}
	return 0
}

func (x *IFOKNNDEECO) GetANJMKKMGELA() uint32 {
	if x != nil {
		return x.ANJMKKMGELA
	}
	return 0
}

func (x *IFOKNNDEECO) GetELFJKDBPEJJ() bool {
	if x != nil {
		return x.ELFJKDBPEJJ
	}
	return false
}

func (x *IFOKNNDEECO) GetHOFALKJFBEI() []*BIDFPLCIODP {
	if x != nil {
		return x.HOFALKJFBEI
	}
	return nil
}

func (x *IFOKNNDEECO) GetNNHDHBMLHEG() uint32 {
	if x != nil {
		return x.NNHDHBMLHEG
	}
	return 0
}

var File_IFOKNNDEECO_proto protoreflect.FileDescriptor

var file_IFOKNNDEECO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x46, 0x4f, 0x4b, 0x4e, 0x4e, 0x44, 0x45, 0x45, 0x43, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x49, 0x44, 0x46, 0x50, 0x4c, 0x43, 0x49, 0x4f, 0x44, 0x50,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x49, 0x46, 0x4f, 0x4b, 0x4e,
	0x4e, 0x44, 0x45, 0x45, 0x43, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4b, 0x45, 0x41, 0x4b, 0x4f,
	0x4c, 0x49, 0x48, 0x49, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4b, 0x45,
	0x41, 0x4b, 0x4f, 0x4c, 0x49, 0x48, 0x49, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4f, 0x44, 0x44,
	0x49, 0x44, 0x44, 0x45, 0x4b, 0x45, 0x4d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a,
	0x4f, 0x44, 0x44, 0x49, 0x44, 0x44, 0x45, 0x4b, 0x45, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4b,
	0x48, 0x43, 0x4a, 0x49, 0x43, 0x49, 0x4b, 0x43, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x4b, 0x4b, 0x48, 0x43, 0x4a, 0x49, 0x43, 0x49, 0x4b, 0x43, 0x4d, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x42, 0x4a, 0x50, 0x45, 0x44, 0x46, 0x4f, 0x41, 0x44, 0x47, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x46, 0x42, 0x4a, 0x50, 0x45, 0x44, 0x46, 0x4f, 0x41, 0x44, 0x47, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x4e, 0x4a, 0x4d, 0x4b, 0x4b, 0x4d, 0x47, 0x45, 0x4c, 0x41, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4e, 0x4a, 0x4d, 0x4b, 0x4b, 0x4d, 0x47, 0x45, 0x4c, 0x41,
	0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4c, 0x46, 0x4a, 0x4b, 0x44, 0x42, 0x50, 0x45, 0x4a, 0x4a, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x45, 0x4c, 0x46, 0x4a, 0x4b, 0x44, 0x42, 0x50, 0x45,
	0x4a, 0x4a, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x4f, 0x46, 0x41, 0x4c, 0x4b, 0x4a, 0x46, 0x42, 0x45,
	0x49, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x49, 0x44, 0x46, 0x50, 0x4c,
	0x43, 0x49, 0x4f, 0x44, 0x50, 0x52, 0x0b, 0x48, 0x4f, 0x46, 0x41, 0x4c, 0x4b, 0x4a, 0x46, 0x42,
	0x45, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4e, 0x48, 0x44, 0x48, 0x42, 0x4d, 0x4c, 0x48, 0x45,
	0x47, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x4e, 0x48, 0x44, 0x48, 0x42, 0x4d,
	0x4c, 0x48, 0x45, 0x47, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IFOKNNDEECO_proto_rawDescOnce sync.Once
	file_IFOKNNDEECO_proto_rawDescData = file_IFOKNNDEECO_proto_rawDesc
)

func file_IFOKNNDEECO_proto_rawDescGZIP() []byte {
	file_IFOKNNDEECO_proto_rawDescOnce.Do(func() {
		file_IFOKNNDEECO_proto_rawDescData = protoimpl.X.CompressGZIP(file_IFOKNNDEECO_proto_rawDescData)
	})
	return file_IFOKNNDEECO_proto_rawDescData
}

var file_IFOKNNDEECO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_IFOKNNDEECO_proto_goTypes = []interface{}{
	(*IFOKNNDEECO)(nil), // 0: IFOKNNDEECO
	(*BIDFPLCIODP)(nil), // 1: BIDFPLCIODP
}
var file_IFOKNNDEECO_proto_depIdxs = []int32{
	1, // 0: IFOKNNDEECO.HOFALKJFBEI:type_name -> BIDFPLCIODP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_IFOKNNDEECO_proto_init() }
func file_IFOKNNDEECO_proto_init() {
	if File_IFOKNNDEECO_proto != nil {
		return
	}
	file_BIDFPLCIODP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_IFOKNNDEECO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IFOKNNDEECO); i {
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
			RawDescriptor: file_IFOKNNDEECO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IFOKNNDEECO_proto_goTypes,
		DependencyIndexes: file_IFOKNNDEECO_proto_depIdxs,
		MessageInfos:      file_IFOKNNDEECO_proto_msgTypes,
	}.Build()
	File_IFOKNNDEECO_proto = out.File
	file_IFOKNNDEECO_proto_rawDesc = nil
	file_IFOKNNDEECO_proto_goTypes = nil
	file_IFOKNNDEECO_proto_depIdxs = nil
}
