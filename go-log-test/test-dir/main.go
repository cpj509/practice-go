package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dat, _ := ioutil.ReadFile("../sample.txt")
	fmt.Println(dat)
	// sql.DB 객체 생성
	db, err := sql.Open("mysql", string(dat))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 복수 Row를 갖는 SQL 쿼리
	var WORK_ID string

	rows, err := db.Query("SELECT WORK_ID FROM t_sdi_work_eqp_info")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&WORK_ID)
		if err != nil {
			log.Fatal(err)
		}
		//dir 생성
		path := fmt.Sprintf("./sample/%v", WORK_ID)
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		// fmt.Println(WORK_ID)
	}

}
