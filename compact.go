// package main

package piscine

// import "fmt"

const N = 6

func Compact(ptr *[]string) int {
	counter := 0
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] == "" || (*ptr)[i] == " " {
			for j := i; j < len(*ptr)-1; j++ {
				(*ptr)[j] = (*ptr)[j+1]
			}
			counter++
			// *ptr = (*ptr)[:len(*ptr)-1]
		} else {
			// counter++
		}
	}
	// for i := 0; i < len(*ptr); i++ {
	// 	if (*ptr)[i] == " " {
	// 		(*ptr)[i] = ""
	// 	}
	// }
	(*ptr) = (*ptr)[:len(*ptr)-counter]
	return len(*ptr)
}

// func main() {
// 	a := make([]string, N)
// 	a[0] = "aasd"
// 	a[2] = "ba sd   "
// 	a[4] = "casd"

// 	for _, v := range a {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("Size after compacting:", Compact(&a))
// 	fmt.Println(len(a))
// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
