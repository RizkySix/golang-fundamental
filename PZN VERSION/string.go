package main

import "fmt"

func main() {
	fmt.Println("String 1")
	fmt.Println("String 2")

	string1 := fmt.Sprintf("Test String")
	fmt.Println("Total String = " , len(string1))
	fmt.Println("Karakter index ke 4 = " , string1[4])
}