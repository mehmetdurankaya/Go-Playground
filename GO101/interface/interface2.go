package main

import (
	"fmt"
	"math"
)

const PI = 3.14

type hesapla interface{
	alan() float64
	cevre() float64
}
type daire struct {
	yaricap float64
}
type dikdortgen struct {
	kisakenar float64
	uzunkenar float64
}

func (d daire) alan() float64 {
	return PI * math.Pow(d.yaricap,2)
}
func(d daire) cevre()float64{
	return 2*PI*d.yaricap
}
func(g dikdortgen)alan()float64{
	return g.kisakenar*g.uzunkenar
}
func(g dikdortgen)cevre()float64{
	return 2*(g.kisakenar+g.uzunkenar)
}
func hesapYap(h hesapla){
	fmt.Println(h)
	fmt.Println(h.alan())
	fmt.Println(h.cevre())

	fmt.Printf("Veri Tpi: %T \n",h )

}
func main() {
	daire1:=daire{5.0}
	dikdortgen1:=dikdortgen{6.0,9.0}

	hesapYap(daire1)
	hesapYap(dikdortgen1)
	
	fmt.Println("Daire1 nesnesinin alanı : ",daire1.alan())
	fmt.Println("Daire1 nesnesinin çevresi: ",daire1.cevre())
}