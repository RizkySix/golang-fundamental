package main

import "fmt"

type oddEven interface {
	modulus(num int) bool
	resultMessage() string
}

type num struct {
	num int
	targetNum int
}

func (n num) modulus(num int) bool {
	mod := num % 2

	if mod  == 0 {
		return true
	}
	return false
}

func (n num) resultMessage() string {
	mod := n.modulus(n.num)
	msg := fmt.Sprintf(`Input number are invalid Odd`)
	if n.num > n.targetNum && mod == true {
		msg = fmt.Sprintf(`Input number %d is even and valid because %d highest then %d` , n.num , n.num , n.targetNum)
	}
	if n.num <= n.targetNum && mod == true {
		msg = fmt.Sprintf(`Input number %d is even but invalid because %d lowest then %d` , n.num , n.num , n.targetNum)
	}

	
	return msg
}

func test(o oddEven) {
	fmt.Println(o.resultMessage())
	fmt.Println("===================")
}

func main(){
	data := num{
		num:8,
		targetNum:4,
	}
	test(data)
	data = num{
		num:4,
		targetNum:9,
	}
	test(data)
	data = num{
		num:3,
		targetNum:1,
	}
	test(data)
}