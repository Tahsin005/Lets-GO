package main

import "fmt"

// Concatter is a closure that concatenates strings.
func concatter () func(string) string {
	doc := ""

	return func(s string) string {
		doc += s + " "
		return doc
	}
}

func main () {
	stringAggregator := concatter()

	stringAggregator("Hello")
	stringAggregator("World")
	stringAggregator("!")
	stringAggregator("How")
	stringAggregator("are")
	stringAggregator("you")
	stringAggregator("doing")
	stringAggregator("today")
	stringAggregator(".")
	stringAggregator("I")
	stringAggregator("hope")
	stringAggregator("you")
	stringAggregator("are")
	stringAggregator("doing")
	stringAggregator("well")
	stringAggregator("!")

	fmt.Println(stringAggregator("Goodbye"))
}