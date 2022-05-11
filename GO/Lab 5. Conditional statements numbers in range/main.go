package main

import "fmt"

/*
Lab 5.
Task:
Write code for solving next:
read 3 float numbers as command-line arguments and if they all belong to the range [-5,5],
 print OK, otherwise print Wrong.
*/
func main() {

	var a, b, c int

	fmt.Println("Type a, b, c nubers: ")

	fmt.Scanln(&a, &b, &c)

	var array = [3]int{a, b, c}

	for i := 0; i < len(array); i++ {
		if array[i] < 5 && array[i] > -5 {
			fmt.Println("OK")
		} else {
			fmt.Println("Wrong")
		}

	}
}

// go run .\main.go
