package parser

import (
	"fmt"

	"github.com/egoholic/serror"
	"github.com/egoholic/spec/signature"
	"github.com/egoholic/spec/signature/parser/token"
)

var sliceStartToken = token.New("[]")

type SliceSignature struct {
	valueType signature.Signature
}

func (parser *Parser) ParseSlice() (sig signature.Signature, err error) {
	parseable := parser.Parseable()
	if !parseable.StartsWith(sliceStartToken) {
		err = serror.New(fmt.Sprintf("can't parse slice signature: `%s`", parseable.String()), "signature starts not as a slice signature")
		return
	}

	sig = &SliceSignature{nil}
	parser.ParseNext()
}

func (s *SliceSignature) GolangSignature() string {
	return "[]" + s.valueType.GolangSignature()
}

func (s *SliceSignature) Matches(signature.Signature) bool {
	return true
}
