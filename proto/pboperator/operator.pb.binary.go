// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: proto/pboperator/operator.proto

package pboperator

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *TransferLeaderRequest) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *TransferLeaderRequest) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *TransferLeaderResponse) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *TransferLeaderResponse) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
