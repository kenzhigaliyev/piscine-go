package piscine

// import "fmt"

func RecursivePower(nb int, power int) int {
	// result := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		nb = nb * RecursivePower(nb, power-1)
	}
	return nb
}

// func main() {
// 	fmt.Println(RecursivePower(4, 3))
// }
