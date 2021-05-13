package main

import (
	"fmt"
	"os"
)

func main() {
	values := os.Args[1:]
	check := []string{"01", "galaxy", "galaxy 01"}
	for _, val := range values {
		for _, str := range check {
			if val == str {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
