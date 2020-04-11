package main

import "fmt"

func main() {
	var a = "this is a variable assignment"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// this is a variable assignment
	// without having to use the `var`
	// declaration.
	f := "oranges"
	fmt.Println(f)
}