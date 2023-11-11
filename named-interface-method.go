package main

import "fmt"

type shape interface {
	area(width int , length int) (area int) //area menerima 2 parameter int dan harus mengembalikan int
}

type item struct {
	name string
	width int
	length int
}

func (i item) area(width , length int) (area int){
	area = width * length
	return
}

func test(s shape , i item){
	fmt.Printf("this %s shape and have area %dcm" , i.name, s.area(i.width , i.length))
}

func main(){
	i := item{
		name : "box",
		width : 15,
		length : 2,
	}

	test(i , i)
}

//full course tentanng named method interfaces pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/57e27288-cfeb-412f-9c75-18748ca11c69