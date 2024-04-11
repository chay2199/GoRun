package main

import (
	"fmt"
)

func main() {
	b := []int{1, 2, 3}
	// start, end and capacity
	c := b[0:2:2]
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c[0] = 9
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 5)
	// c will now point to a different location
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)
}