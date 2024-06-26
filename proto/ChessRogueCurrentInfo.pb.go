// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueCurrentInfo.proto

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

type ChessRogueCurrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoryInfo       *ChessRogueNousStoryInfo     `protobuf:"bytes,10,opt,name=story_info,json=storyInfo,proto3" json:"story_info,omitempty"`
	RogueAvatarInfo *ChessRogueAvatarInfo        `protobuf:"bytes,9,opt,name=rogue_avatar_info,json=rogueAvatarInfo,proto3" json:"rogue_avatar_info,omitempty"`
	RogueVersionId  uint32                       `protobuf:"varint,6,opt,name=rogue_version_id,json=rogueVersionId,proto3" json:"rogue_version_id,omitempty"`
	VirtualItemInfo []*ChessRogueVirtualItemInfo `protobuf:"bytes,4,rep,name=virtual_item_info,json=virtualItemInfo,proto3" json:"virtual_item_info,omitempty"`
	BuffInfo        *ChessRogueBuffInfo          `protobuf:"bytes,12,opt,name=buff_info,json=buffInfo,proto3" json:"buff_info,omitempty"`
	PendingAction   *RogueCommonPendingAction    `protobuf:"bytes,11,opt,name=pending_action,json=pendingAction,proto3" json:"pending_action,omitempty"`
	NousValue       *ChessRogueNousValue         `protobuf:"bytes,1,opt,name=nous_value,json=nousValue,proto3" json:"nous_value,omitempty"`
	MiracleInfo     *ChessRogueMiracleInfo       `protobuf:"bytes,14,opt,name=miracle_info,json=miracleInfo,proto3" json:"miracle_info,omitempty"`
	DiceInfo        *ChessRogueNousDice          `protobuf:"bytes,15,opt,name=dice_info,json=diceInfo,proto3" json:"dice_info,omitempty"`
}

func (x *ChessRogueCurrentInfo) Reset() {
	*x = ChessRogueCurrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueCurrentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueCurrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueCurrentInfo) ProtoMessage() {}

func (x *ChessRogueCurrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueCurrentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueCurrentInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueCurrentInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueCurrentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueCurrentInfo) GetStoryInfo() *ChessRogueNousStoryInfo {
	if x != nil {
		return x.StoryInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueAvatarInfo() *ChessRogueAvatarInfo {
	if x != nil {
		return x.RogueAvatarInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetRogueVersionId() uint32 {
	if x != nil {
		return x.RogueVersionId
	}
	return 0
}

func (x *ChessRogueCurrentInfo) GetVirtualItemInfo() []*ChessRogueVirtualItemInfo {
	if x != nil {
		return x.VirtualItemInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetBuffInfo() *ChessRogueBuffInfo {
	if x != nil {
		return x.BuffInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetPendingAction() *RogueCommonPendingAction {
	if x != nil {
		return x.PendingAction
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetNousValue() *ChessRogueNousValue {
	if x != nil {
		return x.NousValue
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetMiracleInfo() *ChessRogueMiracleInfo {
	if x != nil {
		return x.MiracleInfo
	}
	return nil
}

func (x *ChessRogueCurrentInfo) GetDiceInfo() *ChessRogueNousDice {
	if x != nil {
		return x.DiceInfo
	}
	return nil
}

var File_ChessRogueCurrentInfo_proto protoreflect.FileDescriptor

var file_ChessRogueCurrentInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73,
	0x44, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x37, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4e, 0x6f, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x41, 0x0a, 0x11, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x46, 0x0a, 0x11, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a,
	0x09, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66,
	0x66, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x40, 0x0a, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x6e, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6e, 0x6f, 0x75,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x30, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4e, 0x6f, 0x75, 0x73, 0x44, 0x69, 0x63, 0x65, 0x52, 0x08, 0x64, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueCurrentInfo_proto_rawDescOnce sync.Once
	file_ChessRogueCurrentInfo_proto_rawDescData = file_ChessRogueCurrentInfo_proto_rawDesc
)

func file_ChessRogueCurrentInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueCurrentInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueCurrentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueCurrentInfo_proto_rawDescData)
	})
	return file_ChessRogueCurrentInfo_proto_rawDescData
}

var file_ChessRogueCurrentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueCurrentInfo_proto_goTypes = []interface{}{
	(*ChessRogueCurrentInfo)(nil),     // 0: ChessRogueCurrentInfo
	(*ChessRogueNousStoryInfo)(nil),   // 1: ChessRogueNousStoryInfo
	(*ChessRogueAvatarInfo)(nil),      // 2: ChessRogueAvatarInfo
	(*ChessRogueVirtualItemInfo)(nil), // 3: ChessRogueVirtualItemInfo
	(*ChessRogueBuffInfo)(nil),        // 4: ChessRogueBuffInfo
	(*RogueCommonPendingAction)(nil),  // 5: RogueCommonPendingAction
	(*ChessRogueNousValue)(nil),       // 6: ChessRogueNousValue
	(*ChessRogueMiracleInfo)(nil),     // 7: ChessRogueMiracleInfo
	(*ChessRogueNousDice)(nil),        // 8: ChessRogueNousDice
}
var file_ChessRogueCurrentInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueCurrentInfo.story_info:type_name -> ChessRogueNousStoryInfo
	2, // 1: ChessRogueCurrentInfo.rogue_avatar_info:type_name -> ChessRogueAvatarInfo
	3, // 2: ChessRogueCurrentInfo.virtual_item_info:type_name -> ChessRogueVirtualItemInfo
	4, // 3: ChessRogueCurrentInfo.buff_info:type_name -> ChessRogueBuffInfo
	5, // 4: ChessRogueCurrentInfo.pending_action:type_name -> RogueCommonPendingAction
	6, // 5: ChessRogueCurrentInfo.nous_value:type_name -> ChessRogueNousValue
	7, // 6: ChessRogueCurrentInfo.miracle_info:type_name -> ChessRogueMiracleInfo
	8, // 7: ChessRogueCurrentInfo.dice_info:type_name -> ChessRogueNousDice
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_ChessRogueCurrentInfo_proto_init() }
func file_ChessRogueCurrentInfo_proto_init() {
	if File_ChessRogueCurrentInfo_proto != nil {
		return
	}
	file_ChessRogueNousStoryInfo_proto_init()
	file_ChessRogueAvatarInfo_proto_init()
	file_ChessRogueVirtualItemInfo_proto_init()
	file_RogueCommonPendingAction_proto_init()
	file_ChessRogueBuffInfo_proto_init()
	file_ChessRogueMiracleInfo_proto_init()
	file_ChessRogueNousDice_proto_init()
	file_ChessRogueNousValue_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueCurrentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueCurrentInfo); i {
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
			RawDescriptor: file_ChessRogueCurrentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueCurrentInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueCurrentInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueCurrentInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueCurrentInfo_proto = out.File
	file_ChessRogueCurrentInfo_proto_rawDesc = nil
	file_ChessRogueCurrentInfo_proto_goTypes = nil
	file_ChessRogueCurrentInfo_proto_depIdxs = nil
}
