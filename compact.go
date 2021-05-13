// package main
package piscine

// import "fmt"

const N = 6

func Compact(ptr *[]string) int {
	// var test[] int
	counter := 0
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] == "" {
			counter++
			for j := i; j < len(*ptr)-1; j++ {
				(*ptr)[j] = (*ptr)[j+1]
			}
			*ptr = (*ptr)[:len(*ptr)-1]
		}
	}
	// for i:=0;i<len(*ptr);i++{
	*ptr = (*ptr)[:counter]
	// }

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
// 	// fmt.Println(len(a))
// 	for _, v := range a {
// 		fmt.Println(v)
// 	}
// }
