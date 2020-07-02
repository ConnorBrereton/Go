package main

import "fmt"

func main() {

	// makes a map aka a dictionary
	m := make(map[string]int)

	// sets the key1:42 pair
	m["key1"] = 42
	m["key2"] = 24

	fmt.Println("map:", m)

	// prints the values
	value1 := m["key1"]
	fmt.Println("value1", value1)
	fmt.Println("length:", len(m))

	//delete from map 'm' key2
	delete(m, "key2")
	fmt.Println("map:", m)

	//attempt to access deleted key
	_, prs := m["key2"]
	fmt.Println("value:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}