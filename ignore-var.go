package main

import "fmt"

func main() {
	firstName, _ := getNames() //kita hanya memperdulikat value pertama, dan skip value kedua dg simbol "_"
	fmt.Println("Welcome to Textio,", firstName)
}

// don't edit below this line

func getNames() (string, string) {
	return "John", "Doe" //return 2 value
}
