package parser

import (
	"fmt"

	"github.com/egoholic/serror"

	"github.com/egoholic/spec/signature"
)

var primitives []*tokenData

type tokenData struct {
	token    string
	tokenLen int
	runes    []rune
}

func extendWorldWithType(token string) {
	runes := []rune(token)
	td := &tokenData{token, len(runes), runes}
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

func (td *tokenData) Parse(rawSig string) (sig signature.Signature, err error) {
	var (
		i  int
		r  rune
		sr rune
	)
	var runes = []rune(rawSig)
	for i, r = range td.runes {
		sr = runes[i]
		if r != sr {
			err = serror.New(fmt.Sprintf("can't parse token: `%s`", td.token), fmt.Sprintf("rune `%s` != rune `%s`", r, sr))
			return
		}
	}
	sig = &PrimitiveSignature{td}
	return
}

type PrimitiveSignature struct {
	*tokenData
}

// Implements signature.Signature interface
func (ps *PrimitiveSignature) GolangSignature() string {
	return ps.tokenData.token
}

func (ps *PrimitiveSignature) Matches(sig signature.Signature) bool {
	_, ok := sig.(*PrimitiveSignature)
	if !ok {
		return false
	}
	return ps.GolangSignature() == sig.GolangSignature()
}

func (ps *PrimitiveSignature) Append(sig signature.Signature) (err error) {
	return serror.New(fmt.Sprintf("can't append signature `%s`", sig.GolangSignature()), fmt.Sprintf("primitive signature `%s` can't be extended!", ps.tokenData.token))
}

// \ Implements signature.Signature interface

func ParsePrimitive(rawSig string) (sig signature.Signature, err error) {
	var td *tokenData
	for _, td = range primitives {
		sig, err = td.Parse(rawSig)
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
