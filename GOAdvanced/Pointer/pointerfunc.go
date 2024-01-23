package main

import "fmt"

/*



*/


func isaret(ptr *int){
	*ptr=55 //dereferencing
}

func onlacarp(sayi *int)int{
	// *sayi*=10
	// fmt.Println(*sayi)
	return *sayi*10


}
func main(){
	x:=100

	isaret(&x)
	fmt.Println(x)
	//onlacarp(&x)

	fmt.Println(onlacarp(&x))
}