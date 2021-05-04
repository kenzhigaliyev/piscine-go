// package main

// import "fmt"

package piscine

func FindNextPrime(nb int) int {
	if nb == 1 || nb <= 0 {
		return 2
	}
	for j := 0; j <= 15; j++ {
		result := nb + j
		val := 1
		for i := 2; i <= nb/2; i++ {
			if result%i == 0 {
				val = 0
			}
		}
		if val == 1 {
			return result
		}

	}
	return nb
}

// func main() {
// 	fmt.Println(FindNextPrime(5))
// 	fmt.Println(FindNextPrime(4))
// 	fmt.Println(FindNextPrime(100000000085))
// 	fmt.Println(FindNextPrime(42))
// 	fmt.Println(FindNextPrime(35))
// 	fmt.Println(FindNextPrime(62))
// }
