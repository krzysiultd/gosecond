package main

import (
	"fmt"
)

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42, 0}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	s := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(s), len(s), s)
	}

	// access to slice with for range
	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing a slice
	fmt.Println(x[0:2])
	fmt.Println(x[3:])

	//append to a slice
	x = append(x, 77, 88, 99)
	fmt.Println(x)

	y := []int{123, 234, 567}
	x = append(x, y...)
	fmt.Println(x)

	//deleting from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//make slice - capacity is known and you don't waste compute power to create new array every time slice changes
	abc := make([]int, 5, 8)
	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(cap(abc))

	abc[4] = 50

	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(cap(abc))

	abc = append(abc, 51)

	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(cap(abc))

	abc = append(abc, 52)
	abc = append(abc, 53)
	abc = append(abc, 54)

	fmt.Println(abc)
	fmt.Println(len(abc))
	fmt.Println(cap(abc)) // the array is deleted and double capacity array is created!

	//multi-dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Johnie Walker"}
	fmt.Println(jb)
	mp := []string{"Miss", "Monepenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
	fmt.Println(xp[1][1])

}
