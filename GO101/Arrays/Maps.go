package main

import "fmt"


func main(){

	//Standart oluşturma yöntemi 

	var ogrenci map[int]string=map[int]string{
		10:"Ahmet",
		21:"Mehmet",
		39:"Buse",
		89:"Cem",
	}
	fmt.Println(ogrenci)
	fmt.Println(ogrenci[39])

	ogrenci[21]="Uğur"

	fmt.Println(ogrenci[21])

	//key:int value:string
	plaka:=make(map[int]string)

	plaka[31]="Hatay"
	plaka[01]="Adana"
	plaka[80]="Osmaniye"

	fmt.Println(plaka)
	fmt.Println(plaka[31])
   //key:string value:string

   unvanlar:=make(map[string]string)
   unvanlar["DR"]="DOKTOR"
   unvanlar["PRF"]="PROFÖSÖR"
   unvanlar["DOÇ"]="DOÇENT"
   fmt.Println(unvanlar["PRF"])

   a:=unvanlar["DR"]
   fmt.Println(a)

   //silme
   delete(unvanlar,"DR")
   fmt.Println(unvanlar)

   fmt.Println(len(unvanlar))

   for key,value:=range unvanlar{
	fmt.Printf("Kısaltması:%v Tam Hali:%v\n",key,value)
   }
 

}