package main

import (
	"fmt"
	//"go_training/myhellomodule"
)

type sometThing struct {
	name string
	age  int
}

type somethingElse struct {
	baseValues sometThing
	color      string
}

func (value *somethingElse) print() string {
	return fmt.Sprintf(
		"name: %s\nage: %d\ncolor: %s", value.baseValues.name, value.baseValues.age, value.color)
}

func main() {

	// // fmt.Println(myhellomodule.SayHello("mathys"))
	// // msg := myhellomodule.CreateMessage("hello", "goodbye")
	// fmt.Println(msg)
	// msg.PrintMe()
	// var a int   // a variable of type 'int'
	// var b int   // another 'int'
	// var c *int  // a variable of type 'pointer to int'
	// & in front of variable name is used to retrieve the address of where this variable’s value is stored. That address is what the pointer is going to store.
	// * in front of a type name, means that the declared variable will store an address of another variable of that type (not a value of that type).
	// * in front of a variable of pointer type is used to retrieve a value stored at given address. In Go speak this is called dereferencing.
	// n, _ := fmt.Printf("Hello this %s value is applied by the formatter %q \n", "something", "not sure")
	// // Variable ptrInt is assigned the pointer to the address of the newly generated memory allocation of the empty zero type
	// ptrInt := new(int)

	// fmt.Printf("number of characters printed = %d\n", n)
	// controlFlow()
	// //ampVsStar()

	// fmt.Println("pnrInt")
	// fmt.Println(ptrInt)
	// fmt.Println(*ptrInt)
	// *ptrInt = 5
	// fmt.Println(&ptrInt)
	// fmt.Println(*ptrInt)
	// fmt.Println("pnrInt")
	// // Method updateVar takes the address of where the variable n vaue is stored
	// // updateVar( ptrValue *int) -> *int the address which the variable ptrValue will store when assigned
	// // *ptrValue = 5 -> expression -> assign the integer value of 5 to the dereferenced ptrValue
	// updateVar(&n)
	// fmt.Println(n)

	// newSomeElseThing := new(somethingElse)
	// newSomeElseThing.color = "red"
	// newSomething := new(sometThing)
	// newSomething.age = 10
	// newSomething.name = "hello"
	// newSomeElseThing.baseValues = *newSomething
	// newSomeElseThing.baseValues.name = "brad"
	// fmt.Println(*newSomething)

	// fmt.Println(newSomeElseThing)

	// value := sometThing{age: 100, name: "bob"}
	// fmt.Println(value)

	// newValue := somethingElse{baseValues: sometThing{age: 20, name: "dick"}, color: "blue"}
	// fmt.Println(newValue.print())

}

func controlFlow() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			//fmt.Printf("Current value of i = %d \n", i)
		}
	}
}

// Method updateVar takes the address of where the variable n vaue is stored
// updateVar( ptrValue *int) -> *int the address which the variable ptrValue will store when assigned
// *ptrValue = 5 -> expression -> assign the integer value of 5 to the dereferenced ptrValue
func updateVar(addressOfValue *int) {
	fmt.Println(`updateVar start`)
	fmt.Printf("address of value is %#x\n", &addressOfValue)
	fmt.Printf("dereferenced value of value is %d\n", *addressOfValue)
	*addressOfValue = 20
	fmt.Println(`updateVar end`)
}

func ampVsStar() {
	a := 5
	p := &a // copy by reference
	x := a  // copy by value

	fmt.Println("a = ", a)   // a =  5
	fmt.Println("&a = ", &a) // &a =  5
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  5
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5

	fmt.Println("\n Change *p = 3")
	*p = 3
	fmt.Println("a = ", a)   // a =  3
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  3
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5

	fmt.Println("\n Change a = 888")
	a = 888
	fmt.Println("a = ", a)   // a =  888
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  888
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  5

	fmt.Println("\n Change x = 1")
	x = 1
	fmt.Println("a = ", a)   // a =  888
	fmt.Println("p = ", p)   // p =  0x10414020
	fmt.Println("*p = ", *p) // *p =  888
	fmt.Println("&p = ", &p) // &p =  0x1040c128
	fmt.Println("x = ", x)   // x =  1

	//&p = 3 // error: Cannot
}
