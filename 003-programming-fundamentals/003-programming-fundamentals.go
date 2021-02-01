package main

import (
	"fmt"
	"runtime"
)

const (
	a = iota
	b = iota
	c
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	s := "Hello, moto"

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
