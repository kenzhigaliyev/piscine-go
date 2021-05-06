// package main
package piscine

// import "fmt"

func IsUpper(s string) bool {
	chars := []rune(s)
	for index := range s {
		if chars[index] <= 90 && chars[index] >= 65 {
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsUpper("HELLO"))
// 	fmt.Println(IsUpper("HELLO!"))

// }
