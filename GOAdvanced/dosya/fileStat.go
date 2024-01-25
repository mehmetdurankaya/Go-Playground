package main

/*
	Bir dosyanın bilgilerine nasıl ulaşılır
	os paketinden States metodu kullanılır

*/

import (
	"fmt"
	"os"
)

func main(){

	dosyaBilgi,err:=os.Stat("veri.txt")

	if err!=nil{
		fmt.Println("Hatalı işlem",err)
	}

	fmt.Println("Dosya Adı: ", dosyaBilgi.Name() )
	fmt.Println("Dosya İzinleri: ", dosyaBilgi.Mode() )
	fmt.Println("Dosya Boyutu: ", dosyaBilgi.Size() )
	fmt.Println("Dosya Değiştirme Tarihi: ", dosyaBilgi.ModTime() )
	fmt.Println("Klasör Mü?: ", dosyaBilgi.IsDir() )
	fmt.Println("Sistem Interface: ", dosyaBilgi.Sys())

//eğer dosya boyutu 90kb büyükse bu dosyayı sil gibi
	if dosyaBilgi.Size()>90{
		os.Remove("veri.txt")
	}


}