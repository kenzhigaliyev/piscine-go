// package main

// import "fmt"

package piscine

func Capitalize(s string) string {
	sentence := []rune(s)
	if sentence[0] >= 97 && sentence[0] <= 122 {
		sentence[0] = sentence[0] - 32
	}
	for index := range s {
		if sentence[index] >= 97 && sentence[index] <= 122 {
			sentence[index] = sentence[index]
		} else if (sentence[index] >= 65 && sentence[index] >= 90) && (sentence[index-1] >= 97 && sentence[index-1] <= 122) && (sentence[index-1] >= 65 && sentence[index-1] >= 90) {
			sentence[index] = sentence[index]
		} else if (sentence[index] >= 97 && sentence[index] <= 122) && !(sentence[index-1] >= 97 && sentence[index-1] <= 122) && !(sentence[index-1] >= 65 && sentence[index-1] >= 90) {
			if !(sentence[index-1] >= 48 && sentence[index-1] <= 57) {
				sentence[index] = sentence[index] - 32
			}
		} else {
		}
	}

	return string(sentence)
}

// func main() {
// 	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
// 	fmt.Println(Capitalize("abENxNY}F$6g!"))
// 	fmt.Println(Capitalize("JiNh7HNDza>la"))
// }
