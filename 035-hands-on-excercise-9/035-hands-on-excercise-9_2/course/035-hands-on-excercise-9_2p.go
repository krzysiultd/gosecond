package main

import "fmt"

type person struct {
	first string
	last  string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return fmt.Sprint("My name is ", p.first, " ", p.last)
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{"James", "Bond"}
	saySomething(&p)
}
