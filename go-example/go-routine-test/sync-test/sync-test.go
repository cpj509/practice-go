package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go worker(i, &wg)

	}

}
