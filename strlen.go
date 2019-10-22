package piscine

func PrintStr(str string) {

	for index := range str {
		fmt.Printf(index + 1)
	}
}
