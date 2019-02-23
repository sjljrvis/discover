package main

import (
	"flag"
	"log"
	"net"
	"strconv"
	"strings"

	routines "github.com/sjljrvis/peerfind/routines"
)

// Test is export
var (
	peerChanel  = make(chan net.Conn)
	activePeers = make(map[net.Conn]bool)
	activeIPs   = []string{}
)

type handShake struct {
	IP string
}

func init() {
	log.Println("Starting peer discovery")
}

func main() {

	portFlag := flag.Int("port", 3000, "Port to connect")
	peersFlag := flag.String("peers", "", "list of peers")
	flag.Parse()

	port := *portFlag
	peers := strings.Split(*peersFlag, ",")

	selfAddress := "127.0.0.1:" + strconv.Itoa(port)
	log.Println("Peers can connect to address -> ", "127.0.0.1:", port)

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	defer ln.Close()

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			conn, err := ln.Accept()
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
					log.Println("conn ->", clientConn)
					_, err = clientConn.Write([]byte(selfAddress))
					go routines.ReadFromPeer(clientConn, selfAddress, &activeIPs)
				}
			}
		}()
	}

	for {
		select {
		case conn := <-peerChanel:
			activePeers[conn] = true
			go routines.Read(conn, selfAddress, &activeIPs, &activePeers)
			go routines.Write(conn, selfAddress)
		}
	}

}
