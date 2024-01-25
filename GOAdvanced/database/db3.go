package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func checkErr3(err error){

	if err !=nil{

		log.Fatal(err)
	}
	
}


func main(){

	db,err:=sql.Open("sqlite3","filmler.db") //veri tabanı bağlantısı yapıldı

	defer db.Close()// işlem bittiğinde vekritabanı bağlantısı kapatıldı

	checkErr3(err) // hata olursa handle edildi

	oyuncuAd:=[]string{"Şener Şen","Kemal Sunal"}
	oyuncuId:=[]int {3,4}
	oyuncuYas:=[]int {74,68}

	for i:=0;i<len(oyuncuId);i++{
		veriekle,err:=db.Prepare("INSERT INTO oyuncular(id,ad,yas)VALUES(?,?,?)")
		checkErr3(err)
		_,err=veriekle.Exec(oyuncuId[i],oyuncuAd[i],oyuncuYas[i])
		checkErr3(err)
		
	}


}