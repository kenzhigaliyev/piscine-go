package main

import "github.com/01-edu/z01"

type Door struct {
	state int
}

const (
	CLOSE = 0
	OPEN  = 1
)

func PrintStr(s string) {
	// runes := []rune(s)
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...\n")
	ptrDoor.state = CLOSE
	// return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	return ptrDoor.state == CLOSE
}

func main() {
	// CLOSE := 0
	// OPEN := 0
	// door := &Door{}
	var door Door
	OpenDoor(&door)
	if IsDoorClose(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
