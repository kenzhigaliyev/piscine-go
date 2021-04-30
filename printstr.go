package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, u := range s {
		v := string(u)
		z01.PrintRune(v)
	}
}
