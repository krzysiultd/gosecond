package main

import (
	"fmt"
)

var a int
var b string
var c bool

type dupa int

var pupa dupa = 41
var y int

func main() {
	a = 42
	b = "James Bond"
	c = true

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

	fmt.Println(pupa)
	fmt.Printf("%T\n", pupa)

	pupa = 43
	fmt.Println(pupa)

	y = int(pupa)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
