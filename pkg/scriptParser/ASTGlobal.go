package scriptParser

import (
	"errors"
	"log"
)

// Construct the AST 
// Separate functions help enforce state
func GlobalNode(i *int, node *ASTNode, tokens []ScriptItem) error {
	for *i < len(tokens) {
		switch tokens[*i].Type {
		case FileName:
			errString := "syntax error: files should listed within a batch declaration"
			log.Printf("%s\n", errString)
			return errors.New(errString)

		case Parentheses:
			errString := "syntax error: blocks should not be declared without a batch name identifier"
			log.Printf("%s\n", errString)
			return errors.New(errString)

		case BatchName:
			newNode := &ASTNode{
				Type:     NewBatch,
				Name:     tokens[*i].Val,
				Children: make([]*ASTNode, 0),
			}

			node.Children = append(node.Children, newNode)
			*i++

			contentsNode := &ASTNode{
				Type:     BatchContents,
				Children: make([]*ASTNode, 0),
			}

			newNode.Children = append(newNode.Children, contentsNode)

			treeErr := NewBatchNode(i, contentsNode, tokens)
			if treeErr != nil {
				return treeErr
			}
		}
	}

	return nil
}
