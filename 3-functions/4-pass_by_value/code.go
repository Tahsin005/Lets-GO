package main

import "fmt"

// values are usually passed by value in GO

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	// sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	// fmt.Println("you've sent", sendsSoFar, "messages")

	fmt.Println("Before", sendsSoFar)
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("After", sendsSoFar)
}

// func incrementSends(sendsSoFar, sendsToAdd int) int {
// 	return sendsSoFar + sendsToAdd
// }

func incrementSends(sendsSoFar, sendsToAdd int) {
	sendsSoFar += sendsToAdd
	fmt.Println("After changing it in the function",sendsSoFar)
}
