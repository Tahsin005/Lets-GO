package main

import "fmt"

func getMessageWithRetries() [3]string {
	messages := [3]string {
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}

	return messages
}

func testSend(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i >= len(messages) - 1 {
			fmt.Println("complete failure")
		}
	}
	fmt.Println("=========================")
}

func main() {
	testSend("Bob", 0)
	testSend("Alice", 1)
	testSend("Mangalam", 2)
	testSend("Ozgur", 3)
}
