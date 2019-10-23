package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var a [12]int
	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else {
		for {
			if n == 0 {
				break
			}
			a[n%10]++
			n /= 10
		}
		for index, number := range a {
			if number != 0 {
				for i := 0; i < number; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}
}
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}
