// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetRogueBuffEnhanceInfoScRsp.proto

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

type GetRogueBuffEnhanceInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  uint32                    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ShopInfo *RogueBuffEnhanceShopInfo `protobuf:"bytes,6,opt,name=shop_info,json=shopInfo,proto3" json:"shop_info,omitempty"`
}

func (x *GetRogueBuffEnhanceInfoScRsp) Reset() {
	*x = GetRogueBuffEnhanceInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRogueBuffEnhanceInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRogueBuffEnhanceInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRogueBuffEnhanceInfoScRsp) ProtoMessage() {}

func (x *GetRogueBuffEnhanceInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRogueBuffEnhanceInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRogueBuffEnhanceInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetRogueBuffEnhanceInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRogueBuffEnhanceInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetRogueBuffEnhanceInfoScRsp) GetShopInfo() *RogueBuffEnhanceShopInfo {
	if x != nil {
		return x.ShopInfo
	}
	return nil
}

var File_GetRogueBuffEnhanceInfoScRsp_proto protoreflect.FileDescriptor

var file_GetRogueBuffEnhanceInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x22, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x45, 0x6e,
	0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x45,
	0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x45, 0x6e, 0x68,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x73, 0x68,
	0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescOnce sync.Once
	file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescData = file_GetRogueBuffEnhanceInfoScRsp_proto_rawDesc
)

func file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescData)
	})
	return file_GetRogueBuffEnhanceInfoScRsp_proto_rawDescData
}

var file_GetRogueBuffEnhanceInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRogueBuffEnhanceInfoScRsp_proto_goTypes = []interface{}{
	(*GetRogueBuffEnhanceInfoScRsp)(nil), // 0: GetRogueBuffEnhanceInfoScRsp
	(*RogueBuffEnhanceShopInfo)(nil),     // 1: RogueBuffEnhanceShopInfo
}
var file_GetRogueBuffEnhanceInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetRogueBuffEnhanceInfoScRsp.shop_info:type_name -> RogueBuffEnhanceShopInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetRogueBuffEnhanceInfoScRsp_proto_init() }
func file_GetRogueBuffEnhanceInfoScRsp_proto_init() {
	if File_GetRogueBuffEnhanceInfoScRsp_proto != nil {
		return
	}
	file_RogueBuffEnhanceShopInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetRogueBuffEnhanceInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRogueBuffEnhanceInfoScRsp); i {
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
			RawDescriptor: file_GetRogueBuffEnhanceInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRogueBuffEnhanceInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetRogueBuffEnhanceInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetRogueBuffEnhanceInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetRogueBuffEnhanceInfoScRsp_proto = out.File
	file_GetRogueBuffEnhanceInfoScRsp_proto_rawDesc = nil
	file_GetRogueBuffEnhanceInfoScRsp_proto_goTypes = nil
	file_GetRogueBuffEnhanceInfoScRsp_proto_depIdxs = nil
}