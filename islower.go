// package main

// import "fmt"

package piscine

func IsLower(s string) bool {
	chars := []rune(s)
	for index := range s {
		if chars[index] <= 132 && chars[index] >= 97 {
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsLower("hello"))
// 	fmt.Println(IsLower("hello!"))

// }
