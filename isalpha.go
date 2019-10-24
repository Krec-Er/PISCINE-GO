package piscine

func IsAlpha(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes{
		if (word >= 'A' && word <= 'Z') ||
			word >= '0' && word <= '9') ||
			(word >= 'a' && word <= 'z') {
			continue
		} else {
			return false
		}
	}
	return true
}
