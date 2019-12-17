package func_5_1

import (
	"io"
)

type Node struct {
	Type						NodeType
	Data						string
	Attr						[]Attribute
	FirstChild, NextSiblint		*Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Value string
}

func Parse(r io.Reader) (*Node, error)

