package main

import "fmt"

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	for i := 0; i < 100; i++ {
		go func(i int) {
			c <- i
			if i == 99 {
				close(c)
			}
		}(i)
	}

	return c
}

func receive(c <-chan int) {

	count := 0
	for v := range c {
		fmt.Println(v)
		count++
	}

	fmt.Println("Recieved", count)
}
