package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	funcName := os.Args[0]
		z01.PrintRune(funcName)
	}
	z01.PrintRune('\n')
}
