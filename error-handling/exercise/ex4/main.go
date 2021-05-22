package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {

	f, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(f)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{"123", "123", errors.New("f cannot be less than zero")}
	}
	return 42, nil
}
