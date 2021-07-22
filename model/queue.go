package model

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue {
	return &Queue{nil, nil}
}
