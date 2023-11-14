package main

import "fmt"

func random() any {
	return true
}

func main() {
	result := random()
	//resultString := result.(string)

	//fmt.Println(resultString)

	//resultInt := result.(int) // akan terjadi panic karena type nya sudah dirubah ke string sebelumnya
	//fmt.Println(resultInt)

	//lebih baik lakukan seperti ini
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("unknown", value)

	}

}
