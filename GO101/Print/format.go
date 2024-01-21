package main

import "fmt"

/*
	%d decimal integer type lerde kullanılır
	%f float type lerde kullanılır
	%s string type lerde kullanılır
	%t boolean type lerde kullanılır
	%b binary type lerde kullanılır
	%c karakterler için kullanılır
	%p pointer için kullanılır
	%T veri Type lerini belirlemek için kullanılır
	sprintf() fonksiyonu ile formatı bir değişkene aktarabiliyoruz
	var messagecomplate=fmt.Sprintf("%d yılından %s",year,message)

*/

func main(){
	/*
	var sayi1 int=5

	fmt.Printf("Benim Yaşım:%b", sayi1)
	
		var sayi2 float32=3.1415925

	fmt.Printf("PI SAYISI:%.2f", sayi2)

	var dogruMu bool=true

	fmt.Printf("%t", dogruMu)
	
	

	var isim="Mehmet"
	var sehir="Hatay"
	var yas=46


	fmt.Printf("Benim Adım:%s, Yaşım: %d,  %s' da Yaşıyorum ", isim,yas,sehir)
	

	var sayi1 int=5
	var sayi2 float32=5.6434
	var yanlisMi bool= false
	var str string= "Mehmet Duran"

	fmt.Printf("Veri Type:%T,%T,%T,%T", sayi1,sayi2,yanlisMi,str)
	*/
	var message string= "Merhaba Google"

	var year=2024
	var messagecomplate=fmt.Sprintf("%d yılından %s",year,message)
	fmt.Println(messagecomplate)

}