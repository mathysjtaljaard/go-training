package main

import (
	"fmt"
)

type personError struct {
	First string
	Info  string
}

func (pe personError) Error() string {
	return fmt.Sprintf("error when trying to create person %s", pe.First)
}

func main() {

	pe := personError{
		First: "bando",
		Info:  "bad and not do or",
	}

	foo(pe)
}

func foo(e error) {
	fmt.Println("Error given which was", e, e.(personError).Info)
}
