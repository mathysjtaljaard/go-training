package main

import (
	"fmt"
)

func main() {
	const closeit = 9
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go gen10(c, i, closeit)
	}

	for v := range c {
		fmt.Println(v)
	}
}

func gen10(c chan int, j int, closeIt int) {
	fmt.Println(j, closeIt)
	go func(closeIt int) {
		for i := 0; i < 10; i++ {
			c <- i + j*2
		}
		if j == closeIt {
			close(c)
		}
	}(closeIt)
}
