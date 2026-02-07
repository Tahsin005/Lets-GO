package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Staurday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("The day after tomorrow.")
	default:
		fmt.Println("Too far.")
	}

	// similar to if-else if-else
	switch {
	case today == time.Saturday:
		fmt.Println("Today.")
	case today == time.Sunday:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far.")
	}
}