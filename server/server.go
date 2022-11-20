package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	// 1. start an upd server
	listener, err := net.ListenUDP("udp", &net.UDPAddr{Port: 9527})
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// 2. receive messages from A and B
	peers := make([]*net.UDPAddr, 2, 2)
	buf := make([]byte, 256)
	n, addr, err := listener.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("recevie from %s: %s\n", addr.String(), buf[:n])
	peers[0] = addr

	n, addr, err = listener.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("recevie from %s: %s\n", addr.String(), buf[:n])
	peers[1] = addr

	// 3. exchange messages
	fmt.Printf("begin nat\n")
	_, _ = listener.WriteToUDP([]byte(peers[0].String()), peers[1])
	_, _ = listener.WriteToUDP([]byte(peers[1].String()), peers[0])

	// 4. server can exit
	timer := time.NewTimer(time.Second * 10)
	select {
	case <-timer.C:
		break
	}
}
