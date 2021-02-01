package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("THIS IS IF PRINT", v)
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS IF PRINT", v)
	}

	m["krzysiu"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 33}

	for i, v := range xi {
		fmt.Println(i, v)
	}

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

}
