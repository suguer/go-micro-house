// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: smsService.proto

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

type SmsIndexRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" form:"id" uri:"id"
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: json:"keyword" form:"keyword" uri:"keyword"
	Keyword string `protobuf:"bytes,2,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
	// @inject_tag: json:"status" form:"status" uri:"status"
	Status string `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	// @inject_tag: json:"current" form:"current" uri:"current"
	Current uint32 `protobuf:"varint,4,opt,name=Current,proto3" json:"Current,omitempty"`
	// @inject_tag: json:"pageSize" form:"pageSize" uri:"pageSize"
	PageSize uint32 `protobuf:"varint,5,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	// @inject_tag: json:"user_id" form:"user_id" uri:"user_id"
	UserId uint32 `protobuf:"varint,6,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *SmsIndexRequest) Reset() {
	*x = SmsIndexRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsIndexRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsIndexRequest) ProtoMessage() {}

func (x *SmsIndexRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsIndexRequest.ProtoReflect.Descriptor instead.
func (*SmsIndexRequest) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{0}
}

func (x *SmsIndexRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SmsIndexRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SmsIndexRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SmsIndexRequest) GetCurrent() uint32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *SmsIndexRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SmsIndexRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SmsIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data       []*SmsRecordModel `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	Pagination *Pagination       `protobuf:"bytes,2,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
}

func (x *SmsIndexResponse) Reset() {
	*x = SmsIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsIndexResponse) ProtoMessage() {}

func (x *SmsIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsIndexResponse.ProtoReflect.Descriptor instead.
func (*SmsIndexResponse) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{1}
}

func (x *SmsIndexResponse) GetData() []*SmsRecordModel {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SmsIndexResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"content" form:"content" uri:"content"
	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	// @inject_tag: json:"user_id" form:"user_id" uri:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *SmsRequest) Reset() {
	*x = SmsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsRequest) ProtoMessage() {}

func (x *SmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsRequest.ProtoReflect.Descriptor instead.
func (*SmsRequest) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{2}
}

func (x *SmsRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SmsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SmsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *SmsResponse) Reset() {
	*x = SmsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smsService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsResponse) ProtoMessage() {}

func (x *SmsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smsService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsResponse.ProtoReflect.Descriptor instead.
func (*SmsResponse) Descriptor() ([]byte, []int) {
	return file_smsService_proto_rawDescGZIP(), []int{3}
}

func (x *SmsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_smsService_proto protoreflect.FileDescriptor

var file_smsService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x14, 0x73, 0x6d, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1,
	0x01, 0x0a, 0x0f, 0x53, 0x6d, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x6a, 0x0a, 0x10, 0x53, 0x6d, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e,
	0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x21,
	0x0a, 0x0b, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x32, 0x6b, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6d,
	0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x6d, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b,
	0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_smsService_proto_rawDescOnce sync.Once
	file_smsService_proto_rawDescData = file_smsService_proto_rawDesc
)

func file_smsService_proto_rawDescGZIP() []byte {
	file_smsService_proto_rawDescOnce.Do(func() {
		file_smsService_proto_rawDescData = protoimpl.X.CompressGZIP(file_smsService_proto_rawDescData)
	})
	return file_smsService_proto_rawDescData
}

var file_smsService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_smsService_proto_goTypes = []interface{}{
	(*SmsIndexRequest)(nil),  // 0: pb.SmsIndexRequest
	(*SmsIndexResponse)(nil), // 1: pb.SmsIndexResponse
	(*SmsRequest)(nil),       // 2: pb.SmsRequest
	(*SmsResponse)(nil),      // 3: pb.SmsResponse
	(*SmsRecordModel)(nil),   // 4: pb.SmsRecordModel
	(*Pagination)(nil),       // 5: pb.Pagination
}
var file_smsService_proto_depIdxs = []int32{
	4, // 0: pb.SmsIndexResponse.Data:type_name -> pb.SmsRecordModel
	5, // 1: pb.SmsIndexResponse.Pagination:type_name -> pb.Pagination
	0, // 2: pb.SmsService.Index:input_type -> pb.SmsIndexRequest
	2, // 3: pb.SmsService.Create:input_type -> pb.SmsRequest
	1, // 4: pb.SmsService.Index:output_type -> pb.SmsIndexResponse
	3, // 5: pb.SmsService.Create:output_type -> pb.SmsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_smsService_proto_init() }
func file_smsService_proto_init() {
	if File_smsService_proto != nil {
		return
	}
	file_smsRecordModel_proto_init()
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_smsService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsIndexRequest); i {
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
		file_smsService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsIndexResponse); i {
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
		file_smsService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsRequest); i {
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
		file_smsService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsResponse); i {
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
			RawDescriptor: file_smsService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smsService_proto_goTypes,
		DependencyIndexes: file_smsService_proto_depIdxs,
		MessageInfos:      file_smsService_proto_msgTypes,
	}.Build()
	File_smsService_proto = out.File
	file_smsService_proto_rawDesc = nil
	file_smsService_proto_goTypes = nil
	file_smsService_proto_depIdxs = nil
}