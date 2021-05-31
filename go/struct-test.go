package main

import "fmt"

type Record struct {
	one   string
	two   string
	three string
}

func main() {
	input := []string{"four", "five", "six"}
	record1 := Record{"1", "2.3", "Test"}
	record2 := Record{one: input[0], two: input[1], three: input[2]}

	fmt.Println(record1)
	fmt.Println(record2)
}
