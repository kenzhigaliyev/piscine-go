package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	counter := 0
	for index := range os.Args {
		counter = index
	}
	for new_index := counter; new_index > 0; new_index-- {
		for _, new_val := range os.Args[new_index] {
			z01.PrintRune(new_val)
		}
		z01.PrintRune('\n')
	}
	// z01.PrintRune('\n')
}
