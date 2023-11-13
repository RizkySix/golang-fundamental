package main

import "fmt"

func endApp1() {
	fmt.Println("End App")
}

func endApp2() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

// contoh penggunakan recover yang salah
func runAppWrong(error bool) {
	defer endApp1()

	if error {
		panic("Waduh Error")
	}

} // contoh penggunakan recover yang benar
func runAppCorrect(error bool) {
	defer endApp2()

	if error {
		panic("Waduh Error")
	}

	fmt.Println("Lanjut...")
}

func main() {
	//runAppWrong(true)
	//fmt.Println("Janu ganteng")

	runAppCorrect(true)
	fmt.Println("Janu ganteng")
}
