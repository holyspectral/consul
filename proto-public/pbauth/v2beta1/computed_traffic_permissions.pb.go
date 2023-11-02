// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbauth/v2beta1/computed_traffic_permissions.proto

package authv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type ComputedTrafficPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowPermissions []*Permission `protobuf:"bytes,1,rep,name=allow_permissions,json=allowPermissions,proto3" json:"allow_permissions,omitempty"`
	DenyPermissions  []*Permission `protobuf:"bytes,2,rep,name=deny_permissions,json=denyPermissions,proto3" json:"deny_permissions,omitempty"`
	IsDefault        bool          `protobuf:"varint,3,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
}

func (x *ComputedTrafficPermissions) Reset() {
	*x = ComputedTrafficPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbauth_v2beta1_computed_traffic_permissions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedTrafficPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedTrafficPermissions) ProtoMessage() {}

func (x *ComputedTrafficPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_pbauth_v2beta1_computed_traffic_permissions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedTrafficPermissions.ProtoReflect.Descriptor instead.
func (*ComputedTrafficPermissions) Descriptor() ([]byte, []int) {
	return file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescGZIP(), []int{0}
}

func (x *ComputedTrafficPermissions) GetAllowPermissions() []*Permission {
	if x != nil {
		return x.AllowPermissions
	}
	return nil
}

func (x *ComputedTrafficPermissions) GetDenyPermissions() []*Permission {
	if x != nil {
		return x.DenyPermissions
	}
	return nil
}

func (x *ComputedTrafficPermissions) GetIsDefault() bool {
	if x != nil {
		return x.IsDefault
	}
	return false
}

var File_pbauth_v2beta1_computed_traffic_permissions_proto protoreflect.FileDescriptor

var file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x62, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x1a, 0x28, 0x70, 0x62, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x1a, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x10, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x54, 0x0a, 0x10, 0x64, 0x65, 0x6e, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x64, 0x65, 0x6e, 0x79, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x42, 0xa0,
	0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x42, 0x1f, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2f, 0x70, 0x62, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48,
	0x43, 0x41, 0xaa, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xe2, 0x02, 0x29, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescOnce sync.Once
	file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescData = file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDesc
)

func file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescGZIP() []byte {
	file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescOnce.Do(func() {
		file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescData)
	})
	return file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDescData
}

var file_pbauth_v2beta1_computed_traffic_permissions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbauth_v2beta1_computed_traffic_permissions_proto_goTypes = []interface{}{
	(*ComputedTrafficPermissions)(nil), // 0: hashicorp.consul.auth.v2beta1.ComputedTrafficPermissions
	(*Permission)(nil),                 // 1: hashicorp.consul.auth.v2beta1.Permission
}
var file_pbauth_v2beta1_computed_traffic_permissions_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.auth.v2beta1.ComputedTrafficPermissions.allow_permissions:type_name -> hashicorp.consul.auth.v2beta1.Permission
	1, // 1: hashicorp.consul.auth.v2beta1.ComputedTrafficPermissions.deny_permissions:type_name -> hashicorp.consul.auth.v2beta1.Permission
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbauth_v2beta1_computed_traffic_permissions_proto_init() }
func file_pbauth_v2beta1_computed_traffic_permissions_proto_init() {
	if File_pbauth_v2beta1_computed_traffic_permissions_proto != nil {
		return
	}
	file_pbauth_v2beta1_traffic_permissions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pbauth_v2beta1_computed_traffic_permissions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedTrafficPermissions); i {
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
			RawDescriptor: file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbauth_v2beta1_computed_traffic_permissions_proto_goTypes,
		DependencyIndexes: file_pbauth_v2beta1_computed_traffic_permissions_proto_depIdxs,
		MessageInfos:      file_pbauth_v2beta1_computed_traffic_permissions_proto_msgTypes,
	}.Build()
	File_pbauth_v2beta1_computed_traffic_permissions_proto = out.File
	file_pbauth_v2beta1_computed_traffic_permissions_proto_rawDesc = nil
	file_pbauth_v2beta1_computed_traffic_permissions_proto_goTypes = nil
	file_pbauth_v2beta1_computed_traffic_permissions_proto_depIdxs = nil
}
