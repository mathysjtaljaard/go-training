package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretsAgent struct {
	person
	ltk bool
}

func main() {

	sa := secretsAgent{
		person: person{
			first: "hello",
			last:  "bye",
		},
		ltk: true,
	}

	person1 := person{
		first: "hello",
		last:  "bye",
	}

	fmt.Println(person1)
	fmt.Println(sa.first, sa.last, sa.ltk)
}
