package core

import (
	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/core/pb"
	"github.com/golang/protobuf/proto"
)

type TransactionNode struct {
	Children map[string]*Transaction
	Value    *Transaction
	Size     int
}

func NewTransactionNode(tx *Transaction) *TransactionNode {
	txNode := &TransactionNode{Children: make(map[string]*Transaction)}

	if tx == nil {
		return txNode
	}

	size := tx.GetSize()
	if size == 0 {
		return txNode
	}

	txNode.Value = tx
	txNode.Size = size

	return txNode
}

func (txNode *TransactionNode) GetTipsPerByte() *common.Amount {
	return txNode.Value.Tip.Times(uint64(100000)).Div(uint64(txNode.Size))
}

func (txNode *TransactionNode) ToProto() proto.Message {
	childrenProto := make(map[string]*corepb.Transaction)
	for key, val := range txNode.Children {
		childrenProto[key] = val.ToProto().(*corepb.Transaction)
	}
	return &corepb.TransactionNode{
		Children: childrenProto,
		Value:    txNode.Value.ToProto().(*corepb.Transaction),
		Size:     int64(txNode.Size),
	}
}

func (txNode *TransactionNode) FromProto(pb proto.Message) {
	for key, val := range pb.(*corepb.TransactionNode).Children {
		tx := &Transaction{}
		tx.FromProto(val)
		txNode.Children[key] = tx
	}
	tx := &Transaction{}
	tx.FromProto(pb.(*corepb.TransactionNode).Value)
	txNode.Value = tx
	txNode.Size = int(pb.(*corepb.TransactionNode).Size)
}
