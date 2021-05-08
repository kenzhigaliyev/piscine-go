// package main

// import "fmt"

package piscine

func ConcatParams(args []string) string {
	length := len(args)
	words := ""
	for index, word := range args {
		// for _, letter := range word {
		words = words + word
		if index == length-1 {
		} else {
			words = words + "\n"
		}
		// words[index] = word
		// words[index+1] = '\n'
		// index = index + 2
	}
	return words
}

// func main() {
// 	test := []string{"Hello", "how", "are", "you?"}
// 	fmt.Println(ConcatParams(test))
// }
