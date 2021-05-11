package piscine

// "fmt"
// "piscine"

func ForEach(f func(int), a []int) {
	// numbers := make([]int, 6)
	for i := 0; i < len(a); i++ {
		PrintNbr(a[i])
	}
	// fmt.Print("\n")
}

// func PrintNbr(i int) {
// 	fmt.Print(i)
// 	// return i
// }

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	ForEach(PrintNbr, a)
// }
