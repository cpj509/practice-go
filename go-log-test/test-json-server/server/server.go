package main

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type LogJson struct {
	FlowUUID string `json:"flowUuid"`
	Inputs   struct {
		ServerID             string `json:"serverId"`
		OsbpID               string `json:"osbpId"`
		Workspace            string `json:"workspace"`
		Version              string `json:"version"`
		InputRedfishHost     string `json:"input_redfish_host"`
		InputRedfishUsername string `json:"input_redfish_username"`
		InputRedfishPassword string `json:"input_redfish_password"`
		Kickstart            string `json:"kickstart"`
	} `json:"inputs"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data LogJson
		err := jsoniter.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// print all data
		fmt.Println(data)

	})
	err := http.ListenAndServe(":8515", nil)
	if err != nil {
		fmt.Println(err)
	}
}
