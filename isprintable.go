// package main
package piscine

// import "fmt"

func IsPrintable(s string) bool {
	chars := []rune(s)
	for index := range s {
		if (chars[index] < 32 && chars[index] >= 0) || chars[index] == 92 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrintable("Hello"))
// 	fmt.Println(IsPrintable("Hello\n"))

// }
