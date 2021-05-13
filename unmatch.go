// package main

// import (
// 	"fmt"
// 	// "piscine"
// )

package piscine

func Unmatch(a []int) int {
	// match := true
	for _, val := range a {
		counter := 0
		for _, num := range a {
			if val == num {
				counter++
			}
		}
		if counter == 1 || counter%2 == 1 {
			return val
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
