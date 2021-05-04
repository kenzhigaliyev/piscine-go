// package main

// import "fmt"
package piscine

func FindNextPrime(nb int) int {
	if nb > 0 {
		return 2
	}
	if nb == 1 || nb == 0 {
		nb = FindNextPrime(nb + 1)
	} else {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 {
				nb = FindNextPrime(nb + 1)
			}
		}
	}
	return nb
}

// func main() {
// 	fmt.Println(FindNextPrime(5))
// 	fmt.Println(FindNextPrime(4))
// 	fmt.Println(FindNextPrime(99))
// 	fmt.Println(FindNextPrime(42))
// 	fmt.Println(FindNextPrime(35))
// 	fmt.Println(FindNextPrime(62))
// }