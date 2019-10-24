package piscine

func Index(s string, toFind string) int {
	n := 0
	j := []rune(s)
	for p := range j {
		n = p + 1
	}
	k := 0
	l := []rune(toFind)
	for r := range l {
		n = r + 1
	}

	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		} else {
			return -1
		}
	}
}
