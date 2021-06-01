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
	record2 := Record{
		one:   input[0],
		three: input[1],
		two:   input[2],
	}

	fmt.Println(record1)
	fmt.Println(record2)

	map1 := map[string]int{}

	fmt.Println(map1)

	map2 := map[string]int{
		"Three": 3,
		"Four":  4,
		"One":   5,
	}

	for k, v := range map2 {
		map1[k] = v
	}

	fmt.Println(map1)
}
