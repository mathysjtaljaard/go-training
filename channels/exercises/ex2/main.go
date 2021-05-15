package main

import "fmt"

func main() {
	cd := make(chan int)

	go func() {
		cd <- 42
	}()

	fmt.Println(<-cd)

	fmt.Printf("cd\t%T\n", cd)

	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("cs\t%T\n", cs)
}
