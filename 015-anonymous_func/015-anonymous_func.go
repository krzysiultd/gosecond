package main

import (
	"fmt"
)

func main() {
	func(x string) {
		fmt.Println("anonymous func says", x)
	}("hellooo")
}
