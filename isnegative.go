package piscice

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {
	if nb >= 0 {
		i := 'F'
		z01.PrintRune(i)
		z01.PrintRune('\n')
	} else {
		i := 'T'
		z01.PrintRune(i)
		z01.PrintRune('\n')
	}
}
