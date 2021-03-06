package design

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var CA chan int

func main() {
	wg = sync.WaitGroup{}
	CA = make(chan int)

	wg.Add(2)

	go A()
	go B()
	wg.Wait()
}

func A() {

	for i := 1; i <= 100; i++ {
		CA <- 1
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	wg.Done()

}

func B() {

	for i := 1; i <= 100; i++ {
		<-CA
		if i%2 == 1 {
			fmt.Println(i)
		}

	}
	wg.Done()

}
