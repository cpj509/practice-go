package main

import (

	// "log"

	"os"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {

	var log = &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.TextFormatter{
			DisableQuote: true,
		},
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
	}
	ip := "0.0.0.0"
	port := ":8514"
	addr := ip + port

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	server.ListenUDP(addr)
	server.ListenTCP(addr)
	server.Boot()
	log.Println("start log server. listening port" + port)

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			doc, _ := jsoniter.Marshal(logParts)

			log.Println(string(doc))
		}
	}(channel)

	server.Wait()
}
