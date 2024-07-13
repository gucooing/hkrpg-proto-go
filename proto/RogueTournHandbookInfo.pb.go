// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournHandbookInfo.proto

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

type RogueTournHandbookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HandbookFormulaList    []uint32 `protobuf:"varint,7,rep,packed,name=handbook_formula_list,json=handbookFormulaList,proto3" json:"handbook_formula_list,omitempty"`
	HandbookBuffList       []uint32 `protobuf:"varint,9,rep,packed,name=handbook_buff_list,json=handbookBuffList,proto3" json:"handbook_buff_list,omitempty"`
	EOJECMKIABF            uint32   `protobuf:"varint,11,opt,name=EOJECMKIABF,proto3" json:"EOJECMKIABF,omitempty"`
	TakeHandbookRewardList []uint32 `protobuf:"varint,15,rep,packed,name=take_handbook_reward_list,json=takeHandbookRewardList,proto3" json:"take_handbook_reward_list,omitempty"`
	HandbookAvatarBaseList []uint32 `protobuf:"varint,2,rep,packed,name=handbook_avatar_base_list,json=handbookAvatarBaseList,proto3" json:"handbook_avatar_base_list,omitempty"`
	HandbookMiracleList    []uint32 `protobuf:"varint,8,rep,packed,name=handbook_miracle_list,json=handbookMiracleList,proto3" json:"handbook_miracle_list,omitempty"`
}

func (x *RogueTournHandbookInfo) Reset() {
	*x = RogueTournHandbookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournHandbookInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournHandbookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournHandbookInfo) ProtoMessage() {}

func (x *RogueTournHandbookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournHandbookInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournHandbookInfo.ProtoReflect.Descriptor instead.
func (*RogueTournHandbookInfo) Descriptor() ([]byte, []int) {
	return file_RogueTournHandbookInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournHandbookInfo) GetHandbookFormulaList() []uint32 {
	if x != nil {
		return x.HandbookFormulaList
	}
	return nil
}

func (x *RogueTournHandbookInfo) GetHandbookBuffList() []uint32 {
	if x != nil {
		return x.HandbookBuffList
	}
	return nil
}

func (x *RogueTournHandbookInfo) GetEOJECMKIABF() uint32 {
	if x != nil {
		return x.EOJECMKIABF
	}
	return 0
}

func (x *RogueTournHandbookInfo) GetTakeHandbookRewardList() []uint32 {
	if x != nil {
		return x.TakeHandbookRewardList
	}
	return nil
}

func (x *RogueTournHandbookInfo) GetHandbookAvatarBaseList() []uint32 {
	if x != nil {
		return x.HandbookAvatarBaseList
	}
	return nil
}

func (x *RogueTournHandbookInfo) GetHandbookMiracleList() []uint32 {
	if x != nil {
		return x.HandbookMiracleList
	}
	return nil
}

var File_RogueTournHandbookInfo_proto protoreflect.FileDescriptor

var file_RogueTournHandbookInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6,
	0x02, 0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x0a, 0x15, 0x68, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x13, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x6f, 0x6b, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x6f, 0x6b, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x4f, 0x4a, 0x45, 0x43, 0x4d, 0x4b, 0x49, 0x41, 0x42, 0x46, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x45, 0x4f, 0x4a, 0x45, 0x43, 0x4d, 0x4b, 0x49, 0x41, 0x42, 0x46, 0x12, 0x39, 0x0a,
	0x19, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x16, 0x74, 0x61, 0x6b, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x68, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x16, 0x68, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x61, 0x73, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x5f,
	0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x13, 0x68, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournHandbookInfo_proto_rawDescOnce sync.Once
	file_RogueTournHandbookInfo_proto_rawDescData = file_RogueTournHandbookInfo_proto_rawDesc
)

func file_RogueTournHandbookInfo_proto_rawDescGZIP() []byte {
	file_RogueTournHandbookInfo_proto_rawDescOnce.Do(func() {
		file_RogueTournHandbookInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournHandbookInfo_proto_rawDescData)
	})
	return file_RogueTournHandbookInfo_proto_rawDescData
}

var file_RogueTournHandbookInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournHandbookInfo_proto_goTypes = []interface{}{
	(*RogueTournHandbookInfo)(nil), // 0: RogueTournHandbookInfo
}
var file_RogueTournHandbookInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueTournHandbookInfo_proto_init() }
func file_RogueTournHandbookInfo_proto_init() {
	if File_RogueTournHandbookInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueTournHandbookInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournHandbookInfo); i {
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
			RawDescriptor: file_RogueTournHandbookInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournHandbookInfo_proto_goTypes,
		DependencyIndexes: file_RogueTournHandbookInfo_proto_depIdxs,
		MessageInfos:      file_RogueTournHandbookInfo_proto_msgTypes,
	}.Build()
	File_RogueTournHandbookInfo_proto = out.File
	file_RogueTournHandbookInfo_proto_rawDesc = nil
	file_RogueTournHandbookInfo_proto_goTypes = nil
	file_RogueTournHandbookInfo_proto_depIdxs = nil
}
