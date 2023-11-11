package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("ini angka ke = ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("ini angka ke = ", counter)
	}

	//slice names
	names := []string{
		"Buk",
		"Agus",
		"Kecut",
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//foreach versi golang namanya for range
	for index, value := range names {
		fmt.Println("Index = ", index, " Value = ", value)
	}

	for _, value := range names {
		fmt.Println("Hanya value = ", value)
	}
}
