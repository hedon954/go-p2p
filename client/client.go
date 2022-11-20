package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("./client tag remoteIP remotePort port")
		return
	}

	// parse args
	tag := os.Args[1]
	remoteIP := os.Args[2]
	remotePort, _ := strconv.Atoi(os.Args[3])
	port, _ := strconv.Atoi(os.Args[4])

	// bind port
	localAddr := net.UDPAddr{Port: port}

	// communicate with server
	conn, err := net.DialUDP("udp", &localAddr,
		&net.UDPAddr{IP: net.ParseIP(remoteIP), Port: remotePort})
	if err != nil {
		log.Panic("failed to DialUDP", err)
	}

	// send message to server
	_, _ = conn.Write([]byte("I'm peer: " + tag))

	// get target address from server
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic("failed to ReadFromUPD", err)
	}
	targetAddr := parseAddr(string(buf[:n]))
	fmt.Println("get target addr: ", targetAddr)

	// enter p2p network
	_ = conn.Close()
	p2p(&localAddr, &targetAddr)
}

// parseAddr parses upd address from format ip:port
func parseAddr(add string) net.UDPAddr {
	split := strings.Split(add, ":")
	port, _ := strconv.Atoi(split[1])
	return net.UDPAddr{
		IP:   net.ParseIP(split[0]),
		Port: port,
	}
}

// p2p is a p2p network without server
func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {

	// connect with dst
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println("failed to DialUDP", err)
		return
	}
	defer conn.Close()

	// send messages
	fmt.Println("start to communicate")
	if _, err = conn.Write([]byte("hole\n")); err != nil {
		fmt.Println("send msg err", err)
		return
	}

	// start a goroutine to monitor dst's messages
	go func() {
		buf := make([]byte, 256)
		for {
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("connect finished")
					break
				}
				fmt.Println("err occurs", err)
			}
			if n > 0 {
				fmt.Printf("get message: %s\n", buf[:n])
			}
		}
	}()

	// monitor standard input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("p2p> ")
		data, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic("failed to read from standard input")
		}
		_, err = conn.Write([]byte(data))
		if err != nil {
			fmt.Println("failed to write to dst", err)
		}
	}
}
