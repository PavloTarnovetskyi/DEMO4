package main

import "fmt"

/*
Lab 3.
Task:
Write code for solving next:
define integer variables a and b, read values a and b as command-line arguments and print calculated expressions:
a + b, a - b, a * b, a / b.
*/
var a, b int

func main() {
	fmt.Println("Type a and b value: ")
	fmt.Scanln(&a, &b)
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
}

// go run .\main.go
