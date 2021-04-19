package main

import (
	"fmt"
	"math"
)

func main() {
	num1()
	num2()
	num3()
	num4()
	num5()
	num6()
	num7()
	num8()
	num9()
	num10()
	num11()
}

func num1() {
	fmt.Println("Foo says ", foo())
	x, y := bar()
	fmt.Println("Bar says ", x, y)
	fmt.Println()
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "hello"
}

func num2() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo2(list...))
	fmt.Println(bar2(list[:5]))
	fmt.Println()
}

func foo2(list ...int) int {
	return sum(list)
}

func bar2(list []int) int {
	return sum(list)
}

func sum(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}

	return sum
}

func num3() {
	defer fmt.Println()
	for i := 0; i < 6; i++ {

		if i%2 == 0 {
			defer fmt.Printf("This is number %v defered in the loop\n", i)
		}
		fmt.Printf("This is number %v in the sequence of operations\n", i)
	}
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi; My name is %v %v and I'm %v years of age\n", p.first, p.last, p.age)
}
func num4() {
	p1 := person{
		first: "john",
		last:  "bar",
		age:   20,
	}

	p2 := person{
		first: "Barb",
		last:  "Short",
		age:   40,
	}

	p1.speak()
	p2.speak()
	fmt.Println()
}

type square struct {
	length        float64
	areaFormatter string
}

type circle struct {
	radius        float64
	areaFormatter string
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)

}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) formatter() string {
	return fmt.Sprintf(s.areaFormatter, s.length)

}

func (c circle) formatter() string {
	return fmt.Sprintf(c.areaFormatter, c.radius)
}

type shape interface {
	area() float64
	formatter() string
}

func info(s shape) {
	fmt.Printf("%v %v\n", s.formatter(), s.area())
}

func num5() {

	s := square{
		length:        2,
		areaFormatter: "Area of a Square with sides of lenght %v = ",
	}

	c := circle{
		radius:        2,
		areaFormatter: "Area of a Circle with raduis lenght %v = ",
	}

	info(s)
	info(c)
	fmt.Println()
}

func num6() {

	fmt.Println(func(x int) string {
		return fmt.Sprintf("You gave me the number %v", x)
	}(3))
	fmt.Println()
}

func num7() {

	runMe := func(s string) {
		fmt.Println("You supplied string", s)
	}
	runMe("hello")
	fmt.Println()
}

func num8() {

	multiplyBy100 := func(m int) func(v int) int {
		return func(value int) int {
			return value * m
		}
	}(100)

	multiplyBy := func(m int) func(v int) int {
		return func(value int) int {
			return value * m
		}
	}

	multiplyBy10 := multiplyBy(10)
	fmt.Println("multiply 5 by 100 = ", multiplyBy100(5))
	fmt.Println("multiply 4 by 10 = ", multiplyBy10(4))
	fmt.Println()
}

func num9() {
	printValue := func(v int) {
		fmt.Printf("value given of Type %T is %v\n", v, v)
	}

	iterateValues := func(v []int, f func(v int)) {
		for _, v := range v {
			f(v)
		}
	}

	iterateValues([]int{1, 2, 3, 45, 6}, printValue)
	iterateValues([]int{3, 4, 5, 6, 7, 8, 9}, func(v int) {
		fmt.Println("Anonymous Function called to print numbers", v)
	})
	fmt.Println()
}

func num10() {
	x := 5
	func() {
		x := 0
		x++
		fmt.Println("inner x value ", x)
	}()
	fmt.Println("outer x value ", x)
	fmt.Println()
}

func num11() {
	//teach !!!
	fmt.Println()
}
