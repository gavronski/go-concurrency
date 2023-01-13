package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func updateMessage(s string, wg *sync.WaitGroup) {
	// decrement wait group num by one
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "test1"
	var bankBalance int = 0
	var m sync.Mutex
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

	incomes := []Income{
		{Source: "Job", Amount: 1000},
		{Source: "Gifts", Amount: 5},
		{Source: "Walking dogs", Amount: 50},
	}

	wg.Add(len(incomes))

	for _, income := range incomes {

		go func(income Income) {
			defer wg.Done()
			for i := 1; i <= 52; i++ {
				m.Lock()
				tmp := bankBalance
				tmp += income.Amount
				bankBalance = tmp
				fmt.Printf("Your %d week income %s make bank balance %d \n", i, income.Source, bankBalance)
				m.Unlock()
			}
		}(income)
	}

	wg.Wait()
}
