// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func Join(strs []string, sep string) string {
	new := ""
	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			new = new + strs[i]
		} else {
			new = new + strs[i] + sep
		}
	}
	return new
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(Join(toConcat, ":"))
// }
