package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// * berarti tipe datanya itu harus pointer/referece
func changeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Tabanan", "Bali", ""}
	changeCountry(&address) //gunakan & untuk menandakan/mengirim sebagai pointer/reference
	fmt.Println(address)    // data country akan berubah
}
