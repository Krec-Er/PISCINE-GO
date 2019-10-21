package piscine

func Fibonacci(nb int) int {

	if nb < 0 {
		return -1
	} else if nb == 0 {
		return 0
	} else if nb == 1 {
		return 1
	} else {
		return fibonacci(nb-1) + fibonacci(nb-2)
	}
}
