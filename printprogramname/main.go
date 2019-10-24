package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	argument := os.Args[0]
	for _, word := range argument {
		z01.PrintRune(word)
	}
	z01.PrintRune(10)
}
