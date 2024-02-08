package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	tcpPort := flag.String("tcpport", "515", "TCP port to listen on")
	udpPort := flag.String("udpport", "514", "UDP port to listen on")
	flag.Parse()

	go startUDPServer("0.0.0.0:" + *udpPort)
	startTCPServer("0.0.0.0:" + *tcpPort)
}

func startUDPServer(addr string) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening on UDP port:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("UDP server listening on %s\n", addr)

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("UDP: Received %d bytes from %s: %s\n", n, addr, string(buffer[:n]))
	}
}

func startTCPServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error listening on TCP port:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("TCP server listening on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting TCP connection:", err)
			continue
		}
		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("TCP: Connected to %s\n", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("TCP: Received: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from TCP connection:", err)
	}
}
