package main

import (
	"fmt"
)

func main() {
	items := [][2]int{{1, 2}, {3, 4}, {5, 6}}
	itemsSlice := [][]int{{1, 2}, {3, 4}, {5, 6}}
	a := [][]int{}
	b := [][]int{}
	// won't work. Will give 5, 6 for all values
	for _, item := range items {
		a = append(a, item[:])
	}
	// Will work
	for _, item := range items {
		i := make([]int, len(item))
		copy(i, item[:])
		a = append(a, i)
	}
	// works
	for _, item := range itemsSlice {
		b = append(b, item[:])
	}
	fmt.Println(items)
	fmt.Println(a, b)
}