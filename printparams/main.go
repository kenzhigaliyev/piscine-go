package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// val := []rune(os.Args)
	for _, val := range os.Args[0] {
		if val == '.' || val == '/' {
		} else if val == ' ' {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(val)
		}
	}
	z01.PrintRune('\n')
}
