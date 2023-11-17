package helper

import "fmt"

// penggunaan access modifier, huruf awalnya capital maka bisa diakses diluar package yang sama misal pada main, jika kecil hanya di package itu saja
var version = "1.0.0"      //tidak bisa diakses diluar helper
var Application = "golang" //bisa diakses

func sayGoodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Helo " + name
}

func Contoh() string {
	return fmt.Sprintf("%s %s", sayGoodbye("Jono"), version)
}
