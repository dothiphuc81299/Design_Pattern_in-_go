package prototype

// INode ...
type INode interface {
	Clone() INode
	Print(s string)
}

