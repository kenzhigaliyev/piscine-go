package piscine

func rot14(char rune) rune {
	if char <= 'Z' && char >= 'M' || char <= 'z' && char >= 'm' {
		return char - 12
	} else if char <= 'M' && char >= 'A' || char <= 'm' && char >= 'a' {
		return char + 14
	} else {
	}
	return char
}

func Rot14(s string) string {
	new_string := ""
	for index, val := range s {
		new_string = new_string + string(rot14(val))
		index = index
	}
	return new_string
}

// func main() {
// 	result := Rot14("Hello How are You")

// 	for _, r := range result {
// 		z01.PrintRune(r)
// 	}
// 	z01.PrintRune('\n')
// }
