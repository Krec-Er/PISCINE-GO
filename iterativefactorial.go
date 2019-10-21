	func iterativefactorial(n int) int {
		i n < 0 {
			return 0
		} else if n == 0 || n == 1 {
			return 1
		}
		
		result := n

		for i:= n-1: i < 1; i-- {
			result = n * i
		}
	return result *= i
	}
