package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)


func main(){
	Pisayisi:=math.Pi
	fmt.Println(Pisayisi)

	fmt.Println(math.Pow(2,10))
	fmt.Println(math.Sqrt(64))
	fmt.Println(math.Pow(64,0.5))


	fmt.Println(math.Abs(-55.4))

	fmt.Println(math.Sin(37))
	fmt.Println(math.Cos(37))
	fmt.Println(math.Tan(37))
	fmt.Println(math.Mod(44,3))
	fmt.Println(math.Round(22.5))
	fmt.Println(math.Ceil(28.6))
	fmt.Println(math.Floor(28.6))
	fmt.Println(math.Max(28,89))
	fmt.Println(math.Min(28,89))


	fmt.Println(math.IsNaN(400))
	fmt.Println(math.Log(100))
	fmt.Println(math.Log10(100))


	rand.Seed(time.Now().UnixNano()) //1 ocak 1970 ten sonraki süreyi UnixNano üzerinden hesaplıyoruz
	rastgele:=rand.Intn(100)

	fmt.Println(rastgele)

	



	

	


}