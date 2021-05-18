// package main

package piscine

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

func ListSize(l *List) int {
	point := l.Head
	counter := 0
	for point != nil {
		counter++
		point = point.Next
	}
	return counter
}

// func ListPushFront(l *List, data interface{}) {
// 	val := &NodeL{Data: data}
// 	if l.Head == nil {
// 		l.Head = val
// 		return
// 	}
// 	val.Next = l.Head
// 	l.Head = val
// }

// func main() {
// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "you")
// 	ListPushFront(link, "man")

// 	fmt.Println(ListSize(link))
// }
