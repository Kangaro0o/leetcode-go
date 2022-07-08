package channel

import (
	"fmt"
	"time"
)

func readChan(chanName <-chan int) int {
	msg := <-chanName
	return msg
}

func writeChan(chanName chan<- int) {
	chanName <- 1
}

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func chanRange(chanName chan int) {
	for e := range chanName {
		fmt.Printf("Get element from chan: %d\n", e)
	}
}
