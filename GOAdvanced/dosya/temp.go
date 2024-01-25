package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	geciciKlasor,err := os.MkdirTemp("","udemy")
	if err != nil {
		log.Fatal(err)

}
fmt.Println("Geçici Klasör: ",geciciKlasor)
// işlemler

//defer os.RemoveAll(geciciKlasor) //işlemler bittikten sonra geçici Klasörün silinmesini sağlıyor

geciciDosya,err:=os.CreateTemp(geciciKlasor,"udemykursu")

if err != nil {

	log.Fatal(err)
}

fmt.Println("Geçici Dosya: ", geciciDosya.Name() )

defer os.Remove(geciciDosya.Name())
}