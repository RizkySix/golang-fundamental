package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Tabanan", "Bali", "Indonesia"}
	address2 := address1 //valuenya dicopy dari address 1

	address2.City = "Badung"
	fmt.Println(address1) //tidak berubah
	fmt.Println(address2)

	//berikut adalah penerapan pointer mirip seperti pada PHP menggunakan operator &
	address3 := &address1

	address3.City = "Denpasar"
	fmt.Println(address1) //akan berubah
	fmt.Println(address3) //akan berubah
	fmt.Println(address2) // akan tetap Badung karena address 2 tidak pass by referece/pointer ke address 1
}
