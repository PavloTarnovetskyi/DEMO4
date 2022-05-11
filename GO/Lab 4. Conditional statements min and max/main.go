package main

import "fmt"

/*
Lab 4.
Task:
Write code for solving next:
read 3 integer numbers as command-line arguments and print max and min of them.
*/

func main() {

	var a, b, c int

	fmt.Println("Type a, b, c nubers: ")

	fmt.Scanln(&a, &b, &c)

	var array = [3]int{a, b, c}
	min, max := MinMax(array)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
func MinMax(array [3]int) (min int, max int) {
	min = array[0]
	max = array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// go run .\main.go
