package piscine 

import "github.com/01-edu/z01"

func printcomb() {
	for i:= '0'; i<='9'; i++ {
		for j:= '0'; j<= '9'; j++ {
			for:='0'; k<= '9'; k++ {
				if i<j && j<k {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				if i:= '7' && j:= '8' && k:= '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')}
				}	
		
			}
		}
	}
}
