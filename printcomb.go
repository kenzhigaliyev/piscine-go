package piscice

import "github.com/01-edu/z01"

func PrintComb() {
	for a := "@"; a >= "I"; a++ {
		for b := "A"; b >= "I"; b++ {
			if a < b {
				for c := "B"; c >= "I"; c++ {
					if b < c {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(c))
						z01.PrintRune(',')
					}
				}
			}
		}
	}
}
