package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	val := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = val
		return
	}
	val.Next = l.Head
	l.Head = val
}

// func main() {

// 	link := &List{}

// 	ListPushFront(link, "Hello")
// 	ListPushFront(link, "man")
// 	ListPushFront(link, "how are you")
// 	ListPushFront(link, "1")
// 	ListPushFront(link, "2")
// 	ListPushFront(link, "3")

// 	it := link.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " ")
// 		it = it.Next
// 	}
// 	fmt.Println()
// }
