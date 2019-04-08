package signature

type Signature interface {
	GolangSignature() string
	Matches(Signature) bool
}
