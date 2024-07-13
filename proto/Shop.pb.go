// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Shop.proto

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

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityTakenLevelReward uint64   `protobuf:"varint,4,opt,name=city_taken_level_reward,json=cityTakenLevelReward,proto3" json:"city_taken_level_reward,omitempty"`
	CityExp              uint32   `protobuf:"varint,2,opt,name=city_exp,json=cityExp,proto3" json:"city_exp,omitempty"`
	EndTime              int64    `protobuf:"varint,15,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CityLevel            uint32   `protobuf:"varint,8,opt,name=city_level,json=cityLevel,proto3" json:"city_level,omitempty"`
	ShopId               uint32   `protobuf:"varint,14,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	GoodsList            []*Goods `protobuf:"bytes,9,rep,name=goods_list,json=goodsList,proto3" json:"goods_list,omitempty"`
	BeginTime            int64    `protobuf:"varint,10,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_Shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_Shop_proto_rawDescGZIP(), []int{0}
}

func (x *Shop) GetCityTakenLevelReward() uint64 {
	if x != nil {
		return x.CityTakenLevelReward
	}
	return 0
}

func (x *Shop) GetCityExp() uint32 {
	if x != nil {
		return x.CityExp
	}
	return 0
}

func (x *Shop) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Shop) GetCityLevel() uint32 {
	if x != nil {
		return x.CityLevel
	}
	return 0
}

func (x *Shop) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Shop) GetGoodsList() []*Goods {
	if x != nil {
		return x.GoodsList
	}
	return nil
}

func (x *Shop) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_Shop_proto protoreflect.FileDescriptor

var file_Shop_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x04, 0x53, 0x68,
	0x6f, 0x70, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x14, 0x63, 0x69, 0x74, 0x79, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x69, 0x74,
	0x79, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x69, 0x74,
	0x79, 0x45, 0x78, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Shop_proto_rawDescOnce sync.Once
	file_Shop_proto_rawDescData = file_Shop_proto_rawDesc
)

func file_Shop_proto_rawDescGZIP() []byte {
	file_Shop_proto_rawDescOnce.Do(func() {
		file_Shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_Shop_proto_rawDescData)
	})
	return file_Shop_proto_rawDescData
}

var file_Shop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Shop_proto_goTypes = []interface{}{
	(*Shop)(nil),  // 0: Shop
	(*Goods)(nil), // 1: Goods
}
var file_Shop_proto_depIdxs = []int32{
	1, // 0: Shop.goods_list:type_name -> Goods
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Shop_proto_init() }
func file_Shop_proto_init() {
	if File_Shop_proto != nil {
		return
	}
	file_Goods_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
			RawDescriptor: file_Shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Shop_proto_goTypes,
		DependencyIndexes: file_Shop_proto_depIdxs,
		MessageInfos:      file_Shop_proto_msgTypes,
	}.Build()
	File_Shop_proto = out.File
	file_Shop_proto_rawDesc = nil
	file_Shop_proto_goTypes = nil
	file_Shop_proto_depIdxs = nil
}
