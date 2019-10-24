package piscine

func Index(s string, toFind string) int {
	len := 0
	c := 0
	for i := range toFind {
		if i == i {
			len++
		}
	}
	for _, j := range toFind {
		for i2, j2 := range s {
			if j == j2 {
				if len > 1 {
					for k := 0; k < len; k++ {
						if s[i2+k] == toFind[k] {
							c++
						} else {
							return -1
						}
					}
					if c == len {
						return i2
					}
				} else if len == 1 {
					return i2
				} else {
					return -1
				}
			}
		}
		if c <= 0 {
			return -1
		}
	}
	return len
}
