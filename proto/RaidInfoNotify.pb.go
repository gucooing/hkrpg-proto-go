// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RaidInfoNotify.proto

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

type RaidInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList    *ItemList      `protobuf:"bytes,3,opt,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	RaidId      uint32         `protobuf:"varint,11,opt,name=raid_id,json=raidId,proto3" json:"raid_id,omitempty"`
	Status      RaidStatus     `protobuf:"varint,4,opt,name=status,proto3,enum=RaidStatus" json:"status,omitempty"`
	FinishTime  uint64         `protobuf:"varint,9,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	WorldLevel  uint32         `protobuf:"varint,10,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	DAKHGPHAFPP []*CDIOHFOIDLM `protobuf:"bytes,13,rep,name=DAKHGPHAFPP,proto3" json:"DAKHGPHAFPP,omitempty"`
}

func (x *RaidInfoNotify) Reset() {
	*x = RaidInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RaidInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaidInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaidInfoNotify) ProtoMessage() {}

func (x *RaidInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RaidInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaidInfoNotify.ProtoReflect.Descriptor instead.
func (*RaidInfoNotify) Descriptor() ([]byte, []int) {
	return file_RaidInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RaidInfoNotify) GetItemList() *ItemList {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *RaidInfoNotify) GetRaidId() uint32 {
	if x != nil {
		return x.RaidId
	}
	return 0
}

func (x *RaidInfoNotify) GetStatus() RaidStatus {
	if x != nil {
		return x.Status
	}
	return RaidStatus_RAID_STATUS_NONE
}

func (x *RaidInfoNotify) GetFinishTime() uint64 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *RaidInfoNotify) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *RaidInfoNotify) GetDAKHGPHAFPP() []*CDIOHFOIDLM {
	if x != nil {
		return x.DAKHGPHAFPP
	}
	return nil
}

var File_RaidInfoNotify_proto protoreflect.FileDescriptor

var file_RaidInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x14, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x61, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x44, 0x49, 0x4f, 0x48, 0x46,
	0x4f, 0x49, 0x44, 0x4c, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0e,
	0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x26,
	0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0b, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x41, 0x4b, 0x48, 0x47, 0x50,
	0x48, 0x41, 0x46, 0x50, 0x50, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x44,
	0x49, 0x4f, 0x48, 0x46, 0x4f, 0x49, 0x44, 0x4c, 0x4d, 0x52, 0x0b, 0x44, 0x41, 0x4b, 0x48, 0x47,
	0x50, 0x48, 0x41, 0x46, 0x50, 0x50, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RaidInfoNotify_proto_rawDescOnce sync.Once
	file_RaidInfoNotify_proto_rawDescData = file_RaidInfoNotify_proto_rawDesc
)

func file_RaidInfoNotify_proto_rawDescGZIP() []byte {
	file_RaidInfoNotify_proto_rawDescOnce.Do(func() {
		file_RaidInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RaidInfoNotify_proto_rawDescData)
	})
	return file_RaidInfoNotify_proto_rawDescData
}

var file_RaidInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RaidInfoNotify_proto_goTypes = []interface{}{
	(*RaidInfoNotify)(nil), // 0: RaidInfoNotify
	(*ItemList)(nil),       // 1: ItemList
	(RaidStatus)(0),        // 2: RaidStatus
	(*CDIOHFOIDLM)(nil),    // 3: CDIOHFOIDLM
}
var file_RaidInfoNotify_proto_depIdxs = []int32{
	1, // 0: RaidInfoNotify.item_list:type_name -> ItemList
	2, // 1: RaidInfoNotify.status:type_name -> RaidStatus
	3, // 2: RaidInfoNotify.DAKHGPHAFPP:type_name -> CDIOHFOIDLM
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RaidInfoNotify_proto_init() }
func file_RaidInfoNotify_proto_init() {
	if File_RaidInfoNotify_proto != nil {
		return
	}
	file_RaidStatus_proto_init()
	file_CDIOHFOIDLM_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RaidInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaidInfoNotify); i {
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
			RawDescriptor: file_RaidInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RaidInfoNotify_proto_goTypes,
		DependencyIndexes: file_RaidInfoNotify_proto_depIdxs,
		MessageInfos:      file_RaidInfoNotify_proto_msgTypes,
	}.Build()
	File_RaidInfoNotify_proto = out.File
	file_RaidInfoNotify_proto_rawDesc = nil
	file_RaidInfoNotify_proto_goTypes = nil
	file_RaidInfoNotify_proto_depIdxs = nil
}