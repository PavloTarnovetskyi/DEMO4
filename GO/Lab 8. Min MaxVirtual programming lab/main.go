package main

import "fmt"

/*
Lab 8.
Task:
Implement function getMinMax which takes array as an argument and print max and min of its elements
*/

func findMinAndMax(a [4]int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func main() {

	var a, b, c, d int
	fmt.Println("Type a, b, c and d value: ")
	fmt.Scanln(&a, &b, &c, &d)
	var array = [4]int{a, b, c, d}
	min, max := findMinAndMax(array)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

// go run .\main.go
