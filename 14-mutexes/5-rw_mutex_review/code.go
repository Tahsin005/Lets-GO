package main

import (
	"fmt"
	"sync"
)

var (
	data    = 0
	rwMutex = sync.RWMutex{}
	wg      = sync.WaitGroup{}
)

func reader(id int) {
	defer wg.Done()
	rwMutex.RLock() // Acquire read lock
	fmt.Printf("Reader %d: read data = %d\n", id, data)
	rwMutex.RUnlock() // Release read lock
}

func writer(id int, value int) {
	defer wg.Done()
	rwMutex.Lock() // Acquire write lock
	fmt.Printf("Writer %d: writing data = %d\n", id, value)
	data = value
	rwMutex.Unlock() // Release write lock
}

func main() {
	wg.Add(5)

	go reader(1)
	go writer(1, 10)
	go reader(2)
	go writer(2, 20)
	go reader(3)

	wg.Wait()
}
