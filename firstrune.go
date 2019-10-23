package piscine

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	for _, word := range s {
		return word
	}
	return 0
}
