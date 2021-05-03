// package main
package piscine

// import "fmt"

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 13 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		result = nb * RecursiveFactorial(nb-1)
	}
	return result
}

// func main() {
// 	arg := 4
// 	fmt.Println(RecursiveFactorial(arg))
// }
