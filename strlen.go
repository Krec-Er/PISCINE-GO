package piscine

import "fmt"

func PrintStr(str string) {

	for index := range str {
		fmt.Println(index + 1)
	}
}
