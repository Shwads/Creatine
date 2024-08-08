package scriptParser

import (
	"errors"
	"log"
)

func NewBatchNode(i *int, node *ASTNode, tokens []ScriptItem) error {
	if tokens[*i].Val != "(" {
		errString := "syntax error, expected opening parentheses after batch declaration"
		log.Printf("%s\n", errString)
		return errors.New(errString)
	}

	*i++

	for *i < len(tokens) {
		switch tokens[*i].Type {
		case FileName:
			newNode := &ASTNode{
				Type: File,
				Name: tokens[*i].Val,
			}
			node.Children = append(node.Children, newNode)
			*i++
			break

		case BatchName:
			newNode := &ASTNode{
				Type: Batch,
				Name: tokens[*i].Val,
			}
			node.Children = append(node.Children, newNode)
			*i++
			break

		case Parentheses:
			if tokens[*i].Val == ")" {
				*i++
				return nil
			}
			errString := "syntax error: found '(' expected ')'"
			log.Printf("%s\n", errString)
			return errors.New(errString)
		}
	}

	errString := "syntax err: expected a closing brace ')'"
	log.Printf("%s\n", errString)
	return errors.New(errString)
}
