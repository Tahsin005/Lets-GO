package main

import "fmt"

type Vertexx struct { 
	Lat, Long float64
}

var m map[string]Vertexx

var mm = map[string]Vertexx{
	"Bell Labs": {40.6843, -74.3997},
	"Google":    {37.4220, -122.0841},
}

func main() {
	m = make(map[string]Vertexx)
	m["Bell Labs"] = Vertexx{
		40.6843, -74.3997,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println("----------------")

	fmt.Print("Manipulate values in map")
	intmap := make(map[string]int)

	intmap["tahsin"] = 1
	intmap["ferdous"] = 2
	intmap["tahsin"] = 3

	fmt.Println(intmap["tahsin"])

	delete(intmap, "tahsin")
	fmt.Println(intmap)

	value, ok := intmap["tahsin"]
	fmt.Println(value, ok)

	fmt.Println("----------------")

	intmap["a"] = 1
	intmap["b"] = 2
	intmap["c"] = 3

	for k, v := range intmap {
		fmt.Println(k, v)
	}
}