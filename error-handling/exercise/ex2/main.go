package main

import (
	"encoding/json"
	"errors"
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

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	err = errors.New("Error")
	if err != nil {

		return nil, fmt.Errorf("unable to marshal given json, error is %w", err)
	}

	return bs, nil
}
