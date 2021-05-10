package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println("File name is missing")
		return
	} else if len(arguments) >= 2 {
		fmt.Println("Too many arguments")
		return
	}
	file, err := ioutil.ReadFile("quest8.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(file.Stat())

	// arr := make([]byte, 14)

	// file.Read(arr)

	fmt.Print(string(file))
}
