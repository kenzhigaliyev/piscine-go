package piscine

// func IsPrime(i int) bool {
// 	if i%2 == 0 {
// 		return false
// 	}
// 	return true
// }

func Map(f func(int) bool, a []int) []bool {
	val := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		val[i] = f(a[i])
	}
	return val
}

// func main() {
// 	a := []int{1, 2, 3, 4, 5, 6}
// 	result := Map(IsPrime, a)
// 	fmt.Println(result)
// }
