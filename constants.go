package main

import (
	"fmt"
	"math"
)

const s string = "go is tight"

func main() {
	fmt.Println(s)

	const number = 50000000000
	const Bignum = 3e20 / number
	fmt.Println(Bignum)

	fmt.Println(int64(Bignum))

	fmt.Println(math.Sin(number))
}