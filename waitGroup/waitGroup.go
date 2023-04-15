package waitGroup

import (
	"fmt"
	"sync"
)

var Msg string
var Wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func UpdateMessage(s string, wg *sync.WaitGroup) {
	// decrement wait group num by one
	defer wg.Done()
	Msg = s
}

func PrintMessage() {
	fmt.Println(Msg)
}
