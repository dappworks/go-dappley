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

package main

import (
	"flag"
	"log"

	"github.com/dappley/go-dappley/config"
	"github.com/dappley/go-dappley/consensus"
	"github.com/dappley/go-dappley/core"
	"github.com/dappley/go-dappley/logic"
	"github.com/dappley/go-dappley/network"
	"github.com/dappley/go-dappley/rpc"
	"github.com/dappley/go-dappley/storage"
	logger "github.com/sirupsen/logrus"
)

const (
	genesisAddr     = "121yKAXeG4cw6uaGCBYjWk9yTWmMkhcoDD"
	configFilePath  = "conf/default.conf"
	genesisFilePath = "conf/genesis.conf"
	defaultPassword = "password"
)

func main() {

	var filePath string
	flag.StringVar(&filePath, "f", configFilePath, "Configuration File Path. Default to conf/default.conf")
	flag.Parse()

	logger.SetLevel(logger.DebugLevel)

	//load genesis file information
	genesisConf := config.LoadConfigFromFile(genesisFilePath)
	if genesisConf == nil {
		logger.Error("ERROR: Cannot load genesis configurations from file!Exiting...")
		return
	}
	//load config file information
	conf := config.LoadConfigFromFile(filePath)
	if conf == nil {
		logger.Error("ERROR: Cannot load configurations from file!Exiting...")
		return
	}

	//setup
	db := storage.OpenDatabase(conf.GetNodeConfig().GetDbPath())
	defer db.Close()

	//creat blockchain
	conss, _ := initConsensus(genesisConf)
	bc, err := core.GetBlockchain(db, conss)
	if err != nil {
		bc, err = logic.CreateBlockchain(core.Address{genesisAddr}, db, conss)
		if err != nil {
			log.Panic(err)
		}
	}

	node, err := initNode(conf, bc)
	if err != nil {
		return
	}

	//start rpc server
	server := rpc.NewGrpcServer(node, defaultPassword)
	server.Start(conf.GetNodeConfig().GetRpcPort())
	defer server.Stop()

	//start mining
	minerAddr := conf.GetConsensusConfig().GetMinerAddr()
	conss.Setup(node, minerAddr)
	conss.SetKey(conf.GetConsensusConfig().GetMinerPrivKey())
	logger.Info("Miner Address is ", minerAddr)

	conss.Start()
	defer conss.Stop()

	select {}
}

func initConsensus(conf *config.Config) (core.Consensus, *consensus.Dynasty) {
	//set up consensus
	conss := consensus.NewDpos()
	dynasty := consensus.NewDynastyWithProducers(conf.GetDynastyConfig().GetProducers())
	conss.SetDynasty(dynasty)
	conss.SetTargetBit(0)
	return conss, dynasty
}

func initNode(conf *config.Config, bc *core.Blockchain) (*network.Node, error) {
	//create node
	node := network.NewNode(bc)
	nodeConfig := conf.GetNodeConfig()
	port := nodeConfig.GetListeningPort()
	err := node.Start(int(port))
	if err != nil {
		logger.Error(err)
		logger.Error("ERROR: Invalid Port!Exiting...")
		return nil, err
	}
	seed := nodeConfig.GetSeed()
	if seed != "" {
		node.AddStreamByString(seed)
	}
	return node, nil
}
