// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/pb/transaction.proto

package corepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Transaction struct {
	ID                   []byte      `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Vin                  []*TXInput  `protobuf:"bytes,2,rep,name=Vin,proto3" json:"Vin,omitempty"`
	Vout                 []*TXOutput `protobuf:"bytes,3,rep,name=Vout,proto3" json:"Vout,omitempty"`
	Tip                  uint64      `protobuf:"varint,4,opt,name=Tip,proto3" json:"Tip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_2e1ba3440064bc94, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Transaction) GetVin() []*TXInput {
	if m != nil {
		return m.Vin
	}
	return nil
}

func (m *Transaction) GetVout() []*TXOutput {
	if m != nil {
		return m.Vout
	}
	return nil
}

func (m *Transaction) GetTip() uint64 {
	if m != nil {
		return m.Tip
	}
	return 0
}

type TXInput struct {
	Txid                 []byte   `protobuf:"bytes,1,opt,name=Txid,proto3" json:"Txid,omitempty"`
	Vout                 int32    `protobuf:"varint,2,opt,name=Vout,proto3" json:"Vout,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	PubKey               []byte   `protobuf:"bytes,4,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TXInput) Reset()         { *m = TXInput{} }
func (m *TXInput) String() string { return proto.CompactTextString(m) }
func (*TXInput) ProtoMessage()    {}
func (*TXInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_2e1ba3440064bc94, []int{1}
}
func (m *TXInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXInput.Unmarshal(m, b)
}
func (m *TXInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXInput.Marshal(b, m, deterministic)
}
func (dst *TXInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXInput.Merge(dst, src)
}
func (m *TXInput) XXX_Size() int {
	return xxx_messageInfo_TXInput.Size(m)
}
func (m *TXInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TXInput.DiscardUnknown(m)
}

var xxx_messageInfo_TXInput proto.InternalMessageInfo

func (m *TXInput) GetTxid() []byte {
	if m != nil {
		return m.Txid
	}
	return nil
}

func (m *TXInput) GetVout() int32 {
	if m != nil {
		return m.Vout
	}
	return 0
}

func (m *TXInput) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TXInput) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type TXOutput struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	PubKeyHash           []byte   `protobuf:"bytes,2,opt,name=PubKeyHash,proto3" json:"PubKeyHash,omitempty"`
	Contract             string   `protobuf:"bytes,3,opt,name=Contract,proto3" json:"Contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TXOutput) Reset()         { *m = TXOutput{} }
func (m *TXOutput) String() string { return proto.CompactTextString(m) }
func (*TXOutput) ProtoMessage()    {}
func (*TXOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_2e1ba3440064bc94, []int{2}
}
func (m *TXOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXOutput.Unmarshal(m, b)
}
func (m *TXOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXOutput.Marshal(b, m, deterministic)
}
func (dst *TXOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXOutput.Merge(dst, src)
}
func (m *TXOutput) XXX_Size() int {
	return xxx_messageInfo_TXOutput.Size(m)
}
func (m *TXOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TXOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TXOutput proto.InternalMessageInfo

func (m *TXOutput) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *TXOutput) GetPubKeyHash() []byte {
	if m != nil {
		return m.PubKeyHash
	}
	return nil
}

func (m *TXOutput) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*TXInput)(nil), "corepb.TXInput")
	proto.RegisterType((*TXOutput)(nil), "corepb.TXOutput")
}

func init() {
	proto.RegisterFile("core/pb/transaction.proto", fileDescriptor_transaction_2e1ba3440064bc94)
}

var fileDescriptor_transaction_2e1ba3440064bc94 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xe9, 0xcf, 0xf4, 0x9b, 0x39, 0x53, 0x3e, 0x87, 0x83, 0x48, 0x14, 0x91, 0x5a, 0x5c,
	0x74, 0xd5, 0x01, 0xbd, 0x04, 0x67, 0x61, 0x71, 0xa1, 0xc4, 0x52, 0x5c, 0xb8, 0x49, 0x6b, 0x19,
	0x03, 0x92, 0x84, 0x4c, 0x02, 0xe3, 0xdd, 0x4b, 0xd2, 0x68, 0xdd, 0x9d, 0xf7, 0xa7, 0xef, 0x53,
	0x02, 0xe7, 0x83, 0xd4, 0xe3, 0x56, 0xf5, 0x5b, 0xa3, 0x99, 0x38, 0xb0, 0xc1, 0x70, 0x29, 0x6a,
	0xa5, 0xa5, 0x91, 0x98, 0xb9, 0x48, 0xf5, 0xe5, 0x11, 0xd6, 0xed, 0x1c, 0xe2, 0x7f, 0x88, 0x9b,
	0x1d, 0x89, 0x8a, 0xa8, 0xca, 0x69, 0xdc, 0xec, 0xf0, 0x1a, 0x92, 0x8e, 0x0b, 0x12, 0x17, 0x49,
	0xb5, 0xbe, 0x3d, 0xa9, 0xa7, 0x8f, 0xea, 0xf6, 0xb5, 0x11, 0xca, 0x1a, 0xea, 0x32, 0xbc, 0x81,
	0xb4, 0x93, 0xd6, 0x90, 0xc4, 0x77, 0x36, 0x73, 0xe7, 0xc9, 0x1a, 0x57, 0xf2, 0x29, 0x6e, 0x20,
	0x69, 0xb9, 0x22, 0x69, 0x11, 0x55, 0x29, 0x75, 0x67, 0xb9, 0x87, 0x7f, 0x61, 0x07, 0x11, 0xd2,
	0xf6, 0xc8, 0xdf, 0x03, 0xd7, 0xdf, 0xce, 0xf3, 0xb3, 0x71, 0x11, 0x55, 0x8b, 0x30, 0x72, 0x09,
	0xab, 0x17, 0xbe, 0x17, 0xcc, 0x58, 0x3d, 0x92, 0xc4, 0x97, 0x67, 0x03, 0xcf, 0x20, 0x7b, 0xb6,
	0xfd, 0xe3, 0xf8, 0xe5, 0x29, 0x39, 0x0d, 0xaa, 0x7c, 0x83, 0xe5, 0xcf, 0xcf, 0xe0, 0x29, 0x2c,
	0x3a, 0xf6, 0x69, 0xc7, 0x80, 0x9a, 0x04, 0x5e, 0x01, 0x4c, 0xdd, 0x07, 0x76, 0xf8, 0xf0, 0xc4,
	0x9c, 0xfe, 0x71, 0xf0, 0x02, 0x96, 0xf7, 0x52, 0x18, 0xcd, 0x06, 0xe3, 0xb1, 0x2b, 0xfa, 0xab,
	0xfb, 0xcc, 0xbf, 0xe7, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x93, 0x36, 0x7c, 0x6c,
	0x01, 0x00, 0x00,
}
