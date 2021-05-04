// package main
package piscine

// import "fmt"

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		result := i * i
		if result == nb {
			return i
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(Sqrt(6))
// 	fmt.Println(Sqrt(0))
// 	fmt.Println(Sqrt(100))
// 	fmt.Println(Sqrt(144))
// 	fmt.Println(Sqrt(-4))
// 	fmt.Println(Sqrt(369))
// }
