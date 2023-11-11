package main

import (
	"fmt"
)

//penggunaan type assertion, untuk menemtukan interface expense diimplementasi oleh email type atau sms type
func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok{
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok{
		return s.toPhoneNumber , s.cost()
	}
	return "" , 0.0

	
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}


func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello theres",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})



	avg := estimateYearlyCost(sms{
		isSubscribed:  false,
		body:          "test ini diluar konteks janga dipedulikan!",
		toPhoneNumber: "+155555504444",
	} , 5)
	fmt.Println(avg)
}


//full course tentanng type assertion interface pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/7804ef9a-1a03-42a0-94d1-42ca09b89e43