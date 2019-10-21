package piscine

func Fibonacci(nb int) int {

	if nb < 0 {
		return -1
	} else if nb == 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		return Fibonacci(nb-1) + Fibonacci(nb-2)
	}
}
