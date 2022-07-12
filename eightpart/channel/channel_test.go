package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_readChan(t *testing.T) {
	var mychan = make(chan int, 10)
	writeChan(mychan)
	msg := readChan(mychan)
	msg1 := readChan(mychan)
	fmt.Println(msg)
	fmt.Println(msg1)
}

func Test_addNumberToChan(t *testing.T) {
	chan1 := make(chan int, 10)
	chan2 := make(chan int, 10)

	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	for {
		select {
		case e := <-chan1:
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <-chan2:
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func Test_chanRange(t *testing.T) {
	chan1 := make(chan int, 10)

	//go addNumberToChan(chan1)

	go func() {
		chan1 <- 1
		time.Sleep(10 * time.Second)
	}()
	chanRange(chan1)
}

func TestProcess(t *testing.T) {
	channels := make([]chan int, 10) // 创建 10 个切片

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go Process(channels[i]) // 启动协程
	}

	for i, ch := range channels {
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}

	fmt.Println("go routine done!")
}
