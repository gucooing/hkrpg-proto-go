// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StartCocoonStageScRsp.proto

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

type StartCocoonStageScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleInfo   *SceneBattleInfo `protobuf:"bytes,3,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
	Retcode      uint32           `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PropEntityId uint32           `protobuf:"varint,12,opt,name=prop_entity_id,json=propEntityId,proto3" json:"prop_entity_id,omitempty"`
	CocoonId     uint32           `protobuf:"varint,9,opt,name=cocoon_id,json=cocoonId,proto3" json:"cocoon_id,omitempty"`
	Wave         uint32           `protobuf:"varint,15,opt,name=wave,proto3" json:"wave,omitempty"`
}

func (x *StartCocoonStageScRsp) Reset() {
	*x = StartCocoonStageScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_StartCocoonStageScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCocoonStageScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCocoonStageScRsp) ProtoMessage() {}

func (x *StartCocoonStageScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_StartCocoonStageScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCocoonStageScRsp.ProtoReflect.Descriptor instead.
func (*StartCocoonStageScRsp) Descriptor() ([]byte, []int) {
	return file_StartCocoonStageScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *StartCocoonStageScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

func (x *StartCocoonStageScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *StartCocoonStageScRsp) GetPropEntityId() uint32 {
	if x != nil {
		return x.PropEntityId
	}
	return 0
}

func (x *StartCocoonStageScRsp) GetCocoonId() uint32 {
	if x != nil {
		return x.CocoonId
	}
	return 0
}

func (x *StartCocoonStageScRsp) GetWave() uint32 {
	if x != nil {
		return x.Wave
	}
	return 0
}

var File_StartCocoonStageScRsp_proto protoreflect.FileDescriptor

var file_StartCocoonStageScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x63, 0x6f, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f,
	0x63, 0x6f, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x31,
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x70,
	0x72, 0x6f, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x63, 0x6f, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x63, 0x6f, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x77, 0x61, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x77, 0x61,
	0x76, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_StartCocoonStageScRsp_proto_rawDescOnce sync.Once
	file_StartCocoonStageScRsp_proto_rawDescData = file_StartCocoonStageScRsp_proto_rawDesc
)

func file_StartCocoonStageScRsp_proto_rawDescGZIP() []byte {
	file_StartCocoonStageScRsp_proto_rawDescOnce.Do(func() {
		file_StartCocoonStageScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_StartCocoonStageScRsp_proto_rawDescData)
	})
	return file_StartCocoonStageScRsp_proto_rawDescData
}

var file_StartCocoonStageScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_StartCocoonStageScRsp_proto_goTypes = []interface{}{
	(*StartCocoonStageScRsp)(nil), // 0: StartCocoonStageScRsp
	(*SceneBattleInfo)(nil),       // 1: SceneBattleInfo
}
var file_StartCocoonStageScRsp_proto_depIdxs = []int32{
	1, // 0: StartCocoonStageScRsp.battle_info:type_name -> SceneBattleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_StartCocoonStageScRsp_proto_init() }
func file_StartCocoonStageScRsp_proto_init() {
	if File_StartCocoonStageScRsp_proto != nil {
		return
	}
	file_SceneBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_StartCocoonStageScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCocoonStageScRsp); i {
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
			RawDescriptor: file_StartCocoonStageScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StartCocoonStageScRsp_proto_goTypes,
		DependencyIndexes: file_StartCocoonStageScRsp_proto_depIdxs,
		MessageInfos:      file_StartCocoonStageScRsp_proto_msgTypes,
	}.Build()
	File_StartCocoonStageScRsp_proto = out.File
	file_StartCocoonStageScRsp_proto_rawDesc = nil
	file_StartCocoonStageScRsp_proto_goTypes = nil
	file_StartCocoonStageScRsp_proto_depIdxs = nil
}
