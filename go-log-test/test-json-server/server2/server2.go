package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	FlowUuid string `json:"flowUuid"`
	Inputs   struct {
		ServerId             string `json:"serverId"`
		OsbpId               string `json:"osbpId"`
		Workspace            string `json:"workspace"`
		Version              string `json:"version"`
		InputRedfishHost     string `json:"input_redfish_host"`
		InputRedfishUsername string `json:"input_redfish_username"`
		InputRedfishPassword string `json:"input_redfish_password"`
		Kickstart            string `json:"kickstart"`
	} `json:"inputs"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var m Message
	err := decoder.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(m)
}

func main() {
	http.HandleFunc("/", handlePost)
	http.ListenAndServe(":8515", nil)
}
