package main

import "fmt"
import "os"
import "github.com/chay2199/hello"

func main() {
	fmt.Print(hello.Say(os.Args[1:]))
}