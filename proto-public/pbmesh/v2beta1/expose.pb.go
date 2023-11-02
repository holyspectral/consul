// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/expose.proto

package meshv2beta1

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

// +kubebuilder:validation:Enum=EXPOSE_PATH_PROTOCOL_HTTP;EXPOSE_PATH_PROTOCOL_HTTP2
// +kubebuilder:validation:Type=string
type ExposePathProtocol int32

const (
	// buf:lint:ignore ENUM_ZERO_VALUE_SUFFIX
	ExposePathProtocol_EXPOSE_PATH_PROTOCOL_HTTP  ExposePathProtocol = 0
	ExposePathProtocol_EXPOSE_PATH_PROTOCOL_HTTP2 ExposePathProtocol = 1
)

// Enum value maps for ExposePathProtocol.
var (
	ExposePathProtocol_name = map[int32]string{
		0: "EXPOSE_PATH_PROTOCOL_HTTP",
		1: "EXPOSE_PATH_PROTOCOL_HTTP2",
	}
	ExposePathProtocol_value = map[string]int32{
		"EXPOSE_PATH_PROTOCOL_HTTP":  0,
		"EXPOSE_PATH_PROTOCOL_HTTP2": 1,
	}
)

func (x ExposePathProtocol) Enum() *ExposePathProtocol {
	p := new(ExposePathProtocol)
	*p = x
	return p
}

func (x ExposePathProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExposePathProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_pbmesh_v2beta1_expose_proto_enumTypes[0].Descriptor()
}

func (ExposePathProtocol) Type() protoreflect.EnumType {
	return &file_pbmesh_v2beta1_expose_proto_enumTypes[0]
}

func (x ExposePathProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExposePathProtocol.Descriptor instead.
func (ExposePathProtocol) EnumDescriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_expose_proto_rawDescGZIP(), []int{0}
}

type ExposeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExposePaths []*ExposePath `protobuf:"bytes,1,rep,name=expose_paths,json=exposePaths,proto3" json:"expose_paths,omitempty"`
}

func (x *ExposeConfig) Reset() {
	*x = ExposeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_expose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposeConfig) ProtoMessage() {}

func (x *ExposeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_expose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposeConfig.ProtoReflect.Descriptor instead.
func (*ExposeConfig) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_expose_proto_rawDescGZIP(), []int{0}
}

func (x *ExposeConfig) GetExposePaths() []*ExposePath {
	if x != nil {
		return x.ExposePaths
	}
	return nil
}

type ExposePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerPort  uint32             `protobuf:"varint,1,opt,name=listener_port,json=listenerPort,proto3" json:"listener_port,omitempty"`
	Path          string             `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	LocalPathPort uint32             `protobuf:"varint,3,opt,name=local_path_port,json=localPathPort,proto3" json:"local_path_port,omitempty"`
	Protocol      ExposePathProtocol `protobuf:"varint,4,opt,name=protocol,proto3,enum=hashicorp.consul.mesh.v2beta1.ExposePathProtocol" json:"protocol,omitempty"`
}

func (x *ExposePath) Reset() {
	*x = ExposePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_expose_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposePath) ProtoMessage() {}

func (x *ExposePath) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_expose_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposePath.ProtoReflect.Descriptor instead.
func (*ExposePath) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_expose_proto_rawDescGZIP(), []int{1}
}

func (x *ExposePath) GetListenerPort() uint32 {
	if x != nil {
		return x.ListenerPort
	}
	return 0
}

func (x *ExposePath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ExposePath) GetLocalPathPort() uint32 {
	if x != nil {
		return x.LocalPathPort
	}
	return 0
}

func (x *ExposePath) GetProtocol() ExposePathProtocol {
	if x != nil {
		return x.Protocol
	}
	return ExposePathProtocol_EXPOSE_PATH_PROTOCOL_HTTP
}

var File_pbmesh_v2beta1_expose_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_expose_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x5c, 0x0a, 0x0c,
	0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x0c,
	0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x0b, 0x65,
	0x78, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0a, 0x45,
	0x78, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x4d, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2a, 0x53, 0x0a, 0x12, 0x45, 0x78, 0x70,
	0x6f, 0x73, 0x65, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x1d, 0x0a, 0x19, 0x45, 0x58, 0x50, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12, 0x1e,
	0x0a, 0x1a, 0x45, 0x58, 0x50, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f, 0x50, 0x52,
	0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10, 0x01, 0x42, 0x8c,
	0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x0b, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x73,
	0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02,
	0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02,
	0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x29, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a,
	0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_expose_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_expose_proto_rawDescData = file_pbmesh_v2beta1_expose_proto_rawDesc
)

func file_pbmesh_v2beta1_expose_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_expose_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_expose_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_expose_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_expose_proto_rawDescData
}

var file_pbmesh_v2beta1_expose_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pbmesh_v2beta1_expose_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbmesh_v2beta1_expose_proto_goTypes = []interface{}{
	(ExposePathProtocol)(0), // 0: hashicorp.consul.mesh.v2beta1.ExposePathProtocol
	(*ExposeConfig)(nil),    // 1: hashicorp.consul.mesh.v2beta1.ExposeConfig
	(*ExposePath)(nil),      // 2: hashicorp.consul.mesh.v2beta1.ExposePath
}
var file_pbmesh_v2beta1_expose_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.mesh.v2beta1.ExposeConfig.expose_paths:type_name -> hashicorp.consul.mesh.v2beta1.ExposePath
	0, // 1: hashicorp.consul.mesh.v2beta1.ExposePath.protocol:type_name -> hashicorp.consul.mesh.v2beta1.ExposePathProtocol
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_expose_proto_init() }
func file_pbmesh_v2beta1_expose_proto_init() {
	if File_pbmesh_v2beta1_expose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_expose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposeConfig); i {
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
		file_pbmesh_v2beta1_expose_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposePath); i {
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
			RawDescriptor: file_pbmesh_v2beta1_expose_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_expose_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_expose_proto_depIdxs,
		EnumInfos:         file_pbmesh_v2beta1_expose_proto_enumTypes,
		MessageInfos:      file_pbmesh_v2beta1_expose_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_expose_proto = out.File
	file_pbmesh_v2beta1_expose_proto_rawDesc = nil
	file_pbmesh_v2beta1_expose_proto_goTypes = nil
	file_pbmesh_v2beta1_expose_proto_depIdxs = nil
}
