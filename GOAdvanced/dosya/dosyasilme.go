package main

import (
	"fmt"
	"os"
)

/*
	os.Remove("dosyaadı.uzantısı") metodu kullanılarak silme işlemi yapılıyor

*/


func main(){
		// os.Remove("veri.txt")

		err:=os.Remove("mdk/veri.html")

		if err !=nil{
			fmt.Println("Silme işlemi sırasında hata oluştu",err)
		}
}