package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	// length := len(s)
	// readString := []byte(s)
	for _, u := range s {
		var v := string(u)
		z01.PrintRune(v)
	}
}
