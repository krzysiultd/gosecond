package main

import "fmt"

func main() {
	x := bar()
	fmt.Println(x())

	y := foo()
	y()
}

func bar() func() int {
	return func() int {
		return 44
	}
}

func foo() func() {
	return func() {
		fmt.Println("the number is 42")
	}
}
