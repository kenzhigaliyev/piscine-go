// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func ActiveBits(n int) int {
	binary := 0
	// ramainder := 0
	counter := 1
	bin := 0
	for n != 0 {
		remainder := n % 2
		if remainder == 1 {
			bin++
		}
		n = n / 2
		binary = binary + (remainder * counter)
		counter = counter * 10
	}
	return bin
}

// func main() {
// 	fmt.Println(ActiveBits(14))
// 	fmt.Println(ActiveBits(20))
// 	fmt.Println(ActiveBits(7))
// }
