package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)
type Companies struct{
	CListe [] Sirket
}

type Sirket struct {
	ID            int
	Isim          string
	KurulusTarihi int
	DevamEdiyorMu bool
}


func main() {
	sirket1 := Sirket{}

	sirket1.ID = 1
	sirket1.Isim = "MDK"
	sirket1.KurulusTarihi = 2024
	sirket1.DevamEdiyorMu = true

	sirket1Byte, err := json.Marshal(sirket1)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(sirket1Byte))

	sirket2:=Sirket{22,"MacroSoft",2011,true}
	sirket2Byte,_:=json.Marshal(sirket2)
	fmt.Println(string(sirket2Byte))

	dosya,err:=os.Create("sirketlerim.json")

	if err!=nil{
		log.Fatal(err)
	}
	
	c1:=Companies{}

	c1.CListe=append(c1.CListe,Sirket{ID:543,Isim:"MDK",KurulusTarihi: 2014,DevamEdiyorMu: true})

	c1.CListe=append(c1.CListe,Sirket{ID: 17,Isim: "NK",KurulusTarihi: 2017,DevamEdiyorMu: true})

	jsonWriter:=io.Writer(dosya)
	enc:=json.NewEncoder(jsonWriter)
	json.MarshalIndent(c1," ","\t")
	enc.Encode(c1)

}