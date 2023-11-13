package main

import "fmt"

// dengan loop
func loopFactorial(num int) int {
	result := 1
	for i := num; i > 0; i-- {
		result *= i
	}

	return result
}

// dengan recursive
func recursiveFactorial(num int) int {
	if num == 1 {
		return num
	} else {
		return num * recursiveFactorial(num-1)
	}
}

func main() {
	fmt.Println(loopFactorial(10))
	fmt.Println(recursiveFactorial(10))
}
