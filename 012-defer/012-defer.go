package main

import "fmt"

func main() {
	defer foo() // it will be executed just before exiting main() func
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
