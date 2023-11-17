package main

import (
	"fmt"
	"project-golang/helper"
)

func main() {
	result := helper.SayHello("Janu")
	fmt.Println(result)

	fmt.Println(helper.Application)
	//fmt.Println(helper.version)            // tidak akan bisa
	//fmt.Println(helper.sayGoodbye("Yono")) // tidak akan bisa
	fmt.Println(helper.Contoh())
}
