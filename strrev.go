package piscine

func StrRev(s string) string {

	runes := []rune(s)
	var word string
	count := StrLen(s)
	for i := count - 1; i >= 0; i-- {
		word = word + string(runes[i])
	}
	return word
}
