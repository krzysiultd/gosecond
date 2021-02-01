package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi

}

func (s square) area() float64 {
	return s.length * s.length

}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("The area is", s.area())
}

func main() {
	c := circle{
		radius: 5.2,
	}
	s := square{
		length: 22.1,
	}

	info(c)
	info(s)

}
