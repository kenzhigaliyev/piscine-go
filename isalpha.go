// package main

// import "fmt"

package piscine

func IsAlpha(s string) bool {
	chars := []rune(s)
	for index := range s {
		if (chars[index] <= 122 && chars[index] >= 97) || (chars[index] <= 92 && chars[index] >= 65) || (chars[index] <= 57 && chars[index] >= 48) || chars[index] == 0 {
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))
// 	fmt.Println(IsAlpha(""))

// }
