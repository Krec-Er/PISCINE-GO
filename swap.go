package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a
	*a = b 
	*b = c
}
