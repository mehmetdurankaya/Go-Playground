package main

import "fmt"

/*
	EĞER SAYI 3 VE 3' ÜN KATI İSE FIZZ
	EĞER SAYI 5 VE 5 'İN KATI İSE BUZZ
	EĞER HER İKSİNİNDE ORTAK KATI İSE FIZZBUZZ

*/

func main(){
	for i:=1;i<=100;i++{

		if i%3==0 && i%5==0{
			fmt.Println("FIZZBUZZ")
		}else if i%3==0{
			fmt.Println("FIZZ")
		}else if i%5==0{
			fmt.Println("BUZZ")
		}else{
			fmt.Println(i)
		}
	}

}