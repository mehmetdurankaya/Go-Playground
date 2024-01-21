package main

/*
	ATAMA OPERATÖRLERİ
	=,+=,-=,*=,/=,%=

	İLİŞKİSEL OPERATÖRLER
	==,!=,>,<,>=,<=

	MANTIKSAL OPERATÖRLER
	 &&,||,!

	POİNTER OPERATÖRLERİ
	*,&,|,<<,>>

*/

import "fmt"

func main() {

	var x=28
	var y=4

	fmt.Println("X+Y",x+y) 
	fmt.Println("X-Y",x-y)
	fmt.Println("X*Y",x*y)
	fmt.Println("X/Y",x/y)
	fmt.Println("X%Y",x%y)
	
	x++ //değeri bir artırır 29
	y-- //değeri bir eksiltir 3
	fmt.Println(x)
	fmt.Println(y)

	var sayi1 int=9
	var sayi2 int=18

	var toplam= sayi1+sayi2
	fmt.Println("Sayi1 ve Sayi2 toplamı: ",toplam)

	
}