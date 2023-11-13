package main

import "fmt"

func ups() any {
	//return 1
	//return true
	return "Semua data bisa direturn dengan any"
}

func main() {
	kosong := ups()
	fmt.Println(kosong)
}
