package predicate

import "github.com/egoholic/spec/node"

func IsString(_ []node.Elem, value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func IsNumber(_ []node.Elem, value interface{}) bool {
	_, ok := value.(string)
	return ok
}
