package lib

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) eat() Token {
	if p.pos >= len(p.tokens) {
		return Token{EoF, "EoF"}
	}
	token := p.tokens[p.pos]
	p.pos++

	return token
}

func (p *Parser) peek() Token {
	if p.pos >= len(p.tokens) {
		return Token{EoF, "EoF"}
	}
	return p.tokens[p.pos]
}

func (p *Parser) expect(expectedType TokenType, errorMsg string) Token {
	token := p.eat()
	if token.TokenType != expectedType {
		panic(errorMsg)
	}
	return token
}

func (p *Parser) not_eof() bool {
	if p.peek().TokenType != TokenType(EoF) {
		return true
	}
	return false
}

func (p *Parser) Pars() *Node {
	root := &Node{
		Type: DocumentType,
	}

	for p.not_eof() {
		child := p.parseNode()
		if child != nil {
			child.Parent = root
			root.Children = append(root.Children, child)
		}
	}

	return root
}

func (p *Parser) parseNode() *Node {
	tkType := p.peek().TokenType
	switch tkType {
	case OpenTag: // <
		return p.parseElementNode()
	case Literal:
		return p.parseTextNode()
	default:
		p.eat()
		panic("Unkown Node")
	}
}


func 
