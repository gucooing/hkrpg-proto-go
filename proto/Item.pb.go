// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Item.proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId      uint32 `protobuf:"varint,6,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Level       uint32 `protobuf:"varint,10,opt,name=level,proto3" json:"level,omitempty"`
	Num         uint32 `protobuf:"varint,12,opt,name=num,proto3" json:"num,omitempty"`
	MainAffixId uint32 `protobuf:"varint,3,opt,name=main_affix_id,json=mainAffixId,proto3" json:"main_affix_id,omitempty"`
	Rank        uint32 `protobuf:"varint,7,opt,name=rank,proto3" json:"rank,omitempty"`
	Promotion   uint32 `protobuf:"varint,4,opt,name=promotion,proto3" json:"promotion,omitempty"`
	UniqueId    uint32 `protobuf:"varint,2,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_Item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_Item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Item) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Item) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Item) GetMainAffixId() uint32 {
	if x != nil {
		return x.MainAffixId
	}
	return 0
}

func (x *Item) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *Item) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

func (x *Item) GetUniqueId() uint32 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

var File_Item_proto protoreflect.FileDescriptor

var file_Item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a,
	0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61,
	0x66, 0x66, 0x69, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d,
	0x61, 0x69, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Item_proto_rawDescOnce sync.Once
	file_Item_proto_rawDescData = file_Item_proto_rawDesc
)

func file_Item_proto_rawDescGZIP() []byte {
	file_Item_proto_rawDescOnce.Do(func() {
		file_Item_proto_rawDescData = protoimpl.X.CompressGZIP(file_Item_proto_rawDescData)
	})
	return file_Item_proto_rawDescData
}

var file_Item_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Item_proto_goTypes = []interface{}{
	(*Item)(nil), // 0: Item
}
var file_Item_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Item_proto_init() }
func file_Item_proto_init() {
	if File_Item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
			RawDescriptor: file_Item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Item_proto_goTypes,
		DependencyIndexes: file_Item_proto_depIdxs,
		MessageInfos:      file_Item_proto_msgTypes,
	}.Build()
	File_Item_proto = out.File
	file_Item_proto_rawDesc = nil
	file_Item_proto_goTypes = nil
	file_Item_proto_depIdxs = nil
}
