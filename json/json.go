package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Students struct {
	Ogrenciler []Ogrenci `json; "Ogrenciler"`
}

type Ogrenci struct {
	OgrenciNo    int    `json:"OgrNo"`
	OgrenciAd    string `json:"OgrAd"`
	OgrenciSoyad string `json:"OgrSoyad"`
	OgrenciNot   int    `json:"OgrNot"`
}

func main() {
	dosya, err := os.Open("ogrenciler.json")
	if err!=nil{
		log.Fatal(err)
	}
	defer dosya.Close()

	arr,err:=ioutil.ReadAll(dosya)
	if err!=nil{
		log.Fatal(err)
	}

	var Ogr Students

	json.Unmarshal(arr,&Ogr)

	//fmt.Println(string(arr))

	for i:=0;i<len(Ogr.Ogrenciler);i++{
		fmt.Println("Öğrenci No: ", Ogr.Ogrenciler[i].OgrenciNo)
		fmt.Println("Öğrenci Ad: ", Ogr.Ogrenciler[i].OgrenciAd)
		fmt.Println("Öğrenci No: ", Ogr.Ogrenciler[i].OgrenciNot)
		
	}


}