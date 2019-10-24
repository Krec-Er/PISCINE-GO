package piscine

func ToUpper(s string) string {
	Runes := []rune(s)
	for i := range Runes {
		if Runes[i] >= 'A' && Runes[i] <= 'Z' {
			Runes[i] = Runes[i] + 32
		}
	}
	return string(Runes)
}
