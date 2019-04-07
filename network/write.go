package network

import (
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/peerfind/protos"
)

func (peer *Peer) write(msgChannel chan []byte) {
	for {
		msg := <-msgChannel
		_msg := &protos.Arc{}
		err := proto.Unmarshal(msg, _msg)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}

		switch _msg.GetType() {

		case "handshake":
			log.Println(" this is handshake")
			peer.conn.Write(msg)

		default:
			log.Println("Default message")

		}
	}

}
