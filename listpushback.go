package piscine

func ListPushBack(l *List, data interface{}) {
	if l.Head == nil {
		l.Head = &NodeL{Data: data}
		return
	}

	val := l.Head

	for val.Next != nil {
		val = val.Next
	}
	val.Next = &NodeL{Data: data}
}

// func main() {

// 	link := &List{}

// 	ListPushBack(link, "Hello")
// 	ListPushBack(link, "man")
// 	ListPushBack(link, "how are you")

// 	for link.Head != nil {
// 		fmt.Println(link.Head.Data)
// 		link.Head = link.Head.Next
// 	}
// }
