package main



import (
	"fmt"
	"math"
)
/*
	Built-in function gömülü fonksiyonlar
	user-defined function kullanıcının oluşturduğu fonksiyonlar
	parametresiz
	parametreli
	return dönüş tipli
	multiplereturn çoklu dönüş tipli fonksiyonlara örnek verilmiştir
	isimlendirilmiş dönüş değerleri (return)

*/
const PI=3.14

func merhaba(){    //Parametresiz fonksiyon
	fmt.Println("Merhaba")
}
func parametreliFunc(mesaj string){    //Parametreli fonksiyon
	fmt.Println("Merhaba" , mesaj)
}

func toplama(x int, y int)int{ // return dönüş tipli fonksiyon
	return x+y
}
func onlaCarp(t int)int{
	return t*10
}

func bolme(a,b float32)float32{
	return a/b
}
func ikiyleCarp(g int)int{
	return g*2
}
func ucIslem(x,y int)(int,int,int){//çoklu return
	return x-y,x+y,x*y
}
func maxmin(a,b int)(int,int){//çoklu return
	if a>b{
		return a,b
	}else{
		return b,a
	}
}
func hesaplama(a,b int)(carp int,bol int ){
	carp=a*b
	bol=a/b
	return
}

func daire(yaricap float64)(alan float64, cevre float64){
	alan=PI*math.Pow(yaricap,2.0)
	cevre=2*PI*yaricap
	return
}

func main(){
	/*merhaba()
	parametreliFunc("Mehmet Duran Kaya")
	
	fmt.Println(toplama(7,9))
	
	a:=onlaCarp(10)//fonksiyonu değişkene ataya biliyoruz
	fmt.Println(a)
	
	for i:=0;i<=10;i++{
		fmt.Println(ikiyleCarp(i))
	}
	
	sayi1,sayi2,sayi3:=ucIslem(5,9)

	fmt.Printf("Sonuc: %d \n",sayi1)
	fmt.Printf("Sonuc: %d \n",sayi2)
	fmt.Printf("Sonuc: %d \n",sayi3)
	

	max,min:=maxmin(10,60)
	fmt.Printf("Max= %d Min=%d",max,min)
	

	carpim,bolme:=hesaplama(27,3)

	fmt.Printf("Carpim: %d Bolme: %d",carpim,bolme)
	
	*/

	dairealan,dairecevre:=daire(5.0)
	fmt.Println("Dairenin Alanı: ",dairealan)
	fmt.Println("Dairenin Çevresi: ",dairecevre)

}