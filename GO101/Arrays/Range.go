package main

import "fmt"

func main(){

	
	


	var a[]int=[]int{1,5,6,7,8,44}
	/*
	for i:=0;i<len(a);i++{
		fmt.Println("index: ",i,a[i])
	}
	*/

	for index,eleman:=range a{
		fmt.Printf("%d: %d \n",index,eleman)
	}

	isim:="Mehmet Duran Kaya"

	for i,harf:=range isim{
		fmt.Printf("%d = %c\n",i,harf)
	}

	
}