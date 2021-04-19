package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue

		} else {
			fmt.Printf("%v is uneven \n", i)
		}
	}

	for i := 0; ; i++ {
		if i == 100 {
			println("i == 100")
			break
		}
	}

	value := "this value"
	switch value {
	case "this":
		fmt.Println("this")
	default:
		fmt.Println("default")
	}

	switch value {
	case "1", "two":

	}

}
