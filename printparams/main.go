package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for index, val := range os.Args {
		// if val == '.' || val == '/' || val == '"' {
		// } else if val == ' ' {
		// 	z01.PrintRune('\n')
		// } else {
		// 	z01.PrintRune(val)
		// }
		if index > 0 {
			for _, new_val := range val {
				z01.PrintRune(new_val)
			}
		}
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')
}
