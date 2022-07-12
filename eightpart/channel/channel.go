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

// Process 使用 channel 控制子协程
func Process(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1 //管道中写入一个元素表示当前协程已结束
}
