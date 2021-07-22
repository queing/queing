package model

type Node struct {
	ID    string `json:"id"`
	Order int    `json:"order"`
	Prev  *Node  `json:"-"`
}

func NewNode(id string) *Node {
	return &Node{id, -1, nil}
}
