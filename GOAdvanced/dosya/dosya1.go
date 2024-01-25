package main

import (
	//"fmt"
	"log"
	"os"
)

/*
	O_RDONLY SADECE OKUMA
	O_WRONLY SADECE YAZMA
	O_RDWR OKUMA VE YAZMA
	O_APPEND DOSYAYA EKLEME SONUNA EKLEME
	O_CREATE DOSYA YOKSA OLUSTUR
	O_TRUNC DOSYA SONUNU KESER
	O_EXCL DOSYA YOKSA OLUŞTUR O_CREATE İLE KULLANILIR
	O_SYNC I/O İŞLEMLERİ İÇİN SENKRON ÇALIŞIR


	0000 İZİN YOK
	O700 YÖNETİCİ YAZABİLİR, OKUYABİLİR, OLUŞTURA BİLİR
	0770 YÖNETİCİ VE GRUP YAZABİLİR OKUYABİLİR OLUŞTURABİLİR
	0777 HERKES YAZABİLİR,OKUYABİLİR VE OLUŞTURABİLİR
	0111 OLUŞTURMA
	0222 YAZMA
	0333 YAZMA VE OLUŞTURMA
	0444 OKUYMA
	0555 OKUYMA VE OLUŞTURMA
	0666 OKUYMA VE YAZMA
	0740 YÖNETİCİ YAZABİLİR,OKUYABİLİR VE OLUŞTURABİLİR; GRUP SADECE OKUYABİLİR



	O600 OKUYABİLİR, OLUŞTURABİLİR

*/

func main() {
	dosya, err := os.OpenFile("Veri.txt",os.O_WRONLY|os.O_APPEND,0666)
	defer dosya.Close()
	if err !=nil{
		log.Fatal(err)
	}
	/*
	//fmt.Println("Dosya Açıldı")
	//dosya.WriteString("Mehmet Duran Kaya")

	yaziSlice:=[]byte ("Merhabalar Udemy ")
	
	dosyaYazma,err:=dosya.Write(yaziSlice)

	if err !=nil{
		log.Fatal(err)
	}


	fmt.Printf("Dosyaya %d Byte Boyutunda Yazı Yazıldı\n",dosyaYazma)
	
	*/
	sehirler:=[]string{"Osmaniye","Adana","Hatay","Çanakkale","İzmir","Kütahya"}

	for _,sehir:=range sehirler{
		_,err:=dosya.WriteString(sehir+"\n")
		if err!=nil{
			log.Fatal(err)
		}
	}

}