package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func checkErr4(err error){

	if err !=nil{

		log.Fatal(err)
	}
	
}


func main(){

	db,err:=sql.Open("sqlite3","filmler.db") //veri tabanı bağlantısı yapıldı

	defer db.Close()// işlem bittiğinde vekritabanı bağlantısı kapatıldı

	checkErr4(err) // hata olursa handle edildi

	// GÜNCELLEME

	guncelle,err:=db.Prepare("UPDATE oyuncular SET yas=? WHERE ad=?")
	checkErr4(err)

	res,err:=guncelle.Exec(99,"Orhan")
	fmt.Println(res)

	//affect,err:=res.RowsAffected()  //veri tabanına bağlanmaz ancak veri tabanına gönderilen bilgileri kontrol ettğinden ihtiyaç dışında kullanılması önerilmiyor sık kullanılması hataya neden olabilir veri tabanını beklemeye alabilir
	affect,err:=res.LastInsertId()
	checkErr4(err)
	fmt.Println(affect)





}