package piscine

func CollatzCountdown(start int) int {
	counter := 0
	if start <= 0 {
		return -1
	}
	for start != 1 {
		if start%2 == 0 {
			start = start / 2
			counter++
		} else {
			start = (start * 3) + 1
			counter++
		}
	}
	return counter
}

// func main() {
// 	steps := CollatzCountdown(9999999999999)
// 	fmt.Println(steps)
// }
