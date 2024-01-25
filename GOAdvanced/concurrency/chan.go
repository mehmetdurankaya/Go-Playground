package main

import (
	"fmt"
	"time"
)

/*



 */


func main(){

/*
k:=make(chan string)

go func(){
	k<-"kanaldan gelen veri"
	
}()

mesaj:=<-k


fmt.Println(mesaj)

superK:=make(chan bool,3)

go func(){
	superK<-true
	time.Sleep(1* time.Second)
	superK<-false
	time.Sleep(1* time.Second)
	superK<-true

}()
fmt.Println(<-superK,<-superK,<-superK)
*/

sehirler:=[]string{"Osmaniye","Adana","Kayseri","Ordu","Rize"}

sehirKanal:=make(chan string)
//gönderen
go func (dizi []string){

	for _,sehir:=range dizi{
		sehirKanal<-sehir
	}
}(sehirler)

go func(){
	for i:=0;i<len(sehirler);i++{
		alinan:=<-sehirKanal
		fmt.Println("Kanaldan Gelenler",alinan)
	}
}()

<-time.After(time.Second*3)



}