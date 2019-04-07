package parser

import (
	"strings"

	"github.com/egoholic/spec/signature"
)

type Parser struct {
	rawSig        string
	reader        *strings.Reader
	cursorParsed  int
	cursorCurrent int
}

func New(rawSig string) *Parser {
	return &Parser{rawSig, strings.NewReader(rawSig), 0, 0}
}

func (p *Parser) Parse() (sig signature.Signature, err error) {
	return
}
