package parser

import (
	"fmt"

	"github.com/egoholic/spec/signature"
)

type StructSignature struct {
	name  *string
	fieds map[string]signature.Signature
}

func ParseStruct() (sig signature.Signature, err error) {

}

// Implements signature.Signature interface
func (ss *StructSignature) GolangSignature() string {
	return fmt.Sprintf("struct %s {}", *ss.name)
}

func (ss *StructSignature) Matches(sig signature.Signature) bool {
	return true
}

// \ Implements signature.Signature interface
