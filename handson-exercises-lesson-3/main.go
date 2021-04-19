package main

import "fmt"

func main() {
	//execiseOne()
	//exerciseTwo()
	//exerciseThree()
	//exerciseFour()
	//exerciseFive()
	//exerciseSix()
	//exerciseSeven()
	//exerciseEight()
	//exerciseNine()
	exerciseTen()
}

// 1
func execiseOne() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

// 2
func exerciseTwo() {
	startDecimal := 65
	endDecimal := 90

	for j := startDecimal; j <= endDecimal; j++ {
		fmt.Println(j - 64)
		for i := 1; i <= 3; i++ {
			fmt.Printf("\t%#U\n", j)
		}
	}
}

// 3
func exerciseThree() {
	birthYear := 1982
	count := 0
	for birthYear <= 2021 {
		fmt.Println(birthYear)
		birthYear++
		count++
	}

	println(count)
}

// 4
func exerciseFour() {
	birthYear := 1982
	for {
		fmt.Println(birthYear)
		birthYear++
		if birthYear > 2021 {
			break
		}
	}
}

// 5
func exerciseFive() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v mod 4 = %v\n", i, (i % 4))
	}
}

// 6
func exerciseSix() {
	isIt := true

	if isIt {
		println("it is ")
	}
}

// 7
func exerciseSeven() {

	count := 0
	for {
		if count == 0 {
			println("count is 0")
		} else if count == 1 {
			println("count is 1")
		} else {
			println("count is 3 exiting")
			break
		}
		count++
	}
}

// 8
func exerciseEight() {

	switch {
	default:
		{
			println("switch with no expression")
		}
	}
}

// 9
func exerciseNine() {
	favSport := "Tennis"
	switch favSport {
	case "Tennis":
		{
			println("good tennis")
		}
	}
}

// 10
func exerciseTen() {
	//true
	// false
	//true
	// true
	// false
}
