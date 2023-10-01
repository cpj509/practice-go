// Copyright 2012 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package main

import (
	"fmt"
	"log"
	"net"

	g "github.com/gosnmp/gosnmp"
)

func main() {

	// Default is a pointer to a GoSNMP struct that contains sensible defaults
	// eg port 161, community public, etc
	g.Default.Target = "192.168.5.190"
	g.Default.Community = "public"
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{".1.3.6.1.2.1.25.4.2.1.1"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			b := variable.Value.([]byte)
			fmt.Printf("type: %v \n", variable.Type)
			fmt.Printf("string: %v\n", string(variable.Value.([]byte)))
			fmt.Printf("mac?: %vs\n", net.HardwareAddr(b))
			fmt.Printf("byte: %v\n", variable.Value.([]byte))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
			fmt.Printf("string: %v\n", string(variable.Value.([]byte)))
		}
	}
}
