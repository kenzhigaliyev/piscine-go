package main

import "fmt"

// package piscine

func Capitalize(s string) string {
	sentence := []rune(s)
	length := len(sentence)
	other_chars := true
	check := true
	for index := 0; index < length; index++ {
		if !(sentence[index] <= 122 && sentence[index] >= 97) && !(sentence[index] <= 90 && sentence[index] >= 65) && !(sentence[index] <= 57 && sentence[index] >= 48) {
			other_chars = true
		} else if !(sentence[index] <= 122 && sentence[index] >= 97) && other_chars == true {
			other_chars = false
			for new_index := index + 1; new_index < length && check; new_index++ {
				if sentence[new_index] <= 90 && sentence[new_index] >= 65 {
					sentence[new_index] = sentence[new_index] + 32
					index++
				} else if sentence[index] <= 122 && sentence[index] >= 97 {
					index++
				} else {
					other_chars = true
					check = false
				}
			}
		} else if sentence[index] <= 90 && sentence[index] >= 65 {
			if !other_chars {
				sentence[index] = sentence[index] + 32
			}
		} else if sentence[index] <= 122 && sentence[index] >= 97 {
			if other_chars {
				sentence[index] = sentence[index] - 32
				other_chars = false
			}
		}
	}
	return string(sentence)
}

// for index := range s {
// 	if sentence[index] <= 122 && sentence[index] >= 97 {
// 		sentence[index] = sentence[index] - 32
// 	} else {
// 	}
// 	for index1 := range s[index+1:] {
// 		if sentence[index1] <= 90 && sentence[index1] >= 65 {
// 			sentence[index1] = sentence[index1] + 32
// 		} else if sentence[index1] {
// 			// Capitalize(string(sentence[index:]))
// 		}
// 	}
// }
// 	return string(sentence)
// }

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(Capitalize("abENxNY}F$6g!"))
}
