package main

import (
	"database/sql"
	"log"
	

	_ "github.com/mattn/go-sqlite3" //init bu paketi içe aktar
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db,err:=sql.Open("sqlite3","notlar.db")

	defer db.Close()
	checkErr(err)

	_,err =db.Exec("CREATE TABLE notlar(id INT PRIMARY KEY, notdurumu TEXT,isim TEXT)")
	checkErr(err)

}