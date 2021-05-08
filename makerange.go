// package main
package piscine

// import "fmt"

func MakeRange(min, max int) []int {
	if max > min {
		return nil
	}
	numbers := make([]int, max-min)
	for i := 0; i < (max - min); i++ {
		numbers[i] = min + i
	}
	return numbers
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
