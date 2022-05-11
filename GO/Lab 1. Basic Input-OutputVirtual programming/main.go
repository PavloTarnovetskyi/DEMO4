package main

import "fmt"

/*
Lab 1.
Task:
Write code for solving next:
Output question "How are you?". Read the answer value from console and output: "You are (answer)"
*/

func main() {

	fmt.Println("How are you?")
	var answer string
	fmt.Scan(&answer)
	fmt.Printf("You are %s!", answer)

	var a string
	fmt.Scan(&a)
}

// go run .\main.go
