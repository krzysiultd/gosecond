package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	login := "password1223"

	err = bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println("You can't log in")
		return
	}
	fmt.Println("You're logged in")
}
