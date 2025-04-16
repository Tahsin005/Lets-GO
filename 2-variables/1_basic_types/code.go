package main

import "fmt"

func main () {
	fmt.Println("Variables")

	var smsSendingLimit int
	var costPerSMS float64 = 3.1416
	var hasPermissin bool
	var username string
	var char byte = 'A'
	var unicodeChar rune = '♥'
	congrats := "happy birthday"

	fmt.Printf(
		"\n Integer: %v\n Float: %f\n Boolean: %v\n String: %q\n Byte (in integer or ASCII): %v\n Byte (in character type): %c\n Rune(in interger or unicode): %U\n Rune(in character): %c\n",
		smsSendingLimit,
		costPerSMS,
		hasPermissin,
		username,
		char,
		char,
		unicodeChar,
		unicodeChar,
	)
	fmt.Println(congrats)
}