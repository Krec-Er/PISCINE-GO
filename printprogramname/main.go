package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argument := os.Args
	str := argument[0]
	funcname := []rune(str)
	for _, word := range funcname {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
