package parser

import (
	"github.com/egoholic/spec/signature/parser/parseable"

	"github.com/egoholic/spec/signature"
)

type Parser struct {
	parseable *parseable.Parseable
}

func New(p *parseable.Parseable) *Parser {
	return &Parser{p}
}

func (p *Parser) Parse() (sig signature.Signature, err error) {
	return
}
