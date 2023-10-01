package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	defer wg.Wait()

	wg.Add(2) // 2개만 추가
	go func() {
		wg.Done()
		fmt.Println("one")
	}()

	go func() {
		wg.Done()
		fmt.Println("two")
	}()

	go func() {
		wg.Done()
		fmt.Println("three")
	}()
}
