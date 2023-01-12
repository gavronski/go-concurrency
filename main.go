package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	// decrement wait group num by one
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	msg = "test1"

	var messages = []string{
		"test1",
		"test2",
		"test3",
	}

	// printing always in message varaible order
	for _, e := range messages {
		// put in only one thin in wait group
		wg.Add(1)

		go updateMessage(e, &wg)

		wg.Wait()
		printMessage()

	}
}
