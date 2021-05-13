package piscine

func Abort(a, b, c, d, e int) []int {
	new_array := []int{a, b, c, d, e}
	// sort.Ints(new_array)
	var arr int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if new_array[j] > new_array[i] {
				arr = new_array[j]
				new_array[j] = new_array[i]
				new_array[i] = arr
			}
		}
	}
	return new_array
}

// func main() {
// 	middle := Abort(-8301353824951690532, 2626072374859070259, 6993143751468609912, 1580539742955549786, -8821429897402619750)
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(middle[i])
// 	}
// }
