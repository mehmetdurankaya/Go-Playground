package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

/*
	bir şirkette personel bilgisini tutacak bir xml dosyası oluşturmak amaçlandı



*/

type Sirket struct{
	XMLName xml.Name  `xml:"sirket"`
	Personeller[]Personel `xml:"personel"`


}
type Personel struct{
	XMLName xml.Name`xml:"personel"`
	TCNO int `xml:"tcno"`
	PerAd string `xml:"personelAd"`
	PerSoyad string `xml:"personelSoyad"`
	PerMevki string `xml:"mevki"`
} 


func main (){
	dosya,err:=os.Create("personeller.xml")
	if err!=nil{
		log.Fatal(err)
	}
	defer dosya.Close()
	str:="<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
	dosya.WriteString(str)
	s:=&Sirket{	}
	s.Personeller=append(s.Personeller, Personel{TCNO: 1111111111,PerAd: "Mehmet Duran",PerSoyad: "Kaya",PerMevki: "Full-Stack Developer"})
	s.Personeller=append(s.Personeller, Personel{TCNO: 2222222222,PerAd: "Nizamettin",PerSoyad: "Kaya",PerMevki: "Tester"})
	s.Personeller=append(s.Personeller, Personel{TCNO: 3333333333,PerAd: "Sibel",PerSoyad: "Kaya",PerMevki: "Manager"})

	xmlWriter:=io.Writer(dosya)//dosyaya yaz

	enc:=xml.NewEncoder(xmlWriter)
	enc.Indent("  ","    ")
	enc.Encode(s)




	




}