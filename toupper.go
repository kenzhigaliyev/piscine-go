// package main
package piscine

// import "fmt"

func ToUpper(s string) string {
	letter := []rune(s)
	for index := range s {
		if letter[index] <= 122 && letter[index] >= 97 {
			letter[index] = letter[index] - 32
		}
	}
	return string(letter)
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
