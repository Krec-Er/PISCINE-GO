package piscine

func IsAlpha(str string) bool {
	Runes := []rune(str)
	len := 0
	for i := range Runes {
		i = i
		len++
	}
	for j := 0; j <= len; j++ {
		if (Runes[j] >= 'A' && Runes[j] <= 'Z') ||
			(Runes[j] >= '0' && Runes[j] <= '9') ||
			(Runes[j] >= 'a' && Runes[j] <= 'z') {
			continue
		} else {
			return false
		}
	}
	return true
}
