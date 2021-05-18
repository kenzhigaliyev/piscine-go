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
// 	if l.Head == nil {
// 		l.Head = &NodeL{Data: data}
// 		return
// 	}
// 	val := l.Head
// 	for val.Next != nil {
// 		val = val.Next
// 	}
// 	val.Next = &NodeL{Data: data}
// }

// func PrintList(l *List) {
// 	link := l.Head
// 	for link != nil {
// 		fmt.Print(link.Data, " -> ")
// 		link = link.Next
// 	}
// 	fmt.Println(nil)
// }
package piscine

func ListClear(l *List) {
	l.Head = nil
}

// func main() {
// 	link := &List{}

// 	ListPushBack(link, "I")
// 	ListPushBack(link, 1)
// 	ListPushBack(link, "something")
// 	ListPushBack(link, 2)

// 	fmt.Println("------list------")
// 	PrintList(link)
// 	ListClear(link)
// 	fmt.Println("------updated list------")
// 	PrintList(link)
// }
