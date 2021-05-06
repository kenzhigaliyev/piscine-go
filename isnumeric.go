// package main
package piscine

// import "fmt"

func IsNumeric(s string) bool {
	chars := []rune(s)
	for index := range s {
		if chars[index] <= 57 && chars[index] >= 48 {
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsNumeric("010203"))
// 	fmt.Println(IsNumeric("01,02,03"))
// }
