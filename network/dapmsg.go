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

package network

import (
	logger "github.com/sirupsen/logrus"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/dappley/go-dappley/network/pb"
)

type DapMsg struct {
	cmd            string
	data           []byte
	unixTimeRecvd  int64
	key            string
	uniOrBroadcast int ``
}

func NewDapmsg(cmd string, data []byte, msgKey string, uniOrBroadcast int) *DapMsg {
	return &DapMsg{cmd, data, time.Now().Unix(), msgKey, uniOrBroadcast}
}

func (dm *DapMsg) GetCmd() string {
	return dm.cmd
}

func (dm *DapMsg) GetData() []byte {
	return dm.data
}

func (dm *DapMsg) GetTimestamp() int64 {
	return dm.unixTimeRecvd
}

func (dm *DapMsg) GetFrom() string {
	return dm.key
}

//used to lookup dapmsg cache (key:unix time of command + command in string, value: 1 if received recently, 0 if not).
func (dm *DapMsg) GetKey() string {
	return dm.key
}

func ParseDappMsgFromDappPacket(packet *DappPacket) *DapMsg {
	return ParseDappMsgFromRawBytes(packet.GetData())
}

func ParseDappMsgFromRawBytes(bytes []byte) *DapMsg {
	dmpb := &networkpb.Dapmsg{}

	//unmarshal byte to proto
	if err := proto.Unmarshal(bytes, dmpb); err != nil {
		logger.WithError(err).Warn("Stream: Unable to")
	}

	dm := &DapMsg{}
	dm.FromProto(dmpb)
	return dm
}

func (dm *DapMsg) ToProto() proto.Message {
	return &networkpb.Dapmsg{
		Cmd:              dm.cmd,
		Data:             dm.data,
		UnixTimeReceived: dm.unixTimeRecvd,
		Key:              dm.key,
	}
}

func (dm *DapMsg) FromProto(pb proto.Message) {
	dm.cmd = pb.(*networkpb.Dapmsg).GetCmd()
	dm.data = pb.(*networkpb.Dapmsg).GetData()
	dm.unixTimeRecvd = pb.(*networkpb.Dapmsg).GetUnixTimeReceived()
	dm.key = pb.(*networkpb.Dapmsg).GetKey()

}
