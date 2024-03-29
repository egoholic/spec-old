package parser

import (
	"github.com/egoholic/spec/signature/parser/parseable"

	"github.com/egoholic/spec/signature"
)

type Parser struct {
	parseable *parseable.Parseable
	root      signature.Signature
}

func New(p *parseable.Parseable) *Parser {
	return &Parser{p, nil}
}

// Parses whole raw signature string.
func (p *Parser) Parse() (sig signature.Signature, err error) {
	return
}

func (p *Parser) Parseable() *parseable.Parseable {
	return p.parseable
}
