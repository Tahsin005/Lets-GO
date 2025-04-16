package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	email := "test.user.001@example.com"
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	length := len(email)
	if length > 1 {
		fmt.Println("Email")
	}

	if LENGTH := len(email); LENGTH > 1 {
		fmt.Printf("This variable %v is only available in this scope", LENGTH)
	}
}
