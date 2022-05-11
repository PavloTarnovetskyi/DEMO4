package main

import (
	"fmt"
)

/*
Lab 2.
Task:
Write code for solving next:
Read value as command-line argument and detect if this one is integer. If so, print OK, otherwise print Wrong.
*/

func main() {

	fmt.Println("Type some value: ")

	var input string
	fmt.Scanln(&input)

	var isInteger int

	_, x := fmt.Sscan(input, &isInteger)

	if x == nil {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong")
	}
}

// go run .\main.go
