package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	str := `192.168.0.1:1234`
	fmt.Println("before: ", str)
	str1 := strings.Split(str, ":")
	host, port, _ := net.SplitHostPort(str)

	for a, b := range str1 {
		fmt.Println(a, b)
	}

	fmt.Println(host, port)
}
