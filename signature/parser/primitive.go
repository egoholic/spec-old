package parser

import (
	"fmt"

	"github.com/egoholic/serror"

	"github.com/egoholic/spec/signature"
	"github.com/egoholic/spec/signature/parser/token"
)

var primitives []*token.TData

func extendWorldWithType(tokenStr string) {
	td := token.New(tokenStr)
	primitives = append(primitives, td)
}

func init() {
	// shorter types with the same prefixes should be first.
	// For example:
	//   int64 should be before int.
	extendWorldWithType("string")
	extendWorldWithType("int64")
	extendWorldWithType("int32")
	extendWorldWithType("int")
	extendWorldWithType("float64")
	extendWorldWithType("float")
	extendWorldWithType("bool")
}

func parse(td *token.TData, rawSig string) (sig signature.Signature, err error) {
	var (
		i  int
		r  rune
		sr rune
	)
	var runes = []rune(rawSig)
	for i, r = range td.Runes() {
		sr = runes[i]
		if r != sr {
			err = serror.New(fmt.Sprintf("can't parse token: `%s`", td.String()), fmt.Sprintf("rune `%s` != rune `%s`", r, sr))
			return
		}
	}
	sig = &PrimitiveSignature{td.String()}
	return
}

type PrimitiveSignature struct {
	valueType string
}

// Implements signature.Signature interface
func (ps *PrimitiveSignature) GolangSignature() string {
	return ps.valueType
}

func (ps *PrimitiveSignature) Matches(sig signature.Signature) bool {
	_, ok := sig.(*PrimitiveSignature)
	if !ok {
		return false
	}
	return ps.GolangSignature() == sig.GolangSignature()
}

// \ Implements signature.Signature interface

func ParsePrimitive(rawSig string) (sig signature.Signature, err error) {
	var td *token.TData
	for _, td = range primitives {
		sig, err = parse(td, rawSig)
		if err != nil {
			continue
		}
		if sig != nil {
			return
		}
	}
	err = serror.New("can't parse primitive signature", "given signature is not a signature at all or not a primitive signature")
	return
}
