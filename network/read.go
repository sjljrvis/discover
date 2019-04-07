package network

import (
	"log"

	"github.com/gogo/protobuf/proto"
	protos "github.com/sjljrvis/peerfind/protos"
)

// func read(conn net.Conn, selfAddress string, activeIPs *[]string, activePeers *map[net.Conn]bool) {
// 	defer conn.Close()
// 	data := make([]byte, 1024)
// 	for {
// 		msg, err := conn.Read(data)
// 		if err != nil {
// 			log.Println("Peer disconnected ->", conn.RemoteAddr())
// 			conn.Close()
// 			return
// 		}
// 		log.Println(msg)
// 		// s := string(data[:msg])
// 		// fmt.Println(conn.RemoteAddr(), "->", s)

// 		// *activeIPs = append(*activeIPs, s)
// 		// if len(*activeIPs) > 1 {
// 		// 	fmt.Println("Announcing to peers", *activeIPs)
// 		// 	for peerConn := range *activePeers {
// 		// 		peerConn.Write([]byte(strings.Join(*activeIPs, ",")))
// 		// 	}
// 		// }
// 	}
// }

func (peer *Peer) read(msgChannel chan []byte) {
	defer peer.conn.Close()
	data := make([]byte, 1024)
	for {
		len, err := peer.conn.Read(data)
		if err != nil {
			log.Println("Peer disconnected ->", peer.conn.RemoteAddr())
			peer.conn.Close()
			return
		}
		msg := &protos.Arc{}
		err = proto.Unmarshal(data[0:len], msg)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		log.Println(">>>>>>", msg.GetType(), msg.GetData())
	}
}
