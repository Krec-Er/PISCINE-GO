package piscine

func AlphaCount(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		count++
		if !(str[i] >= 'A') && (str[i] <= 'Z') && !(str[i] >= 'a') && (str[i] <= 'z') {
			count--
		}
	}
	return count
}
