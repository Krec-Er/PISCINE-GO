package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcname := os.Args[1:]
	var arguments int
	for arguments := range funcname {
		arguments = arguments + 1
	}
	for _, argument := range funcname[arguments-1] {
		for _, word := range argument {
			z01.PrintRune(word)
		}
		z01.PrintRune(10)
	}
}
