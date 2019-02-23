package routines

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// ReadFromPeer is here
func ReadFromPeer(conn net.Conn, selfAddress string, activeIPs *[]string) {
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		msg, err := conn.Read(data)
		if err != nil {
			log.Println("Peer disconnected ->", conn.RemoteAddr())
			conn.Close()
			return
		}
		s := string(data[:msg])
		log.Println(conn.RemoteAddr(), "-> @@", s)
		*activeIPs = append(*activeIPs, s)
		if len(*activeIPs) > 1 {
			fmt.Println("Announcing to peers-->", activeIPs)
			for _, ip := range *activeIPs {
				fmt.Println(ip)
			}
		}
	}
}

// Read is here
func Read(conn net.Conn, selfAddress string, activeIPs *[]string, activePeers *map[net.Conn]bool) {
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		msg, err := conn.Read(data)
		if err != nil {
			log.Println("Peer disconnected ->", conn.RemoteAddr())
			conn.Close()
			return
		}
		s := string(data[:msg])
		fmt.Println(conn.RemoteAddr(), "->", s)

		*activeIPs = append(*activeIPs, s)
		if len(*activeIPs) > 1 {
			fmt.Println("Announcing to peers", *activeIPs)
			for peerConn := range *activePeers {
				peerConn.Write([]byte(strings.Join(*activeIPs, ",")))
			}
		}
	}
}

// Write from connections
func Write(conn net.Conn, selfAddress string) {

}
