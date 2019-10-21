package piscine

func IterativePower(nb int, power int) int {

	if nb == 1 || power == 0 {
		return 1
	} else if power < 0 || nb == 0 {
		return 0
	}
	result := 1

	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
}
