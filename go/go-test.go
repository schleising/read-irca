package main

import "fmt"

func main() {
	var testMap = make(map[string]string)

	testMap["one"] = "first Entry"
	testMap["two"] = "Second Entry"
	testMap["three"] = "Third Entry"
	testMap["four"] = "Forth Entry"

	fmt.Println(testMap)
	fmt.Println()
	fmt.Println(testMap["three"])
}
