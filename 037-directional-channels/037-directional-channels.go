package main

import "fmt"

func main() {
	c := make(chan int)
	cs := make(chan<- int) //send
	cr := make(<-chan int) //receive

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	//general to specific assigns
	cr = c
	cs = c
}
