package main

import (
	"fmt"
	"time"
)

// Interface
type message interface {
	getMessage() string
}

type BirthDayMessage struct {
	birthdayTime time.Time
	recipientName string
}

type SendingReport struct {
	reportName string
	reportTime time.Time
}

func (bm BirthDayMessage) getMessage() string {
	return fmt.Sprintf("Happy Birthday %s! Your birthday is on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

func (sr SendingReport) getMessage() string {
	return fmt.Sprintf("Your report named (%s) have been sent at %s", sr.reportName, sr.reportTime.Format(time.RFC3339))
}

func SendMessage(m message) {
	fmt.Println(m.getMessage())
	fmt.Println("====================================")
}

func test(m message) {
	SendMessage(m)
	fmt.Println("====================================")
}

func main () {
	fmt.Println("Interfaces")

	birthdayMessage := BirthDayMessage{
		birthdayTime: time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
		recipientName: "Tahsin Ferdous",
	}

	sendingReport := SendingReport{
		reportName: "First Report",
		reportTime: time.Date(2023, 10, 01, 0, 0, 0, 0, time.UTC),
	}

	test(birthdayMessage)
	test(sendingReport)
}