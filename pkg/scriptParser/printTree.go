package scriptParser

import "fmt"

func (tree *ASTNode) printTree(indent int) {
	for x := 0; x < indent; x++ {
		fmt.Print("\t")
	}

	fmt.Printf("%s\n", *tree)

	for _, child := range (*tree).Children {
		child.printTree(indent + 1)
	}
}
