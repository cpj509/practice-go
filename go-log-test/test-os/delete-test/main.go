package main

import (
	"fmt"
	"os"
)

func main() {
	//delete file
	err := os.Remove("test.txt")
	if err != nil {
		fmt.Println("asd")
		fmt.Println(err)
		return
	}
	fmt.Println("Delete file successfully")

}
