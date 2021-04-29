package main

import (
	"fmt"
)

func main() {
	for i := 97; i <= 123; i++ {
		if i > 122 {
			fmt.Println()
		} else {
			fmt.Print(string(i))
		}
	}

}
