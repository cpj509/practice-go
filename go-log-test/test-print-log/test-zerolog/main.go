package main

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog"

	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenTCP("0.0.0.0:8514")
	server.ListenUDP("0.0.0.0:8514")
	server.Boot()

	logger := zerolog.New(os.Stdout)
	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			doc, _ := json.Marshal(logParts)
			logger.Info().Msgf("msg", doc)
		}
	}(channel)

	server.Wait()
}
