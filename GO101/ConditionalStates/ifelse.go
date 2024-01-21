package main


import "fmt"
/*
if(kosul){
	//yapılacak işlem
}else{
	//yapılacak işlem
}
*/



func main(){
/*
	isim:="Mehmet Duran Kaya"

	if isim=="Mehmet Duran Kaya"{
		fmt.Println("Hoşgeldin Mehmet Duran Kaya")
	}else if isim=="Turan"{
		fmt.Println("Hoşgeldin Turan")
	}else{
		fmt.Println("Kimsiniz sizi tanımıyorum")
	}
	
*/
	yas:=19

	fmt.Println("18 yaşındank küçükler giremez")

	if yas>=18{
		fmt.Println("Mekana hoşgeldiniz")
	}else if yas<18 && yas>=14{
		fmt.Println("Ailenle girebilirsin")
	}else{
		fmt.Println("Mekana giremezsin")
	}

}