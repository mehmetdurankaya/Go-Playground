package main

import (
	"log"
	"os"
)

func main(){
	dosya,err:=os.Create("../veri.html")// ../ önceki klasöre git ../../ iki öncekine git
	if err !=nil{
		log.Fatal(err)
	}
	dosya.Close()
}