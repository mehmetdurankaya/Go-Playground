package main

import(
	"os"
	"bufio"
	"fmt"
	"strconv"
	"math"
)

func main(){
	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Print("A sayısını giriniz: ")
	scanner.Scan()
	a,_:=strconv.ParseFloat(scanner.Text(),64)

	fmt.Print("B sayısını girniiz: ")

	scanner.Scan()
	b,_:=strconv.ParseFloat(scanner.Text(),64)
	
	fmt.Print("C sayısını griniz: ")
	scanner.Scan()
	c,_:=strconv.ParseFloat(scanner.Text(),64)

	delta:=(math.Pow(b,2))-4*a*c

	kokbir:=(-b-(math.Pow(delta,0.5)))/2*a
	kokiki:=(-b+(math.Pow(delta,0.5)))/2*a

	fmt.Printf("1.Kök: %f\n2.Kök: %f\n Delta: %f",kokbir,kokiki,delta)


}
