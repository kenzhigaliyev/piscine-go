// package main

package piscine

// import "fmt"

func IterativeFactorial(nb int) int {
	if nb <= 0 || nb < 13 {
		return 0
	} else {
		result := 1
		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	}
}

// func main() {
// 	arg := 1132414124124124
// 	fmt.Println(IterativeFactorial(arg))
// }
