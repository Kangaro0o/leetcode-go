package design

import (
	"sync"
	"testing"
)

func TestChannel(t *testing.T) {
	wg = sync.WaitGroup{}
	CA = make(chan int)

	wg.Add(2)

	go A()
	go B()
	wg.Wait()

}
