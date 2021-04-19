package main

import "fmt"

func main() {
	value := 42
	execiseOne(value)
	exerciseTwo()
	exerciseThree()
	exerciseFour()
	exerciseFive()
	exerciseSix()
	exerciseSeven()
}

// 1
func execiseOne(value int) {
	printHex(value)
	printBoolean(value)
	printDecimal(value)
	fmt.Println()
}

func printHex(value int) {
	fmt.Printf("Hex value is %#X\n", value)
}

func printBoolean(value int) {
	fmt.Printf("Boolean value is %#b\n", value)
}

func printDecimal(value int) {
	fmt.Printf("Decimal value is %#d\n", value)
}

// 2
func exerciseTwo() {
	//operators
	x := 2
	y := 2

	print("==", (x == y), x, y)
	print("<=", (x <= y), x, y)
	print(">=", (x >= y), x, y)
	print(">", (x > y), x, y)
	print("<", (x < y), x, y)

}

func print(operator string, result bool, x int, y int) {
	fmt.Printf("%v %v %v = %v", x, operator, y, result)
	fmt.Println()
}

// 3
func exerciseThree() {
	const typeConst int = 4
	const untypedConst = "hello"

	fmt.Printf("typed constant value %v and type of %T\n", typeConst, typeConst)
	fmt.Printf("unTyped constant value %v and type of %T\n", untypedConst, untypedConst)
	fmt.Println()
}

// 4
func exerciseFour() {
	notShifted := 42
	fmt.Println("Not Shifted value")
	execiseOne(notShifted)
	fmt.Println("Shifted value")
	shifted := notShifted << 1
	execiseOne(shifted)
}

// 5
func exerciseFive() {
	rawString := `a raw string literal is the back tick
	 value and allows one to write out 
	 information as 	with tabs and newline \n 
	 characters inbedded into the value`
	//"Raw String Literal" //WRONG
	// A raw string literal is the back tick
	fmt.Printf("value %v and type %T\n", rawString, rawString)
}

// 6
func exerciseSix() {
	// This does not work... I believe it's associated with grouping
	// const year2018 = (2018 + iota)
	// const year2019 = (2018 + iota)
	// const year2020 = (2018 + iota)
	// const year2021 = (2018 + iota)
	// -> result is every output is 2018

	// Correct Grouping plays a role in how IOTA is used
	const (
		year2018 = (2018 + iota)
		year2019 = (2018 + iota)
		year2020 = (2018 + iota)
		year2021 = (2018 + iota)
	)

	fmt.Print(year2018, year2019, year2020, year2021)
	fmt.Println()
}

// 7
func exerciseSeven() {
	// take the quiz
	var a = 6
	fmt.Print(a)
	b := 6
	fmt.Print(b)
}
