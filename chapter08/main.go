package main

import (
	"fmt"
)

// you can name you return value. So this is an example of a naked return
func doIt() (a int) {
	defer func() {
		a = 3
	}()

	a = 15
	return
}

func main() {
	a := doIt()
	fmt.Println(a)
}