package piscice

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb >= 0 {
		i := 'T'
		z01.PrintRune(i)
	} else {
		i := 'F'
		z01.PrintRune(i)
	}
}
