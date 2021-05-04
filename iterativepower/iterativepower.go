// package main

package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return result
	} else {
		for i := 1; i <= power; i++ {
			result = result * nb
		}
		return result
	}
}

// func main() {
// 	fmt.Println(IterativePower(4, 3))
// }
