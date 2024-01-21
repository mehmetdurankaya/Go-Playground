package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	//TYPE CASTING  uyumlu tipleri dönüştürür
	var toplam int=984
	var sayi int=17

	var sonuc =float32(toplam) /float32(sayi)
	fmt.Println(sonuc)

	var s1=int(sonuc)
	fmt.Println(s1)
	

	//TYPE CONVERSION  uyumsuz tipleri paket kullanarak dönüşüm sağlıyor  
	sayıyı stringe strconv.Ito() 
	stringi sayıya strconv.Atoi()

	var str="1"
	//var sonuc=str+5 // string bir değişkenle sayısal bir değişken arasında matematiksel işlem yapılamaz string tipinden sayısal tipe dönüştürülmesi gerekir
	var sayi,_=strconv.Atoi(str)
	fmt.Println(sayi)

	var sonuc= sayi+7
	fmt.Println(sonuc)
	
	*/
	// Int String' e Dönüşüm

	var sayi=1

	var str=strconv.Itoa(sayi)
	fmt.Println("Sayıyı stringe çevirdik:",str)


}