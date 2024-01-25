package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func checkErr5(err error){

	if err !=nil{

		log.Fatal(err)
	}
	
}


func main(){

	db,err:=sql.Open("sqlite3","filmler.db") //veri tabanı bağlantısı yapıldı

	defer db.Close()// işlem bittiğinde vekritabanı bağlantısı kapatıldı

	checkErr5(err) // hata olursa handle edildi

	// GÜNCELLEME

	silme,err:=db.Prepare("DELETE FROM oyuncular WHERE id=?")
	checkErr5(err)

	res,err:=silme.Exec(1)
	checkErr5(err)

	affect,err:=res.RowsAffected()
	checkErr5(err)

	fmt.Println("Silinen Satır Sayısı: ",affect)






}