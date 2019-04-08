package token

type TData struct {
	token    string
	tokenLen int
	runes    []rune
}

func New(token string) *TData {
	runes := []rune(token)
	return &TData{token, len(runes), runes}
}

func (t *TData) String() string {
	return t.token
}

func (t *TData) Len() int {
	return t.tokenLen
}

func (t *TData) Runes() []rune {
	return t.runes
}
