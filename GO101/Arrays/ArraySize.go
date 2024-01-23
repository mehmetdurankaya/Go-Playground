package main

import "fmt"

/*
	Array statik yapılıdır. List dinamik yapılıdır.


*/

func main() {
	/*
	 arr:=[3]int{1,2,3,4,5,6} //index 3 is out of bounds (>= 3) hatası
	fmt.Println(arr)
	
	 arr:=[...]int{1,2,3,4,5,6} // [...] eleman sayısı kadar array açar
	fmt.Println(arr)
	
	arr:=[10]int{1,2,3,4,5,6}
	arr[15]=99 //index 15 out of bounds [0:6] 15 index olmadığından alınan hata
	fmt.Println()
	*/
	arr:=[...]int{1,2,3,4,5,6}
	//len() komutu kaç eleaman olduğunu gösterir
	fmt.Println(len(arr))

	//cap komutu bir dizinin kaç elaman alabileceğini söyler
	fmt.Println(cap(arr))

	for i:=0;i<=len(arr)-1;i++{
		fmt.Println(arr[i])
	}
}