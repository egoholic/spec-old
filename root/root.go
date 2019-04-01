package root

import (
	"github.com/egoholic/spec/node"
	"github.com/egoholic/validation"
)

type Root struct {
	node *node.Node
}

func New(name, description string) *Root {
	n := node.New(name, description)
	return &Root{n}
}

func (r *Root) GetNode() *node.Node {
	return r.node
}

func (r *Root) Description() string {
	return r.node.Description()
}

func (r *Root) Validate(value interface{}) (result *validation.Node, ok bool) {
	return r.node.Validate(value)
}

func (r *Root) JSON() string {
	return r.node.JSON()
}

func (r *Root) GolangSignature() string {
	return r.node.GolangSignature()
}
