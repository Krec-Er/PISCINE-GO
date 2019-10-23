package piscine

func LastRune(s string) rune {
	a := []rune(s)
	var len int
	for index := range a {
		len = index + 1
	}
	return a[len-1]
}
