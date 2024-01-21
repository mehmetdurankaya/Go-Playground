package main

import (
	"bufio" // buffer input output
	"fmt"
	"os"      // operation system
	"strconv" //string conversion
)
func main(){
	/*
	scanner:=bufio.NewScanner(os.Stdin)//standart input device

	fmt.Print("Lütfen Birşeyler Yazınız: ")
	scanner.Scan()

	dataentry:=scanner.Text()
	fmt.Printf("Bunu Yazdınız: %s",dataentry)
	
	*/
	scanner:=bufio.NewScanner(os.Stdin)//standart input device

	fmt.Print("Hangi yıl doğdunuz: ")
	scanner.Scan()

	data,_:=strconv.ParseInt(scanner.Text(),10,64)


	fmt.Printf("Şuanda %d yaşındasın.", 2024-data)


}