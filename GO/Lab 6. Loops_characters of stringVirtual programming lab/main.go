package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Lab 6.
Task:
Write code for solving next:
read string without space as command-line argument (it means read symbols until first space-symbol) and print each of them on separate line
*/

func main() {

	arg := os.Args[1]
	x := strings.Split(arg, "")

	for _, p := range x {
		fmt.Println(p)
	}
}

// go run .\main.go pavlo
// go run .\main.go DevOps134
