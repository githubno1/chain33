// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// merkle avl tree
type LeafNode struct {
	Key    []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Height int32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Size   int32  `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *LeafNode) Reset()                    { *m = LeafNode{} }
func (m *LeafNode) String() string            { return proto.CompactTextString(m) }
func (*LeafNode) ProtoMessage()               {}
func (*LeafNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *LeafNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LeafNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LeafNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *LeafNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type InnerNode struct {
	LeftHash  []byte `protobuf:"bytes,1,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash []byte `protobuf:"bytes,2,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height    int32  `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Size      int32  `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *InnerNode) Reset()                    { *m = InnerNode{} }
func (m *InnerNode) String() string            { return proto.CompactTextString(m) }
func (*InnerNode) ProtoMessage()               {}
func (*InnerNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *InnerNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *InnerNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *InnerNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *InnerNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MAVLProof struct {
	LeafHash   []byte       `protobuf:"bytes,1,opt,name=leafHash,proto3" json:"leafHash,omitempty"`
	InnerNodes []*InnerNode `protobuf:"bytes,2,rep,name=innerNodes" json:"innerNodes,omitempty"`
	RootHash   []byte       `protobuf:"bytes,3,opt,name=rootHash,proto3" json:"rootHash,omitempty"`
}

func (m *MAVLProof) Reset()                    { *m = MAVLProof{} }
func (m *MAVLProof) String() string            { return proto.CompactTextString(m) }
func (*MAVLProof) ProtoMessage()               {}
func (*MAVLProof) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *MAVLProof) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

func (m *MAVLProof) GetInnerNodes() []*InnerNode {
	if m != nil {
		return m.InnerNodes
	}
	return nil
}

func (m *MAVLProof) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

type StoreNode struct {
	Key       []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	LeftHash  []byte `protobuf:"bytes,3,opt,name=leftHash,proto3" json:"leftHash,omitempty"`
	RightHash []byte `protobuf:"bytes,4,opt,name=rightHash,proto3" json:"rightHash,omitempty"`
	Height    int32  `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	Size      int32  `protobuf:"varint,6,opt,name=size" json:"size,omitempty"`
}

func (m *StoreNode) Reset()                    { *m = StoreNode{} }
func (m *StoreNode) String() string            { return proto.CompactTextString(m) }
func (*StoreNode) ProtoMessage()               {}
func (*StoreNode) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *StoreNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *StoreNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StoreNode) GetLeftHash() []byte {
	if m != nil {
		return m.LeftHash
	}
	return nil
}

func (m *StoreNode) GetRightHash() []byte {
	if m != nil {
		return m.RightHash
	}
	return nil
}

func (m *StoreNode) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StoreNode) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type LocalDBSet struct {
	KV []*KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
}

func (m *LocalDBSet) Reset()                    { *m = LocalDBSet{} }
func (m *LocalDBSet) String() string            { return proto.CompactTextString(m) }
func (*LocalDBSet) ProtoMessage()               {}
func (*LocalDBSet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *LocalDBSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type LocalDBList struct {
	Prefix    []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Key       []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
	Count     int32  `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
}

func (m *LocalDBList) Reset()                    { *m = LocalDBList{} }
func (m *LocalDBList) String() string            { return proto.CompactTextString(m) }
func (*LocalDBList) ProtoMessage()               {}
func (*LocalDBList) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *LocalDBList) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *LocalDBList) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalDBList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *LocalDBList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type LocalDBGet struct {
	Keys [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *LocalDBGet) Reset()                    { *m = LocalDBGet{} }
func (m *LocalDBGet) String() string            { return proto.CompactTextString(m) }
func (*LocalDBGet) ProtoMessage()               {}
func (*LocalDBGet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *LocalDBGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type LocalReplyValue struct {
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *LocalReplyValue) Reset()                    { *m = LocalReplyValue{} }
func (m *LocalReplyValue) String() string            { return proto.CompactTextString(m) }
func (*LocalReplyValue) ProtoMessage()               {}
func (*LocalReplyValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *LocalReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type StoreSet struct {
	StateHash []byte      `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	KV        []*KeyValue `protobuf:"bytes,2,rep,name=KV" json:"KV,omitempty"`
}

func (m *StoreSet) Reset()                    { *m = StoreSet{} }
func (m *StoreSet) String() string            { return proto.CompactTextString(m) }
func (*StoreSet) ProtoMessage()               {}
func (*StoreSet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *StoreSet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreSet) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type StoreGet struct {
	StateHash []byte   `protobuf:"bytes,1,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Keys      [][]byte `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (m *StoreGet) Reset()                    { *m = StoreGet{} }
func (m *StoreGet) String() string            { return proto.CompactTextString(m) }
func (*StoreGet) ProtoMessage()               {}
func (*StoreGet) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *StoreGet) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *StoreGet) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type StoreReplyValue struct {
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (m *StoreReplyValue) Reset()                    { *m = StoreReplyValue{} }
func (m *StoreReplyValue) String() string            { return proto.CompactTextString(m) }
func (*StoreReplyValue) ProtoMessage()               {}
func (*StoreReplyValue) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *StoreReplyValue) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*LeafNode)(nil), "types.LeafNode")
	proto.RegisterType((*InnerNode)(nil), "types.InnerNode")
	proto.RegisterType((*MAVLProof)(nil), "types.MAVLProof")
	proto.RegisterType((*StoreNode)(nil), "types.StoreNode")
	proto.RegisterType((*LocalDBSet)(nil), "types.LocalDBSet")
	proto.RegisterType((*LocalDBList)(nil), "types.LocalDBList")
	proto.RegisterType((*LocalDBGet)(nil), "types.LocalDBGet")
	proto.RegisterType((*LocalReplyValue)(nil), "types.LocalReplyValue")
	proto.RegisterType((*StoreSet)(nil), "types.StoreSet")
	proto.RegisterType((*StoreGet)(nil), "types.StoreGet")
	proto.RegisterType((*StoreReplyValue)(nil), "types.StoreReplyValue")
}

