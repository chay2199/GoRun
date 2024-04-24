package main

import "fmt"

func main() {
	ch := make(chan int, 1) // works
	//ch := make(chan int) // deadlock
	ch <- 1
	b, ok := <-ch
	fmt.Println(b, ok) // 1 true

	close(ch)

	c, ok := <-ch
	fmt.Println(c, ok) // 0 false
}
