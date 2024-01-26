package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Kitaplar struct{
XMLName xml.Name `xml:"Kitaplar"`
KitapListesi []Kitap `xml:"Kitap"`




}

type Kitap struct{

XMLName xml.Name `xml:"Kitap"`
Isim string `xml:"isim"`
Yazar string`xml:"yazar"`
Yayinevi string`xml:"yayinevi"`
Fiyat string`xml:"fiyat"`

}

func main(){
	dosya,err:=os.Open("kitaplar.xml")
	if err !=nil{
		log.Fatal(err)
	}
	defer dosya.Close()

	arr,err:=ioutil.ReadAll(dosya)
	if err != nil{
		log.Fatal(err)
	}
	// var ogrenci Ogr
	// xml.Unmarshal(arr,ogrenci)

	// fmt.Println(string(arr))

	Kitaplar:=&Kitaplar{}

	err=xml.Unmarshal(arr,Kitaplar)
	if err!=nil{
		log.Fatal(err)
	}

	for _, degerler:=range Kitaplar.KitapListesi{
		fmt.Println("Kitap Adı: ",degerler.Isim)
		fmt.Println("Yazar: ",degerler.Yazar)
	}

}