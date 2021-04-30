package piscice

import "github.com/01-edu/z01"

func PrintComb() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					e := a + b
					f := c + d
					if e < f {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(' ')
						z01.PrintRune(rune(c))
						z01.PrintRune(rune(d))
						z01.PrintRune(',')
						z01.PrintRune(' ')
						sdfsdf
					}
				}
			}
		}
	}
}
