package main

import "fmt"

/*
	func(alıcı_adı  Türü) fonksiyon_adı(parametre_listesi)(dönüş_türü){
		metot gövdesi
	}

	func(o ogrenci) AdGoster()string{

	}


*/

type Ogrenci struct{
	isim string
	yas int
	notlar []int
	gectiMi bool

}
func(o Ogrenci) AdGoster()string{
	return o.isim
}

func(ogr Ogrenci) BilgileriGoster(){

		fmt.Println("Öğrenci Adı:",ogr.isim)
		fmt.Println("Öğrenci Yaşı:",ogr.yas)
		fmt.Println("Öğrenci Notları:",ogr.notlar)
		fmt.Println("Öğrenci Geçti mi:",ogr.gectiMi)
	

	}

func (not Ogrenci) ortalamaHesapla()float64{
	toplam:=0
	for _,i:=range not.notlar{
		toplam+=i
	}
	return float64(toplam)/float64(len(not.notlar))
}

func(gec Ogrenci)gectimiKaldimi()string{
	
	if gec.ortalamaHesapla()<=50{	
		return "Kaldı"
		
	}
	return "Geçti"
	
}
func main(){

	o1:=Ogrenci{
		isim:"Mehmet",
		yas:47,
		notlar:[]int{90,95,100,100},
		gectiMi:true,
	}
	
	//fmt.Println("Ogrencinin adı :", o1.AdGoster())
	//o1.BilgileriGoster()
	//fmt.Println(o1.ortalamaHesapla())

	fmt.Printf("Öğrencinin ortalaması: %f\n Öğrenci Durumu: %s",o1.ortalamaHesapla(),o1.gectimiKaldimi())	
	

}