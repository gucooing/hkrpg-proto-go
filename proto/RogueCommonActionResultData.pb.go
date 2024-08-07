// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueCommonActionResultData.proto

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

type RogueCommonActionResultData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResultData:
	//
	//	*RogueCommonActionResultData_GetItemList
	//	*RogueCommonActionResultData_RemoveItemList
	//	*RogueCommonActionResultData_GetBuffList
	//	*RogueCommonActionResultData_RemoveBuffList
	//	*RogueCommonActionResultData_GetMiracleList
	//	*RogueCommonActionResultData_EJHDINBOAHL
	//	*RogueCommonActionResultData_LKFDEOOPJCM
	//	*RogueCommonActionResultData_BHAPKIBNMAL
	//	*RogueCommonActionResultData_MAHNLJALGKG
	//	*RogueCommonActionResultData_FGNGKEPMPDN
	//	*RogueCommonActionResultData_HGNNEGPPLCD
	//	*RogueCommonActionResultData_GetFormulaList
	//	*RogueCommonActionResultData_KDFGDFIIOMK
	//	*RogueCommonActionResultData_KLMOPHIFMHP
	//	*RogueCommonActionResultData_FPMNHKBFPEB
	//	*RogueCommonActionResultData_LJHEHDDAPCF
	//	*RogueCommonActionResultData_GetKeywordList
	//	*RogueCommonActionResultData_JMAEGNOGIPL
	ResultData isRogueCommonActionResultData_ResultData `protobuf_oneof:"result_data"`
}

func (x *RogueCommonActionResultData) Reset() {
	*x = RogueCommonActionResultData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueCommonActionResultData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueCommonActionResultData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueCommonActionResultData) ProtoMessage() {}

func (x *RogueCommonActionResultData) ProtoReflect() protoreflect.Message {
	mi := &file_RogueCommonActionResultData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueCommonActionResultData.ProtoReflect.Descriptor instead.
func (*RogueCommonActionResultData) Descriptor() ([]byte, []int) {
	return file_RogueCommonActionResultData_proto_rawDescGZIP(), []int{0}
}

func (m *RogueCommonActionResultData) GetResultData() isRogueCommonActionResultData_ResultData {
	if m != nil {
		return m.ResultData
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetItemList() *RogueCommonMoney {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetItemList); ok {
		return x.GetItemList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetRemoveItemList() *RogueCommonMoney {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_RemoveItemList); ok {
		return x.RemoveItemList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetBuffList() *RogueCommonBuff {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetBuffList); ok {
		return x.GetBuffList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetRemoveBuffList() *RogueCommonBuff {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_RemoveBuffList); ok {
		return x.RemoveBuffList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetMiracleList() *RogueCommonMiracle {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetMiracleList); ok {
		return x.GetMiracleList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetEJHDINBOAHL() *IHHGLFKEHNP {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_EJHDINBOAHL); ok {
		return x.EJHDINBOAHL
	}
	return nil
}

func (x *RogueCommonActionResultData) GetLKFDEOOPJCM() *EFLONNCOBBG {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_LKFDEOOPJCM); ok {
		return x.LKFDEOOPJCM
	}
	return nil
}

func (x *RogueCommonActionResultData) GetBHAPKIBNMAL() *MIPJKEPDMOF {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_BHAPKIBNMAL); ok {
		return x.BHAPKIBNMAL
	}
	return nil
}

func (x *RogueCommonActionResultData) GetMAHNLJALGKG() *NKICHKIGFNE {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_MAHNLJALGKG); ok {
		return x.MAHNLJALGKG
	}
	return nil
}

func (x *RogueCommonActionResultData) GetFGNGKEPMPDN() *LCBJFJPBJCB {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_FGNGKEPMPDN); ok {
		return x.FGNGKEPMPDN
	}
	return nil
}

func (x *RogueCommonActionResultData) GetHGNNEGPPLCD() *MNBGAMAKDOK {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_HGNNEGPPLCD); ok {
		return x.HGNNEGPPLCD
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetFormulaList() *RogueCommonFormula {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetFormulaList); ok {
		return x.GetFormulaList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetKDFGDFIIOMK() *GNBGFJCHFAN {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_KDFGDFIIOMK); ok {
		return x.KDFGDFIIOMK
	}
	return nil
}

func (x *RogueCommonActionResultData) GetKLMOPHIFMHP() *IOPMLAKEOCJ {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_KLMOPHIFMHP); ok {
		return x.KLMOPHIFMHP
	}
	return nil
}

func (x *RogueCommonActionResultData) GetFPMNHKBFPEB() *HHDOEPJGGNE {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_FPMNHKBFPEB); ok {
		return x.FPMNHKBFPEB
	}
	return nil
}

func (x *RogueCommonActionResultData) GetLJHEHDDAPCF() *MBCGGNJCEOE {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_LJHEHDDAPCF); ok {
		return x.LJHEHDDAPCF
	}
	return nil
}

func (x *RogueCommonActionResultData) GetGetKeywordList() *RogueCommonKeyword {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_GetKeywordList); ok {
		return x.GetKeywordList
	}
	return nil
}

