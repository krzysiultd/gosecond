package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":49},{"First":"Ms","Last":"Moneypenny","Age":32}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data", people)

	for i, v := range people {
		fmt.Println("---PERSON NO", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
}
