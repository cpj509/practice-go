package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	session := &gosnmp.GoSNMP{
		Target:    "127.0.0.1",
		Port:      161,
		Community: "test1",
		Version:   gosnmp.Version2c,
		Timeout:   10 * time.Second,
		MaxOids:   10000,
	}
	service := []string{"amsd"}
	err := session.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
		os.Exit(1)
	}
	defer session.Conn.Close()

	// HOST-RESOURCES-MIB::hrSWRunName
	oid := ".1.3.6.1.2.1.25.4.2.1.2"

	result, err := session.BulkWalkAll(oid)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, variable := range result {
		if contains(service, string(variable.Value.([]byte))) {
			// .1.3.6.1.2.1.25.4.2.1.2.serviceNo에서 serviceNo부분을 긁어옴.
			serviceNo := strings.TrimPrefix(variable.Name, oid+".")

			// for testing
			fmt.Println("Extracted value:", serviceNo)

			// HOST-RESOURCES-MIB::hrSWRunStatus
			result2, err := session.BulkWalkAll("1.3.6.1.2.1.25.4.2.1.7." + serviceNo)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}

			for _, variable2 := range result2 {
				fmt.Println(variable2.Value.(int))
				if variable2.Value.(int) == 4 {
					fmt.Println("Failed due to value 4.")
					continue
				}
			}

			fmt.Println("Success!")
			return
		}
	}

	fmt.Println("Failed to find desired value.")
}

func contains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}
