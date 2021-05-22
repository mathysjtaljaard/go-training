package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"shaken", "stirred"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
