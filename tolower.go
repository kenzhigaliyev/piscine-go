// package main

// import "fmt"

package piscine

func ToLower(s string) string {
	letter := []rune(s)
	for index := range s {
		if letter[index] <= 90 && letter[index] >= 65 {
			letter[index] = letter[index] + 32
		}
	}
	return string(letter)
}

// func main() {
// 	fmt.Println(ToLower("Hello! How are you?"))
// }
