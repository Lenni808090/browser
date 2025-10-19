package main

import (
	"browser/lib"
	"fmt"
)

func printNode(node *lib.Node, depth int) {
	prefix := ""
	for i := 0; i < depth; i++ {
		prefix += "  " // indent
	}

	switch node.Type {
	case lib.ElementType:
		fmt.Printf("%sElement: <%s>\n", prefix, node.Data)
		for _, attr := range node.Attr {
			fmt.Printf("%s  Attr: %s=%q\n", prefix, attr.Key, attr.Value)
		}
	case lib.TextType:
		fmt.Printf("%sText: %q\n", prefix, node.Data)
	case lib.DocumentType:
		fmt.Printf("%sDocument\n", prefix)
	}

	for _, child := range node.Children {
		printNode(child, depth+1)
	}
}

func main() {
	input := `<div class=example>Hello <b>World</b>!</div>`

	lexer := lib.NewLexer(input)
	tokens := lexer.Tokenize()

	fmt.Println("Tokens:")
	for _, tok := range tokens {
		fmt.Printf("Type: %v, Value: %q\n", tok.TokenType, tok.Value)
	}

	parser := lib.NewParser(tokens)
	ast := parser.Pars()

	fmt.Println("\nAST:")
	printNode(ast, 0)
}
