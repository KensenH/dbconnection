package main

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"

	"net/http"
	"io/ioutil"
    "os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Dbconfig struct {
	User 	string `json:"username"`
	Passwd	string `json:"password"`
	Net		string `json:"net"`
	Addr	string `json:"addr"`
	Dbname	string `json:"dbname"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	
	// jsonFile, err := os.Open("../vault/secrets/dbconfig.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	var dbconfig Dbconfig
	// json.Unmarshal(byteValue, &dbconfig)
	dbconfig.User = os.Getenv("USERNAME")
	dbconfig.Passwd = os.Getenv("PASSWORD")
	dbconfig.Net = os.Getenv("NET")
	dbconfig.Addr = os.Getenv("ADDR")
	dbconfig.Dbname = os.Getenv("DBNAME")

	fmt.Println(dbconfig)

	var fail = 0
	cfg := mysql.Config{
		User:   dbconfig.User,
		Passwd: dbconfig.Passwd,
		Net:    dbconfig.Net,
		Addr:   dbconfig.Addr,
		DbName: dbconfig.Dbname,
	}

	// Get a database handle.
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		fail = 1
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		fail = 1
	}

	if fail == 1 {
		fmt.Println("Failed!")
		fmt.Fprintln(w, "Failed!")
	}else{
		fmt.Println("Connected!")
		fmt.Fprintln(w, "Connected!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	// Capture connection properties.
}