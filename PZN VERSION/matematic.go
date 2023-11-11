package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a + b

	fmt.Println(c)

	//augmented namanya
	d := 10
	d += 20
	d += 5
	
	fmt.Println(d)

	//unary operator
	e := 10
	e++
	e++

	fmt.Println(e)

}