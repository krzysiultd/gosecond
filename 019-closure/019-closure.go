package main

import "fmt"

var x int

func main() {
	var x int
	fmt.Println(x)
	x++

	// code block in code block; 'y' is not accesible from the outside of {}
	{
		y := 42
		fmt.Println(y)
	}

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
