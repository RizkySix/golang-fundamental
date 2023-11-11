package main

import "fmt"

func main() {
	/* var person map[string]string = map[string] string{}
	person["name"] = "Rizky"
	person["address"] = "Tabanan" */

	person := map[string]string{
		"name" : "Rizky",
		"address" : "Tabanan",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	person["old"] = "18"
	fmt.Println(person)

	//akan mengeprint string kosong '' karena tidak ada key "salah"
	fmt.Println(person["salah"])


	//map dapat didelete seperti ini dan dibuat tanpa default value seperti ini 
	book := make(map[string]string) //map[string]string{} atau begni langsung

	book["title"] = "Kurcaci-Ngaceng"
	book["author"] = "janu"
	book["wrong"] = "salahhh"
	
	fmt.Println(book)
	delete(book, "wrong")

	fmt.Println(book)
}