// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: sdk/logical/plugin.proto

package logical

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PluginEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VaultVersion is the version of the Vault server
	VaultVersion string `protobuf:"bytes,1,opt,name=vault_version,json=vaultVersion,proto3" json:"vault_version,omitempty"`
	// VaultVersionPrerelease is the prerelease information of the Vault server
	VaultVersionPrerelease string `protobuf:"bytes,2,opt,name=vault_version_prerelease,json=vaultVersionPrerelease,proto3" json:"vault_version_prerelease,omitempty"`
	// VaultVersionMetadata is the version metadata of the Vault server
	VaultVersionMetadata string `protobuf:"bytes,3,opt,name=vault_version_metadata,json=vaultVersionMetadata,proto3" json:"vault_version_metadata,omitempty"`
	// VaultBuildDate is the build date of the Vault server
	VaultBuildDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=vault_build_date,json=vaultBuildDate,proto3" json:"vault_build_date,omitempty"`
	// VaultBuiltinPublicKeys contains the builtin public keys of the Vault server
	VaultBuiltinPublicKeys []string `protobuf:"bytes,5,rep,name=vault_builtin_public_keys,json=vaultBuiltinPublicKeys,proto3" json:"vault_builtin_public_keys,omitempty"`
}

func (x *PluginEnvironment) Reset() {
	*x = PluginEnvironment{}
	mi := &file_sdk_logical_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginEnvironment) ProtoMessage() {}

func (x *PluginEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_logical_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginEnvironment.ProtoReflect.Descriptor instead.
func (*PluginEnvironment) Descriptor() ([]byte, []int) {
	return file_sdk_logical_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginEnvironment) GetVaultVersion() string {
	if x != nil {
		return x.VaultVersion
	}
	return ""
}

func (x *PluginEnvironment) GetVaultVersionPrerelease() string {
	if x != nil {
		return x.VaultVersionPrerelease
	}
	return ""
}

func (x *PluginEnvironment) GetVaultVersionMetadata() string {
	if x != nil {
		return x.VaultVersionMetadata
	}
	return ""
}

func (x *PluginEnvironment) GetVaultBuildDate() *timestamppb.Timestamp {
	if x != nil {
		return x.VaultBuildDate
	}
	return nil
}

func (x *PluginEnvironment) GetVaultBuiltinPublicKeys() []string {
	if x != nil {
		return x.VaultBuiltinPublicKeys
	}
	return nil
}

var File_sdk_logical_plugin_proto protoreflect.FileDescriptor

var file_sdk_logical_plugin_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x6f, 0x67, 0x69,
	0x63, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x11, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x18, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x16, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x44, 0x0a, 0x10, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x42,
	0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sdk_logical_plugin_proto_rawDescOnce sync.Once
	file_sdk_logical_plugin_proto_rawDescData = file_sdk_logical_plugin_proto_rawDesc
)

func file_sdk_logical_plugin_proto_rawDescGZIP() []byte {
	file_sdk_logical_plugin_proto_rawDescOnce.Do(func() {
		file_sdk_logical_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_logical_plugin_proto_rawDescData)
	})
	return file_sdk_logical_plugin_proto_rawDescData
}

var file_sdk_logical_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sdk_logical_plugin_proto_goTypes = []any{
	(*PluginEnvironment)(nil),     // 0: logical.PluginEnvironment
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_sdk_logical_plugin_proto_depIdxs = []int32{
	1, // 0: logical.PluginEnvironment.vault_build_date:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sdk_logical_plugin_proto_init() }
func file_sdk_logical_plugin_proto_init() {
	if File_sdk_logical_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sdk_logical_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sdk_logical_plugin_proto_goTypes,
		DependencyIndexes: file_sdk_logical_plugin_proto_depIdxs,
		MessageInfos:      file_sdk_logical_plugin_proto_msgTypes,
	}.Build()
	File_sdk_logical_plugin_proto = out.File
	file_sdk_logical_plugin_proto_rawDesc = nil
	file_sdk_logical_plugin_proto_goTypes = nil
	file_sdk_logical_plugin_proto_depIdxs = nil
}
