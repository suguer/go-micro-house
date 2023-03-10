// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: contractService.proto

package service

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

type ContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"data" form:"data" uri:"data"
	Data *ContractModel `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	// @inject_tag: json:"level" form:"level" uri:"level"
	Level string `protobuf:"bytes,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *ContactRequest) Reset() {
	*x = ContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRequest) ProtoMessage() {}

func (x *ContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contractService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRequest.ProtoReflect.Descriptor instead.
func (*ContactRequest) Descriptor() ([]byte, []int) {
	return file_contractService_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRequest) GetData() *ContractModel {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ContactRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type ContactDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ContractModel `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Code int32          `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *ContactDetailResponse) Reset() {
	*x = ContactDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactDetailResponse) ProtoMessage() {}

func (x *ContactDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contractService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactDetailResponse.ProtoReflect.Descriptor instead.
func (*ContactDetailResponse) Descriptor() ([]byte, []int) {
	return file_contractService_proto_rawDescGZIP(), []int{1}
}

func (x *ContactDetailResponse) GetData() *ContractModel {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ContactDetailResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type ContractIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id" uri:"id"
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: json:"type" form:"type" uri:"type"
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
	// @inject_tag: json:"keyword" form:"keyword" uri:"keyword"
	Keyword string `protobuf:"bytes,3,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
	// @inject_tag: json:"status" form:"status" uri:"status"
	Status string `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"group_id" form:"group_id" uri:"group_id"
	GroupId uint32 `protobuf:"varint,5,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	// @inject_tag: json:"current" form:"current" uri:"current"
	Current uint32 `protobuf:"varint,6,opt,name=Current,proto3" json:"Current,omitempty"`
	// @inject_tag: json:"pageSize" form:"pageSize" uri:"pageSize"
	PageSize uint32 `protobuf:"varint,7,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	// @inject_tag: json:"user_id" form:"user_id" uri:"user_id"
	UserId uint32 `protobuf:"varint,8,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *ContractIndexRequest) Reset() {
	*x = ContractIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractIndexRequest) ProtoMessage() {}

func (x *ContractIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contractService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractIndexRequest.ProtoReflect.Descriptor instead.
func (*ContractIndexRequest) Descriptor() ([]byte, []int) {
	return file_contractService_proto_rawDescGZIP(), []int{2}
}

func (x *ContractIndexRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContractIndexRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ContractIndexRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ContractIndexRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContractIndexRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ContractIndexRequest) GetCurrent() uint32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ContractIndexRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ContractIndexRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type ContractIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*ContractModel `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	Pagination *Pagination      `protobuf:"bytes,2,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *ContractIndexResponse) Reset() {
	*x = ContractIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contractService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractIndexResponse) ProtoMessage() {}

func (x *ContractIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contractService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractIndexResponse.ProtoReflect.Descriptor instead.
func (*ContractIndexResponse) Descriptor() ([]byte, []int) {
	return file_contractService_proto_rawDescGZIP(), []int{3}
}

func (x *ContractIndexResponse) GetData() []*ContractModel {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ContractIndexResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_contractService_proto protoreflect.FileDescriptor

var file_contractService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x52, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6e, 0x0a, 0x15,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x0a,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb5, 0x02, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contractService_proto_rawDescOnce sync.Once
	file_contractService_proto_rawDescData = file_contractService_proto_rawDesc
)

func file_contractService_proto_rawDescGZIP() []byte {
	file_contractService_proto_rawDescOnce.Do(func() {
		file_contractService_proto_rawDescData = protoimpl.X.CompressGZIP(file_contractService_proto_rawDescData)
	})
	return file_contractService_proto_rawDescData
}

var file_contractService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contractService_proto_goTypes = []interface{}{
	(*ContactRequest)(nil),        // 0: pb.ContactRequest
	(*ContactDetailResponse)(nil), // 1: pb.ContactDetailResponse
	(*ContractIndexRequest)(nil),  // 2: pb.ContractIndexRequest
	(*ContractIndexResponse)(nil), // 3: pb.ContractIndexResponse
	(*ContractModel)(nil),         // 4: pb.ContractModel
	(*Pagination)(nil),            // 5: pb.Pagination
}
var file_contractService_proto_depIdxs = []int32{
	4, // 0: pb.ContactRequest.Data:type_name -> pb.ContractModel
	4, // 1: pb.ContactDetailResponse.Data:type_name -> pb.ContractModel
	4, // 2: pb.ContractIndexResponse.Data:type_name -> pb.ContractModel
	5, // 3: pb.ContractIndexResponse.Pagination:type_name -> pb.Pagination
	2, // 4: pb.ContractService.Index:input_type -> pb.ContractIndexRequest
	0, // 5: pb.ContractService.Create:input_type -> pb.ContactRequest
	0, // 6: pb.ContractService.Instance:input_type -> pb.ContactRequest
	0, // 7: pb.ContractService.Update:input_type -> pb.ContactRequest
	0, // 8: pb.ContractService.Delete:input_type -> pb.ContactRequest
	3, // 9: pb.ContractService.Index:output_type -> pb.ContractIndexResponse
	1, // 10: pb.ContractService.Create:output_type -> pb.ContactDetailResponse
	1, // 11: pb.ContractService.Instance:output_type -> pb.ContactDetailResponse
	1, // 12: pb.ContractService.Update:output_type -> pb.ContactDetailResponse
	1, // 13: pb.ContractService.Delete:output_type -> pb.ContactDetailResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_contractService_proto_init() }
func file_contractService_proto_init() {
	if File_contractService_proto != nil {
		return
	}
	file_contractModel_proto_init()
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contractService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRequest); i {
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
		file_contractService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactDetailResponse); i {
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
		file_contractService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractIndexRequest); i {
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
		file_contractService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractIndexResponse); i {
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
			RawDescriptor: file_contractService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contractService_proto_goTypes,
		DependencyIndexes: file_contractService_proto_depIdxs,
		MessageInfos:      file_contractService_proto_msgTypes,
	}.Build()
	File_contractService_proto = out.File
	file_contractService_proto_rawDesc = nil
	file_contractService_proto_goTypes = nil
	file_contractService_proto_depIdxs = nil
}
