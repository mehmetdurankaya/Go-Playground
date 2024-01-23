package main

import (
	"fmt"
	"time" // time package
)
func main(){
	
	zaman:=time.Now()
	
	fmt.Println(zaman)
	/*
	fmt.Println(zaman.Day())
	fmt.Println(zaman.Month())
	fmt.Println(zaman.Year())
	
	fmt.Println()
	fmt.Println(zaman.Hour())
	fmt.Println(zaman.Minute())
	fmt.Println(zaman.Second())
	
	zamanstring:=zaman.String()
	
	fmt.Println(zamanstring)
	// sleep beklemeye alıyor
	fmt.Println("Merhaba")
	time.Sleep(5*time.Second)//5 saniye bekletiyor
	fmt.Println("Merhaba")

	//UTC
	*/
	zamanUTC:=time.Now().UTC()

	fmt.Println(zamanUTC)
	// Date Kendi Tarihimizi oluşturmak

	dogumgunu:=time.Date(1977,6,13,05,30,0,0,time.Local)
	fmt.Println(dogumgunu)
	fmt.Printf("%v-%d-%v\n",dogumgunu.Day(),dogumgunu.Month(),dogumgunu.Year())

	baslangic:=time.Date(1915,3,18,0,0,0,0,time.UTC)
	bugun:=time.Now().UTC()
	fark:=bugun.Sub(baslangic).Hours()/24

	fmt.Println(fark)

	kacgunoldu:=bugun.Sub(dogumgunu).Hours()/24
	fmt.Println(kacgunoldu)

	var d time.Duration=1000000000


	fmt.Println(d.Hours())

	//Tarih Ekleme

	suankizaman:=time.Now()
	yarin:=suankizaman.AddDate(0,0,1)
	fmt.Println(yarin)
	ikigunsonra:=suankizaman.Add(time.Hour*48)

	fmt.Println(ikigunsonra)

	//Kendimize özel zaman biçimi
	hafta:=time.Hour*7*24
	fmt.Println(hafta)

}