package main

import "fmt"

func intSequence() func() int {
	i := 0
	
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSequence()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intSequence()
	fmt.Println(newInt())
}