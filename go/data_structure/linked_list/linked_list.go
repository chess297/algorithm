package linked_list

type Node struct {
	e    any
	next *Node
}

func NewNode(e any, next Node) *Node {
	return &Node{
		e: e, next: &next,
	}
}

type LinkedList struct {
	dummyHead *Node
	size      int
}

func NewLinkedList(head *Node, size int) LinkedList {
	return LinkedList{
		dummyHead: &Node{
			e:    nil,
			next: head,
		}, size: size,
	}
}

func (l *LinkedList) Add(index int, e int) {
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &Node{
		e:    e,
		next: prev.next,
	}
	l.size++
}

func (l *LinkedList) AddFirst(e int) {
	l.Add(0, e)
}

func (l *LinkedList) AddLast(e int) {
	l.Add(l.size-1, e)
}

func (l *LinkedList) Get(e int) *Node {
	for cur := l.dummyHead.next; cur.next != nil; cur = cur.next {
		if cur.e == e {
			return cur
		}
	}
	return nil
}

func (l *LinkedList) Remove(e int) {

}
