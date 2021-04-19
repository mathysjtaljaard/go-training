package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	problem1()
	problem2()
	problem3() //-- note if running this and want to program to finish
	// don't use -race in command
	problem4()
	problem5()
	problem6()
}

func problem1() {

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("Function 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Function 2")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("waiting for grandma")
}

type person struct {
	first, last string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Printf("My full name is %s %s\n", p.first, p.last)
}

func saySomething(h human) {
	h.speak()
}

func problem2() {
	al := person{
		first: "Ali",
		last:  "Baba",
	}

	saySomething(&al)
	//saySomething(al)
	//cannot use al (variable of type person) as human value in argument to saySomething: missing method speak (speak has pointer receiver)compilerInvalidIfaceAssign
}

func problem3() {
	counter := 0

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func problem4() {

	counter := 0

	const gr = 100
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			mux.Lock()
			v := counter
			v++
			counter = v
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func problem5() {

	var counter int64

	const gr = 100
	var wg sync.WaitGroup

	wg.Add(gr)
	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func problem6() {

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}

// func problem7() {
// 	//NA
// }
