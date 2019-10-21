package piscine

func IterativeFactorial (nb int) int
		
		i nb < 0 {
			return 0
		} else if nb == 0 || nb == 1 {
			return 1
		}
		
		result := nb

		for i:= nb-1: i < 1; i-- {
			result = nb * i
		}
	return result
	}
