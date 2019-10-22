package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	for index, _ := range str {
		z01.PrintRune(index + 1)
	}
}
