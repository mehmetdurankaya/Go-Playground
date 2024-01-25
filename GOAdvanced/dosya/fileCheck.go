package main

import (
	"fmt"
	"os"
)

/*

	Dosya kontrolü os paketinin IsNotExist(err) metodu ile yapılıyor parametre olarak hata nesnesini alıyor sonucu tru/false olarak döndürüyor
	false dosya mevcut
	true dosya mevcut degil


*/

func main(){

_,err:=os.Stat("yeni.txt")

if err !=nil{
	fmt.Println("HATA: ", err)
}

dosyaKontrol:=os.IsNotExist(err)//dosya mevcut değil mi 
fmt.Println(dosyaKontrol)

if dosyaKontrol== true{
	os.Create("data.txt")
}

}