package piscine

func StrRev(s string) string {
	a := 0
	for _, letter := range s {
		_ = letter
		a++
	}
	for b := 0; b < a/2; b++ {
		rune[a], rune[a-1-b] = rune[a-1-b], rune[b]
	}
	new := string(rune)
	return new
}
