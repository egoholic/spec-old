package parser

import (
	"fmt"
	"strings"

	"github.com/egoholic/serror"
	"github.com/egoholic/spec/signature"
)

var MAP_START_TOKEN = [4]rune{'m', 'a', 'p', '['}

type MapSignature struct {
	keyType                  signature.Signature
	valueType                signature.Signature
	rawKeyType, rawValueType string
}

// Implements signature.Signature interface
func (ms *MapSignature) GolangSignature() string {
	return fmt.Sprintf("%s]%s", ms.rawKeyType, ms.rawValueType)
}

func (ms *MapSignature) Matches(sig signature.Signature) bool {
	return true
}

// \ Implements signature.Signature interface
func ParseMap(sigRdr *strings.Reader) (sig signature.Signature, err error) {
	var (
		r            rune
		sr           rune
		keyType      signature.Signature
		valueType    signature.Signature
		rawKeyType   []rune
		rawValueType []rune
	)
	// check if it starts as a map
	for _, r = range MAP_START_TOKEN {
		sr, _, err = sigRdr.ReadRune()
		if err != nil {
			return
		}
		if r != sr {
			fmt.Printf("parse: %s | %s", r, sr)
			err = serror.New("can't parse a map", "given signature does not represent a map")
			return
		}
	}
	openBracketsCount := 1
	// parses key type which is between '[' and ']'
	for openBracketsCount > 0 {
		r, _, err = sigRdr.ReadRune()
		if err != nil {
			return
		}
		if r == '[' {
			openBracketsCount++
		} else if r == ']' {
			openBracketsCount--
		}
		rawKeyType = append(rawKeyType, r)
	}
	keyParser := New(string(rawKeyType))
	valueParser := New(string(rawValueType))
	keyType, err = keyParser.Parse()
	valueType, err = valueParser.Parse()
	sig = &MapSignature{keyType, valueType, string(rawKeyType), string(rawValueType)}

	return
}
