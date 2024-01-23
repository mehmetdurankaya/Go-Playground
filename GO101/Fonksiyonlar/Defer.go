package main

/*
	Defer içerisinde bulunduğu fonksiyonların işlemleri tamamlanana kadar
	bulunduğu satırı geçiktirir
	programlama dilleri yukarıdan aşağıya doğru satır satır çalışırlar aşağıdaki örneklerde önce çalışması gereken kod satrı defer yapısı ile geciktirilerek fonksiyonun işlemi bittikten sonra devreye girer ve bulunduğu satırı çalıştırır
	veri tabanı işlemlerinde klasör işlemlerinde sıklıklar kullanılır
	örneğin veritabanı bağlantısı yapıldığında bağlantınının işlem sonunda kapatılması gerekiyor en başta defer yapısı kullanılırsa fonksiyon işlemini bitirdikten sonra otomatik olarak devreye girerek bulunduğu satırı çalıştırıp veritabanı bağlantısını kesecektir
*/
import "fmt"

func start(){
	defer fmt.Println("PC AÇILDI")
	fmt.Println("PC AÇILIYOR")
}
func stop(){
	defer fmt.Println("PC KAPANDI")
	fmt.Println("PC KAPANIYOR")
	
}
func sira(){
	defer fmt.Println("Birinci")//en başta tanımlanan sonda çalışır
	defer fmt.Println("İkinci")//sıralama sondan başa doğru ilerler
	defer fmt.Println("Üçüncü")//1. olarak çalışır
	fmt.Println("fonksiyon başlıyor")
}
func main(){

//start()
//stop()
sira()


}