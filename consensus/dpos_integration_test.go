// +build integration

// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//

package consensus

import (
	"testing"
	"time"

	"github.com/dappley/go-dappley/core/block_producer_info"

	"github.com/dappley/go-dappley/logic/blockchain_logic"
	"github.com/dappley/go-dappley/logic/transaction_pool"

	"github.com/dappley/go-dappley/core"
	"github.com/dappley/go-dappley/core/account"
	"github.com/dappley/go-dappley/network"
	"github.com/dappley/go-dappley/storage"
	"github.com/dappley/go-dappley/util"
	"github.com/stretchr/testify/assert"
)

func TestDpos_Start(t *testing.T) {
	cbAddr := account.NewAddress("dPGZmHd73UpZhrM6uvgnzu49ttbLp4AzU8")
	keystr := "5a66b0fdb69c99935783059bb200e86e97b506ae443a62febd7d0750cd7fac55"

	producer := block_producer_info.NewBlockProducerInfo(cbAddr.String())
	dpos := NewDPOS(producer)
	dpos.SetKey(keystr)

	miners := []string{cbAddr.String()}
	dynasty := NewDynasty(miners, 2, 2)
	dpos.SetDynasty(dynasty)

	dpos.Start()
	//wait for all producer gets a chance to produce
	time.Sleep(time.Second * 2 * 2)
	dpos.Stop()

	assert.Equal(t, len(dpos.notifierCh), 1)
}

func TestDpos_MultipleMiners(t *testing.T) {
	const (
		timeBetweenBlock = 2
		dposRounds       = 3
	)

	miners := []string{
		"dPGZmHd73UpZhrM6uvgnzu49ttbLp4AzU8",
		"dQEooMsqp23RkPsvZXj3XbsRh9BUyGz2S9",
		"dastXXWLe5pxbRYFhcyUq8T3wb5srWkHKa",
		"dUuPPYshbBgkzUrgScEHWvdGbSxC8z4R12",
		"dPGD4t6ibpmyKZnXH1TNbbPw98EDaaZq8C",
	}
	keystrs := []string{
		"5a66b0fdb69c99935783059bb200e86e97b506ae443a62febd7d0750cd7fac55",
		"bb23d2ff19f5b16955e8a24dca34dd520980fe3bddca2b3e1b56663f0ec1aa7e",
		"300c0338c4b0d49edc66113e3584e04c6b907f9ded711d396d522aae6a79be1a",
		"da9282440fae188c371165e01615a2e1b14af68b3eaae51e6608c0bd86d4e6a6",
		"7c918ed7660d55759b7fc42b25f26bdab3caf8fc07586b2659a26470fb8dfc69",
	}
	dynasty := NewDynasty(miners, len(miners), timeBetweenBlock)
	var dposArray []*DPOS
	var nodeArray []*network.Node

	for i, miner := range miners {
		dpos := NewDPOS(miner)
		dpos.SetKey(keystrs[i])
		dpos.SetDynasty(dynasty)
		bc := blockchain_logic.CreateBlockchain(account.NewAddress(miners[0]), storage.NewRamStorage(), dpos, transaction_pool.NewTransactionPool(nil, 128), nil, 100000)
		pool := core.NewBlockPool()

		node := network.NewNode(bc.GetDb(), nil)
		node.Start(21200+i, "")
		nodeArray = append(nodeArray, node)

		bm := blockchain_logic.NewBlockchainManager(bc, pool, node)

		dpos.Setup(miner, bm)
		dpos.SetKey(keystrs[i])
		dposArray = append(dposArray, dpos)
	}

	for i := range miners {
		for j := range miners {
			if i != j {
				nodeArray[i].GetNetwork().ConnectToSeed(nodeArray[j].GetHostPeerInfo())
			}
		}

		dposArray[i].Start()
	}

	time.Sleep(time.Second*time.Duration(dynasty.GetDynastyTime()*dposRounds) + time.Second/2)

	for i := range miners {
		dposArray[i].Stop()
		nodeArray[i].Stop()
	}
	//Waiting block sync to other nodes
	time.Sleep(time.Second * 2)
	for i := range miners {
		v := dposArray[i]
		util.WaitDoneOrTimeout(func() bool {
			return !v.IsProducingBlock()
		}, 20)
	}

	for i := range miners {
		assert.Equal(t, uint64(dynasty.dynastyTime*dposRounds/timeBetweenBlock), dposArray[i].bm.Getblockchain().GetMaxHeight())
	}
}

func TestDPOS_UpdateLIB(t *testing.T) {
	const (
		timeBetweenBlock = 2
		dposRounds       = 3
	)

	miners := []string{
		"dPGZmHd73UpZhrM6uvgnzu49ttbLp4AzU8",
		"dQEooMsqp23RkPsvZXj3XbsRh9BUyGz2S9",
		"dastXXWLe5pxbRYFhcyUq8T3wb5srWkHKa",
		"dUuPPYshbBgkzUrgScEHWvdGbSxC8z4R12",
		"dPGD4t6ibpmyKZnXH1TNbbPw98EDaaZq8C",
	}
	keystrs := []string{
		"5a66b0fdb69c99935783059bb200e86e97b506ae443a62febd7d0750cd7fac55",
		"bb23d2ff19f5b16955e8a24dca34dd520980fe3bddca2b3e1b56663f0ec1aa7e",
		"300c0338c4b0d49edc66113e3584e04c6b907f9ded711d396d522aae6a79be1a",
		"da9282440fae188c371165e01615a2e1b14af68b3eaae51e6608c0bd86d4e6a6",
		"7c918ed7660d55759b7fc42b25f26bdab3caf8fc07586b2659a26470fb8dfc69",
	}
	dynasty := NewDynasty(miners, len(miners), timeBetweenBlock)

	var dposArray []*DPOS
	var nodeArray []*network.Node

	for i, miner := range miners {
		dpos := NewDPOS()
		dpos.SetDynasty(dynasty)
		bc := blockchain_logic.CreateBlockchain(account.NewAddress(miners[0]), storage.NewRamStorage(), dpos, transaction_pool.NewTransactionPool(nil, 128), nil, 100000)
		pool := core.NewBlockPool()

		node := network.NewNode(bc.GetDb(), nil)
		node.Start(22200+i, "")
		nodeArray = append(nodeArray, node)

		bm := blockchain_logic.NewBlockchainManager(bc, pool, node)

		dpos.Setup(miner, bm)
		dpos.SetKey(keystrs[i])
		dposArray = append(dposArray, dpos)
	}

	for i := range miners {
		for j := range miners {
			if i != j {
				nodeArray[i].GetNetwork().ConnectToSeed(nodeArray[j].GetHostPeerInfo())
			}
		}

		dposArray[i].Start()
	}

	time.Sleep(time.Second*time.Duration(dynasty.dynastyTime*dposRounds) + time.Second/2)

	for i := range miners {
		dposArray[i].Stop()
		nodeArray[i].Stop()
	}

	//Waiting block sync to other nodes
	time.Sleep(time.Second * 2)
	for i := range miners {
		v := dposArray[i]
		util.WaitDoneOrTimeout(func() bool {
			return !v.IsProducingBlock()
		}, 20)
	}

	block0, _ := dposArray[0].bm.Getblockchain().GetLIB()
	assert.NotEqual(t, 0, block0.GetHeight())

	for i := range miners {
		block, _ := dposArray[i].bm.Getblockchain().GetLIB()
		assert.Equal(t, block0.GetHash(), block.GetHash())
	}
}
