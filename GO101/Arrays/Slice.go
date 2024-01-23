package main

import "fmt"

/*
	SLICE dilim slice_a:=array_a[2:5] dizinin 2. indexinden 5.indexine kadar(5 dahil değildir) olan alanı dilimle
	[:5] 5. indexe kadar
	[0:] 0. indexden sona kadar
	[:] tüm diziyi al

*/
func main(){
	/*
	
	
	
	dizi:=[5]int{1,3,5,7,9}
	

	var dilim[]int =dizi[1:4] //[1:4]1.index dahil 4.index' e kadar 

	fmt.Println(dilim)
	*/
	a:=[5]string{"Go","Python","Kotlin","PHP","Javascript"}
	//b:=a[1:3]
	//b:=a[1:]
	b:=a[1:3]
	//Slice lar go programlama dilinde uzunluğunun iki katı olarak tanımlanır
	fmt.Println("Slice:",b)
	fmt.Println("Slice Kapasitesi:",cap(b))
	fmt.Println("Slice Uzunluğu:",len(b))
	fmt.Println("Array:",a)


}