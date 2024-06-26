// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueActionResultData.proto

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

type RogueActionResultData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResultAction:
	//
	//	*RogueActionResultData_AddItemList
	//	*RogueActionResultData_RemoveItemList
	//	*RogueActionResultData_AddBuffList
	//	*RogueActionResultData_RemoveBuffList
	//	*RogueActionResultData_AddMiracleList
	ResultAction isRogueActionResultData_ResultAction `protobuf_oneof:"result_action"`
}

func (x *RogueActionResultData) Reset() {
	*x = RogueActionResultData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueActionResultData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueActionResultData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueActionResultData) ProtoMessage() {}

func (x *RogueActionResultData) ProtoReflect() protoreflect.Message {
	mi := &file_RogueActionResultData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueActionResultData.ProtoReflect.Descriptor instead.
func (*RogueActionResultData) Descriptor() ([]byte, []int) {
	return file_RogueActionResultData_proto_rawDescGZIP(), []int{0}
}

func (m *RogueActionResultData) GetResultAction() isRogueActionResultData_ResultAction {
	if m != nil {
		return m.ResultAction
	}
	return nil
}

func (x *RogueActionResultData) GetAddItemList() *RogueVirtualItemData {
	if x, ok := x.GetResultAction().(*RogueActionResultData_AddItemList); ok {
		return x.AddItemList
	}
	return nil
}

func (x *RogueActionResultData) GetRemoveItemList() *RogueVirtualItemData {
	if x, ok := x.GetResultAction().(*RogueActionResultData_RemoveItemList); ok {
		return x.RemoveItemList
	}
	return nil
}

func (x *RogueActionResultData) GetAddBuffList() *RogueBuffData {
	if x, ok := x.GetResultAction().(*RogueActionResultData_AddBuffList); ok {
		return x.AddBuffList
	}
	return nil
}

func (x *RogueActionResultData) GetRemoveBuffList() *RogueBuffData {
	if x, ok := x.GetResultAction().(*RogueActionResultData_RemoveBuffList); ok {
		return x.RemoveBuffList
	}
	return nil
}

func (x *RogueActionResultData) GetAddMiracleList() *RogueMiracleData {
	if x, ok := x.GetResultAction().(*RogueActionResultData_AddMiracleList); ok {
		return x.AddMiracleList
	}
	return nil
}

type isRogueActionResultData_ResultAction interface {
	isRogueActionResultData_ResultAction()
}

type RogueActionResultData_AddItemList struct {
	AddItemList *RogueVirtualItemData `protobuf:"bytes,11,opt,name=add_item_list,json=addItemList,proto3,oneof"`
}

type RogueActionResultData_RemoveItemList struct {
	RemoveItemList *RogueVirtualItemData `protobuf:"bytes,14,opt,name=remove_item_list,json=removeItemList,proto3,oneof"`
}

type RogueActionResultData_AddBuffList struct {
	AddBuffList *RogueBuffData `protobuf:"bytes,1586,opt,name=add_buff_list,json=addBuffList,proto3,oneof"`
}

type RogueActionResultData_RemoveBuffList struct {
	RemoveBuffList *RogueBuffData `protobuf:"bytes,1574,opt,name=remove_buff_list,json=removeBuffList,proto3,oneof"`
}

type RogueActionResultData_AddMiracleList struct {
	AddMiracleList *RogueMiracleData `protobuf:"bytes,516,opt,name=add_miracle_list,json=addMiracleList,proto3,oneof"`
}

func (*RogueActionResultData_AddItemList) isRogueActionResultData_ResultAction() {}

func (*RogueActionResultData_RemoveItemList) isRogueActionResultData_ResultAction() {}

func (*RogueActionResultData_AddBuffList) isRogueActionResultData_ResultAction() {}

func (*RogueActionResultData_RemoveBuffList) isRogueActionResultData_ResultAction() {}

func (*RogueActionResultData_AddMiracleList) isRogueActionResultData_ResultAction() {}

var File_RogueActionResultData_proto protoreflect.FileDescriptor

var file_RogueActionResultData_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x3b, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x0b, 0x61, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00,
	0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0xb2, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x42, 0x75, 0x66, 0x66, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x42,
	0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xa6, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x75, 0x66, 0x66,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x10, 0x61, 0x64, 0x64, 0x5f, 0x6d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x84, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueActionResultData_proto_rawDescOnce sync.Once
	file_RogueActionResultData_proto_rawDescData = file_RogueActionResultData_proto_rawDesc
)

func file_RogueActionResultData_proto_rawDescGZIP() []byte {
	file_RogueActionResultData_proto_rawDescOnce.Do(func() {
		file_RogueActionResultData_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueActionResultData_proto_rawDescData)
	})
	return file_RogueActionResultData_proto_rawDescData
}

var file_RogueActionResultData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueActionResultData_proto_goTypes = []interface{}{
	(*RogueActionResultData)(nil), // 0: RogueActionResultData
	(*RogueVirtualItemData)(nil),  // 1: RogueVirtualItemData
	(*RogueBuffData)(nil),         // 2: RogueBuffData
	(*RogueMiracleData)(nil),      // 3: RogueMiracleData
}
var file_RogueActionResultData_proto_depIdxs = []int32{
	1, // 0: RogueActionResultData.add_item_list:type_name -> RogueVirtualItemData
	1, // 1: RogueActionResultData.remove_item_list:type_name -> RogueVirtualItemData
	2, // 2: RogueActionResultData.add_buff_list:type_name -> RogueBuffData
	2, // 3: RogueActionResultData.remove_buff_list:type_name -> RogueBuffData
	3, // 4: RogueActionResultData.add_miracle_list:type_name -> RogueMiracleData
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_RogueActionResultData_proto_init() }
func file_RogueActionResultData_proto_init() {
	if File_RogueActionResultData_proto != nil {
		return
	}
	file_RogueBuffData_proto_init()
	file_RogueMiracleData_proto_init()
	file_RogueVirtualItemData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueActionResultData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueActionResultData); i {
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
	file_RogueActionResultData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueActionResultData_AddItemList)(nil),
		(*RogueActionResultData_RemoveItemList)(nil),
		(*RogueActionResultData_AddBuffList)(nil),
		(*RogueActionResultData_RemoveBuffList)(nil),
		(*RogueActionResultData_AddMiracleList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueActionResultData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueActionResultData_proto_goTypes,
		DependencyIndexes: file_RogueActionResultData_proto_depIdxs,
		MessageInfos:      file_RogueActionResultData_proto_msgTypes,
	}.Build()
	File_RogueActionResultData_proto = out.File
	file_RogueActionResultData_proto_rawDesc = nil
	file_RogueActionResultData_proto_goTypes = nil
	file_RogueActionResultData_proto_depIdxs = nil
}
