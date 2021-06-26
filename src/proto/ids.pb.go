// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: dist/proto/ids.proto

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

type GetIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetIdReq) Reset() {
	*x = GetIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdReq) ProtoMessage() {}

func (x *GetIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdReq.ProtoReflect.Descriptor instead.
func (*GetIdReq) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{0}
}

func (x *GetIdReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIdRes) Reset() {
	*x = GetIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdRes) ProtoMessage() {}

func (x *GetIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdRes.ProtoReflect.Descriptor instead.
func (*GetIdRes) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{1}
}

func (x *GetIdRes) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetIdInMapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EId  uint64 `protobuf:"varint,1,opt,name=e_id,json=eId,proto3" json:"e_id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetIdInMapReq) Reset() {
	*x = GetIdInMapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdInMapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdInMapReq) ProtoMessage() {}

func (x *GetIdInMapReq) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdInMapReq.ProtoReflect.Descriptor instead.
func (*GetIdInMapReq) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{2}
}

func (x *GetIdInMapReq) GetEId() uint64 {
	if x != nil {
		return x.EId
	}
	return 0
}

func (x *GetIdInMapReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetIdInMapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIdInMapRes) Reset() {
	*x = GetIdInMapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdInMapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdInMapRes) ProtoMessage() {}

func (x *GetIdInMapRes) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdInMapRes.ProtoReflect.Descriptor instead.
func (*GetIdInMapRes) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{3}
}

func (x *GetIdInMapRes) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetIdArrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Length uint64 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *GetIdArrReq) Reset() {
	*x = GetIdArrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdArrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdArrReq) ProtoMessage() {}

func (x *GetIdArrReq) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdArrReq.ProtoReflect.Descriptor instead.
func (*GetIdArrReq) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{4}
}

func (x *GetIdArrReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetIdArrReq) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type GetIdArrRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIdArrRes) Reset() {
	*x = GetIdArrRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdArrRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdArrRes) ProtoMessage() {}

func (x *GetIdArrRes) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdArrRes.ProtoReflect.Descriptor instead.
func (*GetIdArrRes) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{5}
}

func (x *GetIdArrRes) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetIdArrInMapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EId    uint64 `protobuf:"varint,1,opt,name=e_id,json=eId,proto3" json:"e_id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Length uint64 `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *GetIdArrInMapReq) Reset() {
	*x = GetIdArrInMapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdArrInMapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdArrInMapReq) ProtoMessage() {}

func (x *GetIdArrInMapReq) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdArrInMapReq.ProtoReflect.Descriptor instead.
func (*GetIdArrInMapReq) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{6}
}

func (x *GetIdArrInMapReq) GetEId() uint64 {
	if x != nil {
		return x.EId
	}
	return 0
}

func (x *GetIdArrInMapReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetIdArrInMapReq) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type GetIdArrInMapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIdArrInMapRes) Reset() {
	*x = GetIdArrInMapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dist_proto_ids_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdArrInMapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdArrInMapRes) ProtoMessage() {}

func (x *GetIdArrInMapRes) ProtoReflect() protoreflect.Message {
	mi := &file_dist_proto_ids_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdArrInMapRes.ProtoReflect.Descriptor instead.
func (*GetIdArrInMapRes) Descriptor() ([]byte, []int) {
	return file_dist_proto_ids_proto_rawDescGZIP(), []int{7}
}

func (x *GetIdArrInMapRes) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_dist_proto_ids_proto protoreflect.FileDescriptor

var file_dist_proto_ids_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x36, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x70,
	0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x04, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x49, 0x64, 0x41, 0x72, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41,
	0x72, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72,
	0x72, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x12, 0x11, 0x0a, 0x04, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49,
	0x64, 0x41, 0x72, 0x72, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0x53, 0x0a, 0x0a,
	0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x05, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x09, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x09,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x49, 0x64, 0x41, 0x72, 0x72, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72, 0x72, 0x52, 0x65,
	0x73, 0x32, 0x76, 0x0a, 0x0f, 0x49, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x4d,
	0x61, 0x70, 0x12, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72, 0x72, 0x49, 0x6e,
	0x4d, 0x61, 0x70, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72, 0x72, 0x49, 0x6e,
	0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x41, 0x72,
	0x72, 0x49, 0x6e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dist_proto_ids_proto_rawDescOnce sync.Once
	file_dist_proto_ids_proto_rawDescData = file_dist_proto_ids_proto_rawDesc
)

func file_dist_proto_ids_proto_rawDescGZIP() []byte {
	file_dist_proto_ids_proto_rawDescOnce.Do(func() {
		file_dist_proto_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_dist_proto_ids_proto_rawDescData)
	})
	return file_dist_proto_ids_proto_rawDescData
}

var file_dist_proto_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dist_proto_ids_proto_goTypes = []interface{}{
	(*GetIdReq)(nil),         // 0: GetIdReq
	(*GetIdRes)(nil),         // 1: GetIdRes
	(*GetIdInMapReq)(nil),    // 2: GetIdInMapReq
	(*GetIdInMapRes)(nil),    // 3: GetIdInMapRes
	(*GetIdArrReq)(nil),      // 4: GetIdArrReq
	(*GetIdArrRes)(nil),      // 5: GetIdArrRes
	(*GetIdArrInMapReq)(nil), // 6: GetIdArrInMapReq
	(*GetIdArrInMapRes)(nil), // 7: GetIdArrInMapRes
}
var file_dist_proto_ids_proto_depIdxs = []int32{
	0, // 0: IdServices.GetId:input_type -> GetIdReq
	4, // 1: IdServices.GetIdArr:input_type -> GetIdArrReq
	2, // 2: IdInMapServices.GetIdInMap:input_type -> GetIdInMapReq
	6, // 3: IdInMapServices.GetIdArrInMap:input_type -> GetIdArrInMapReq
	1, // 4: IdServices.GetId:output_type -> GetIdRes
	5, // 5: IdServices.GetIdArr:output_type -> GetIdArrRes
	3, // 6: IdInMapServices.GetIdInMap:output_type -> GetIdInMapRes
	7, // 7: IdInMapServices.GetIdArrInMap:output_type -> GetIdArrInMapRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dist_proto_ids_proto_init() }
func file_dist_proto_ids_proto_init() {
	if File_dist_proto_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dist_proto_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdReq); i {
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
		file_dist_proto_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdRes); i {
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
		file_dist_proto_ids_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdInMapReq); i {
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
		file_dist_proto_ids_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdInMapRes); i {
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
		file_dist_proto_ids_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdArrReq); i {
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
		file_dist_proto_ids_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdArrRes); i {
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
		file_dist_proto_ids_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdArrInMapReq); i {
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
		file_dist_proto_ids_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdArrInMapRes); i {
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
			RawDescriptor: file_dist_proto_ids_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_dist_proto_ids_proto_goTypes,
		DependencyIndexes: file_dist_proto_ids_proto_depIdxs,
		MessageInfos:      file_dist_proto_ids_proto_msgTypes,
	}.Build()
	File_dist_proto_ids_proto = out.File
	file_dist_proto_ids_proto_rawDesc = nil
	file_dist_proto_ids_proto_goTypes = nil
	file_dist_proto_ids_proto_depIdxs = nil
}
