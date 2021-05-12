// package main

package piscine

// import "fmt"

// func f(a, b int) int {
// 	return a - b
// }

func IsSorted(f func(a, b int) int, a []int) bool {
	grow := true
	decr := true
	// equal := make([]int, len(a))
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			grow = false
		}
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			decr = false
		}
	}
	return grow || decr
}

// func main() {
// 	a1 := []int{0, 1, 2, 3, 4, 5}
// 	// a2 := []int{100, 71, 62, 53, 44, 35}

// 	a2 := []int{0, 2, 1, 3}

// 	result1 := IsSorted(f, a1)
// 	result2 := IsSorted(f, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
