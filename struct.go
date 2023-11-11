package main

import "fmt"

type messageToSend struct {
	message string
	phoneNumber int
	
}

// don't edit below this line

func test(msg messageToSend) {
	
	fmt.Printf("Sending message: '%s' to: %v\n", msg.message, msg.phoneNumber)
	
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}

//untuk melihat full penjelasan ke link ini : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/81362a78-09b3-459e-bd16-c0312d3ef2d2