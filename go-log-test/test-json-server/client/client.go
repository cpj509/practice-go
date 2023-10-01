package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8515")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	msg := map[string]interface{}{
		"inputs": map[string]interface{}{
			"serverId":  "2470001",
			"workspace": "Default",
			"version":   "0.0.0",
		},
	}
	err = json.NewEncoder(conn).Encode(msg)
	if err != nil {
		fmt.Println("Error encoding JSON:", err.Error())
		return
	}
	fmt.Println("Sent message:", msg)

	// print serverId
	fmt.Println(msg["inputs"].(map[string]interface{})["serverId"])
}