func (x *RogueCommonActionResultData) GetJMAEGNOGIPL() *GLLJMPJJAFD {
	if x, ok := x.GetResultData().(*RogueCommonActionResultData_JMAEGNOGIPL); ok {
		return x.JMAEGNOGIPL
	}
	return nil
}

type isRogueCommonActionResultData_ResultData interface {
	isRogueCommonActionResultData_ResultData()
}

type RogueCommonActionResultData_GetItemList struct {
	GetItemList *RogueCommonMoney `protobuf:"bytes,6,opt,name=get_item_list,json=getItemList,proto3,oneof"`
}

type RogueCommonActionResultData_RemoveItemList struct {
	RemoveItemList *RogueCommonMoney `protobuf:"bytes,5,opt,name=remove_item_list,json=removeItemList,proto3,oneof"`
}

type RogueCommonActionResultData_GetBuffList struct {
	GetBuffList *RogueCommonBuff `protobuf:"bytes,202,opt,name=get_buff_list,json=getBuffList,proto3,oneof"`
}

type RogueCommonActionResultData_RemoveBuffList struct {
	RemoveBuffList *RogueCommonBuff `protobuf:"bytes,101,opt,name=remove_buff_list,json=removeBuffList,proto3,oneof"`
}

type RogueCommonActionResultData_GetMiracleList struct {
	GetMiracleList *RogueCommonMiracle `protobuf:"bytes,173,opt,name=get_miracle_list,json=getMiracleList,proto3,oneof"`
}

type RogueCommonActionResultData_EJHDINBOAHL struct {
	EJHDINBOAHL *IHHGLFKEHNP `protobuf:"bytes,363,opt,name=EJHDINBOAHL,proto3,oneof"`
}

type RogueCommonActionResultData_LKFDEOOPJCM struct {
	LKFDEOOPJCM *EFLONNCOBBG `protobuf:"bytes,1465,opt,name=LKFDEOOPJCM,proto3,oneof"`
}

type RogueCommonActionResultData_BHAPKIBNMAL struct {
	BHAPKIBNMAL *MIPJKEPDMOF `protobuf:"bytes,1875,opt,name=BHAPKIBNMAL,proto3,oneof"`
}

type RogueCommonActionResultData_MAHNLJALGKG struct {
	MAHNLJALGKG *NKICHKIGFNE `protobuf:"bytes,835,opt,name=MAHNLJALGKG,proto3,oneof"`
}

type RogueCommonActionResultData_FGNGKEPMPDN struct {
	FGNGKEPMPDN *LCBJFJPBJCB `protobuf:"bytes,142,opt,name=FGNGKEPMPDN,proto3,oneof"`
}

type RogueCommonActionResultData_HGNNEGPPLCD struct {
	HGNNEGPPLCD *MNBGAMAKDOK `protobuf:"bytes,1609,opt,name=HGNNEGPPLCD,proto3,oneof"`
}

type RogueCommonActionResultData_GetFormulaList struct {
	GetFormulaList *RogueCommonFormula `protobuf:"bytes,1837,opt,name=get_formula_list,json=getFormulaList,proto3,oneof"`
}

type RogueCommonActionResultData_KDFGDFIIOMK struct {
	KDFGDFIIOMK *GNBGFJCHFAN `protobuf:"bytes,1953,opt,name=KDFGDFIIOMK,proto3,oneof"`
}

