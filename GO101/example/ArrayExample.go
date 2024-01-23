package main

import "fmt"

func main(){

	var dizi1 [2][3]int
	var dizi2 [2][3]int

	

	var dizisonuc [2][3]int

	fmt.Print("1.Dizinin Elemanları nı Giriniz:\n")

	for k,r:=range dizi1{
		for l:=range r{
			fmt.Scan(&dizi1[k][l])
		}
	}

	fmt.Print("2.Dizinin Elemanlarını Giriniz:\n")

	for m,rr:=range dizi2{
		for n:=range rr{
			fmt.Scan(&dizi2[m][n])
		}
	}
	fmt.Print("Dizilerin Çarpımı Sonucu:\n")

	for i,satir:=range dizisonuc{
		for j:=range satir{
			dizisonuc[i][j]=dizi1[i][j]*dizi2[i][j]
			fmt.Print(dizisonuc[i][j],"\t")
		}

	}

	fmt.Println()
}