package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	gosnmp.Default.Target = "192.168.5.190"
	gosnmp.Default.Community = "public"
	gosnmp.Default.Timeout = time.Duration(10 * time.Second) // Timeout better suited to walking
	err := gosnmp.Default.Connect()
	if err != nil {
		fmt.Printf("Connect err: %v\n", err)
		os.Exit(1)
	}
	defer gosnmp.Default.Conn.Close()

	err = gosnmp.Default.BulkWalk(".1.3.6.1.2.1.25.4.2.1.2", printValue)
	if err != nil {
		fmt.Printf("Walk Error: %v\n", err)
		os.Exit(1)
	}

	// oid = ".1.3.6.1.2.1.25.4.2.1.7"

	// result, err := snmpClient.Get([]string{oid})
	// if err != nil {
	// 	fmt.Printf("Error retrieving value for OID: %v", err)
	// 	return
	// }

	// for i, variable := range result.Variables {
	// 	if variable.Type == gosnmp.Integer {
	// 		if variable.Value.(int) == 1 {
	// 			fmt.Printf("Service is running (Index: %d)\n", i)
	// 		} else {
	// 			fmt.Printf("Service is not running (Index: %d)\n", i)
	// 		}
	// 	}
	// }
}

func printValue(pdu gosnmp.SnmpPDU) error {
	fmt.Printf("%s = ", pdu.Name)

	switch pdu.Type {
	case gosnmp.OctetString:
		b := pdu.Value.([]byte)
		fmt.Printf("STRING: %s\n", string(b))
		fmt.Printf("TYPE %v: STRING: %s\n", pdu.Type.String(), string(b))
	default:
		fmt.Printf("TYPE %v: %d\n", pdu.Type, gosnmp.ToBigInt(pdu.Value))
	}
	return nil
}
