// package main

// import "fmt"

package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	length := len(s)
	letters := ""
	for index, word := range s {
		if index == length-1 {
			letters = letters + string(word)
			words = append(words, letters)
		} else if word != ' ' {
			letters = letters + string(word)
		} else {
			words = append(words, letters)
			letters = ""
		}
	}
	return words
}

// func main() {
// 	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
// }
