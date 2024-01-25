package main

/*
	önclelikli olarak okunacak dosyayı "kaynak" olarak veriyoruz
	ioutil paketinden ReadFile metoduyla okunmasını ve kopyalanmasını sağlıyoruz

	os paketinin Create metoduyla yeni bir dosya oluşturuyoruz daha önce oluşturulan bir dosyaya da yazılabilir

	ioutil paketinin WriteFile metoduyla yeni oluşturulan dosyaya kopyalamış olduğmuz verileri aktarıyoruz

	defer kullanarak dosyayı kapatıyoruz

*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	kaynak, err := ioutil.ReadFile("veri.txt")
	if err !=nil{
		log.Fatal(err)
	}

	yeniDosya,err2:=os.Create("yenidosya.txt")

	if err2 !=nil{
		log.Fatal(err2)
	}

	err3:=ioutil.WriteFile("yenidosya.txt",kaynak,0666)
	if err3 !=nil{
		fmt.Println("oluştururken hata oluştu",err3)
	}
	defer yeniDosya.Close()
}