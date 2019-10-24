package piscine

func IsLower(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes {
		if (word >= 'a' && word <= 'z') {
			continue
		} else {
			return false
		}
	}
	return true
}
