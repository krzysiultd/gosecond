package main

import "fmt"

func main() {
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
	for2()
	for3()

}

func for2() {
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}

func for3() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
