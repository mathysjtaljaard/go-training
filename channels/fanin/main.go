package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	// send
	go send(even, odd)

	go receive(even, odd, fanin)

	for fin := range fanin {
		fmt.Println("FanIn", fin)
	}

	fmt.Println("Later gang")
}

func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
	fmt.Println("Closing Even")
	close(even)
	fmt.Println("Closing Odd")
	close(odd)

}

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v + 5
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v + 1
		}
		wg.Done()
	}()
	wg.Wait()

	close(fanin)
}
