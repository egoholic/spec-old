package node

import (
	"github.com/egoholic/validation"
)

type Elem interface {
	GetNode() *Node
}

type Node struct {
	name       string
	decription string
	predicates []Predicate
	children   map[string]*Node
}

type Predicate func(ancestors []Elem, value interface{}) bool

func New(name, description string) *Node {
	predicates := make([]Predicate, 10)
	children := map[string]*Node{}
	return &Node{name, description, predicates, children}
}

func (n *Node) GetNode() *Node {
	return n
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Description() string {
	return n.decription
}

func (n *Node) Validate(value interface{}) (result *validation.Node, ok bool) {
	return nil, true
}

func (n *Node) JSON() string {
	return ""
}

func (n *Node) GolangSignature() string {
	return ""
}
