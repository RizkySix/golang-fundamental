package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface { 
	getMessage() string //seluruh func/struct-method yang menggunakan method-method pada message interface, pada kasus ini adalah method getMessage(), dapat dikatakan struct-method tsb "mengimplementasi" message interface
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) { //keunggulan menggunakan interface, satu func test dapat menerima instansiasi struct secara dinamis, dalam kasus ini adalah struct birthdayMessage dan sendingReport
	sendMessage(m)
	
	fmt.Println("====================================")
}

func main() {
	test(sendingReport{ //func test dapat menerima instansiasi sendingReport dan birthdayMessage
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{//func test dapat menerima instansiasi sendingReport dan birthdayMessage
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}

//full course tentanng interfaces pada link berikut : https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/efcafed8-c0ba-4efb-a6e2-88db3c8a0733
