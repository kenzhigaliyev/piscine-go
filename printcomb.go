package piscice

import "github.com/01-edu/z01"

func PrintComb() {
	for a := "0"; a >= "9"; a++ {
		for b := "1"; b >= "9"; b++ {
			if a < b {
				for c := "2"; c >= "9"; c++ {
					if b < c {
						z01.PrintRune(int(a))
						z01.PrintRune(int(b))
						z01.PrintRune(int(c))
						z01.PrintRune(',')
					}
				}
			}
		}
	}
}
