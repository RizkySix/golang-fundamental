package main

import (
	"fmt"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}
	return true
}

// don't touch below this line

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}

//nested struct pada go full course dapat dilihat disini : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/59d90390-f479-4472-bb16-9af599285229/525b8c5b-ab89-4588-8117-f91e81b778bc