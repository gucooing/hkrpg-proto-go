// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GNAPKGLAENF.proto

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

type GNAPKGLAENF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EBJKHDAIKCK uint32         `protobuf:"varint,15,opt,name=EBJKHDAIKCK,proto3" json:"EBJKHDAIKCK,omitempty"`
	CMDFNKEOJBH CMDAHLBMKLC    `protobuf:"varint,7,opt,name=CMDFNKEOJBH,proto3,enum=CMDAHLBMKLC" json:"CMDFNKEOJBH,omitempty"`
	OKJMPIDJPEB []*GHBFMPHPLGC `protobuf:"bytes,5,rep,name=OKJMPIDJPEB,proto3" json:"OKJMPIDJPEB,omitempty"`
	OGHFKDANBNJ uint32         `protobuf:"varint,6,opt,name=OGHFKDANBNJ,proto3" json:"OGHFKDANBNJ,omitempty"`
	NFEBFNFCLAD NGKADOPKNMP    `protobuf:"varint,2,opt,name=NFEBFNFCLAD,proto3,enum=NGKADOPKNMP" json:"NFEBFNFCLAD,omitempty"`
}

func (x *GNAPKGLAENF) Reset() {
	*x = GNAPKGLAENF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GNAPKGLAENF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GNAPKGLAENF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GNAPKGLAENF) ProtoMessage() {}

func (x *GNAPKGLAENF) ProtoReflect() protoreflect.Message {
	mi := &file_GNAPKGLAENF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GNAPKGLAENF.ProtoReflect.Descriptor instead.
func (*GNAPKGLAENF) Descriptor() ([]byte, []int) {
	return file_GNAPKGLAENF_proto_rawDescGZIP(), []int{0}
}

func (x *GNAPKGLAENF) GetEBJKHDAIKCK() uint32 {
	if x != nil {
		return x.EBJKHDAIKCK
	}
	return 0
}

func (x *GNAPKGLAENF) GetCMDFNKEOJBH() CMDAHLBMKLC {
	if x != nil {
		return x.CMDFNKEOJBH
	}
	return CMDAHLBMKLC_PAGE_DESC_NONE
}

func (x *GNAPKGLAENF) GetOKJMPIDJPEB() []*GHBFMPHPLGC {
	if x != nil {
		return x.OKJMPIDJPEB
	}
	return nil
}

func (x *GNAPKGLAENF) GetOGHFKDANBNJ() uint32 {
	if x != nil {
		return x.OGHFKDANBNJ
	}
	return 0
}

func (x *GNAPKGLAENF) GetNFEBFNFCLAD() NGKADOPKNMP {
	if x != nil {
		return x.NFEBFNFCLAD
	}
	return NGKADOPKNMP_PAGE_NONE
}

var File_GNAPKGLAENF_proto protoreflect.FileDescriptor

var file_GNAPKGLAENF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x4e, 0x41, 0x50, 0x4b, 0x47, 0x4c, 0x41, 0x45, 0x4e, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x48, 0x42, 0x46, 0x4d, 0x50, 0x48, 0x50, 0x4c, 0x47, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4d, 0x44, 0x41, 0x48, 0x4c, 0x42, 0x4d,
	0x4b, 0x4c, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x47, 0x4b, 0x41, 0x44,
	0x4f, 0x50, 0x4b, 0x4e, 0x4d, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a,
	0x0b, 0x47, 0x4e, 0x41, 0x50, 0x4b, 0x47, 0x4c, 0x41, 0x45, 0x4e, 0x46, 0x12, 0x20, 0x0a, 0x0b,
	0x45, 0x42, 0x4a, 0x4b, 0x48, 0x44, 0x41, 0x49, 0x4b, 0x43, 0x4b, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x45, 0x42, 0x4a, 0x4b, 0x48, 0x44, 0x41, 0x49, 0x4b, 0x43, 0x4b, 0x12, 0x2e,
	0x0a, 0x0b, 0x43, 0x4d, 0x44, 0x46, 0x4e, 0x4b, 0x45, 0x4f, 0x4a, 0x42, 0x48, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x43, 0x4d, 0x44, 0x41, 0x48, 0x4c, 0x42, 0x4d, 0x4b, 0x4c,
	0x43, 0x52, 0x0b, 0x43, 0x4d, 0x44, 0x46, 0x4e, 0x4b, 0x45, 0x4f, 0x4a, 0x42, 0x48, 0x12, 0x2e,
	0x0a, 0x0b, 0x4f, 0x4b, 0x4a, 0x4d, 0x50, 0x49, 0x44, 0x4a, 0x50, 0x45, 0x42, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x48, 0x42, 0x46, 0x4d, 0x50, 0x48, 0x50, 0x4c, 0x47,
	0x43, 0x52, 0x0b, 0x4f, 0x4b, 0x4a, 0x4d, 0x50, 0x49, 0x44, 0x4a, 0x50, 0x45, 0x42, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x47, 0x48, 0x46, 0x4b, 0x44, 0x41, 0x4e, 0x42, 0x4e, 0x4a, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x47, 0x48, 0x46, 0x4b, 0x44, 0x41, 0x4e, 0x42, 0x4e, 0x4a,
	0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x46, 0x45, 0x42, 0x46, 0x4e, 0x46, 0x43, 0x4c, 0x41, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x4e, 0x47, 0x4b, 0x41, 0x44, 0x4f, 0x50, 0x4b,
	0x4e, 0x4d, 0x50, 0x52, 0x0b, 0x4e, 0x46, 0x45, 0x42, 0x46, 0x4e, 0x46, 0x43, 0x4c, 0x41, 0x44,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GNAPKGLAENF_proto_rawDescOnce sync.Once
	file_GNAPKGLAENF_proto_rawDescData = file_GNAPKGLAENF_proto_rawDesc
)

