package main

import "fmt"

func main() {
	dbIP1 := "https://192.168.0.1"

	dbID1 := 1

	dbHOSTNAME := ""

	logIP := "192.168.0.1"
	logHOSTNAME := "testname"

	if dbHOSTNAME == "" {
		if dbIP1 == fmt.Sprintf("https://%v", logIP) {
			fmt.Println(dbID1)
		}
	} else {
		if dbHOSTNAME == logHOSTNAME {
			fmt.Println(dbID1)
		}
	}

}
