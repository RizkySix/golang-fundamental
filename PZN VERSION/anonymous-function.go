package main

import "fmt"

// alias
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu di block", name)
	} else {
		fmt.Println("Halo", name)
	}
}

func main() {
	//bisa dimasukan kedalam variable dulu
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Janu", blacklist)

	//atau langsung kedalam paramternya
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
