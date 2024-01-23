package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.ToLower("MEHMET DURAN KAYA"))
	fmt.Println(strings.ToUpper("mehmet duran kaya"))

	//Index
	fmt.Println(strings.Index("Türkiye","r"))
	

	//Count
	fmt.Println(strings.Count("Adana","a"))
	fmt.Println(strings.Count("Adana",""))


	//Contains
	fmt.Println(strings.Contains("info@gmail.com","@"))

	//Compare
	//Return 0, if str1==str2
	//Return 1, if str1 > str2
	//Return -1, if str1 < str2
	k1:="Merhaba"
	k2:="Merhaba Dünya"
	k3:="Merhaba"

	fmt.Println(strings.Compare(k1,k2))
	fmt.Println(strings.Compare(k2,k3))
	fmt.Println(strings.Compare(k1,k3))

	//Replace
	fmt.Println(strings.Replace("Olmak yada Olmamak","k", "z",2))

	//ReplaceAll
	fmt.Println(strings.ReplaceAll("Hey Hey Hey","Hey", "Mey"))

	//Fields
	data:="username password   email      date"// aradaki boşlukları yok sayar
	fields:=strings.Fields(data)
	fmt.Printf("%q\n", fields)

	//Split
	fmt.Println(strings.Split("a,b,c,d", ","))
	fmt.Println(strings.Split("13-06-1977", "-"))
	//Spilt Slice
	tarih:="13-06-1977"
	ayir:=strings.Split(tarih,"-")
	fmt.Println(ayir[2])

	//SplitAfter
	sira:="a,b,c,d"
	ayri:=strings.SplitAfter(sira,",")
	fmt.Println(ayri)
	fmt.Println(ayri[1])
}