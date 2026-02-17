package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func readFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return  nil
	}

	defer file.Close()

	data, err := io.ReadAll(file) 
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func processData(data []int) {
	start := time.Now()
	defer func ()  {
		fmt.Println(time.Since(start))	
	}()

	for _, d := range data {
		fmt.Println(d)
		time.Sleep(time.Millisecond * 100)
	}

}

func safeOperation() {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}	
	}()

	panic("Something went wrong!")

	fmt.Println("Cannot reach this code")
}

func main() {
	// closing opened file
	if err := readFile("index.html"); err != nil {
		fmt.Println(err)
	}

	// for performance benchmark
	data := []int{1, 2, 3, 4, 5}
	processData(data)

	// recover from panic
	safeOperation()
}

