package piscine

// "fmt"

// func IsNumeric(some string) bool {
// 	// chars:=make([]string,len(some))
// 	for i := 0; i < len(some); i++ {
// 		if some[i] < '0' || some[i] > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

func Any(f func(string) bool, a []string) bool {
	// vals:=make([]string,len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			return true
		}
	}
	return false
}

// func main() {
// 	a1 := []string{"Hello", "how", "are", "you"}
// 	a2 := []string{"This", "is", "4", "you"}

// 	result1 := Any(IsNumeric, a1)
// 	result2 := Any(IsNumeric, a2)

// 	fmt.Println(result1)
// 	fmt.Println(result2)
// }
