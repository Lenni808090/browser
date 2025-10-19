package lexer_html

import (
	"strings"
	"unicode"
)

type Lexer struct {
	chars []rune
	pos   int
}

func NewLexer(s string) *Lexer {
	return &Lexer{
		chars: []rune(s),
		pos:   0,
	}
}

func (l *Lexer) next() (rune, bool) {
	if l.pos >= len(l.chars) {
		return 0, true
	}
	ch := l.chars[l.pos]
	l.pos++
	return ch, false
}

func (l *Lexer) peek() (rune, bool) {
	if l.pos >= len(l.chars) {
		return 0, true
	}
	ch := l.chars[l.pos]
	return ch, false

}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	var insideTag bool
	for {
		ch, eof := l.next()
		if eof {
			break
		}
		switch ch {
		case ' ', '\t', '\n', '\r':
			continue
		case '<':
			ch, _ := l.peek()
			if ch == '/' {
				l.next()
				tokens = append(tokens, token(OpenEndTag, "</"))
			} else {
				tokens = append(tokens, token(OpenTag, "<"))
			}
			insideTag = true
		case '>':
			tokens = append(tokens, token(CloseTag, ">"))
			insideTag = false
		case '/':
			ch, _ := l.peek()
			if ch == '>' {
				l.next()
				tokens = append(tokens, token(SelfCloseTag, "/>"))
				insideTag = false
			}
		case '"':
			str := l.consumeQuotedString()
			tokens = append(tokens, token(Identifier, str))
		case '=':
			tokens = append(tokens, token(Equals, "="))
		default:
			if insideTag {
				word := l.consumeWord(ch)
				tokens = append(tokens, token(Identifier, word))
			} else {
				text := l.consumeText(ch)
				tokens = append(tokens, token(Literal, text))
			}

		}
	}
	tokens = append(tokens, token(EoF, "EoF"))
	return tokens
}

func (l *Lexer) consumeQuotedString() string {
	var sb strings.Builder
	for {
		ch, eof := l.next()
		if eof {
			break
		}
		if ch == '"' {
			break
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (l *Lexer) consumeWord(firstChar rune) string {
	var word strings.Builder
	word.WriteRune(firstChar)
	for {
		ch, eof := l.peek()
		if eof || !(unicode.IsDigit(ch) || unicode.IsLetter(ch) || ch == '-' || ch == '_') {
			break
		}
		l.next()
		word.WriteRune(ch)
	}

	return word.String()
}

func (l *Lexer) consumeText(firstChar rune) string {
	var text strings.Builder
	text.WriteRune(firstChar)
	for {
		ch, eof := l.peek()
		if eof || ch == '<' {
			break
		}
		l.next()
		text.WriteRune(ch)
	}
	return text.String()
}
