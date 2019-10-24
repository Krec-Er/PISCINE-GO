package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	Runes := []rune(str)
	for _, word := range Runes {
		if word >= 0 && word <= 31 {
			continue
		} else if {
			return true
		}
	}
	return false
}
