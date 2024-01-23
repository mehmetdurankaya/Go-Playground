package main


/*
	Değişken sayıda parametreyi kabul ediyor bazı durumlarda fonksiyonun kaç parametre alacağını bilemeye biliriz bu yüzden arguman sayısı kadar parametre alan variadic function ları kullanıyoruz
	fonksiyon çağrıldığında içine girilen değer kadar parametre oluşturuyor

	arguman: fonksiyon çağrıldığında gönderilen değerlere denir
*/
import "fmt"

func merhaba(hello ...string){
	fmt.Println(hello)
}

func sehirler(param...string){
	for _,sehir:=range param{
		fmt.Println(sehir)
	}
}

func toplama(sayilar ...int)int{
	sonuc:=0
	for _,sayi:=range sayilar{
		sonuc+=sayi
	}
	return sonuc
}

func main(){
	/*
	merhaba("Merhaba Mehmet","Duran","Kaya")
	
	sehirler("Istanbul","Ankara","Hatay")
	*/ 
	son:=toplama(2,3,4,5,4,78)
	fmt.Println(son)
}