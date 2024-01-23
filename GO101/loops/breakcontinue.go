package main

import "fmt"


func main(){
/*
	//Break döngüyü kırmak için kullanılır
	for i:=0;i<10;i++{
		if i==5{
			break
		}
		fmt.Println(i)
	}
	
*/
	//continue atlama noktası oluşturuyor aşağıda 5' e eşit olduğunda 1 adım atlıyor ve döngünün başına dönüp devam ediyor
	for i:=0;i<10;i++{
		if i==5{
			continue
		}
		fmt.Println(i)
	}

}