package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(even, odd, quit)

	receive(even, odd, quit)

}

func send(even, odd, quit chan<- int) {
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

	quit <- 0
	close(quit)
}

func receive(even, odd, quit <-chan int) {

	for {
		select {
		case e, ok := <-even:

			fmt.Println("even", e, "is ok", ok)

		case o, ok := <-odd:
			if ok {
				fmt.Println("odd", o, "is ok", ok)
			}
		case q, ok := <-quit:
			fmt.Println("quitting", q, "is ok", ok)
			return
		}
	}
}
