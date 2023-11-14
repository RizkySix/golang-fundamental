package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// new hanya bisa digunakan ke data kosong
	address1 := new(Address) // ini adalah pointer/reference ke stuct Address kosong
	address2 := address1     //address2 juga akan menjadi pointer/reference yang mengacu ke sumber yang sama yaitu Address kosong

	address2.Country = "Indonesia"

	// Country kedua address akan berubah karena mengacu ke pointer/reference yang sama
	fmt.Println(address1)
	fmt.Println(address2)
}
