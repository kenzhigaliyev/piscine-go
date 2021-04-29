package piscice

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			if a < b {
				for c := b + 1; c <= '9'; c++ {
					if b < c {
						if a == '7' && b == '8' && c == '9' {
							z01.PrintRune(rune(a))
							z01.PrintRune(rune(b))
							z01.PrintRune(rune(c))
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(rune(a))
							z01.PrintRune(rune(b))
							z01.PrintRune(rune(c))
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
