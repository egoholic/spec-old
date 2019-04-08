package parser

import (
	"fmt"

	"github.com/egoholic/serror"
	"github.com/egoholic/spec/signature"
	"github.com/egoholic/spec/signature/parser/parseable"
	"github.com/egoholic/spec/signature/parser/token"
)

var sliceStartToken = token.New("[]")

type SliceSignature struct {
	valueType signature.Signature
}

func ParseSlice(rawSig string) (sig signature.Signature, err error) {
	p := parseable.New(rawSig)
	if !p.StartsWith(sliceStartToken) {
		err = serror.New(fmt.Sprintf("can't parse slice signature: `%s`", rawSig), "signature starts not as a slice signature")
		return
	}

	valueTypeParser := New(p)
}

func (s *SliceSignature) GolangSignature() string {
	return "[]" + s.valueType.GolangSignature()
}

func (s *SliceSignature) Matches(signature.Signature) bool {
	return true
}
