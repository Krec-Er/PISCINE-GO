package piscine

func Capitalize(s string) string {
	Runes := []rune(s)
	len := 0
	for i := range Runes {
		i = i
		len++
	}
	if Runes[0] >= 'a' && Runes[0] <= 'z' {
		Runes[0] = Runes[0] - 32
	}

	for j := 1; j < len; j++ {
		if Runes[j] >= 'a' && Runes[j] <= 'z' {
			if (Runes[j-1] >= 'A' && Runes[j-1] <= 'Z') || (Runes[j-1] >= '0' && Runes[j-1] <= '9') || (Runes[j-1] >= 'a' && Runes[j-1] <= 'z') {
				Runes[j] = Runes[j]
			} else {
				Runes[j] = Runes[j] - 32
			}
		} else if Runes[j] >= 'A' && Runes[j] <= 'Z' {
			if (Runes[j-1] >= 'A' && Runes[j-1] <= 'Z') || (Runes[j-1] >= '0' && Runes[j-1] <= '9') || (Runes[j-1] >= 'a' && Runes[j-1] <= 'z') {
				Runes[j] = Runes[j] + 32
			} else {
				Runes[j] = Runes[j]
			}

		} else {
			Runes[j] = Runes[j]
		}
	}
	return string(Runes)
}
