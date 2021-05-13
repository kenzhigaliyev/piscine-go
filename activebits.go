package main

// package piscine

import (
	"fmt"
	// "piscine"
)

func ActiveBits(n int) int {
	binary := 0
	for n != 0 {
		n = n / 2
		binary++
	}
	return binary
}

func main() {
	fmt.Println(ActiveBits(14))
}
