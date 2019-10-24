package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes {
		if (word >= 0 && word <= 32) {
					continue
		} else {
			return false
		}
	}
	return true
}
