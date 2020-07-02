package main

import "fmt"

func main() {

	//Construct an empty array
	//of size=3
	s := make([]string, 3)
	fmt.Println("empty:", s)

	//Enter values into the
	//constructed array.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended:", s)

	//Make an array of size 's'
	//and copy the data into it.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//Take a slice from the 'c'
	//array.
	slice_1 := c[2:5]
	fmt.Println("slice 1:", slice_1)

	//Single line implementation of
	//a slice of strings.
	t := []string{"g", "h", "i"}
	fmt.Println("declared:", t)

	//Make a 2D array
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j:= 0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}