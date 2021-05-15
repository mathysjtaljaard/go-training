package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- i
			if i == 99 {
				close(c)
			}
		}(i)

	}
	go func() {
		q <- 1
		close(q)
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case q, ok := <-q:
			if ok {
				fmt.Println("Value of Q", q)
			} else {
				return
			}
		case c, ok := <-c:
			if ok {
				fmt.Println("Value of c", c)
			} else {
				return
			}
		}
	}
}
