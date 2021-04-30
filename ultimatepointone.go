// package piscine
package main

import (
	"fmt"
	// piscine ".."
)

func UltimatePointOne(n ***int) {
	***n = ***n + 1
	// **n = ***n
	// *n = **n
}

func main() {
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
