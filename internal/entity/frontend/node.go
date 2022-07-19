package frontend

import "github.com/hhandhuan/ku-bbs/internal/model"

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
