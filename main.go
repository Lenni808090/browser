package main

import (
	"browser/lib"
	"fmt"
)

func main() {
	input := `<div class=example>Hello <b>World</b>!</div>`

	lexer := lib.NewLexer(input)
	tokens := lexer.Tokenize()

	for _, tok := range tokens {
		fmt.Printf("Type: %v, Value: %q\n", tok.TokenType, tok.Value)
	}
}
