// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: userPlatformModels.proto

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

type UserPlatformModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: json:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @inject_tag: json:"openid"
	Openid string `protobuf:"bytes,3,opt,name=Openid,proto3" json:"Openid,omitempty"`
	// @inject_tag: json:"token"
	Token string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	// @inject_tag: json:"platform"
	Platform string `protobuf:"bytes,5,opt,name=Platform,proto3" json:"Platform,omitempty"`
	// @inject_tag: json:"created_at"
	CreatedAt string `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	// @inject_tag: json:"updated_at"
	UpdatedAt string `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *UserPlatformModel) Reset() {
	*x = UserPlatformModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userPlatformModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPlatformModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPlatformModel) ProtoMessage() {}

func (x *UserPlatformModel) ProtoReflect() protoreflect.Message {
	mi := &file_userPlatformModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPlatformModel.ProtoReflect.Descriptor instead.
func (*UserPlatformModel) Descriptor() ([]byte, []int) {
	return file_userPlatformModels_proto_rawDescGZIP(), []int{0}
}

func (x *UserPlatformModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserPlatformModel) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserPlatformModel) GetOpenid() string {
	if x != nil {
		return x.Openid
	}
	return ""
}

func (x *UserPlatformModel) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserPlatformModel) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *UserPlatformModel) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserPlatformModel) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_userPlatformModels_proto protoreflect.FileDescriptor

var file_userPlatformModels_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xc1,
	0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x70, 0x65, 0x6e, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x70,
	0x65, 0x6e, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userPlatformModels_proto_rawDescOnce sync.Once
	file_userPlatformModels_proto_rawDescData = file_userPlatformModels_proto_rawDesc
)

func file_userPlatformModels_proto_rawDescGZIP() []byte {
	file_userPlatformModels_proto_rawDescOnce.Do(func() {
		file_userPlatformModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_userPlatformModels_proto_rawDescData)
	})
	return file_userPlatformModels_proto_rawDescData
}

var file_userPlatformModels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_userPlatformModels_proto_goTypes = []interface{}{
	(*UserPlatformModel)(nil), // 0: pb.UserPlatformModel
}
var file_userPlatformModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_userPlatformModels_proto_init() }
func file_userPlatformModels_proto_init() {
	if File_userPlatformModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userPlatformModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPlatformModel); i {
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
			RawDescriptor: file_userPlatformModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userPlatformModels_proto_goTypes,
		DependencyIndexes: file_userPlatformModels_proto_depIdxs,
		MessageInfos:      file_userPlatformModels_proto_msgTypes,
	}.Build()
	File_userPlatformModels_proto = out.File
	file_userPlatformModels_proto_rawDesc = nil
	file_userPlatformModels_proto_goTypes = nil
	file_userPlatformModels_proto_depIdxs = nil
}
