package piscine

// package main

// import (
// 	"fmt"
// 	// "piscine"
// )

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b

	*mod = a % b
}

// func main() {
// 	a := 13
// 	b := 2
// 	var div int
// 	var mod int
// 	piscine.DivMod(a, b, &div, &mod)
// 	fmt.Println(div)
// 	fmt.Println(mod)
// }
