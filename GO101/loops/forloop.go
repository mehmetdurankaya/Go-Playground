package main

import "fmt"

func main(){
	/*
		Golang while döngüsü kullanılmıyor for döngüsünün iki farklı kullanımı var
		
		//for döngüsü diğer
		x:=0 // diğer programlama dillerindeki while döngüsü gibi
		for x<5{
		fmt.Println(x)
		x+=1
	}
	*********************************************
		for i:=0;i<5;i++{}  //2. şekil
	x:=0
	for x<5{
		fmt.Println(x)
		x+=1
	}
	*********************************************
	

	for x:=0;x<=10;x++{
		fmt.Println(x)
	}
	
	

	//çift sayıları bulma

	for i:=0;i<=100;i++{
		if i%2==0{
			fmt.Println(i)
		}
	}
	*/

	toplam:=0
	for i:=1;i<10;i++{
		toplam+=i
	}
	fmt.Println(toplam)// 1+2+3+4+5+6...
}