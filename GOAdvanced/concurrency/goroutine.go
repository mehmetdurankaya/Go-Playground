package main

/*
	go routine lerde geçiktirme kullanılmalıdır

*/

import (
	"fmt"
	"sync"
	//"time"
)



func yazdir(s string) {
	for i:=0;i<=5;i++{
		fmt.Println(s)
	}
	
}

func main() {

	/*
	//Thread 1 MB Goroutine 2KB

	go yazdir("Merhaba")
	time.Sleep(1 * time.Second)
	yazdir("Dunya!")
   */

   wg:=sync.WaitGroup{}
   wg.Add(1)
   go func(){
	for i:=0;i<5;i++{
		fmt.Println("Goroutine")
	}
	wg.Done()
}()

wg.Wait()

fmt.Println("Ana Thread")

}