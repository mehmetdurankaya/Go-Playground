package main
/*
	Fibonacci
	0,1 1,2 ,3 ,5, 8, 13, 21 ,34, 55, 89, 144, 233, 377, ...

	f(n)=f(n-1)+(f(n-2))


*/
import "fmt"
func fibo(n int)int{
	a:=0
	b:=1
	for i:=0;i<n;i++{
		gecici:=a
		a=b
		b=gecici+a

	}
	return a
}

func fibore(n int)int{
	if n<=1{
		return n
	}
	return fibore(n-1)+fibore(n-2)
}

func main(){
		/*
	for n:=0;n<=20;n++{
		sonuc:=fibo(n)
		fmt.Printf("Fibonacci:%d = %d \n", n,sonuc)
	}
	*/
	for i:=0;i<=20;i++{
		sonuc:=fibo(i)
		fmt.Printf("Fibonacci:%d = %d \n", i,sonuc)
	}
}