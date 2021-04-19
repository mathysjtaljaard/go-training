package main

import "fmt"

var x2 int
var y2 string
var z2 bool

var x3 = 42
var y3 = "James Bond"
var z3 = true

type customType int

var x4 customType

var y5 int

func main() {
	exercise1(1)
	exercise2(2)
	exercise3(3)
	exercise4(4)
	exercise5(5)
	exercise6(6)
}

func printExerciseStatement(number int) string {
	return fmt.Sprintf("\nExercise %d", number)
}

func exercise1(number int) {
	println(printExerciseStatement(number))
	x := 42
	y := "James Bond"
	z := true
	println(x, y, z)
	println(z)
	println(y)
	println(z)
}

func exercise2(number int) {
	println(printExerciseStatement(number))
	println(x2, y2, z2)
	println(`answer of compiler assigned values to variables x2,y2,z2 is Zero Type`)

}

func exercise3(number int) {
	println(printExerciseStatement(number))
	s := fmt.Sprintf("x3 = %d\ny3 = %s\nz3 = %t", x3, y3, z3)
	println(s)
}

func exercise4(number int) {
	println(printExerciseStatement(number))

	fmt.Printf("Type of x4 is %T\n", x4)
	fmt.Printf("Value of x4 is %v\n", x4)
	x4 = 42
	fmt.Printf("Type of x4 is %T\n", x4)
	fmt.Printf("Value of x4 is %v\n", x4)
}

func exercise5(number int) {
	println(printExerciseStatement(number))
	x4 = 39
	fmt.Printf("Type of x4 is %T\n", x4)
	fmt.Printf("Value of x4 is %v\n", x4)

	//NOTE: shorthand declaration operator redeclares a new variable
	// if you don't want to redaclare the variable, use assignment
	// y5 := int(x4) redeclares y5 to be an int of value 4
	y5 = int(x4)
	fmt.Printf("Type of y5 is %T\n", y5)
	fmt.Printf("Value of y5 is %v\n", y5)
}

func exercise6(number int) {
	println(printExerciseStatement(number))
	//println(x6)
}
