package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Say hay", name)
}

func getName(name string) string {
	return "Aku adalah " + name
}

// return multiple value
func getFullName() (string, string) {
	return "Mamah", "Hot"
}

// named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Mama"
	middleName = "Agus"
	lastName = "Kecut"

	return firstName, middleName, lastName
}

func main() {
	sayHello("Joko")
	sayHello("Widodo")

	fmt.Println(getName("Tocil"))

	//ambil data return multiple value
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	//hanya ambil salah satu return multiple value
	onlyFirstName, _ := getFullName()
	fmt.Println(onlyFirstName)

	name1, name2, name3 := getCompleteName()
	fmt.Println(name1, name2, name3)
}
