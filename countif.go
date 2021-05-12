// package main

package piscine

// "piscine"

// func IsNumeric(some string) bool {
// 	num := false
// 	for i := 0; i < len(some); i++ {
// 		if some[i] >= '0' && some[i] <= '9' {
// 			num = true
// 		}
// 	}
// 	return num
// }

func CountIf(f func(string) bool, tab []string) int {
	// slicing:=make([]string,len(tab))
	counter := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			counter++
		}
	}
	return counter
}

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := CountIf(IsNumeric, tab1)
// 	answer2 := CountIf(IsNumeric, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
