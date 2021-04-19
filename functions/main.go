package main

import "fmt"

type human interface {
	speak() string
}
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() string {
	return fmt.Sprint("My name is ", sa.first, "\n")
}

func (p person) speak() string {
	return fmt.Sprint("My name is ", p.first, "\n")
}

func poly(h human) {
	fmt.Printf("%T human passed %v\n", h, h)
}

func plainMultiplier(x int) {
	fmt.Println("plain multiplier called with value ", x)
}

func main() {
	fmt.Print(`hello
	
	`)
	test()
	multiplyBy10 := multiplier(10)

	fmt.Println(multiplyBy10(10))
	sliceInt := []int{1, 23, 4, 5, 6, 7, 8}
	fmt.Println(counter(sliceInt...))

	deferfunc("")
	foo("")

	sa1 := secretAgent{
		person: person{
			first: "Hello",
			last:  "Barney",
		},
		ltk: true,
	}

	p1 := person{
		first: "last",
		last:  "first",
	}

	fmt.Println(sa1)
	fmt.Println(sa1.speak())
	fmt.Printf("%T\n", sa1)
	trainedSA := trainSecretAgent("bob", "marley", true)

	fmt.Println(trainedSA)
	fmt.Println(trainedSA.first)
	fmt.Printf("%T\n", trainedSA)
	fmt.Println(trainedSA.speak())

	fmt.Println(p1)
	poly(sa1)
	poly(p1)

	printListFormatter := specFormatter(" Give me the values %v")
	sArr := []string{"a", "b"}

	fmt.Println(printListFormatter(sArr...))

	runCallback(100, plainMultiplier)

	fmt.Println(fib(3))
}

func test() {
	fmt.Println("hello")
}

func multiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func counter(x ...int) int {
	///
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func deferfunc(s string) {
	fmt.Println("hello up")
	defer fmt.Println("hello up 1")
	fmt.Println("hello up 2")
}

func foo(s string) {
	fmt.Println("bango bingo")
}

/// methods are the synonim for using the receiver
// func (r receiver) identifier(params) (return(s)) {code ...}
func trainSecretAgent(first string, last string, ltk bool) secretAgent {
	return secretAgent{
		ltk: ltk,
		person: person{
			first: first,
			last:  last,
		},
	}
}

// interface allows us to define a behavior

// Function expression
func specFormatter(s string) func(...string) string {
	return func(s ...string) string {
		return fmt.Sprintf(" %v ", s)
	}
}

// callback
func runCallback(value int, f func(int)) {
	count := 10 * value
	f(count)
}

// 1 1 2 3 5 8 13 21 34 (sum(n-1) + sum(n - 2)) N > 1

// func  printfibseq(numberOfValues int) {

// }

// func sumOfFibNumbers(numbersToSum int) int {

// 	sum := 2

// 	return 0
// }
var numList []uint64

func fib(numberOfValues int) uint64 {
	fmt.Println("numbers to calculate ", numberOfValues)

	sum := sumFib(0, numberOfValues, 0)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(numList)
	return sum
}

func sumFib(pos int, end int, sum uint64) uint64 {
	if pos == 0 || pos == 1 {
		numList = append(numList, 1)
		if pos == end-1 {
			return sum + 1
		}
		return sum + sumFib(pos+1, end, 1)
	}

	curSum := (numList[pos-1] + (numList[pos-2]))
	numList = append(numList, curSum)

	if pos == (end - 1) {
		return sum + curSum
	}

	return sum + sumFib(pos+1, end, curSum)
}
