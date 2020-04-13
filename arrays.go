package main

import "fmt"

func main() {
	var a[5]int //Create an array w/ 5 slots.
	fmt.Println("empty array:", a)

	a[4] = 100 //Set slot 4 to value 100
	fmt.Println("set array a:", a)
	fmt.Println("get array:", a[4])

	fmt.Println("size of array:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("set array b:", b)

	var twoDim [2][3]int //Demonstrate how to print the 2D array out.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2D array:", twoDim)
}