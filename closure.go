package main

import "fmt"

func main(){
	number1 := 20
	number2 := 10
	closure := callback(func(number1 ,number2 int) int {
		return number1*number2
	} , number1 , number2 )

	fmt.Println(closure)
}

func callback(call func(int , int) int , number1 , number2 int) int {
	return call(number1 , number2) + 100
}

