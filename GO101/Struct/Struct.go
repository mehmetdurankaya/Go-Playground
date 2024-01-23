package main

/*
	Struct (Class)
	Kendimize ait yapıların inşaa edilmesi
	Kendi veri tipimizi oluşturmamızı sağlayan yapılardır.

*/

import "fmt"

type ogrenci struct{
 //field fieldtype
	isim string
	soyisim string
	yas int
	notlar []int
	isMezunMu bool
}
	func main(){

		//********** nesne oluşturmanın 1.şekli**********
		var o1 ogrenci
		o1.isim = "Mehmet Duran"
		o1.soyisim = "Kaya"
		o1.yas=47
		o1.notlar=[]int{90,95,100}
		o1.isMezunMu=true

		fmt.Printf("Öğrencinin Adı: %s \n Öğrencinin Soyadı: %s \n Yaşı: %d \n Notları: %v \n Mezuniyet Durumu: %v ",o1.isim,o1.soyisim,o1.yas,o1.notlar,o1.isMezunMu)
		//************************************************
		//********** nesne oluşturmanın 2.şekli**********
		
		fmt.Println()
		o2:=ogrenci{
			isim:"Nizamettin",
			soyisim:"Kaya",
			yas:17,
			notlar:[]int{90,95,100},
			isMezunMu:false,
		}
		fmt.Println(o2)
		//************************************************
		//********** nesne oluşturmanın 3.şekli**********

		o3:=ogrenci{"Fatma","Kaya",19,[]int{90,95,95},false}
		fmt.Println(o3)



	}