type RogueCommonActionResultData_KLMOPHIFMHP struct {
	KLMOPHIFMHP *IOPMLAKEOCJ `protobuf:"bytes,1077,opt,name=KLMOPHIFMHP,proto3,oneof"`
}

type RogueCommonActionResultData_FPMNHKBFPEB struct {
	FPMNHKBFPEB *HHDOEPJGGNE `protobuf:"bytes,499,opt,name=FPMNHKBFPEB,proto3,oneof"`
}

type RogueCommonActionResultData_LJHEHDDAPCF struct {
	LJHEHDDAPCF *MBCGGNJCEOE `protobuf:"bytes,1885,opt,name=LJHEHDDAPCF,proto3,oneof"`
}

type RogueCommonActionResultData_GetKeywordList struct {
	GetKeywordList *RogueCommonKeyword `protobuf:"bytes,602,opt,name=get_keyword_list,json=getKeywordList,proto3,oneof"`
}

type RogueCommonActionResultData_JMAEGNOGIPL struct {
	JMAEGNOGIPL *GLLJMPJJAFD `protobuf:"bytes,1746,opt,name=JMAEGNOGIPL,proto3,oneof"`
}

func (*RogueCommonActionResultData_GetItemList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_RemoveItemList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetBuffList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_RemoveBuffList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetMiracleList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_EJHDINBOAHL) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_LKFDEOOPJCM) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_BHAPKIBNMAL) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_MAHNLJALGKG) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_FGNGKEPMPDN) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_HGNNEGPPLCD) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetFormulaList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_KDFGDFIIOMK) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_KLMOPHIFMHP) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_FPMNHKBFPEB) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_LJHEHDDAPCF) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_GetKeywordList) isRogueCommonActionResultData_ResultData() {}

func (*RogueCommonActionResultData_JMAEGNOGIPL) isRogueCommonActionResultData_ResultData() {}

var File_RogueCommonActionResultData_proto protoreflect.FileDescriptor

