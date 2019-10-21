package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		res := 0
		for i := 0; i*i <= nb; i++ {
			res = i
		}
		if res*res != nb {
			return 0
		} else {
			return res
		}

	} else {
		return 0
	}
}
