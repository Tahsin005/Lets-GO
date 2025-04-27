package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
			case <-snapshotTicker:
				takeSnapshot()
			case <-saveAfter:
				saveSnapshot()
				return
			default:
				waitForData()
				time.Sleep(500 * time.Millisecond)
		}
	}
}

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond) // returns a channel that sends a value every 800ms
	saveAfter := time.After(2800 * time.Millisecond) // returns a channel that sends a value after 2800ms
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func main() {
	test()
}
