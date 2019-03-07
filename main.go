package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/sjljrvis/peerfind/network"
)

// Test is export
var (
	peerChanel  = make(chan net.Conn)
	activePeers = make(map[net.Conn]bool)
	activeIPs   = []string{}
)

func init() {
	log.Println("Starting peer discovery")
}

func main() {

	portFlag := flag.Int("port", 3000, "Port to connect")
	peersFlag := flag.String("peers", "", "list of peers")
	flag.Parse()

	port := *portFlag
	peers := strings.Split(*peersFlag, ",")

	log.Println("Peers can connect to address -> ", "127.0.0.1:", port)

	network.Init(port, peers)

}
