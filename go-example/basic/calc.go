package calc

import "fmt"

// Sum -
func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func Hello() {
	fmt.Println("hello go test!")
}
