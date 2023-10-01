package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dat, _ := os.ReadFile("../sample-db.txt")

	// sql.DB 객체 생성
	db, err := sql.Open("mysql", string(dat))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 복수 Row를 갖는 SQL 쿼리

	// var EQP_NM string
	// var SVR_IP string
	// var UNQ_IDNT_NO int
	// rows, err := db.Query("SELECT EQP_NM, SVR_IP, UNQ_IDNT_NO FROM t_sdi_work_eqp_info WHERE WORK_ID = 'O230117150410' AND UNQ_IDNT_NO >= ?", 1)
	var name string

	err = db.QueryRow("SELECT uuid FROM t_sdi_nodes WHERE vendor = ?", "HPsE").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ASd")
	fmt.Println(name)
	var uuid string
	var mgmt_ip string

	rows, err := db.Query("SELECT uuid, mgmt_ip FROM t_sdi_nodes WHERE vendor = ?", "asd")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	for rows.Next() {
		err := rows.Scan(&uuid, &mgmt_ip)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uuid, mgmt_ip)

	}
}
