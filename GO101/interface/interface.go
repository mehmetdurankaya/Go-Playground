package main

import "fmt"

/*
	interface:Nesnelerin ilişkili olduğu metotları yapıları toplu bir şekilde kullanmammızı sağlayan yapılardır
	nesne: oluşturulan struct lardan türetilen değişkenlerdir.

	farklı veri türlerinin aynı şekilde davranmasını istediğimiz noktalarda kullandığımız bir yapıdır
	ortak davranışların yönetilmesini kolaylaştırır
	tanımlama:
	type interface_adi inteface{
		//Method imzaları
		func1()int
		func2()float64
	}

*/

type Insan string
type Kedi string


type yuruyen interface{
	Yuru() string
}

func (i Insan)Yuru()string{
	return "Ben bir insanım ve yürüyorum"
}
func (k Kedi)Yuru()string{
	return "Ben bir kediyim ve yürüyorum"
}

//Interface'den oluşturulan function
func yurume(y yuruyen){
	fmt.Println(y.Yuru())
}

func main(){

var insan Insan
var kedi Kedi
yurume(insan)
yurume(kedi)
}