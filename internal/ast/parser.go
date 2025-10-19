package ast

import (
	lx "browser/internal/lexer_html"
	"fmt"
)

type Parser struct {
	tokens []lx.Token
	pos    int
}

func NewParser(tokens []lx.Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) eat() lx.Token {
	if p.pos >= len(p.tokens) {
		return lx.Token{TokenType: lx.EoF, Value: "EoF"}
	}
	token := p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) at() lx.Token {
	if p.pos == len(p.tokens) {
		return lx.Token{TokenType: lx.EoF, Value: "EoF"}
	}
	return p.tokens[p.pos-1]
}

func (p *Parser) peek() lx.Token {
	if p.pos >= len(p.tokens) {
		return lx.Token{TokenType: lx.EoF, Value: "EoF"}
	}
	return p.tokens[p.pos]
}

func (p *Parser) expect(expectedType lx.TokenType, errorMsg string) lx.Token {
	token := p.eat()
	if token.TokenType != expectedType {
		panic(errorMsg)
	}
	return token
}

func (p *Parser) not_eof() bool {
	return p.peek().TokenType != lx.TokenType(lx.EoF)
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
	tok := p.peek()

	switch tok.TokenType {
	case lx.OpenTag:
		return p.parseElementNode()
	case lx.Literal:
		return p.parseTextNode()
	default:
		badTok := p.eat()
		panic(fmt.Sprintf("Unknown Node: Type=%v, Value=%q", badTok.TokenType, badTok.Value))
	}
}

func (p *Parser) parseElementNode() *Node {
	p.eat() //<

	tag := p.expect(lx.Identifier, "tag expected").Value

	var attr []Attr
	for p.peek().TokenType == lx.Identifier {
		key := p.eat().Value
		p.expect(lx.Equals, "equals expected after key")
		value := p.expect(lx.Identifier, "value needed after key").Value
		attr = append(attr, Attr{Key: key, Value: value})
	}

	node := NewElementNode(tag, attr)

	if p.peek().TokenType == lx.SelfCloseTag {
		p.eat()
		return node
	}

	p.expect(lx.CloseTag, "> expected")

	for {
		tok := p.peek()
		if tok.TokenType == lx.OpenEndTag {
			p.eat()
			endTag := p.expect(lx.Identifier, "tag name expected")
			if endTag.Value != node.Data {
				panic("tags don't match")
			}
			p.expect(lx.CloseTag, "expected > after closing tag")
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
