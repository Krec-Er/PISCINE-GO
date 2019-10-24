package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcname := os.Args[1:]
	var index int
	for index := range funcname {
		index = index
	}
	for i, j := 0, index; i < j; i, j = i+1, j-1 {
		funcname[i], funcname[j] = funcname[j], funcname[i]
	}
	for _, argument := range funcname {
		for _, word := range argument {
			z01.PrintRune(word)
		}
		z01.PrintRune(10)
	}
}
