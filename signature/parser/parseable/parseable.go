package parseable

import "github.com/egoholic/spec/signature/parser/token"

type Parseable struct {
	rawSig   string
	runes    []rune
	position int
}

func New(rawSig string) *Parseable {
	return &Parseable{rawSig, []rune(rawSig), 0}
}

func (p *Parseable) Runes() []rune {
	return p.runes[p.position:]
}

func (p *Parseable) String() string {
	return p.rawSig
}

func (p *Parseable) StartsWith(td *token.TData) bool {
	tokenRunes := td.Runes()
	if len(tokenRunes) > len(p.runes) {
		return false
	}
	var (
		i int
		r rune
	)
	for i, r = range p.Runes() {
		if r != tokenRunes[i] {
			return false
		}
	}
	p.position = i
	return true
}
