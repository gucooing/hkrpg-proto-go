// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetFriendChallengeDetailScRsp.proto

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

type GetFriendChallengeDetailScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid         uint32                     `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Retcode     uint32                     `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ChallengeId uint32                     `protobuf:"varint,13,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	GCAONKNELPB []*DisplayAvatarDetailInfo `protobuf:"bytes,7,rep,name=GCAONKNELPB,proto3" json:"GCAONKNELPB,omitempty"`
}

func (x *GetFriendChallengeDetailScRsp) Reset() {
	*x = GetFriendChallengeDetailScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetFriendChallengeDetailScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFriendChallengeDetailScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFriendChallengeDetailScRsp) ProtoMessage() {}

func (x *GetFriendChallengeDetailScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetFriendChallengeDetailScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFriendChallengeDetailScRsp.ProtoReflect.Descriptor instead.
func (*GetFriendChallengeDetailScRsp) Descriptor() ([]byte, []int) {
	return file_GetFriendChallengeDetailScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetFriendChallengeDetailScRsp) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *GetFriendChallengeDetailScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetFriendChallengeDetailScRsp) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *GetFriendChallengeDetailScRsp) GetGCAONKNELPB() []*DisplayAvatarDetailInfo {
	if x != nil {
		return x.GCAONKNELPB
	}
	return nil
}

var File_GetFriendChallengeDetailScRsp_proto protoreflect.FileDescriptor

var file_GetFriendChallengeDetailScRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x65, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x43, 0x41, 0x4f, 0x4e, 0x4b, 0x4e,
	0x45, 0x4c, 0x50, 0x42, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x47, 0x43, 0x41, 0x4f, 0x4e, 0x4b, 0x4e, 0x45, 0x4c, 0x50,
	0x42, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetFriendChallengeDetailScRsp_proto_rawDescOnce sync.Once
	file_GetFriendChallengeDetailScRsp_proto_rawDescData = file_GetFriendChallengeDetailScRsp_proto_rawDesc
)

func file_GetFriendChallengeDetailScRsp_proto_rawDescGZIP() []byte {
	file_GetFriendChallengeDetailScRsp_proto_rawDescOnce.Do(func() {
		file_GetFriendChallengeDetailScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetFriendChallengeDetailScRsp_proto_rawDescData)
	})
	return file_GetFriendChallengeDetailScRsp_proto_rawDescData
}

var file_GetFriendChallengeDetailScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetFriendChallengeDetailScRsp_proto_goTypes = []interface{}{
	(*GetFriendChallengeDetailScRsp)(nil), // 0: GetFriendChallengeDetailScRsp
	(*DisplayAvatarDetailInfo)(nil),       // 1: DisplayAvatarDetailInfo
}
var file_GetFriendChallengeDetailScRsp_proto_depIdxs = []int32{
	1, // 0: GetFriendChallengeDetailScRsp.GCAONKNELPB:type_name -> DisplayAvatarDetailInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetFriendChallengeDetailScRsp_proto_init() }
func file_GetFriendChallengeDetailScRsp_proto_init() {
	if File_GetFriendChallengeDetailScRsp_proto != nil {
		return
	}
	file_DisplayAvatarDetailInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetFriendChallengeDetailScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFriendChallengeDetailScRsp); i {
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
			RawDescriptor: file_GetFriendChallengeDetailScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetFriendChallengeDetailScRsp_proto_goTypes,
		DependencyIndexes: file_GetFriendChallengeDetailScRsp_proto_depIdxs,
		MessageInfos:      file_GetFriendChallengeDetailScRsp_proto_msgTypes,
	}.Build()
	File_GetFriendChallengeDetailScRsp_proto = out.File
	file_GetFriendChallengeDetailScRsp_proto_rawDesc = nil
	file_GetFriendChallengeDetailScRsp_proto_goTypes = nil
	file_GetFriendChallengeDetailScRsp_proto_depIdxs = nil
}
