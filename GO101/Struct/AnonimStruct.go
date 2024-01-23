package main

/*
	Bir kez kullanılacak struct yapılarını anonim struct olarak tanımlaya biliyoruz
	anonim fonksiyonlara benziyorlar bir isimleri yoktu r
*/

import "fmt"

type adres struct {
	sehir string
	ulke  string
}

type kisi struct {
	isim      string
	yas       int
	kisiadres adres
}
	

func main() {
kurucu:=struct{
		isim string
		sermaye int		
	}{isim:"Tahsin",sermaye:750000}
fmt.Println(kurucu)

}