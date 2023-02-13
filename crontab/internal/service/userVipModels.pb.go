// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: userVipModels.proto

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

type UserVipModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: json:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @inject_tag: json:"start_at"
	StartAt string `protobuf:"bytes,3,opt,name=StartAt,proto3" json:"StartAt,omitempty"`
	// @inject_tag: json:"end_at"
	EndAt string `protobuf:"bytes,4,opt,name=EndAt,proto3" json:"EndAt,omitempty"`
	// @inject_tag: json:"level"
	Level string `protobuf:"bytes,5,opt,name=Level,proto3" json:"Level,omitempty"`
	// @inject_tag: json:"created_at"
	CreatedAt string `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @inject_tag: json:"updated_at"
	UpdatedAt float32 `protobuf:"fixed32,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	// @inject_tag: json:"status"
	Status string `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UserVipModel) Reset() {
	*x = UserVipModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userVipModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserVipModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVipModel) ProtoMessage() {}

func (x *UserVipModel) ProtoReflect() protoreflect.Message {
	mi := &file_userVipModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVipModel.ProtoReflect.Descriptor instead.
func (*UserVipModel) Descriptor() ([]byte, []int) {
	return file_userVipModels_proto_rawDescGZIP(), []int{0}
}

func (x *UserVipModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserVipModel) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserVipModel) GetStartAt() string {
	if x != nil {
		return x.StartAt
	}
	return ""
}

func (x *UserVipModel) GetEndAt() string {
	if x != nil {
		return x.EndAt
	}
	return ""
}

func (x *UserVipModel) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *UserVipModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserVipModel) GetUpdatedAt() float32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *UserVipModel) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_userVipModels_proto protoreflect.FileDescriptor

var file_userVipModels_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x56, 0x69, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x69, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x64,
	0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1b, 0x5a, 0x19,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_userVipModels_proto_rawDescOnce sync.Once
	file_userVipModels_proto_rawDescData = file_userVipModels_proto_rawDesc
)

func file_userVipModels_proto_rawDescGZIP() []byte {
	file_userVipModels_proto_rawDescOnce.Do(func() {
		file_userVipModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_userVipModels_proto_rawDescData)
	})
	return file_userVipModels_proto_rawDescData
}

var file_userVipModels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_userVipModels_proto_goTypes = []interface{}{
	(*UserVipModel)(nil), // 0: pb.UserVipModel
}
var file_userVipModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userVipModels_proto_init() }
func file_userVipModels_proto_init() {
	if File_userVipModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userVipModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserVipModel); i {
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
			RawDescriptor: file_userVipModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userVipModels_proto_goTypes,
		DependencyIndexes: file_userVipModels_proto_depIdxs,
		MessageInfos:      file_userVipModels_proto_msgTypes,
	}.Build()
	File_userVipModels_proto = out.File
	file_userVipModels_proto_rawDesc = nil
	file_userVipModels_proto_goTypes = nil
	file_userVipModels_proto_depIdxs = nil
}
