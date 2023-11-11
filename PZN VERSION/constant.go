package main

import "fmt"

func main() {
	const firstName string = "Buk Agus"
	const lastName = "Janu"

	//ini akan error
	/* firstName = "Asem"
	lastName = "MSM" */

	//multi constant
	const (
		name string = "Ma agus"
		childName = "Asem"
	)

	fmt.Println(name)
	fmt.Println(childName)
}