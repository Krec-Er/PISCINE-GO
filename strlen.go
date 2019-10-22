package piscine

func PrintStr(str string) int {
	a := []rune(str)
	var len int
	for i := range a {
		len = i + 1
	}
	return len
}
