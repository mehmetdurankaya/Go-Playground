package main

import "fmt"


func main(){
	/*
	for i:=1;i<=5;i++{
		for j:=1;j<=5;j++{
			fmt.Printf("İ:%d J:%d \n",i,j)
		}
	*/
	for i:=1;i<=10;i++{  //çarpım tablosu
		fmt.Println("*******************")
		for j:=1;j<=10;j++{
			fmt.Printf("%d X %d = %d\n",i,j,i*j)
		}
	}

	
}
