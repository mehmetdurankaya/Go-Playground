package main

import (
	"log"
	"os"
)

/*
	Klasör oluşturma
*/


func main(){

	err:=os.Mkdir("../xml",os.ModePerm)

	if err !=nil{
		log.Fatal(err)
	}

}