// package main
package piscine

// import "fmt"

func IsPrime(nb int) bool {
	if nb == 1 && nb <= 0 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrime(5))
// 	fmt.Println(IsPrime(4))
// 	fmt.Println(IsPrime(7))
// 	fmt.Println(IsPrime(15))
// }
