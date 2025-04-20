package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

type Person struct {
	name	string
	age		int
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %d\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	m1 := messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	}
	test(m1)

	m2 := messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	}
	test(m2)

	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})

	fmt.Println("====================================")

	p1 := Person{
		name: "MD. Tahsin Ferdous",
		age:  23,
	}

	p2 := Person{
		name: "Pau Cubersi",
		age:  17,
	}

	fmt.Println(p1, p2)
}
