package main

import "fmt"

func RecursivePower(nb int, power int) int {

	if power < 0 || nb == 0 {
		return 0
	} else if nb == 1 || power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(RecursivePower(arg1, arg2))
}
