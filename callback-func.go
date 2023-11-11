package main

import "fmt"

func main(){
	number1 := 10
	number2 := 10
	fmt.Println(callback(kali , number1 , number2))
}

func kali(number1 int, number2 int) int{
	return number1 * number2
}

func callback(call func(int , int) int , number1 int , number2 int) int {
	return call(number1 , number2) + 100
}