var file_RogueCommonActionResultData_proto_rawDesc = []byte{
	0x0a, 0x21, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x4b, 0x49, 0x43, 0x48, 0x4b, 0x49, 0x47, 0x46, 0x4e, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x43, 0x42, 0x4a, 0x46, 0x4a, 0x50, 0x42,
	0x4a, 0x43, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4e, 0x42, 0x47, 0x41,
	0x4d, 0x41, 0x4b, 0x44, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x46,
	0x4c, 0x4f, 0x4e, 0x4e, 0x43, 0x4f, 0x42, 0x42, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x48, 0x44, 0x4f, 0x45, 0x50, 0x4a, 0x47,
	0x47, 0x4e, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4e, 0x42, 0x47, 0x46,
	0x4a, 0x43, 0x48, 0x46, 0x41, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x4c,
	0x4c, 0x4a, 0x4d, 0x50, 0x4a, 0x4a, 0x41, 0x46, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x4d, 0x49, 0x50, 0x4a, 0x4b, 0x45, 0x50, 0x44, 0x4d, 0x4f, 0x46, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x42, 0x43, 0x47, 0x47, 0x4e, 0x4a, 0x43, 0x45, 0x4f,
	0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x49, 0x48, 0x48, 0x47, 0x4c, 0x46, 0x4b, 0x45, 0x48, 0x4e, 0x50, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x4f, 0x50, 0x4d, 0x4c, 0x41, 0x4b, 0x45, 0x4f, 0x43,
	0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x92, 0x08, 0x0a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x10, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x65, 0x74,
	0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42,
	0x75, 0x66, 0x66, 0x48, 0x00, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x62, 0x75, 0x66,
	0x66, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x75, 0x66, 0x66, 0x48, 0x00,
	0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x75, 0x66, 0x66, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0xad, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x45, 0x4a, 0x48, 0x44, 0x49, 0x4e, 0x42, 0x4f, 0x41, 0x48,
	0x4c, 0x18, 0xeb, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x48, 0x48, 0x47, 0x4c,
	0x46, 0x4b, 0x45, 0x48, 0x4e, 0x50, 0x48, 0x00, 0x52, 0x0b, 0x45, 0x4a, 0x48, 0x44, 0x49, 0x4e,
	0x42, 0x4f, 0x41, 0x48, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x4c, 0x4b, 0x46, 0x44, 0x45, 0x4f, 0x4f,
	0x50, 0x4a, 0x43, 0x4d, 0x18, 0xb9, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x46,
	0x4c, 0x4f, 0x4e, 0x4e, 0x43, 0x4f, 0x42, 0x42, 0x47, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x4b, 0x46,
	0x44, 0x45, 0x4f, 0x4f, 0x50, 0x4a, 0x43, 0x4d, 0x12, 0x31, 0x0a, 0x0b, 0x42, 0x48, 0x41, 0x50,
	0x4b, 0x49, 0x42, 0x4e, 0x4d, 0x41, 0x4c, 0x18, 0xd3, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4d, 0x49, 0x50, 0x4a, 0x4b, 0x45, 0x50, 0x44, 0x4d, 0x4f, 0x46, 0x48, 0x00, 0x52, 0x0b,
	0x42, 0x48, 0x41, 0x50, 0x4b, 0x49, 0x42, 0x4e, 0x4d, 0x41, 0x4c, 0x12, 0x31, 0x0a, 0x0b, 0x4d,
	0x41, 0x48, 0x4e, 0x4c, 0x4a, 0x41, 0x4c, 0x47, 0x4b, 0x47, 0x18, 0xc3, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x4b, 0x49, 0x43, 0x48, 0x4b, 0x49, 0x47, 0x46, 0x4e, 0x45, 0x48,
	0x00, 0x52, 0x0b, 0x4d, 0x41, 0x48, 0x4e, 0x4c, 0x4a, 0x41, 0x4c, 0x47, 0x4b, 0x47, 0x12, 0x31,
	0x0a, 0x0b, 0x46, 0x47, 0x4e, 0x47, 0x4b, 0x45, 0x50, 0x4d, 0x50, 0x44, 0x4e, 0x18, 0x8e, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x43, 0x42, 0x4a, 0x46, 0x4a, 0x50, 0x42, 0x4a,
	0x43, 0x42, 0x48, 0x00, 0x52, 0x0b, 0x46, 0x47, 0x4e, 0x47, 0x4b, 0x45, 0x50, 0x4d, 0x50, 0x44,
	0x4e, 0x12, 0x31, 0x0a, 0x0b, 0x48, 0x47, 0x4e, 0x4e, 0x45, 0x47, 0x50, 0x50, 0x4c, 0x43, 0x44,
	0x18, 0xc9, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4e, 0x42, 0x47, 0x41, 0x4d,
	0x41, 0x4b, 0x44, 0x4f, 0x4b, 0x48, 0x00, 0x52, 0x0b, 0x48, 0x47, 0x4e, 0x4e, 0x45, 0x47, 0x50,
	0x50, 0x4c, 0x43, 0x44, 0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x75, 0x6c, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xad, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x75,
	0x6c, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x44, 0x46, 0x47, 0x44, 0x46,
	0x49, 0x49, 0x4f, 0x4d, 0x4b, 0x18, 0xa1, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47,
	0x4e, 0x42, 0x47, 0x46, 0x4a, 0x43, 0x48, 0x46, 0x41, 0x4e, 0x48, 0x00, 0x52, 0x0b, 0x4b, 0x44,
	0x46, 0x47, 0x44, 0x46, 0x49, 0x49, 0x4f, 0x4d, 0x4b, 0x12, 0x31, 0x0a, 0x0b, 0x4b, 0x4c, 0x4d,
	0x4f, 0x50, 0x48, 0x49, 0x46, 0x4d, 0x48, 0x50, 0x18, 0xb5, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x4f, 0x50, 0x4d, 0x4c, 0x41, 0x4b, 0x45, 0x4f, 0x43, 0x4a, 0x48, 0x00, 0x52,
	0x0b, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x48, 0x49, 0x46, 0x4d, 0x48, 0x50, 0x12, 0x31, 0x0a, 0x0b,
	0x46, 0x50, 0x4d, 0x4e, 0x48, 0x4b, 0x42, 0x46, 0x50, 0x45, 0x42, 0x18, 0xf3, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x48, 0x44, 0x4f, 0x45, 0x50, 0x4a, 0x47, 0x47, 0x4e, 0x45,
	0x48, 0x00, 0x52, 0x0b, 0x46, 0x50, 0x4d, 0x4e, 0x48, 0x4b, 0x42, 0x46, 0x50, 0x45, 0x42, 0x12,
	0x31, 0x0a, 0x0b, 0x4c, 0x4a, 0x48, 0x45, 0x48, 0x44, 0x44, 0x41, 0x50, 0x43, 0x46, 0x18, 0xdd,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x42, 0x43, 0x47, 0x47, 0x4e, 0x4a, 0x43,
	0x45, 0x4f, 0x45, 0x48, 0x00, 0x52, 0x0b, 0x4c, 0x4a, 0x48, 0x45, 0x48, 0x44, 0x44, 0x41, 0x50,
	0x43, 0x46, 0x12, 0x40, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0xda, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x4a, 0x4d, 0x41, 0x45, 0x47, 0x4e, 0x4f, 0x47,
	0x49, 0x50, 0x4c, 0x18, 0xd2, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x4c, 0x4c,
	0x4a, 0x4d, 0x50, 0x4a, 0x4a, 0x41, 0x46, 0x44, 0x48, 0x00, 0x52, 0x0b, 0x4a, 0x4d, 0x41, 0x45,
	0x47, 0x4e, 0x4f, 0x47, 0x49, 0x50, 0x4c, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueCommonActionResultData_proto_rawDescOnce sync.Once
	file_RogueCommonActionResultData_proto_rawDescData = file_RogueCommonActionResultData_proto_rawDesc
)

