package main

import "fmt"

/*
	bir işlem yaparken yeni bir değişken yani yeni bir memory bellek alanı kullanmak yerine mevcut yerdeki değer üzerinde işlemin yapılmasını sağlıyoruz
*/

type Vatandas struct{
	isim string
	tcno int
}
func (v Vatandas) AyarlaValue(){// method on value
	v.isim="Mehmet"
	v.tcno=44555444

}
func (v *Vatandas) AyarlaPointer(){// method on pointer
	v.isim="Cem"
	v.tcno=99999999
}
func main(){
	vat1:=Vatandas{
		isim:"Mehmet Duran",
		tcno:123456789,
	}
	var ptr *Vatandas=&vat1

	fmt.Println(ptr)
	fmt.Println(&(ptr.isim))
	fmt.Println(&(ptr.tcno))
	fmt.Println(*ptr)
	ptr.isim="Mehmet Duran Kaya"


	fmt.Println(vat1)

	fmt.Println("****************")
	fmt.Println(vat1)
	ptr.AyarlaPointer()
	fmt.Println(vat1)
	



}