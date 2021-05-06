// package main

package piscine

func TrimAtoi(s string) int {
	values := []rune(s)
	// var newarr[] int
	a := 0
	b := 0
	for index := range s {
		if a == 0 && values[index] == '-' {
			b = 1
		}
		if values[index] >= '0' && values[index] <= '9' {
			a = (a + int(values[index]-48)) * 10
		}
	}
	if b == 1 {
		a = -a / 10
	} else {
		a = a / 10
	}
	return a
}

// func main() {
// 	fmt.Println(TrimAtoi("12345"))
// 	fmt.Println(TrimAtoi("str123ing45"))
// 	fmt.Println(TrimAtoi("012 345"))
// 	fmt.Println(TrimAtoi("Hello World!"))
// 	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
// 	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
// }
