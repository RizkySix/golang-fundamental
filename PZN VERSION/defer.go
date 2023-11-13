package main

import "fmt"

func logging() {
	fmt.Println("Record logging here...")
}

func runApplication() {
	defer logging() //karena menggunakan deffer, maka akan selalu dieksekusi paling akhir walaupun terdapat error pada saat proses eksekusi runApplication()
	fmt.Println("Run Application OK")
}

func main() {
	runApplication()
}
