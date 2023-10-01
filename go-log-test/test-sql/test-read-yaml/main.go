package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

// TestConfig
type TestConfig struct {
	Database Database `yaml:"database"`
}

// Database
type Database struct {
	Ems string `yaml:"ems"`
}

func Config(fileName string) (*TestConfig, error) {

	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	t := &TestConfig{}
	err = yaml.Unmarshal(buf, t)
	if err != nil {
		//err
	}

	return t, nil
}

func main() {

	t, err := Config("./test-sample.yml")
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", t.Database.Ems)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT uuid FROM t_sdi_nodes WHERE vendor = ?", "HPE").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
