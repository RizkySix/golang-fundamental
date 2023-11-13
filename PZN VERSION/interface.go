package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name    string
	Address string
}

type Animal struct {
	Name string
}

// struct method
func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

func main() {
	person := Person{Name: "Malika"}
	sayHello(person)

	animal := Animal{Name: "Buaya"}
	sayHello(animal)
}
