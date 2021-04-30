package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, u := range s {
		var v [u]int
		z01.PrintRune(v)
	}
}
