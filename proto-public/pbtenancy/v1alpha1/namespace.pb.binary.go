// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbtenancy/v1alpha1/namespace.proto

package tenancyv1alpha1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *Namespace) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *Namespace) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}