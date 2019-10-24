package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args
	strings := []string(arguments)
	for index, strs := range strings {
		if index >= 1 {
			runes := []rune(strs)
			for _, runess := range runes {
				z01.PrintRune(runess)
			}
			z01.PrintRune(10)
		}
	}
}
