package main

import "fmt"

func abc(t int) int {
	return t*10
}

func main() {
	any := func(toplam int) int {
		return toplam / 3 * 4 * 5
	}(10)
	fmt.Println(any)
	fmt.Println(abc(any))
}