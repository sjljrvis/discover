package main

import (
	"flag"
	"log"
	"strings"

	"github.com/sjljrvis/peerfind/network"
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
