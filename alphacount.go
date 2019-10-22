package piscine

func AlphaCount(str string) int {
    counter := 0
	slice := []rune(str)
  	for _, word	:= range slice {			 
        if word >= 65 && word <= 92 || word >= 97 && word <= 122 {
             counter ++
        }
    }
    return counter
}

