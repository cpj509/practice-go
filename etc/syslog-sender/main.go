package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"time"
)

func main() {
	ip := flag.String("ip", "localhost", "The IP address of the syslog server")
	port := flag.String("port", "514", "The port of the syslog server")
	protocol := flag.String("protocol", "udp", "The protocol for syslog communication (udp/tcp)")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", *ip, *port)
	writer, err := syslog.Dial(*protocol, address, syslog.LOG_INFO|syslog.LOG_DAEMON, "syslog_test")
	if err != nil {
		log.Fatalf("syslog.Dial() failed: %s", err)
	}

	// Send a dummy syslog message every 5 seconds
	for {
		msg := fmt.Sprintf("Dummy syslog message sent at %s", time.Now().Format(time.RFC3339))
		err = writer.Info(msg)
		if err != nil {
			log.Fatalf("writer.Info() failed: %s", err)
		}
		fmt.Printf("Sent syslog message: %s\n", msg)
		time.Sleep(5 * time.Second)
	}
}
