// package main

package piscine

// import (
// 	"fmt"
// 	// "piscine"
// )

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	for i := range a {
		max := true
		for j := i; j < len(a); j++ {
			if a[i] < a[j] {
				max = false
			}
		}
		if max {
			return a[i]
		}
	}
	return 0
}

// func main() {
// 	a := []int{23, 123, 1, 11, 55, 93}
// 	max := Max(a)
// 	fmt.Println(max)
// }
