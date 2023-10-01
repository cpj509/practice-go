package main

import (
	"fmt"
	"net"
	"time"

	// "log"

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

	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			client := logParts["client"]
			clientIP, _, _ := net.SplitHostPort(fmt.Sprintf("%v", client))
			date := time.Now().Format("20060102")
			var filename string = fmt.Sprintf("./log/%v-%v.log", clientIP, date)
			log.SetOutput(&lumberjack.Logger{
				Filename:   filename,
				MaxSize:    100, // megabytes
				MaxBackups: 5,
				// MaxAge:     1, //days
				// Compress:   true, // disabled by default
				LocalTime: true,
			})
			log.Println(logParts)
		}
	}(channel)

	server.Wait()
}