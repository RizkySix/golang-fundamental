package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var rizky Customer
	rizky.Name = "Rizky Janu"
	rizky.Address = "Kuala Lumpur"
	rizky.Age = 18
	fmt.Println(rizky)
	fmt.Println(rizky.Name)
	fmt.Println(rizky.Address)
	fmt.Println(rizky.Age)

	//bisa juga seperti ini
	mamah := Customer{
		Name:    "Mamah Agus",
		Address: "Indonesia",
		Age:     46,
	}

	fmt.Println(mamah)

	//atau seperti ini tapi urutan fieldnya harus sesuai
	mamah2 := Customer{"Mamah Rica", "Indonesia", 46}

	fmt.Println(mamah2)

	//struct method
	mamah.sayHello("Adik")
}
