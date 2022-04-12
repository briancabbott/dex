// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: server/internal/types.proto

// Package internal holds protobuf types used by the server.

package internal

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

// RefreshToken is a message that holds refresh token data used by dex.
type RefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshId string `protobuf:"bytes,1,opt,name=refresh_id,json=refreshId,proto3" json:"refresh_id,omitempty"`
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *RefreshToken) Reset() {
	*x = RefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_internal_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshToken) ProtoMessage() {}

func (x *RefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_server_internal_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshToken.ProtoReflect.Descriptor instead.
func (*RefreshToken) Descriptor() ([]byte, []int) {
	return file_server_internal_types_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshToken) GetRefreshId() string {
	if x != nil {
		return x.RefreshId
	}
	return ""
}

func (x *RefreshToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// IDTokenSubject represents both the userID and connID which is returned
// as the "sub" claim in the ID Token.
type IDTokenSubject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ConnId string `protobuf:"bytes,2,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
}

func (x *IDTokenSubject) Reset() {
	*x = IDTokenSubject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_internal_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDTokenSubject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDTokenSubject) ProtoMessage() {}

func (x *IDTokenSubject) ProtoReflect() protoreflect.Message {
	mi := &file_server_internal_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDTokenSubject.ProtoReflect.Descriptor instead.
func (*IDTokenSubject) Descriptor() ([]byte, []int) {
	return file_server_internal_types_proto_rawDescGZIP(), []int{1}
}

func (x *IDTokenSubject) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IDTokenSubject) GetConnId() string {
	if x != nil {
		return x.ConnId
	}
	return ""
}

var File_server_internal_types_proto protoreflect.FileDescriptor

var file_server_internal_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x0e,
	0x49, 0x44, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x72, 0x69, 0x61, 0x6e, 0x63, 0x61, 0x62, 0x62, 0x6f, 0x74, 0x74, 0x2f, 0x64, 0x65, 0x78, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_internal_types_proto_rawDescOnce sync.Once
	file_server_internal_types_proto_rawDescData = file_server_internal_types_proto_rawDesc
)

func file_server_internal_types_proto_rawDescGZIP() []byte {
	file_server_internal_types_proto_rawDescOnce.Do(func() {
		file_server_internal_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_internal_types_proto_rawDescData)
	})
	return file_server_internal_types_proto_rawDescData
}

var file_server_internal_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_internal_types_proto_goTypes = []interface{}{
	(*RefreshToken)(nil),   // 0: internal.RefreshToken
	(*IDTokenSubject)(nil), // 1: internal.IDTokenSubject
}
var file_server_internal_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_internal_types_proto_init() }
func file_server_internal_types_proto_init() {
	if File_server_internal_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_internal_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshToken); i {
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
		file_server_internal_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDTokenSubject); i {
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
			RawDescriptor: file_server_internal_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_internal_types_proto_goTypes,
		DependencyIndexes: file_server_internal_types_proto_depIdxs,
		MessageInfos:      file_server_internal_types_proto_msgTypes,
	}.Build()
	File_server_internal_types_proto = out.File
	file_server_internal_types_proto_rawDesc = nil
	file_server_internal_types_proto_goTypes = nil
	file_server_internal_types_proto_depIdxs = nil
}