func init() { proto.RegisterFile("db.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0xab, 0xda, 0x40,
	0x14, 0x25, 0x89, 0x91, 0xe4, 0x2a, 0x28, 0x43, 0x29, 0x41, 0x84, 0x86, 0xac, 0xec, 0xa2, 0x52,
	0xda, 0x6d, 0x37, 0x2d, 0x05, 0x2b, 0xa6, 0xa5, 0x44, 0xc8, 0xb2, 0x10, 0xe3, 0x4d, 0x0d, 0xc6,
	0x4c, 0x3a, 0x19, 0x4b, 0xd3, 0x7f, 0xf2, 0xfe, 0xed, 0x63, 0x3e, 0x62, 0xf2, 0xe0, 0xf9, 0xc4,
	0xdd, 0x9c, 0xe3, 0xf5, 0x9e, 0x8f, 0x4b, 0xc0, 0xd9, 0xef, 0x96, 0x15, 0xa3, 0x9c, 0x12, 0x9b,
	0x37, 0x15, 0xd6, 0xb3, 0x71, 0x4a, 0x4f, 0x27, 0x5a, 0x2a, 0x32, 0xf8, 0x05, 0x4e, 0x88, 0x49,
	0xf6, 0x83, 0xee, 0x91, 0x4c, 0xc1, 0x3a, 0x62, 0xe3, 0x19, 0xbe, 0xb1, 0x18, 0x47, 0xe2, 0x49,
	0x5e, 0x81, 0xfd, 0x37, 0x29, 0xce, 0xe8, 0x99, 0x92, 0x53, 0x80, 0xbc, 0x86, 0xe1, 0x01, 0xf3,
	0xdf, 0x07, 0xee, 0x59, 0xbe, 0xb1, 0xb0, 0x23, 0x8d, 0x08, 0x81, 0x41, 0x9d, 0xff, 0x47, 0x6f,
	0x20, 0x59, 0xf9, 0x0e, 0xfe, 0x80, 0xbb, 0x2e, 0x4b, 0x64, 0x52, 0x60, 0x06, 0x4e, 0x81, 0x19,
	0xff, 0x96, 0xd4, 0x07, 0xad, 0x72, 0xc1, 0x64, 0x0e, 0x2e, 0x13, 0x5b, 0xe4, 0x8f, 0x4a, 0xae,
	0x23, 0xee, 0x92, 0x3c, 0x83, 0xfb, 0xfd, 0x73, 0x1c, 0xfe, 0x64, 0x94, 0x66, 0x4a, 0x32, 0xc9,
	0x9e, 0x4a, 0x2a, 0x4c, 0xde, 0x03, 0xe4, 0xad, 0xb7, 0xda, 0x33, 0x7d, 0x6b, 0x31, 0xfa, 0x30,
	0x5d, 0xca, 0x96, 0x96, 0x17, 0xd3, 0x51, 0x6f, 0x46, 0x6c, 0x63, 0x94, 0x2a, 0x8f, 0x96, 0xda,
	0xd6, 0xe2, 0xe0, 0xc1, 0x00, 0x77, 0xcb, 0x29, 0xc3, 0xbb, 0xba, 0xec, 0x57, 0x62, 0xbd, 0x54,
	0xc9, 0xe0, 0x7a, 0x25, 0xf6, 0xb3, 0x95, 0x0c, 0x7b, 0x95, 0xbc, 0x03, 0x08, 0x69, 0x9a, 0x14,
	0x5f, 0xbf, 0x6c, 0x91, 0x93, 0x37, 0x60, 0x6e, 0x62, 0x9d, 0x77, 0xa2, 0xf3, 0x6e, 0xb0, 0x89,
	0x85, 0xa1, 0xc8, 0xdc, 0xc4, 0xc1, 0x11, 0x46, 0x7a, 0x3c, 0xcc, 0x6b, 0x2e, 0x94, 0x2a, 0x86,
	0x59, 0xfe, 0x4f, 0xc7, 0xd1, 0xa8, 0xcd, 0x68, 0x76, 0x19, 0xe7, 0xe0, 0xee, 0x73, 0x86, 0x29,
	0xcf, 0x69, 0xa9, 0x2f, 0xd5, 0x11, 0xa2, 0x81, 0x94, 0x9e, 0x4b, 0xae, 0xaf, 0xa5, 0x40, 0xe0,
	0x5f, 0xbc, 0xad, 0x50, 0xba, 0x3f, 0x62, 0xa3, 0xae, 0x31, 0x8e, 0xe4, 0x3b, 0x78, 0x0b, 0x13,
	0x39, 0x11, 0x61, 0x55, 0x28, 0x97, 0xc2, 0x92, 0xec, 0xaf, 0x1d, 0xd4, 0x28, 0x58, 0x83, 0x23,
	0x6f, 0x20, 0x62, 0xce, 0xc1, 0xad, 0x79, 0xc2, 0xb1, 0x77, 0xfb, 0x8e, 0xb8, 0x5d, 0xc2, 0x27,
	0xbd, 0x6a, 0x75, 0x73, 0xd5, 0x15, 0xcf, 0xf2, 0xdf, 0xb7, 0x3d, 0xef, 0x86, 0xf2, 0x4b, 0xfc,
	0xf8, 0x18, 0x00, 0x00, 0xff, 0xff, 0x24, 0x96, 0x7a, 0xe3, 0xaa, 0x03, 0x00, 0x00,
}