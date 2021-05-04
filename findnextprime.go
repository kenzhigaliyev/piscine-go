// package main

// import "fmt"

package piscine

func FindNextPrime(nb int) int {
	if nb == 1 || nb == 2 || nb <= 0 {
		return 2
	}
	if nb == 3 {
		return 3
	}
	if nb%2 == 0 || nb%3 == 0 {
		return FindNextPrime(nb + 1)
	}
	for i := 5; i*i < nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}

// func main() {
// 	fmt.Println(FindNextPrime(1))
// 	fmt.Println(FindNextPrime(2))
// 	fmt.Println(FindNextPrime(3))
// 	fmt.Println(FindNextPrime(4))
// 	fmt.Println(FindNextPrime(5))
// 	fmt.Println(FindNextPrime(4))
// 	fmt.Println(FindNextPrime(1000009292920))
// 	fmt.Println(FindNextPrime(42))
// 	fmt.Println(FindNextPrime(35))
// 	fmt.Println(FindNextPrime(62))
// }
