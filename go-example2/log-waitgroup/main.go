package main

import (
	"os"
	"strconv"
	"sync"
	"time"
)

var logChannel chan string
var wg sync.WaitGroup
var logWg sync.WaitGroup

func logSetup(logFile string) {

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		f, _ := os.Create(logFile)
		defer f.Close()
	}

	logChannel = make(chan string, 100)

	logWg.Add(1)

	go func() {
		defer logWg.Done()
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

	wg.Add(2)
	go func() {
		for i := 1; i <= 20; i++ {
			n := strconv.Itoa(i)
			println(n)
			logChannel <- n
		}
		wg.Done()
	}()

	go func() {
		for i := 100; i <= 120; i++ {
			logChannel <- strconv.Itoa(i)
		}
		wg.Done()
	}()

	wg.Wait()
	close(logChannel)
	logWg.Wait()
}
