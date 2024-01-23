package main

/*
	Golang'da iota, genellikle sabit değerlerin ardışık olarak artan değerler almasını sağlamak için kullanılan
	özel bir anahtardır. iota, genellikle const bloğu içinde kullanılır ve bu blok içindeki ilk sabit değere 0, sonraki değere 1, bir sonrakine 2 ve böyle devam ederek otomatik olarak artan değerleri temsil eder.


*/

import (
	"fmt"
	"time"
)

type ustunf float64// varsayılan veri tipinin yanında kendi veri tipimizi belirliyoruz 

func(y ustunf) carpBes()ustunf{
	return y*5.2
}

type aylar int // aylar adında int tipinde bir veri tipi oluşturdum

const ( //bir sabit oluşturdum ve bu sabite oluşturmuş olduğum veri tipini referans aldım
	Ocak aylar = 1+iota // bu satırda ilk deger atadım ve 1 den başlayarak numaralandır demiş oldum
	Subat
	Mart
	Nisan
	Mayıs
	Haziran
	Temmuz
	Agustos
	Eylul
	Ekim
	Kasim
	Aralik
)

func main() {

	var u1 ustunf = 5.4

	fmt.Printf("%T,%v \n",u1,u1)

	var u2 float64=6.77
	fmt.Println(u1+ustunf(u2))//casting

	y1:=ustunf(50.5)
	fmt.Println(y1.carpBes())
	_,ay,_:=time.Now().Date()
	fmt.Println(ay)
	fmt.Println(aylar(ay))

}