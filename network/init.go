package network

import (
	"log"
	"net"
	"strconv"
)

var (
	peerChanel  = make(chan net.Conn)
	activePeers = make(map[net.Conn]bool)
	activeIPs   = []string{}
)

// Init connections and discovery here
func Init(port int, peers []string) {

	selfAddress := "127.0.0.1:" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			peerChanel <- conn
		}
	}()

	if len(peers) > 0 {
		go func() {
			for _, peer := range peers {
				if len(peer) > 0 {
					clientConn, err := net.Dial("tcp", peer)
					if err != nil {
						log.Println("Peer disconnected ->", clientConn.RemoteAddr())
						clientConn.Close()
						return
					}
					_, err = clientConn.Write([]byte(selfAddress))
					go read(clientConn, selfAddress, &activeIPs, &activePeers)
				}
			}
		}()
	}

	for {
		select {
		case conn := <-peerChanel:
			activePeers[conn] = true
			go read(conn, selfAddress, &activeIPs, &activePeers)
			go write(conn, selfAddress)
		}
	}

}
