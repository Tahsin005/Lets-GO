package main

import "fmt"

func main() {
	const name = "MD. Tahsin Ferdous"
	const openRate = 30.5342

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent", name, openRate)


	fmt.Println(msg)
}
