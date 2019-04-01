package spec

import (
	"github.com/egoholic/spec/root"
)

func New(name, description string) *root.Root {
	return root.New(name, description)
}
