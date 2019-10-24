package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes {
		if word < 32 {
			return false
		} 
	}
	return true
}
