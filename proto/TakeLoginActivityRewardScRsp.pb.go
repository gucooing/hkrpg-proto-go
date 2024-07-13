// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeLoginActivityRewardScRsp.proto

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

type TakeLoginActivityRewardScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  uint32    `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Id       uint32    `protobuf:"varint,11,opt,name=id,proto3" json:"id,omitempty"`
	Reward   *ItemList `protobuf:"bytes,8,opt,name=reward,proto3" json:"reward,omitempty"`
	TakeDays uint32    `protobuf:"varint,15,opt,name=take_days,json=takeDays,proto3" json:"take_days,omitempty"`
}

func (x *TakeLoginActivityRewardScRsp) Reset() {
	*x = TakeLoginActivityRewardScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeLoginActivityRewardScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeLoginActivityRewardScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeLoginActivityRewardScRsp) ProtoMessage() {}

func (x *TakeLoginActivityRewardScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeLoginActivityRewardScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeLoginActivityRewardScRsp.ProtoReflect.Descriptor instead.
func (*TakeLoginActivityRewardScRsp) Descriptor() ([]byte, []int) {
	return file_TakeLoginActivityRewardScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeLoginActivityRewardScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TakeLoginActivityRewardScRsp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TakeLoginActivityRewardScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *TakeLoginActivityRewardScRsp) GetTakeDays() uint32 {
	if x != nil {
		return x.TakeDays
	}
	return 0
}

var File_TakeLoginActivityRewardScRsp_proto protoreflect.FileDescriptor

var file_TakeLoginActivityRewardScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x54, 0x61, 0x6b, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x1c, 0x54, 0x61, 0x6b, 0x65, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x21, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x6b, 0x65, 0x44, 0x61, 0x79, 0x73, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TakeLoginActivityRewardScRsp_proto_rawDescOnce sync.Once
	file_TakeLoginActivityRewardScRsp_proto_rawDescData = file_TakeLoginActivityRewardScRsp_proto_rawDesc
)

func file_TakeLoginActivityRewardScRsp_proto_rawDescGZIP() []byte {
	file_TakeLoginActivityRewardScRsp_proto_rawDescOnce.Do(func() {
		file_TakeLoginActivityRewardScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeLoginActivityRewardScRsp_proto_rawDescData)
	})
	return file_TakeLoginActivityRewardScRsp_proto_rawDescData
}

var file_TakeLoginActivityRewardScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeLoginActivityRewardScRsp_proto_goTypes = []interface{}{
	(*TakeLoginActivityRewardScRsp)(nil), // 0: TakeLoginActivityRewardScRsp
	(*ItemList)(nil),                     // 1: ItemList
}
var file_TakeLoginActivityRewardScRsp_proto_depIdxs = []int32{
	1, // 0: TakeLoginActivityRewardScRsp.reward:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TakeLoginActivityRewardScRsp_proto_init() }
func file_TakeLoginActivityRewardScRsp_proto_init() {
	if File_TakeLoginActivityRewardScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeLoginActivityRewardScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeLoginActivityRewardScRsp); i {
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
			RawDescriptor: file_TakeLoginActivityRewardScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeLoginActivityRewardScRsp_proto_goTypes,
		DependencyIndexes: file_TakeLoginActivityRewardScRsp_proto_depIdxs,
		MessageInfos:      file_TakeLoginActivityRewardScRsp_proto_msgTypes,
	}.Build()
	File_TakeLoginActivityRewardScRsp_proto = out.File
	file_TakeLoginActivityRewardScRsp_proto_rawDesc = nil
	file_TakeLoginActivityRewardScRsp_proto_goTypes = nil
	file_TakeLoginActivityRewardScRsp_proto_depIdxs = nil
}
