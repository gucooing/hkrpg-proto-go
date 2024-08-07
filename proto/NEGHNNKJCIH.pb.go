// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NEGHNNKJCIH.proto

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

type NEGHNNKJCIH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicInfo   *BEDDPMJHDJB `protobuf:"bytes,1,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	NBAJBDFFLHL *OBJPLDJBJPP `protobuf:"bytes,2,opt,name=NBAJBDFFLHL,proto3" json:"NBAJBDFFLHL,omitempty"`
	PlayerInfo  *NDKCBDPOLAH `protobuf:"bytes,3,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
}

func (x *NEGHNNKJCIH) Reset() {
	*x = NEGHNNKJCIH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NEGHNNKJCIH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NEGHNNKJCIH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NEGHNNKJCIH) ProtoMessage() {}

func (x *NEGHNNKJCIH) ProtoReflect() protoreflect.Message {
	mi := &file_NEGHNNKJCIH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NEGHNNKJCIH.ProtoReflect.Descriptor instead.
func (*NEGHNNKJCIH) Descriptor() ([]byte, []int) {
	return file_NEGHNNKJCIH_proto_rawDescGZIP(), []int{0}
}

func (x *NEGHNNKJCIH) GetBasicInfo() *BEDDPMJHDJB {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *NEGHNNKJCIH) GetNBAJBDFFLHL() *OBJPLDJBJPP {
	if x != nil {
		return x.NBAJBDFFLHL
	}
	return nil
}

func (x *NEGHNNKJCIH) GetPlayerInfo() *NDKCBDPOLAH {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

var File_NEGHNNKJCIH_proto protoreflect.FileDescriptor

var file_NEGHNNKJCIH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x45, 0x47, 0x48, 0x4e, 0x4e, 0x4b, 0x4a, 0x43, 0x49, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x44, 0x4b, 0x43, 0x42, 0x44, 0x50, 0x4f, 0x4c, 0x41, 0x48,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42, 0x45, 0x44, 0x44, 0x50, 0x4d, 0x4a, 0x48,
	0x44, 0x4a, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x42, 0x4a, 0x50, 0x4c,
	0x44, 0x4a, 0x42, 0x4a, 0x50, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a,
	0x0b, 0x4e, 0x45, 0x47, 0x48, 0x4e, 0x4e, 0x4b, 0x4a, 0x43, 0x49, 0x48, 0x12, 0x2b, 0x0a, 0x0a,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x42, 0x45, 0x44, 0x44, 0x50, 0x4d, 0x4a, 0x48, 0x44, 0x4a, 0x42, 0x52, 0x09,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x42, 0x41,
	0x4a, 0x42, 0x44, 0x46, 0x46, 0x4c, 0x48, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4f, 0x42, 0x4a, 0x50, 0x4c, 0x44, 0x4a, 0x42, 0x4a, 0x50, 0x50, 0x52, 0x0b, 0x4e, 0x42,
	0x41, 0x4a, 0x42, 0x44, 0x46, 0x46, 0x4c, 0x48, 0x4c, 0x12, 0x2d, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4e, 0x44, 0x4b, 0x43, 0x42, 0x44, 0x50, 0x4f, 0x4c, 0x41, 0x48, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NEGHNNKJCIH_proto_rawDescOnce sync.Once
	file_NEGHNNKJCIH_proto_rawDescData = file_NEGHNNKJCIH_proto_rawDesc
)

func file_NEGHNNKJCIH_proto_rawDescGZIP() []byte {
	file_NEGHNNKJCIH_proto_rawDescOnce.Do(func() {
		file_NEGHNNKJCIH_proto_rawDescData = protoimpl.X.CompressGZIP(file_NEGHNNKJCIH_proto_rawDescData)
	})
	return file_NEGHNNKJCIH_proto_rawDescData
}

var file_NEGHNNKJCIH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NEGHNNKJCIH_proto_goTypes = []interface{}{
	(*NEGHNNKJCIH)(nil), // 0: NEGHNNKJCIH
	(*BEDDPMJHDJB)(nil), // 1: BEDDPMJHDJB
	(*OBJPLDJBJPP)(nil), // 2: OBJPLDJBJPP
	(*NDKCBDPOLAH)(nil), // 3: NDKCBDPOLAH
}
var file_NEGHNNKJCIH_proto_depIdxs = []int32{
	1, // 0: NEGHNNKJCIH.basic_info:type_name -> BEDDPMJHDJB
	2, // 1: NEGHNNKJCIH.NBAJBDFFLHL:type_name -> OBJPLDJBJPP
	3, // 2: NEGHNNKJCIH.player_info:type_name -> NDKCBDPOLAH
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_NEGHNNKJCIH_proto_init() }
func file_NEGHNNKJCIH_proto_init() {
	if File_NEGHNNKJCIH_proto != nil {
		return
	}
	file_NDKCBDPOLAH_proto_init()
	file_BEDDPMJHDJB_proto_init()
	file_OBJPLDJBJPP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NEGHNNKJCIH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NEGHNNKJCIH); i {
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
			RawDescriptor: file_NEGHNNKJCIH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NEGHNNKJCIH_proto_goTypes,
		DependencyIndexes: file_NEGHNNKJCIH_proto_depIdxs,
		MessageInfos:      file_NEGHNNKJCIH_proto_msgTypes,
	}.Build()
	File_NEGHNNKJCIH_proto = out.File
	file_NEGHNNKJCIH_proto_rawDesc = nil
	file_NEGHNNKJCIH_proto_goTypes = nil
	file_NEGHNNKJCIH_proto_depIdxs = nil
}
