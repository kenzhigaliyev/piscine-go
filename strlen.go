package piscine

func StrLen(s string) int {
	a := 0
	for _, letter := range s {
		a++
	}
	return a
}
