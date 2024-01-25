package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func checkErr7(err error){

	if err !=nil{

		log.Fatal(err)
	}
	
}


func main(){

	db,err:=sql.Open("sqlite3","filmler.db") //veri tabanı bağlantısı yapıldı

	defer db.Close()// işlem bittiğinde vekritabanı bağlantısı kapatıldı

	checkErr7(err) // hata olursa handle edildi

	guncelle,err:=db.Prepare("UPDATE filmler SET butce=?")

	tx,BGerr:=db.Begin()
	checkErr7(BGerr)

	_,txerr:=tx.Stmt(guncelle).Exec(200)
	if txerr !=nil{
		fmt.Println("Geri alınıyor...ROLLBACK")
		tx.Rollback()
	}else{
		fmt.Println("Commit Edildi")
		tx.Commit()
	}
	




}