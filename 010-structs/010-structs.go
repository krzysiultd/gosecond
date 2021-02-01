package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   49,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	sa1 := secretAgent{
		person: person{
			first: "Dr",
			last:  "No",
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(sa1)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
	fmt.Println(sa1.person.first, sa1.last) // need to remember you can have 'first' for both structs so be careful!
}
