package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// simple program to replace strings
func main() {
	word := "énum"
	fmt.Println(len(word))

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old_word, new_word := os.Args[1], os.Args[2]
	scan :=  bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		s := strings.Split(scan.Text(), old_word)
		fmt.Println(s, len(s))
		t := strings.Join(s, new_word)
		fmt.Println(t)
	}
}