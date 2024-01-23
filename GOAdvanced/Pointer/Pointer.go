package main

import "fmt"

/*
	Pointer bellek adresini tutan yapılar
	belleği direkt yönetmemizi sağlar
	değişken gibi tanımlanır ancak & ampersan işareti kullanılır

*/
func main(){
	x:=10
	fmt.Println(x)
	fmt.Println(&x)

	p:=&x
	fmt.Println(*p) 
	// * asterix pointerlerin bellek içerisindeki tuttukları değerleri verir
	*p=44
	fmt.Println(x) 

	fmt.Printf("%T,%v\n",x,x)
	fmt.Printf("%T,%v\n",p,p)

	
}

