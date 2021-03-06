// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: github.com/dappley/go-dappley/tool/db_inspect/pb/pretty_block.proto

package db_inspect_pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Vin      []*TXInput  `protobuf:"bytes,2,rep,name=vin,proto3" json:"vin,omitempty"`
	Vout     []*TXOutput `protobuf:"bytes,3,rep,name=vout,proto3" json:"vout,omitempty"`
	Tip      string      `protobuf:"bytes,4,opt,name=tip,proto3" json:"tip,omitempty"`
	GasLimit string      `protobuf:"bytes,5,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasPrice string      `protobuf:"bytes,6,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transaction) GetVin() []*TXInput {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *Transaction) GetVout() []*TXOutput {
	if x != nil {
		return x.Vout
	}
	return nil
}

func (x *Transaction) GetTip() string {
	if x != nil {
		return x.Tip
	}
	return ""
}

func (x *Transaction) GetGasLimit() string {
	if x != nil {
		return x.GasLimit
	}
	return ""
}

func (x *Transaction) GetGasPrice() string {
	if x != nil {
		return x.GasPrice
	}
	return ""
}

type TXInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Txid      string `protobuf:"bytes,1,opt,name=txid,proto3" json:"txid,omitempty"`
	Vout      int32  `protobuf:"varint,2,opt,name=vout,proto3" json:"vout,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PublicKey string `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *TXInput) Reset() {
	*x = TXInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXInput) ProtoMessage() {}

func (x *TXInput) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXInput.ProtoReflect.Descriptor instead.
func (*TXInput) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{1}
}

func (x *TXInput) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *TXInput) GetVout() int32 {
	if x != nil {
		return x.Vout
	}
	return 0
}

func (x *TXInput) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *TXInput) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type TXOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value         string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	PublicKeyHash string `protobuf:"bytes,2,opt,name=public_key_hash,json=publicKeyHash,proto3" json:"public_key_hash,omitempty"`
	Contract      string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *TXOutput) Reset() {
	*x = TXOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TXOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TXOutput) ProtoMessage() {}

func (x *TXOutput) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TXOutput.ProtoReflect.Descriptor instead.
func (*TXOutput) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{2}
}

func (x *TXOutput) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TXOutput) GetPublicKeyHash() string {
	if x != nil {
		return x.PublicKeyHash
	}
	return ""
}

func (x *TXOutput) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{3}
}

func (x *Block) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Block) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash         string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PreviousHash string `protobuf:"bytes,2,opt,name=previous_hash,json=previousHash,proto3" json:"previous_hash,omitempty"`
	Nonce        int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp    int64  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature    string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Height       uint64 `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{4}
}

func (x *BlockHeader) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockHeader) GetPreviousHash() string {
	if x != nil {
		return x.PreviousHash
	}
	return ""
}

func (x *BlockHeader) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *BlockHeader) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *BlockHeader) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *BlockHeader) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type Utxo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount        string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	PublicKeyHash string `protobuf:"bytes,2,opt,name=public_key_hash,json=publicKeyHash,proto3" json:"public_key_hash,omitempty"`
	Txid          string `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	TxIndex       uint32 `protobuf:"varint,4,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
}

func (x *Utxo) Reset() {
	*x = Utxo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Utxo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Utxo) ProtoMessage() {}

func (x *Utxo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Utxo.ProtoReflect.Descriptor instead.
func (*Utxo) Descriptor() ([]byte, []int) {
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP(), []int{5}
}

func (x *Utxo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Utxo) GetPublicKeyHash() string {
	if x != nil {
		return x.PublicKeyHash
	}
	return ""
}

func (x *Utxo) GetTxid() string {
	if x != nil {
		return x.Txid
	}
	return ""
}

func (x *Utxo) GetTxIndex() uint32 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

var File_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto protoreflect.FileDescriptor

var file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDesc = []byte{
	0x0a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x70,
	0x70, 0x6c, 0x65, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x61, 0x70, 0x70, 0x6c, 0x65, 0x79, 0x2f,
	0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x2f,
	0x70, 0x62, 0x2f, 0x70, 0x72, 0x65, 0x74, 0x74, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x5f, 0x70, 0x62, 0x22, 0xbe, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x70,
	0x62, 0x2e, 0x54, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x2b,
	0x0a, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x58, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x70, 0x12, 0x1a, 0x0a,
	0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x73,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x6e, 0x0a, 0x07, 0x54, 0x58, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x78, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x64, 0x0a, 0x08, 0x54, 0x58, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x7b, 0x0a, 0x05, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x5f, 0x70, 0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x64, 0x62, 0x5f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x75, 0x0a, 0x04, 0x55,
	0x74, 0x78, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x78, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x78, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescOnce sync.Once
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescData = file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDesc
)

func file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescGZIP() []byte {
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescOnce.Do(func() {
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescData)
	})
	return file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDescData
}

var file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_goTypes = []interface{}{
	(*Transaction)(nil), // 0: db_inspect_pb.Transaction
	(*TXInput)(nil),     // 1: db_inspect_pb.TXInput
	(*TXOutput)(nil),    // 2: db_inspect_pb.TXOutput
	(*Block)(nil),       // 3: db_inspect_pb.Block
	(*BlockHeader)(nil), // 4: db_inspect_pb.BlockHeader
	(*Utxo)(nil),        // 5: db_inspect_pb.Utxo
}
var file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_depIdxs = []int32{
	1, // 0: db_inspect_pb.Transaction.vin:type_name -> db_inspect_pb.TXInput
	2, // 1: db_inspect_pb.Transaction.vout:type_name -> db_inspect_pb.TXOutput
	4, // 2: db_inspect_pb.Block.header:type_name -> db_inspect_pb.BlockHeader
	0, // 3: db_inspect_pb.Block.transactions:type_name -> db_inspect_pb.Transaction
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_init() }
func file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_init() {
	if File_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXInput); i {
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
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TXOutput); i {
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
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
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
		file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Utxo); i {
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
			RawDescriptor: file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_goTypes,
		DependencyIndexes: file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_depIdxs,
		MessageInfos:      file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_msgTypes,
	}.Build()
	File_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto = out.File
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_rawDesc = nil
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_goTypes = nil
	file_github_com_dappley_go_dappley_tool_db_inspect_pb_pretty_block_proto_depIdxs = nil
}
