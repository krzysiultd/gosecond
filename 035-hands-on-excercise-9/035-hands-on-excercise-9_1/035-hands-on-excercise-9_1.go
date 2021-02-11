package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("hi")
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	bar()
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	go grr()
	wg.Wait()
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("this is foo i: ", i)
	}
	wg.Done()
}

func grr() {
	for i := 0; i <= 10; i++ {
		fmt.Println("this is grr i: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("this is bar i: ", i)
	}
}
