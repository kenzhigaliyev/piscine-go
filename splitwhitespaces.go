package main

import "fmt"

// package piscine

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
			if letters == "" {
			} else {
				words = append(words, letters)
				letters = ""
			}
		}
	}
	return words
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", SplitWhiteSpaces("H.`G9wbN:}e>  szb,&88gQy]m< Jb?6YpS(Q}2&3 PJ\\/4Ssdi|xnG UH>\\L#N-,IqH{ n^SbU5!6WN^+n k `Sc&#+namj+ na7hvk\\Ti#L$H"))
}
