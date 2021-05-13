package piscine

const N = 6

func Compact(ptr *[]string) int {
	// var test[] int
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] == string("") {
			for j := i; j < len(*ptr)-1; j++ {
				(*ptr)[j] = (*ptr)[j+1]
			}
			*ptr = (*ptr)[:len(*ptr)-1]
		}
	}
	return len(*ptr)
}

// func main() {
// 	a := make([]string, N)
// 	a[0] = "a"
// 	a[2] = "b"
// 	a[4] = "c"

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("Size after compacting:", Compact(&a))

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