func file_RogueCommonActionResultData_proto_rawDescGZIP() []byte {
	file_RogueCommonActionResultData_proto_rawDescOnce.Do(func() {
		file_RogueCommonActionResultData_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueCommonActionResultData_proto_rawDescData)
	})
	return file_RogueCommonActionResultData_proto_rawDescData
}

var file_RogueCommonActionResultData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueCommonActionResultData_proto_goTypes = []interface{}{
	(*RogueCommonActionResultData)(nil), // 0: RogueCommonActionResultData
	(*RogueCommonMoney)(nil),            // 1: RogueCommonMoney
	(*RogueCommonBuff)(nil),             // 2: RogueCommonBuff
	(*RogueCommonMiracle)(nil),          // 3: RogueCommonMiracle
	(*IHHGLFKEHNP)(nil),                 // 4: IHHGLFKEHNP
	(*EFLONNCOBBG)(nil),                 // 5: EFLONNCOBBG
	(*MIPJKEPDMOF)(nil),                 // 6: MIPJKEPDMOF
	(*NKICHKIGFNE)(nil),                 // 7: NKICHKIGFNE
	(*LCBJFJPBJCB)(nil),                 // 8: LCBJFJPBJCB
	(*MNBGAMAKDOK)(nil),                 // 9: MNBGAMAKDOK
	(*RogueCommonFormula)(nil),          // 10: RogueCommonFormula
	(*GNBGFJCHFAN)(nil),                 // 11: GNBGFJCHFAN
	(*IOPMLAKEOCJ)(nil),                 // 12: IOPMLAKEOCJ
	(*HHDOEPJGGNE)(nil),                 // 13: HHDOEPJGGNE
	(*MBCGGNJCEOE)(nil),                 // 14: MBCGGNJCEOE
	(*RogueCommonKeyword)(nil),          // 15: RogueCommonKeyword
	(*GLLJMPJJAFD)(nil),                 // 16: GLLJMPJJAFD
}
var file_RogueCommonActionResultData_proto_depIdxs = []int32{
	1,  // 0: RogueCommonActionResultData.get_item_list:type_name -> RogueCommonMoney
	1,  // 1: RogueCommonActionResultData.remove_item_list:type_name -> RogueCommonMoney
	2,  // 2: RogueCommonActionResultData.get_buff_list:type_name -> RogueCommonBuff
	2,  // 3: RogueCommonActionResultData.remove_buff_list:type_name -> RogueCommonBuff
	3,  // 4: RogueCommonActionResultData.get_miracle_list:type_name -> RogueCommonMiracle
	4,  // 5: RogueCommonActionResultData.EJHDINBOAHL:type_name -> IHHGLFKEHNP
	5,  // 6: RogueCommonActionResultData.LKFDEOOPJCM:type_name -> EFLONNCOBBG
	6,  // 7: RogueCommonActionResultData.BHAPKIBNMAL:type_name -> MIPJKEPDMOF
	7,  // 8: RogueCommonActionResultData.MAHNLJALGKG:type_name -> NKICHKIGFNE
	8,  // 9: RogueCommonActionResultData.FGNGKEPMPDN:type_name -> LCBJFJPBJCB
	9,  // 10: RogueCommonActionResultData.HGNNEGPPLCD:type_name -> MNBGAMAKDOK
	10, // 11: RogueCommonActionResultData.get_formula_list:type_name -> RogueCommonFormula
	11, // 12: RogueCommonActionResultData.KDFGDFIIOMK:type_name -> GNBGFJCHFAN
	12, // 13: RogueCommonActionResultData.KLMOPHIFMHP:type_name -> IOPMLAKEOCJ
	13, // 14: RogueCommonActionResultData.FPMNHKBFPEB:type_name -> HHDOEPJGGNE
	14, // 15: RogueCommonActionResultData.LJHEHDDAPCF:type_name -> MBCGGNJCEOE
	15, // 16: RogueCommonActionResultData.get_keyword_list:type_name -> RogueCommonKeyword
	16, // 17: RogueCommonActionResultData.JMAEGNOGIPL:type_name -> GLLJMPJJAFD
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_RogueCommonActionResultData_proto_init() }
func file_RogueCommonActionResultData_proto_init() {
	if File_RogueCommonActionResultData_proto != nil {
		return
	}
	file_NKICHKIGFNE_proto_init()
	file_LCBJFJPBJCB_proto_init()
	file_MNBGAMAKDOK_proto_init()
	file_EFLONNCOBBG_proto_init()
	file_RogueCommonBuff_proto_init()
	file_HHDOEPJGGNE_proto_init()
	file_GNBGFJCHFAN_proto_init()
	file_GLLJMPJJAFD_proto_init()
	file_MIPJKEPDMOF_proto_init()
	file_RogueCommonFormula_proto_init()
	file_RogueCommonMoney_proto_init()
	file_MBCGGNJCEOE_proto_init()
	file_RogueCommonKeyword_proto_init()
	file_IHHGLFKEHNP_proto_init()
	file_IOPMLAKEOCJ_proto_init()
	file_RogueCommonMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueCommonActionResultData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueCommonActionResultData); i {
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
	file_RogueCommonActionResultData_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RogueCommonActionResultData_GetItemList)(nil),
		(*RogueCommonActionResultData_RemoveItemList)(nil),
		(*RogueCommonActionResultData_GetBuffList)(nil),
		(*RogueCommonActionResultData_RemoveBuffList)(nil),
		(*RogueCommonActionResultData_GetMiracleList)(nil),
		(*RogueCommonActionResultData_EJHDINBOAHL)(nil),
		(*RogueCommonActionResultData_LKFDEOOPJCM)(nil),
		(*RogueCommonActionResultData_BHAPKIBNMAL)(nil),
		(*RogueCommonActionResultData_MAHNLJALGKG)(nil),
		(*RogueCommonActionResultData_FGNGKEPMPDN)(nil),
		(*RogueCommonActionResultData_HGNNEGPPLCD)(nil),
		(*RogueCommonActionResultData_GetFormulaList)(nil),
		(*RogueCommonActionResultData_KDFGDFIIOMK)(nil),
		(*RogueCommonActionResultData_KLMOPHIFMHP)(nil),
		(*RogueCommonActionResultData_FPMNHKBFPEB)(nil),
		(*RogueCommonActionResultData_LJHEHDDAPCF)(nil),
		(*RogueCommonActionResultData_GetKeywordList)(nil),
		(*RogueCommonActionResultData_JMAEGNOGIPL)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueCommonActionResultData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueCommonActionResultData_proto_goTypes,
		DependencyIndexes: file_RogueCommonActionResultData_proto_depIdxs,
		MessageInfos:      file_RogueCommonActionResultData_proto_msgTypes,
	}.Build()
	File_RogueCommonActionResultData_proto = out.File
	file_RogueCommonActionResultData_proto_rawDesc = nil
	file_RogueCommonActionResultData_proto_goTypes = nil
	file_RogueCommonActionResultData_proto_depIdxs = nil
}
