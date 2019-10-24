package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcname := os.Args[1:]
	var len int
	for index := range funcname {
		len = index + 1
	}
	for i := 0; i <= (len-1), i++ {
		funcname[i] = funname [len-1-i]
	} 
	for _, argument := range funcname {
		for _, word := range argument {
			z01.PrintRune(word)
		}
		z01.PrintRune(10)
	}
}
