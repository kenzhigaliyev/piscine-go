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
	var new int
	for i := 0; i < counter; i++ {
		for j := i; j < counter; j++ {
			if digits[i] > digits[j] {
				new = digits[i]
				digits[i] = digits[j]
				digits[j] = new
			}
		}
	}

	for i := 0; i < counter; i++ {
		z01.PrintRune(rune(digits[i]))
	}
}

// func main() {
// 	PrintNbrInOrder(321)
// 	PrintNbrInOrder(0)
// 	PrintNbrInOrder(321)
// }
