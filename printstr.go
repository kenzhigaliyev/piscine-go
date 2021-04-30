package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	// st := []s
	length := len(s)
	runes := []rune(s)
	for i:=0;i<=length:i++ {
		z01.PrintRune(runes)
	}
}
