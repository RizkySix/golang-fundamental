package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	contoh := getHello
	misal := getHello

	fmt.Println(contoh("Mamah"))
	fmt.Println(misal("Asem"))
}
