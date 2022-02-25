package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	dbconfig.User = "PharmanetBois"
	dbconfig.Passwd = "d3v3l0p8015"
	dbconfig.Net = "tcp"
	dbconfig.Addr = "34.87.44.167"
	dbconfig.Dbname = "q_life"

	fmt.Println(dbconfig)

	var fail = 0
	// cfg := mysql.Config{
	// 	User:   dbconfig.User,
	// 	Passwd: dbconfig.Passwd,
	// 	Net:    dbconfig.Net,
	// 	Addr:   dbconfig.Addr,
	// 	DBName: dbconfig.Dbname,
	// }

	var conn_string = "PharmanetBois:d3v3l0p8015@tcp(34.87.44.167:3306)/q_life?parseTime=true"
	var conn_string2 = "accounting_integra:d3v3l0p8015@tcp(35.198.226.66:3306)/ecommerce_data?parseTime=true"
	var err error
	// Get a database handle.
	db, err = sql.Open("mysql", conn_string)
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
		// fmt.Println("Failed!")
		fmt.Fprintln(w, "Failed!")
	}else{
		// fmt.Println("Connected!")
		fmt.Fprintln(w, "Connected!")
	}

	rows, err := db.Query("select remedyname FROM m_remedy limit 5")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var title string
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, title)
		log.Println(title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()


//second

	db, err = sql.Open("mysql", conn_string2)
	if err != nil {
		fmt.Println(err)
		fail = 1
	}

	pingErr = db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		fail = 1
	}

	if fail == 1 {
		// fmt.Println("Failed!")
		fmt.Fprintln(w, "Failed!")
	}else{
		// fmt.Println("Connected!")
		fmt.Fprintln(w, "Connected!")
	}

	rows, err = db.Query("select company from data_api_tokopedia limit 5")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, title)
		log.Println(title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
	// Capture connection properties.
}