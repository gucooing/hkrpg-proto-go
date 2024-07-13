// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMapRotationDataScRsp.proto

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

type GetMapRotationDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IICFOONOLAI int32              `protobuf:"varint,14,opt,name=IICFOONOLAI,proto3" json:"IICFOONOLAI,omitempty"`
	HMCAFEJAPJK uint32             `protobuf:"varint,10,opt,name=HMCAFEJAPJK,proto3" json:"HMCAFEJAPJK,omitempty"`
	RotaterData []*RotaterData     `protobuf:"bytes,15,rep,name=rotater_data,json=rotaterData,proto3" json:"rotater_data,omitempty"`
	CLHMAFCHJAF bool               `protobuf:"varint,7,opt,name=CLHMAFCHJAF,proto3" json:"CLHMAFCHJAF,omitempty"`
	Retcode     uint32             `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EnergyInfo  *RotatorEnergyInfo `protobuf:"bytes,1,opt,name=energy_info,json=energyInfo,proto3" json:"energy_info,omitempty"`
	ChargerInfo []*ChargerInfo     `protobuf:"bytes,11,rep,name=charger_info,json=chargerInfo,proto3" json:"charger_info,omitempty"`
	MapInfo     *IJJHKDNFKMD       `protobuf:"bytes,4,opt,name=map_info,json=mapInfo,proto3" json:"map_info,omitempty"`
}

func (x *GetMapRotationDataScRsp) Reset() {
	*x = GetMapRotationDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMapRotationDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMapRotationDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMapRotationDataScRsp) ProtoMessage() {}

func (x *GetMapRotationDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMapRotationDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMapRotationDataScRsp.ProtoReflect.Descriptor instead.
func (*GetMapRotationDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetMapRotationDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMapRotationDataScRsp) GetIICFOONOLAI() int32 {
	if x != nil {
		return x.IICFOONOLAI
	}
	return 0
}

func (x *GetMapRotationDataScRsp) GetHMCAFEJAPJK() uint32 {
	if x != nil {
		return x.HMCAFEJAPJK
	}
	return 0
}

func (x *GetMapRotationDataScRsp) GetRotaterData() []*RotaterData {
	if x != nil {
		return x.RotaterData
	}
	return nil
}

func (x *GetMapRotationDataScRsp) GetCLHMAFCHJAF() bool {
	if x != nil {
		return x.CLHMAFCHJAF
	}
	return false
}

func (x *GetMapRotationDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMapRotationDataScRsp) GetEnergyInfo() *RotatorEnergyInfo {
	if x != nil {
		return x.EnergyInfo
	}
	return nil
}

func (x *GetMapRotationDataScRsp) GetChargerInfo() []*ChargerInfo {
	if x != nil {
		return x.ChargerInfo
	}
	return nil
}

func (x *GetMapRotationDataScRsp) GetMapInfo() *IJJHKDNFKMD {
	if x != nil {
		return x.MapInfo
	}
	return nil
}

var File_GetMapRotationDataScRsp_proto protoreflect.FileDescriptor

var file_GetMapRotationDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x49, 0x4a, 0x4a, 0x48, 0x4b, 0x44, 0x4e, 0x46, 0x4b, 0x4d, 0x44, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x49, 0x49, 0x43, 0x46, 0x4f, 0x4f, 0x4e, 0x4f, 0x4c, 0x41, 0x49, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x49, 0x49, 0x43, 0x46, 0x4f, 0x4f, 0x4e, 0x4f, 0x4c, 0x41, 0x49, 0x12,
	0x20, 0x0a, 0x0b, 0x48, 0x4d, 0x43, 0x41, 0x46, 0x45, 0x4a, 0x41, 0x50, 0x4a, 0x4b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4d, 0x43, 0x41, 0x46, 0x45, 0x4a, 0x41, 0x50, 0x4a,
	0x4b, 0x12, 0x2f, 0x0a, 0x0c, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4c, 0x48, 0x4d, 0x41, 0x46, 0x43, 0x48, 0x4a, 0x41,
	0x46, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x4c, 0x48, 0x4d, 0x41, 0x46, 0x43,
	0x48, 0x4a, 0x41, 0x46, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x33,
	0x0a, 0x0b, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x4a, 0x4a, 0x48, 0x4b, 0x44, 0x4e,
	0x46, 0x4b, 0x4d, 0x44, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMapRotationDataScRsp_proto_rawDescOnce sync.Once
	file_GetMapRotationDataScRsp_proto_rawDescData = file_GetMapRotationDataScRsp_proto_rawDesc
)

func file_GetMapRotationDataScRsp_proto_rawDescGZIP() []byte {
	file_GetMapRotationDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetMapRotationDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMapRotationDataScRsp_proto_rawDescData)
	})
	return file_GetMapRotationDataScRsp_proto_rawDescData
}

var file_GetMapRotationDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMapRotationDataScRsp_proto_goTypes = []interface{}{
	(*GetMapRotationDataScRsp)(nil), // 0: GetMapRotationDataScRsp
	(*RotaterData)(nil),             // 1: RotaterData
	(*RotatorEnergyInfo)(nil),       // 2: RotatorEnergyInfo
	(*ChargerInfo)(nil),             // 3: ChargerInfo
	(*IJJHKDNFKMD)(nil),             // 4: IJJHKDNFKMD
}
var file_GetMapRotationDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetMapRotationDataScRsp.rotater_data:type_name -> RotaterData
	2, // 1: GetMapRotationDataScRsp.energy_info:type_name -> RotatorEnergyInfo
	3, // 2: GetMapRotationDataScRsp.charger_info:type_name -> ChargerInfo
	4, // 3: GetMapRotationDataScRsp.map_info:type_name -> IJJHKDNFKMD
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GetMapRotationDataScRsp_proto_init() }
func file_GetMapRotationDataScRsp_proto_init() {
	if File_GetMapRotationDataScRsp_proto != nil {
		return
	}
	file_IJJHKDNFKMD_proto_init()
	file_RotaterData_proto_init()
	file_RotatorEnergyInfo_proto_init()
	file_ChargerInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMapRotationDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMapRotationDataScRsp); i {
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
			RawDescriptor: file_GetMapRotationDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMapRotationDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetMapRotationDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetMapRotationDataScRsp_proto_msgTypes,
	}.Build()
	File_GetMapRotationDataScRsp_proto = out.File
	file_GetMapRotationDataScRsp_proto_rawDesc = nil
	file_GetMapRotationDataScRsp_proto_goTypes = nil
	file_GetMapRotationDataScRsp_proto_depIdxs = nil
}