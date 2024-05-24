// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeInfo.proto

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

type ChallengeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId     uint32              `protobuf:"varint,14,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ScoreTwo        uint32              `protobuf:"varint,4,opt,name=score_two,json=scoreTwo,proto3" json:"score_two,omitempty"`
	Status          ChallengeStatus     `protobuf:"varint,6,opt,name=status,proto3,enum=ChallengeStatus" json:"status,omitempty"`
	ExtraLineupType ExtraLineupType     `protobuf:"varint,12,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	StoryInfo       *ChallengeStoryInfo `protobuf:"bytes,9,opt,name=story_info,json=storyInfo,proto3" json:"story_info,omitempty"`
	RoundCount      uint32              `protobuf:"varint,2,opt,name=round_count,json=roundCount,proto3" json:"round_count,omitempty"`
	Score           uint32              `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ChallengeInfo) Reset() {
	*x = ChallengeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeInfo) ProtoMessage() {}

func (x *ChallengeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeInfo.ProtoReflect.Descriptor instead.
func (*ChallengeInfo) Descriptor() ([]byte, []int) {
	return file_ChallengeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeInfo) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ChallengeInfo) GetScoreTwo() uint32 {
	if x != nil {
		return x.ScoreTwo
	}
	return 0
}

func (x *ChallengeInfo) GetStatus() ChallengeStatus {
	if x != nil {
		return x.Status
	}
	return ChallengeStatus_CHALLENGE_UNKNOWN
}

func (x *ChallengeInfo) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *ChallengeInfo) GetStoryInfo() *ChallengeStoryInfo {
	if x != nil {
		return x.StoryInfo
	}
	return nil
}

func (x *ChallengeInfo) GetRoundCount() uint32 {
	if x != nil {
		return x.RoundCount
	}
	return 0
}

func (x *ChallengeInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_ChallengeInfo_proto protoreflect.FileDescriptor

var file_ChallengeInfo_proto_rawDesc = []byte{
	0x0a, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02,
	0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x77, 0x6f, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x11, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e,
	0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeInfo_proto_rawDescOnce sync.Once
	file_ChallengeInfo_proto_rawDescData = file_ChallengeInfo_proto_rawDesc
)

func file_ChallengeInfo_proto_rawDescGZIP() []byte {
	file_ChallengeInfo_proto_rawDescOnce.Do(func() {
		file_ChallengeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeInfo_proto_rawDescData)
	})
	return file_ChallengeInfo_proto_rawDescData
}

var file_ChallengeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeInfo_proto_goTypes = []interface{}{
	(*ChallengeInfo)(nil),      // 0: ChallengeInfo
	(ChallengeStatus)(0),       // 1: ChallengeStatus
	(ExtraLineupType)(0),       // 2: ExtraLineupType
	(*ChallengeStoryInfo)(nil), // 3: ChallengeStoryInfo
}
var file_ChallengeInfo_proto_depIdxs = []int32{
	1, // 0: ChallengeInfo.status:type_name -> ChallengeStatus
	2, // 1: ChallengeInfo.extra_lineup_type:type_name -> ExtraLineupType
	3, // 2: ChallengeInfo.story_info:type_name -> ChallengeStoryInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChallengeInfo_proto_init() }
func file_ChallengeInfo_proto_init() {
	if File_ChallengeInfo_proto != nil {
		return
	}
	file_ExtraLineupType_proto_init()
	file_ChallengeStoryInfo_proto_init()
	file_ChallengeStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeInfo); i {
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
			RawDescriptor: file_ChallengeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeInfo_proto_goTypes,
		DependencyIndexes: file_ChallengeInfo_proto_depIdxs,
		MessageInfos:      file_ChallengeInfo_proto_msgTypes,
	}.Build()
	File_ChallengeInfo_proto = out.File
	file_ChallengeInfo_proto_rawDesc = nil
	file_ChallengeInfo_proto_goTypes = nil
	file_ChallengeInfo_proto_depIdxs = nil
}
