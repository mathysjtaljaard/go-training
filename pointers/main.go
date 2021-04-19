package main

import "fmt"

func main() {
	a := int(42)
	fmt.Println(a)
	// address of a which is a type of *int
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)
	b := &a
	// dereferce the address of a by using the * operatore
	fmt.Println(b)
	fmt.Println(*b)

	//Pass by reference given address dereference type and operate on
	*b++
	fmt.Println("dereference b", *b)
	fmt.Println("value of a", a)

	p := newPerson("human", "being")
	fmt.Println(p)
	p.sayName()
	p.updateName("borg")
	p.sayName()
	//when to use pointers

	borg := p.assimulate()
	fmt.Printf("%T\n", borg)
	borg.isCollective()

	p2 := newPerson("welcome", "brad")
	fmt.Println(p2)
	p2.sayName()
	p2.updateName("Billy")
	p2.sayName()

	borg2 := p.assimulate()
	fmt.Printf("%T\n", borg2)
	borg2.isCollective()

	p.sayName()
	borg.isCollective()
}

type person struct {
	first string
	last  string
}

type borg struct {
	person        *person
	isAssimulated bool
}

func newPerson(f string, l string) *person {
	return &person{
		first: f,
		last:  l,
	}
}

func (p person) sayName() {
	fmt.Println("Hello, me is", p.first, p.last)
}

func (b borg) isCollective() {
	fmt.Printf("The collectives name is %v %v. we are assimulated %v\n", b.person.first, b.person.last, b.isAssimulated)
}

// Want to update the reference of the person instance
func (p *person) updateName(f string) {
	p.first = f
}

func (p person) assimulate() *borg {
	return &borg{
		person:        newPerson(p.first, p.last),
		isAssimulated: true,
	}
}
