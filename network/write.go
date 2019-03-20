package network

import (
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/peerfind/protos"
)

func (peer *Peer) write() {

	msg := &protos.HandShake{
		IpAddr: peer.conn.LocalAddr().String(),
		Type:   "handshake",
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Println("Marshalling Error", err)
	}
	log.Println(data)
	peer.conn.Write(data)

}
