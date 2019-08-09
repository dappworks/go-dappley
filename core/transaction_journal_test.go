// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either pubKeyHash 3 of the License, or
// (at your option) any later pubKeyHash.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//
package core

import (
	"github.com/dappley/go-dappley/common"
	"github.com/dappley/go-dappley/core/transaction"
	"github.com/dappley/go-dappley/core/transaction_base"
	"github.com/dappley/go-dappley/util"
	"testing"

	"github.com/dappley/go-dappley/storage"
	"github.com/stretchr/testify/assert"
)

var tx1 = transaction.Transaction{
	ID:       util.GenerateRandomAoB(1),
	Vin:      transaction_base.GenerateFakeTxInputs(),
	Vout:     transaction_base.GenerateFakeTxOutputs(),
	Tip:      common.NewAmount(5),
	GasLimit: common.NewAmount(0),
	GasPrice: common.NewAmount(0),
}

func TestJournalPutAndGet(t *testing.T) {
	db := storage.NewRamStorage()
	vin := transaction_base.TXInput{tx1.ID, 1, nil, nil}
	vout, err := GetTxOutput(vin, db)
	// Expect transaction logs have been successfully saved
	assert.Nil(t, err)
	assert.Equal(t, vout.PubKeyHash, tx1.Vout[1].PubKeyHash)
}
