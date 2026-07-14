package design_browser_history

type Node struct {
	Value string
	Next  *Node
	Prev  *Node
}

type BrowserHistory struct {
	Current *Node
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		Current: &Node{
			Value: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &Node{
		Value: url,
		Prev:  this.Current,
	}

	this.Current.Next = newNode

	this.Current = newNode
}

func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.Current.Prev != nil {
		this.Current = this.Current.Prev
		steps--
	}

	return this.Current.Value
}

func (this *BrowserHistory) Forward(steps int) string {
	for steps > 0 && this.Current.Next != nil {
		this.Current = this.Current.Next
		steps--
	}

	return this.Current.Value
}
