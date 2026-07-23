package clone_graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	hashMap := make(map[*Node]*Node)

	var cloneNode func(node *Node) *Node
	cloneNode = func(node *Node) *Node {
		if _, exist := hashMap[node]; exist {
			return hashMap[node]
		}

		newNode := &Node{Val: node.Val}
		hashMap[node] = newNode

		for _, neighbor := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, cloneNode(neighbor))
		}

		return newNode
	}

	return cloneNode(node)
}
