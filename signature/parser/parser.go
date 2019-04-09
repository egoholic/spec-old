package parser

import (
	"github.com/egoholic/spec/signature/parser/parseable"

	"github.com/egoholic/spec/signature"
)

type Parser struct {
	parseable  *parseable.Parseable
	root       signature.Signature
	lastParsed signature.Signature
}

func New(p *parseable.Parseable) *Parser {
	return &Parser{p}
}

// Parses whole raw signature string.
func (p *Parser) Parse() (sig signature.Signature, err error) {
	return
}

// Parses next node in a raw signature string.
func (p *Parser) ParseNext() (sig signature.Signature, err error) {

}

func (p *Parser) Parseable() *parseable.Parseable {
	return p.parseable
}
