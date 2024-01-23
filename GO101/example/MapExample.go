package main


import "fmt"

func main(){

	sozluk:=make(map[string]string)

	var eng string

	var tr string

	for i:=1;i<=5;i++{
		fmt.Print("ingilizce Kelime Giriniz")
		fmt.Scan(&eng)

		fmt.Print("Türkçe Karşılğığını giriniz")
		fmt.Scan(&tr)

		sozluk[eng]=tr
	}

	fmt.Println(sozluk)

	fmt.Print("Silmek istediğiniz ingilizce kelimeyi giriniz: ")
	fmt.Scan(&eng)
	delete(sozluk,eng)
	for k,v:=range sozluk{
		fmt.Printf("ENG:%s  \t TR:%s  \n",k,v)
	}
}