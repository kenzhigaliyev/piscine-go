package piscine

// package main

// import (
// 	"fmt"

// 	"piscine"
// )

func StrLen(s string) int {
	a := 0
	for _, letter := range s {
		a++
		_ = letter
	}
	// a := len(s)
	return a
}

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }
