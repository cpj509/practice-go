package main

import (
	"os"
	"strconv"
	"time"
)

var logChannel chan string

func logSetup(logFile string) {
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ := os.Create(logFile)
		defer f.Close()
	}

	logChannel = make(chan string, 100)
	go func() {

		for msg := range logChannel {
			f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			f.WriteString(time.Now().String() + " " + msg + "\n")
			f.Close()

		}
	}()
}

func main() {
	logSetup("./logfile.log")

	go func() {
		for i := 1; i <= 20; i++ {
			n := strconv.Itoa(i)
			println(n)
			logChannel <- n
		}
	}()

	go func() {
		for i := 100; i <= 120; i++ {
			logChannel <- strconv.Itoa(i)
		}
	}()

	time.Sleep(2 * time.Second)

	close(logChannel)
}
