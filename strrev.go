package piscine

func StrRev(s string) string {
	a := 0
	r := []rune(s)
	for _, letter := range s {
		_ = letter
		a++
	}
	for b := 0; b < a/2; b++ {
		x := r[b]
		y := r[a-1-b]
		r[b] = y
		r[a-1-b] = x
	}
	new := string(r)
	return new
}
