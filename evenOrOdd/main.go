package main

import "fmt"

func main() {

	myInts := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenOdd string

	for _, num := range myInts {
		if num%2 == 0 {
			evenOdd = "even"
		} else {
			evenOdd = "odd"
		}
		fmt.Printf("Number: %d - Even/Odd: %s\n", num, evenOdd)
	}
}
