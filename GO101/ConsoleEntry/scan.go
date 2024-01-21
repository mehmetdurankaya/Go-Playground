package main

import 	"fmt"



func main(){
/*
	var isim string
	var yas int

	fmt.Print("Lütfen adınızı giriniz:")
	fmt.Scan(&isim)

	fmt.Print("Lütfen yaşınızı giriniz:")
	fmt.Scan(&yas)

	fmt.Printf("Adınız %s Yaşınız %d",isim,yas)
*/
 var isim string
 fmt.Print("İsminizi Giriniz: ")
 fmt.Scanf("%s",& isim)

 fmt.Println("Merhaba",isim)


}