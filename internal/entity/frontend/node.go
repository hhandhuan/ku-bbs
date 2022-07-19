package frontend

import "github.com/huhaophp/hblog/internal/model"

type Node struct {
	model.Nodes
}

type Nodes struct {
	List []Node
}

type NodeTree struct {
	Item  Node
	Child Nodes
}
