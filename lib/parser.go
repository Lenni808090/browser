package lib

import "fmt"

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

func (p *Parser) at() Token {
	if p.pos == len(p.tokens) {
		return Token{EoF, "EoF"}
	}
	return p.tokens[p.pos-1]
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
	tok := p.peek() // look at the next token

	switch tok.TokenType {
	case OpenTag: // <
		return p.parseElementNode()
	case Literal:
		return p.parseTextNode()
	default:
		badTok := p.eat()
		panic(fmt.Sprintf("Unknown Node: Type=%v, Value=%q", badTok.TokenType, badTok.Value))
	}
}

func (p *Parser) parseElementNode() *Node {
	p.eat() //<
	tag := p.expect(Identifier, "tag expected").Value

	var attr []Attr
	for p.peek().TokenType == Identifier {
		key := p.eat().Value
		p.expect(Equals, "equals eypected after value")
		value := p.expect(Identifier, "value needed after key").Value
		attr = append(attr, Attr{Key: key, Value: value})
	}
	node := NewElementNode(tag, attr)

	if p.peek().TokenType == SelfCloseTag {
		p.eat()
		return node
	}

	p.expect(CloseTag, "> expected")

	for {
		tok := p.peek()
		if tok.TokenType == OpenEndTag {
			p.eat()
			endTag := p.expect(Identifier, "tag name expected")
			if endTag.Value != node.Data {
				panic("tags dont match")
			}
			p.expect(CloseTag, "expected > aftzer closing tag")
			break
		}

		child := p.parseNode()
		if child != nil {
			child.Parent = node
			node.Children = append(node.Children, child)
		}
	}

	return node
}

func (p *Parser) parseTextNode() *Node {
	tok := p.eat()
	return &Node{Type: TextType, Data: tok.Value}
}
