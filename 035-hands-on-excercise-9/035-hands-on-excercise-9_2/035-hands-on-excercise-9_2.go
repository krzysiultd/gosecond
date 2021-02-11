package main

import "fmt"

type mobilePhone struct {
	number int64
}

type telephone interface {
	call() string
}

func (m *mobilePhone) call() string {
	s := fmt.Sprint("+48", m.number)
	return s
}

func message(t telephone) {
	fmt.Println(t.call())
}

func main() {
	m1 := mobilePhone{882033921}
	message(&m1)
}
