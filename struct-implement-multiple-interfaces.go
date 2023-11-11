package main

import (
	"fmt"
)


//struct email mengimplementasi interface expense dan printer, karena memiliki method-struct cost() dan print()
func (e email) cost() float64 {
	if e.isSubscribed == true {
		return float64(len(e.body)) * 0.01
	}else{
		return float64(len(e.body)) * 0.05
	}
}

func (e email) print() {
	fmt.Println(e.body)
	
}

// don't touch below this line

type expense interface {
	cost() float64
}


type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func test(e expense, p printer) { //menerima paramter dari instance expense dan printer interface
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(e, e) //cukup mengirim struct email saja, struct email sudah memenuhi syarat karena memiliki method-struct cost() dan print()
	e = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(e, e)
}

//full course tentanng multiple interfaces pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/59b4509c-0cb1-487f-8b3f-68357629e60e