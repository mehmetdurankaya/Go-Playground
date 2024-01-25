package main

import "fmt"

/*
	iki den fazla işlemin birbirlerini beklemesi sonucu karşılaşılan hata
	sen çalış ben bekliyorum der her iki tarafta ancak iki tarafta haraket etmez ve kilitlenme olur birbirlerini bloke ettiklerinden dolayı bu hataya deadlock diyoruz.

	çözüm yollarından biri kanal kapasitesini bir arttırmaktır

*/

func main(){
    /*
	ch:=make(chan int,2)

	ch<-1
	fmt.Println(<-ch)
	
	*/

	kanal:=make(chan string ,5)

	kanal<-"AUDI"
	kanal<-"FORD"
	kanal<-"KIA"
	kanal<-"LADA"

	fmt.Printf("Kanal Kapasitesi%d\n",cap(kanal))
	fmt.Printf("Gelen Veri Boyutu%d\n",len(kanal))
	fmt.Printf("Alınan Veri%s\n",<-kanal)
	fmt.Printf("Yeni Gelen Veri Boyutu%d\n",len(kanal))
	

}