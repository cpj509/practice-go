package main

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mcuadros/go-syslog.v2"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
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
			var filename strings.Builder
			filename.WriteString(fmt.Sprintf("./log/%v.log", logParts["hostname"]))
			log.SetOutput(&lumberjack.Logger{
				Filename:   filename.String(),
				MaxSize:    1, // megabytes
				MaxBackups: 3,
				MaxAge:     1, //days
				LocalTime:  true,
			})
			log.Println(logParts)
		}
	}(channel)

	server.Wait()
}
