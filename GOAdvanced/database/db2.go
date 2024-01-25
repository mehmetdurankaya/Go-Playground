package main

/*
	VERİ TABANINDAN VERİ ÇEKME
	db, err := sql.Open("sqlite3", "filmler.db")


*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" //init bu paketi içe aktar
)

func checkErr2(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "filmler.db")
	defer db.Close()
	checkErr2(err)

	satirlar,err:=db.Query("SELECT * FROM filmler")
	checkErr2(err)


		var id int
		var filmAd string
		var filmKategori string
		var yonetmenAd string
		var yapimyili int
		var butce int

	for satirlar.Next(){
		err=satirlar.Scan(&id,&filmAd,&filmKategori,&yonetmenAd,&yapimyili,&butce)
		checkErr2(err)
		// fmt.Println(id)
		// fmt.Println(filmAd)
		// fmt.Println(butce)

		fmt.Printf(" %s  %s  %s   %s   %s    %s\n","ID","Film Adı","Film Kategori","Yönetmen","Yapım Yılı","Bütçe")
		fmt.Printf(" %d  %s  %s   %s   %d    %d",id,filmAd,filmKategori,yonetmenAd,yapimyili,butce)
	

	}




}