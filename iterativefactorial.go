package piscine

func IterativeFactorial (nb int) int {
		
		if nb < 0 {
			return 0
		} else if nb == 0 || nb == 1 {
			return 1
		}
		
		result := 1

		for i:= 1; i <= nb; i++ {
			result = result * i
		}
	return result
	}
