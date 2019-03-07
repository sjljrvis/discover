package network

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func read(conn net.Conn, selfAddress string, activeIPs *[]string, activePeers *map[net.Conn]bool) {
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
