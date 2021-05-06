// package main
package piscine

// import "fmt"

func AlphaCount(s string) int {
	letter := []rune(s)
	counter := 0
	for index := range s {
		if (int(letter[index]) <= 90 && int(letter[index]) >= 65) || (int(letter[index]) <= 122 && int(letter[index]) >= 97) {
			counter++
		}
	}
	return counter
}

// func main() {
// 	s := "Hello 78 World!   123dasadsad 4455 /"
// 	nb := AlphaCount(s)
// 	fmt.Println(nb)
// }
