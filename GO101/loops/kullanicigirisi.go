package main

import (
	"bufio"
	"fmt"
	"os"
)


func main(){
	k_adi:="MDK"
	k_sifre:="12345"

	giris_hakki:=3

	fmt.Println("****LOGIN****")

	for true{
		scanner:=bufio.NewScanner(os.Stdin)
		fmt.Print("Kullanıcı Adını Giriniz: ")
		scanner.Scan()
		giris_kadi:=scanner.Text()

		fmt.Print(" Kullanıcı Şifresini Giriniz: ")
		scanner.Scan()
		giris_ksifre:=scanner.Text()

		if k_adi != giris_kadi && k_sifre !=giris_ksifre{
			fmt.Println("Kullanıcı adı ve şifreniz yanlış")
			giris_hakki--
		}else if k_adi== giris_kadi && k_adi!=giris_ksifre{
			fmt.Println(" Şifreniz yanlıştır")
			giris_hakki--
		}else if k_adi!=giris_kadi && k_adi == giris_kadi{
			fmt.Println("Kullanıcı adı yanlış")
			giris_hakki--
		}else if k_sifre != giris_ksifre && k_sifre == giris_ksifre{
			fmt.Println("Kullanıcı adı yanlış")
			giris_hakki--
		}else{
			fmt.Println("Hosgeldiniz")
			break
		}
		if giris_hakki == 0 {
			fmt.Println("Üzgünüm giriş Hakkınız Sona erdi.")
			break
		}
	}
	
}