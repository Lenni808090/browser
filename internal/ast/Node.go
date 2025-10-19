package ast

type Node struct {
	Type     NodeType
	Data     string
	Attr     []Attr
	Children []*Node
	Parent   *Node
}

type Attr struct {
	Key   string
	Value string
}

type NodeType int

const (
	DocumentType NodeType = iota
	ElementType
	TextType
)

func (t NodeType) String() string {
	switch t {
	case DocumentType:
		return "DocumentType"
	case ElementType:
		return "ElementType"
	case TextType:
		return "TextType"
	default:
		return "Unkown"
	}
}

func NewElementNode(tag string, attributes []Attr) *Node {
	return &Node{
		Type: ElementType,
		Data: tag,
		Attr: attributes,
	}
}

func NewTextNode(text string) *Node {
	return &Node{
		Data: text,
		Type: TextType,
	}
}
