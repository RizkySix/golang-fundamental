package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Buk"
	names[1] = "Agus"
	names[2] = "Asem"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//tidak seperti php penambahan index tidak bisa dinamis
	//names[3] = "hot" //ini akan error
	//names[0] = 1 //ini jg akan error karena default tipe data nya string

	//fmt.Println(names[0])

	//memberi array value secara langsung
	var values = [3]int {100,90,85}

	fmt.Println(values)


	//untuk membuat panjang array yang panjang tidak pasti diawal
	var array = [...]int{
		100,
		40,
		50,
		35,
		20,
	}

	fmt.Println(len(array))
}