// package main

package piscine

// import "fmt"

func Index(s string, toFind string) int {
	first := []rune(s)
	letter := []rune(toFind)

	counter := 0
	for index1 := range toFind {
		if (int(letter[index1]) <= 90 && int(letter[index1]) >= 65) || (int(letter[index1]) <= 122 && int(letter[index1]) >= 97) {
			counter++
		} else {
			counter++
		}
	}

	counter1 := 0
	for index2 := range s {
		if (int(first[index2]) <= 90 && int(first[index2]) >= 65) || (int(first[index2]) <= 122 && int(first[index2]) >= 97) {
			counter1++
		} else {
			counter1++
		}
	}

	for i := 0; i <= counter1-counter; i++ {
		if s[i:i+counter] == toFind {
			return i
		}
	}
	return -1
}

// func main() {
// 	fmt.Println(Index("Hello!", "l"))
// 	fmt.Println(Index("Salut!", "Salut"))
// 	fmt.Println(Index("nSSYMr^5v0TQQ", "nSSYMr^5v0TQ"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
