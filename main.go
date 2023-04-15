package main

import (
	"app/deadLock"
	"app/waitGroup"
	"fmt"
	"sync"
)

func main() {
	waitGroup.Msg = "test1"
	var bankBalance int = 0
	var m sync.Mutex

	var messages = []string{
		"test1",
		"test2",
		"test3",
	}

	// printing always in message varaible order
	for _, e := range messages {
		// put in only one thing in wait group
		waitGroup.Wg.Add(1)

		go waitGroup.UpdateMessage(e, &waitGroup.Wg)

		waitGroup.Wg.Wait()
		waitGroup.PrintMessage()
	}

	incomes := []waitGroup.Income{
		{Source: "Job", Amount: 1000},
		{Source: "Gifts", Amount: 5},
		{Source: "Walking dogs", Amount: 50},
	}

	waitGroup.Wg.Add(len(incomes))

	for _, income := range incomes {

		go func(income waitGroup.Income) {
			defer waitGroup.Wg.Done()
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

	waitGroup.Wg.Wait()

	//deadlock

	collection := deadLock.NewCollection()
	collection.Add("test", "test")
}
