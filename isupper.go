package piscine

func IsUpper(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes {
		if word >= 'A' && word <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
