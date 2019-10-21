package piscine

func IterativePower(nb int, power int) int {

	if nb < 0  {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else if power < 0 {
		return 0
	} else if power == 0 || power == 1 {
		return 1
	}

	result := nb

	for i := 1; i <= power - 1; i++ {
		result = result * nb
	}
	return result
}
