package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	hashMap := make(map[*Node]*Node)

	current := head
	for current != nil {
		hashMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	current = head
	for current != nil {
		hashMap[current].Next = hashMap[current.Next]
		hashMap[current].Random = hashMap[current.Random]
		current = current.Next
	}

	return hashMap[head]
}
