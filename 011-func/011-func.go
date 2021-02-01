package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := sum(xi...)
	fmt.Println("Total is", x)

}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item", i, "we are adding", v, "and the sum is", sum)
	}
	return sum
}
