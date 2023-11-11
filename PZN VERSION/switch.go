package main

import "fmt"

func main() {
	name := "MaAgusKecuts"

	switch name {
	case "MaAgusAsem":
		fmt.Println("Ya ini si mama asem")
	case "MaAgusKecut":
		fmt.Println("Ya ini si mama kecut")
	default:
		fmt.Println("Mama siapa nihh")
	}

	//short switch expreesion
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu mesum si mama", length+1)
	case false:
		fmt.Println("Kurang mesum si mama")
	}

}
