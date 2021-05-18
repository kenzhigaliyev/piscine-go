// package main

// import (
// 	"fmt"
// )

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

// func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}

	val := l.Head
	for val.Next != nil {
		val = val.Next
	}
	val.Next = &NodeL{Data: data}
// }

package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	position := 0
	val := l
	for val != nil {
		if position == pos {
			return val
		}
		position++
		val = val.Next
	}
	return val
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "hello")
// 	ListPushBack(link, "how are")
// 	ListPushBack(link, "you")
// 	ListPushBack(link, 1)

// 	fmt.Println(ListAt(link.Head, 3).Data)
// 	fmt.Println(ListAt(link.Head, 1).Data)
// 	fmt.Println(ListAt(link.Head, 7))
// }
