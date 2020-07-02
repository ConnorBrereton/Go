package main

import "fmt"

// simple function that sums two
// integers and returns the sum.
func plus(a int, b int) int {
	return a + b
}

// simple function that sums
// three intergers and returns
// the sum.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	result := plus(1, 2)
	fmt.Println("1+2=", result)

	res := plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}