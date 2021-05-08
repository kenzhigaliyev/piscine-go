// package main

package piscine

// import "fmt"

func AppendRange(min, max int) []int {
	numbers := make([]int, 0)
	if max > min {
		for i := 0; i < max-min; i++ {
			numbers = append(numbers, min+i)
		}
	} else {
		numbers = nil
	}
	return numbers
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 5))
// 	fmt.Println(AppendRange(0, 0))

// }