func file_GNAPKGLAENF_proto_rawDescGZIP() []byte {
	file_GNAPKGLAENF_proto_rawDescOnce.Do(func() {
		file_GNAPKGLAENF_proto_rawDescData = protoimpl.X.CompressGZIP(file_GNAPKGLAENF_proto_rawDescData)
	})
	return file_GNAPKGLAENF_proto_rawDescData
}

var file_GNAPKGLAENF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GNAPKGLAENF_proto_goTypes = []interface{}{
	(*GNAPKGLAENF)(nil), // 0: GNAPKGLAENF
	(CMDAHLBMKLC)(0),    // 1: CMDAHLBMKLC
	(*GHBFMPHPLGC)(nil), // 2: GHBFMPHPLGC
	(NGKADOPKNMP)(0),    // 3: NGKADOPKNMP
}
var file_GNAPKGLAENF_proto_depIdxs = []int32{
	1, // 0: GNAPKGLAENF.CMDFNKEOJBH:type_name -> CMDAHLBMKLC
	2, // 1: GNAPKGLAENF.OKJMPIDJPEB:type_name -> GHBFMPHPLGC
	3, // 2: GNAPKGLAENF.NFEBFNFCLAD:type_name -> NGKADOPKNMP
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GNAPKGLAENF_proto_init() }
func file_GNAPKGLAENF_proto_init() {
	if File_GNAPKGLAENF_proto != nil {
		return
	}
	file_GHBFMPHPLGC_proto_init()
	file_CMDAHLBMKLC_proto_init()
	file_NGKADOPKNMP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GNAPKGLAENF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GNAPKGLAENF); i {
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
			RawDescriptor: file_GNAPKGLAENF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GNAPKGLAENF_proto_goTypes,
		DependencyIndexes: file_GNAPKGLAENF_proto_depIdxs,
		MessageInfos:      file_GNAPKGLAENF_proto_msgTypes,
	}.Build()
	File_GNAPKGLAENF_proto = out.File
	file_GNAPKGLAENF_proto_rawDesc = nil
	file_GNAPKGLAENF_proto_goTypes = nil
	file_GNAPKGLAENF_proto_depIdxs = nil
}
