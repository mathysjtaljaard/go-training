package main

import "fmt"

func main() {
	num1()
	num2()
}

func num1() {
	a := 43
	fmt.Println("address of a is", &a)
	println()
}

type person struct {
	first string
	last  string
}

func (p *person) changeMe(f string, l string) {
	if len(f) != 0 {
		p.first = f
	}

	if len(l) != 0 {
		p.last = l
	}
}

func changeMe(p *person) {
	p.changeMe("bubba", "")
}

func num2() {

	p := person{
		first: "james",
		last:  "borg",
	}
	fmt.Println("Person", p)
	p.changeMe("", "borgetti")
	fmt.Println("Person", p)
	p.changeMe("Jameson", "")
	fmt.Println("Person", p)
	changeMe(&p)
	fmt.Println("Person", p)
	println()
}
