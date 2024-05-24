// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMissionStatusCsReq.proto

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

type GetMissionStatusCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainMissionIdList  []uint32 `protobuf:"varint,12,rep,packed,name=main_mission_id_list,json=mainMissionIdList,proto3" json:"main_mission_id_list,omitempty"`
	MissionEventIdList []uint32 `protobuf:"varint,9,rep,packed,name=mission_event_id_list,json=missionEventIdList,proto3" json:"mission_event_id_list,omitempty"`
	SubMissionIdList   []uint32 `protobuf:"varint,6,rep,packed,name=sub_mission_id_list,json=subMissionIdList,proto3" json:"sub_mission_id_list,omitempty"`
}

func (x *GetMissionStatusCsReq) Reset() {
	*x = GetMissionStatusCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMissionStatusCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionStatusCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionStatusCsReq) ProtoMessage() {}

func (x *GetMissionStatusCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetMissionStatusCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionStatusCsReq.ProtoReflect.Descriptor instead.
func (*GetMissionStatusCsReq) Descriptor() ([]byte, []int) {
	return file_GetMissionStatusCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *GetMissionStatusCsReq) GetMainMissionIdList() []uint32 {
	if x != nil {
		return x.MainMissionIdList
	}
	return nil
}

func (x *GetMissionStatusCsReq) GetMissionEventIdList() []uint32 {
	if x != nil {
		return x.MissionEventIdList
	}
	return nil
}

func (x *GetMissionStatusCsReq) GetSubMissionIdList() []uint32 {
	if x != nil {
		return x.SubMissionIdList
	}
	return nil
}

var File_GetMissionStatusCsReq_proto protoreflect.FileDescriptor

var file_GetMissionStatusCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x73,
	0x75, 0x62, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x75, 0x62, 0x4d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMissionStatusCsReq_proto_rawDescOnce sync.Once
	file_GetMissionStatusCsReq_proto_rawDescData = file_GetMissionStatusCsReq_proto_rawDesc
)

func file_GetMissionStatusCsReq_proto_rawDescGZIP() []byte {
	file_GetMissionStatusCsReq_proto_rawDescOnce.Do(func() {
		file_GetMissionStatusCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMissionStatusCsReq_proto_rawDescData)
	})
	return file_GetMissionStatusCsReq_proto_rawDescData
}

var file_GetMissionStatusCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMissionStatusCsReq_proto_goTypes = []interface{}{
	(*GetMissionStatusCsReq)(nil), // 0: GetMissionStatusCsReq
}
var file_GetMissionStatusCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetMissionStatusCsReq_proto_init() }
func file_GetMissionStatusCsReq_proto_init() {
	if File_GetMissionStatusCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetMissionStatusCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionStatusCsReq); i {
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
			RawDescriptor: file_GetMissionStatusCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMissionStatusCsReq_proto_goTypes,
		DependencyIndexes: file_GetMissionStatusCsReq_proto_depIdxs,
		MessageInfos:      file_GetMissionStatusCsReq_proto_msgTypes,
	}.Build()
	File_GetMissionStatusCsReq_proto = out.File
	file_GetMissionStatusCsReq_proto_rawDesc = nil
	file_GetMissionStatusCsReq_proto_goTypes = nil
	file_GetMissionStatusCsReq_proto_depIdxs = nil
}