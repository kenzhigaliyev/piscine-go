// package main
package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	var digits []int
	counter := 0
	number := 0
	if n <= 0 {
		z01.PrintRune('0')
	}
	for n != 0 {
		number = n % 10
		digits = append(digits, number+48)
		n = n / 10
		counter++
	}
	for i := 0; i < counter; i++ {

		z01.PrintRune(rune(digits[i]))
	}
	// for i:=0; i<=counter;i++{
	// 	if digits
	// }
}

// func main() {
// 	PrintNbrInOrder(321)
// 	PrintNbrInOrder(0)
// 	PrintNbrInOrder(321)
// }
