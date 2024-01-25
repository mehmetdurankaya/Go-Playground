package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func checkErr6(err error){

	if err !=nil{

		log.Fatal(err)
	}
	
}


func main(){

	db,err:=sql.Open("sqlite3","filmler.db") //veri tabanı bağlantısı yapıldı

	defer db.Close()// işlem bittiğinde vekritabanı bağlantısı kapatıldı

	checkErr6(err) // hata olursa handle edildi

	satirlar,err:=db.Query("SELECT * FROM filmler")

	checkErr6(err)
	var id int
	var filmAd string
	var filmKategory string
	var yonetmenAd string
	var yapimyili int
	var butce int

	for satirlar.Next(){
		err=satirlar.Scan(&id,&filmAd,&filmKategory,&yonetmenAd,&yapimyili,&butce)
		checkErr6(err)
		fmt.Println("id: ",id, "film Adı: ",filmAd,"film Kategori: ","Yönetmen Ad: ", yonetmenAd,"Yapım Yılı: ", yapimyili,"Bütçe: ",butce  )
	}

	




}