package main

import "fmt"

func main() {
	name := "MaAgusAsem"

	if name == "MaAgusKecut" {
		fmt.Println("Ya dia kecut")
	} else if name == "MaAgusAsem" {
		fmt.Println("Ya dia asem")
	} else {
		fmt.Println("Salah eh")
	}

	//short if expreesion
	if leght := len(name); leght > 5 {
		fmt.Println("Mama terlalu mesum", leght+1)
	} else {
		fmt.Println("Mama kurang mesum")
	}

}
