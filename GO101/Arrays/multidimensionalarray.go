package main

import "fmt"

func main(){

	/*
	// 2 Boyutlu Dizi
	var dizi=[2][2]int{{0,1},{2,4}}//iki satır iki sütunluk bir dizi

	fmt.Println(dizi)
	fmt.Println(dizi[0][1])
	fmt.Println(dizi[1][0])

	*/
	var dizi2=[5][2]int{
		{0,0},
		{1,2},
		{2,4},
		{3,6},
		{4,8}}
		fmt.Println(dizi2[3][0])//3
		fmt.Println(dizi2[4][1])//8
		fmt.Println(dizi2[3][1])//6
		fmt.Println("*******************")
		for i:=0;i<5;i++{
			for j:=0;j<2;j++{
				fmt.Printf("Dizi 2D[%d][%d]= %d\n",i,j,dizi2[i][j])
			}
		}

		//3 Boyutlu Diziler

		var dizi3D=[3][3][3]int{
			{{1,2,3},{1,2,3},{1,2,3}},
			{{1,2,3},{1,2,3},{1,2,3}},
			{{1,2,3},{1,2,3},{1,2,3}}}

			fmt.Println(dizi3D)


			fmt.Println(dizi3D[0][2][1])//2



}