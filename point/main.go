package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	line := "x = a, y = b"
	for _, letter := range line {
		if letter == 'a' {
			// for points.x != 4 {
			z01.PrintRune(rune((points.x-(points.x%10))/10 + 48))
			z01.PrintRune(rune((points.x % 10) + 48))
			// }
		} else if letter == 'b' {
			// for points.y != 2 {
			z01.PrintRune(rune((points.y-(points.y%10))/10 + 48))
			z01.PrintRune(rune((points.y % 10) + 48))

			// }
		} else {
			z01.PrintRune(letter)
		}
		// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
	}
	z01.PrintRune('\n')
}
