// package main
package piscine

// import "github.com/01-edu/z01"

func NRune(s string, n int) rune {
	letter := []rune(s)
	// counter := 0
	if n < 0 {
		return 0
	}
	for index := range s {
		if index == n-1 {
			return letter[index]
		}
	}
	return 0
}

// func main() {
// 	z01.PrintRune(NRune("Hello!", 3))
// 	z01.PrintRune(NRune("Salut!", 2))
// 	z01.PrintRune(NRune("Bye!", -1))
// 	z01.PrintRune(NRune("Bye!", 5))
// 	z01.PrintRune(NRune("Ola!", 4))
// 	z01.PrintRune('\n')
// }
