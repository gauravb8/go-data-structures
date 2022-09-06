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
func (list *DLL) addNode(newNode *DLLNode) {
	if list.front == nil && list.rear == nil {
		// empty list
		list.front = newNode
		list.rear = newNode
	} else {
		newNode.prev = list.rear
		list.rear.next = newNode
		list.rear = newNode
	}
}

// Removes node from the list
func (list *DLL) removeNode(node *DLLNode) {
	if node == list.front && node == list.rear {
		// single node in the list
		list.front = nil
		list.rear = nil
	} else if node == list.front {
		list.front = list.front.next
	} else if node == list.rear {
		list.rear = list.rear.prev
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}
