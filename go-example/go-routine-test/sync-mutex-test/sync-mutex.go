package main

import (
	"fmt"
	"sync"
	"time"
)

var globalValue int

func action(num int, mu *sync.Mutex) {
	mu.Lock()
	globalValue += num
	mu.Unlock()
	time.Sleep(time.Microsecond * 1)
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	startTime := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i

		go func() {
			action(i, &mu)
			defer wg.Done()
		}()
	}

	wg.Wait()

	endTime := time.Since(startTime)

	fmt.Println("global value", globalValue)
	fmt.Println("end time : ", endTime)
}
