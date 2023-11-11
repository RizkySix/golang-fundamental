package main

import "fmt"

func main() {
	type noKtp string

	var ktpJanu noKtp = "6969696969"

	var contoh string = "222222222"
	var contohKtp noKtp = noKtp(contoh)

	fmt.Println(ktpJanu)
	fmt.Println(contohKtp)
}