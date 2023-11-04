package cache

type HistoryItem struct {
	Value string
	Next  *HistoryItem
	Prev  *HistoryItem
}

type HistoryList struct {
	Head *HistoryItem
	Tail *HistoryItem
}

func (ll *HistoryList) Append(value string) {
	newNode := &HistoryItem{Value: value, Next: nil, Prev: ll.Tail}
	if ll.Tail != nil {
		ll.Tail.Next = newNode
	}
	ll.Tail = newNode
	if ll.Head == nil {
		ll.Head = newNode
	}
}

func (ll *HistoryList) PopFront() (string, bool) {
	if ll.Head == nil {
		return "", false
	}
	value := ll.Head.Value
	ll.Head = ll.Head.Next
	if ll.Head == nil {
		ll.Tail = nil
	} else {
		ll.Head.Prev = nil
	}
	return value, true
}

func (ll *HistoryList) MoveToEnd(node *HistoryItem) {
	if node == ll.Tail {
		return
	}
	if node == ll.Head {
		ll.Head = node.Next
		ll.Head.Prev = nil
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
	}
	ll.Tail.Next = node
	node.Prev = ll.Tail
	node.Next = nil
	ll.Tail = node
}

func (ll *HistoryList) Find(value string) (*HistoryItem, bool) {
	current := ll.Head
	for current != nil {
		if current.Value == value {
			return current, true
		}
		current = current.Next
	}
	return nil, false
}

func (ll *HistoryList) Remove(item *HistoryItem) {
	current := ll.Head
	for current != nil {
		if current.Value == item.Value {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				ll.Head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				ll.Tail = current.Prev
			}
			return
		}
		current = current.Next
	}
}
