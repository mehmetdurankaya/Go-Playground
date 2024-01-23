package main



import "fmt"

type calisan struct {
	isim      string
	yas       int
	maas      int
	isHaveKid bool
}
type mudur struct{
	calisan
	unvan string
}

func main() {

	a1 := calisan{
		isim:      "Ahmet",
		yas:       46,
		maas:      4250,
		isHaveKid: true,
	}
	fmt.Println(a1)

	y1:=mudur{
		calisan: calisan{
			isim:      "Mehmet",
			yas:       47,
			maas:      50000,
			isHaveKid: true,
		},		
		unvan: "Bilgiişlem Müdürü",
	}
	fmt.Println(y1)

}