package network

import (
	"log"
	"net"
	"strconv"
)

// Peer Struct
type Peer struct {
	conn net.Conn
	msg  chan []byte
}

var (
	peerChannel = make(chan net.Conn)
	register    = make(chan *Peer)
	activePeers = make(map[net.Conn]bool)
	msgChannel  = make(chan []byte)
	activeIPs   = []string{}
)

func connector(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		peerChannel <- conn
	}
}

// Init connections and discovery here
func Init(port int, peers []string) {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))

	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}

	go connector(listener)

	if len(peers) > 0 {
		go func() {
			for _, peer := range peers {
				if len(peer) > 0 {
					clientConn, err := net.Dial("tcp", peer)
					if err != nil {
						log.Println("Peer disconnected ->", clientConn.RemoteAddr())
						clientConn.Close()
						activePeers[clientConn] = false
						return
					}
					peerChannel <- clientConn
				}
			}
		}()
	}

	for {
		select {
		case conn := <-peerChannel:
			activePeers[conn] = true
			peer := &Peer{conn: conn}
			register <- peer
			go peer.read()
			go peer.write()
		}
	}

}
