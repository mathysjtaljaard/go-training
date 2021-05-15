package main

import "fmt"

func main() {
	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("about to exit")
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		fmt.Println("sending ", i)
		c <- i
	}
	close(c)
}

func bar(c <-chan int) {
	for v := range c {
		fmt.Println("pulling ", v)
	}
}
