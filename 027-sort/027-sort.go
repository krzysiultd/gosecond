package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 5, 1, 12, 77, 34, 54, 993}
	xs := []string{"James Bond", "Miss Moneypenny", "Dr No"}

	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(sort.StringsAreSorted(xs))
	sort.Strings(xs)
	fmt.Println(xs)
	fmt.Println(sort.StringsAreSorted(xs))
}
