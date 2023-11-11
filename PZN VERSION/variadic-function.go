package main

import "fmt"

func manualSumAll(baseNum int, numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return baseNum + total
}

// ketimbang manual menggunakan tipe paramnya slice seperti diatas, lebih baik gunakan variadic param
func sumAll(baseNum int, numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return baseNum + total
}

func main() {
	//manual membuat slice
	manualTotal := manualSumAll(150, []int{10, 10, 10, 10, 10, 10})
	fmt.Println(manualTotal)

	//menggunakan variadic param
	total := sumAll(150, 10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	//anggap jika sudah terlanjur memiliki data dalam bentuk slice, kita msih dapat mengirimnya ke variadic function
	numbers := []int{10, 10, 10, 10, 10, 10} //slice
	total = sumAll(150, numbers...)
	fmt.Println(total)
}
