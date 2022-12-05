package main

import (
	"fmt"
	"sync"
	"time"
)

var globalValue int

func action(num int, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	globalValue += num
	mu.Unlock()
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	defer wg.Wait()
	startTime := time.Now()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go action(i, &wg, &mu)
	}

	endTime := time.Since(startTime)

	fmt.Println("global value", globalValue)
	fmt.Println("end time : ", endTime)
}
