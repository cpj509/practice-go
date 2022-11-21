package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"

	snmplib "github.com/deejross/go-snmplib"
	_ "github.com/go-sql-driver/mysql"
)

type snmpHandler struct{}
type Info struct {
	Address   string
	Community string
}

func (h snmpHandler) OnError(addr net.Addr, err error) {
	log.Println(addr.String(), err)
}

func (h snmpHandler) OnTrap(addr net.Addr, trap snmplib.Trap) {
	prettyPrint, _ := json.MarshalIndent(trap, "", "\t")
	log.Println(string(prettyPrint))
	var info Info
	json.Unmarshal([]byte(prettyPrint), &info)

	db, _ := sql.Open("mysql", "root:1111@tcp(localhost:3306)/snmp_test")
	defer db.Close()
	fmt.Printf("%v, %v\n", info.Address, info.Community)
	result, err := db.Exec("INSERT INTO test_table01(address, community) VALUES(?, ?)", info.Address, info.Community)
	if err != nil {
		fmt.Println(err)
	}

	n, RowsAffectedErr := result.RowsAffected()
	if RowsAffectedErr != nil {
		fmt.Println(RowsAffectedErr)
	}
	fmt.Println(n)

}

func main() {

	port := 9999
	server, err := snmplib.NewTrapServer("0.0.0.0", port)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Listening for traps on port %v", port)
	server.ListenAndServe(snmpHandler{})

}
