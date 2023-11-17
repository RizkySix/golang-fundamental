package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Jangan menggunakan 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", result)
	} else {
		fmt.Println("Error", err.Error())
	}
}
