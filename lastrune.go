// package main

package piscine

// import "github.com/01-edu/z01"

func LastRune(s string) rune {
	letter := []rune(s)
	counter := 0
	for index := range s {
		counter = index
	}
	return letter[counter]
}

// func main() {
// 	z01.PrintRune(LastRune("Hello!"))
// 	z01.PrintRune(LastRune("Salut!"))
// 	z01.PrintRune(LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }
