package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Printf("%T", i)
}

func foo() string {
	s := "Hello"
	return s
}

func bar() func() int {
	return func() int {
		return 451
	}
}
