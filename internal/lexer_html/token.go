package lexer_html

type TokenType int

const (
	OpenTag TokenType = iota
	CloseTag
	OpenEndTag
	SelfCloseTag
	Identifier
	Literal
	Equals
	Quotes
	EoF
)

func (t TokenType) String() string {
	switch t {
	case OpenTag:
		return "OpenTag"
	case CloseTag:
		return "CloseTag"
	case OpenEndTag:
		return "OpenEndTag"
	case SelfCloseTag:
		return "SelfCloseTag"
	case Identifier:
		return "Identifier"
	case Equals:
		return "Equals"
	case Quotes:
		return "Quotes"
	case Literal:
		return "Literal"
	case EoF:
		return "EoF"
	default:
		return "Unkown"
	}
}

type Token struct {
	TokenType TokenType
	Value     string
}

func token(tokenType TokenType, value string) Token {
	return Token{
		TokenType: tokenType,
		Value:     value,
	}
}
