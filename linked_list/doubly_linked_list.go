package linked_list

type DLLNode struct {
	val  int
	next *DLLNode
	prev *DLLNode
}

type DLL struct {
	front *DLLNode
	rear  *DLLNode
}

// Adds node to the rear of the list
func (list *DLL) addNode(value int) *DLLNode {
	newNode := &DLLNode{
		val: value,
	}
	if list.front == nil {
		list.front = newNode
		list.rear = newNode
	} else {
		list.rear.next = newNode
		list.rear = newNode
	}

	return newNode
}

// Removes node from the list
func (list *DLL) removeNode(node *DLLNode) {
	if node == list.front {
		list.front = list.front.next
	} else if node == list.rear {
		list.rear = list.rear.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}
