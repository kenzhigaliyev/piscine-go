package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	length := len(s)
	readString := []byte(s)
	for i := 0; i <= length; i++ {
		a := readString[i]
		z01.PrintRune(a)
	}
}
