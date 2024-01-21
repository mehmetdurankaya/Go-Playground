package main

import "fmt"

func main(){
	/*
	\n bir alt satıra geçer
	\t boşluk bırakmayı sağlar
	"    \"....."\  "   çift tırnak içerisinde çift tırnak kullanmak için kaçış operatörü yani \ kullanılır. çift tırnak içerisinde çift tırnak kullanmamız gerekirse öncesine ve sonrasına ters slash \ kullanmalıyız. go dilinde tek tırnağın kaçış işareti bulunmamaktadır. 
	*/

	fmt.Println("Merhaba\nDünya") 
	fmt.Println("Hello\tWorld") 


	fmt.Println("Hatay'ın \"Künefesi\" Ünlüdür") 

	fmt.Println("Hatay'ın"+"Künefesi"+"Ünlüdür") // birleştirir
	fmt.Println("Hatay'ın","Künefesi","Ünlüdür") // aralara boşluk bırakır


	fmt.Println("Lorem İpsum"+ // çoklu satır birleştirme
	"Merhaba Dünya" +
	"Hello World") 

	//String birleştirme
	isim:="Mehmet Duran"
	soyisim:="Kaya"
	sehir:="Hatay"

	fmt.Println(isim,soyisim,sehir)

	//fmt.Println(len(isim))
	var stringUzunluk=len(isim)

	fmt.Println(stringUzunluk)

}