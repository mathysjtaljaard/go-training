package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	// getMeSpecs()
	// wg.Add(2)
	// go foo()
	// go bar()

	// getMeSpecs()
	// runtime.ReadTrace()
	// wg.Wait()
	// raceCond()
	// raceMutex()
	raceAtomic()
}

func foo() {
	noDry("foo")
	wg.Done()
}

func bar() {
	noDry("bar")
	wg.Done()
}

func noDry(functName string) {
	for i := 0; i < 10; i++ {
		fmt.Println(functName, i)
	}
}

func getMeSpecs() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func raceCond() {
	getMeSpecs()

	counter := 0

	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	getMeSpecs()
	fmt.Println(counter)
}

func raceMutex() {

	getMeSpecs()

	counter := 0

	const gs = 100
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			getMeSpecs()
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	getMeSpecs()
	fmt.Println(counter)

}

func raceAtomic() {
	getMeSpecs()

	var counter int64

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter:\t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	getMeSpecs()
	fmt.Println(counter)
}
