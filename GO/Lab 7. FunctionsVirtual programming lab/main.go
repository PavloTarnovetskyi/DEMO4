package main

import (
	"fmt"
)

/*
Lab 7.
Task:
Implement function basicOperations which takes two integer numbers as arguments and returns calculated expressions (see Lab 3):
a + b, a - b, a * b, a / b
*/

func basicOperations(x float64, y float64) {

	fmt.Println("a + b = ", x+y)
	fmt.Println("a - b = ", x-y)
	fmt.Println("a * b = ", x*y)
	fmt.Println("a / b = ", x/y)

}

func main() {

	var a, b float64
	fmt.Println("Type a and b value: ")
	fmt.Scanln(&a, &b)
	basicOperations(a, b)
}

// go run .\main.go
