package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcname := os.Args[1:]
	for _, argument := range funcname {
		for _, word := range argument {
			z01.PrintRune(word)
		}
		z01.PrintRune(10)
	}
}
