package scriptParser

type ASTNode struct {
	Type     NodeType
	Name     string
	Children []*ASTNode
}
