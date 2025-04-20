package main

import "fmt"

type MessageToSend struct {
	message   string
	sender    User
	recipient User
}

type User struct {
	name   string
	number int
}

func canSendMessage(message MessageToSend) bool {
	if message.recipient.number == 0 {
		return false
	}
	if message.sender.number == 0 {
		return false
	}
	if message.recipient.name == "" {
		return false
	}
	if message.sender.name == "" {
		return false
	}

	return true
}

func sendMessage(message MessageToSend) {
	fmt.Printf("Sending %s from %s : %d to %s : %d\n",
		message.message,
		message.sender.name,
		message.sender.number,
		message.recipient.name,
		message.recipient.number,
	)

	if canSendMessage(message) {
		fmt.Println("Message sent successfully")
	} else {
		fmt.Println("Failed to send message")
	}
	fmt.Println("====================================")

}

func main() {
	message1 := MessageToSend{
		message: "Hello, How are you?",
		sender: User{
			name:   "MD. Tahsin Ferdous",
			number: 123456789,
		},
		recipient: User{
			name:   "MD. Azizul Hakim Aziz",
			number: 987654321,
		},
	}

	message2 := MessageToSend{
		message: "Hello, How are you?",	
		sender: User{
			name:   "MD. Tahsin Ferdous",		
			number: 123456789,
		},
		recipient: User{
			name:   "MD. Azizul Hakim Aziz",
			number: 0,
		},
	}

	sendMessage(message1)
	sendMessage(message2)

}
