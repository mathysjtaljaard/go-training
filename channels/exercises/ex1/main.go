package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	d := make(chan int, 1)

	d <- 43

	fmt.Println(<-d)

	e := make(chan int, 100)

	for i := 0; i < 100; i++ {
		e <- i
	}
	close(e)

	for k := range e {
		fmt.Println(k)
	}
}
