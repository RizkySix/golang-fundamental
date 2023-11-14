package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Tabanan", "Bali", "Indonesia"}
	address2 := &address1 //valuenya dicopy dari address 1

	address2.City = "Badung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("=================================")

	address2 = &Address{"Malbau", "Jateng", "Indonesia"}
	fmt.Println(address1) // akan tetap menjadi Badung
	fmt.Println(address2) // akan berubah ke data terbaru
	fmt.Println("=================================")

	//asterisk operator akan merubah seluruh variabel yang terkait dengan value reference lama dari address3 ke sumber yang terbaru
	address3 := &address1
	address4 := address3
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("=================================")

	//gunakan operator *
	*address3 = Address{"Sukarja", "Toronto", "France"}
	fmt.Println(address1) // akan berubah karena address1 merupakan value reference awal address3
	fmt.Println(address4) // akan berubah karena address4 merupakan reference address 3 yang dimana sumbernya dari address 1
	fmt.Println(address3) // akan berubah

	fmt.Println(address2) // tidak akan terpengaruh karena address2 tidak lagi reference ke address1, valuenya telah di rewrate pada line 18
}
