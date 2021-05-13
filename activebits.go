// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func ActiveBits(n int) int {
	binary := 0
	if n%2 == 0 {
		for n != 1 {
			n = n / 2
			binary++
		}
		return binary
	}
	for n != 0 {
		n = n / 2
		binary++
	}
	return binary
}

// func main() {
// 	fmt.Println(ActiveBits(14))
// 	fmt.Println(ActiveBits(7))
// }
