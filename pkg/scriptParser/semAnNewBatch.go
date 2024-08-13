package scriptParser

import (
	"errors"
	"fmt"
	"log"
)

/*
Nodes with Type NewBatch currently only have one child node of type BatchContents
check with a switch to easily extend funcionality and add new tags to scripts 
- Add global headers? Headers that will be added to every request in the batch
- Add the ability to update the url for every request in the batch for changing endpoints?
*/
func semAnNewBatch(root *ASTNode, batches map[string][]string) error {
	for _, child := range root.Children {
		switch child.Type {
		case BatchContents:
			if _, ok := batches[root.Name]; ok {
				errString := fmt.Sprintf("syntax error: %s redeclared in this scope.", root.Name)
				log.Printf("%s When processing new batch.\n", errString)
				return errors.New(errString)
			}
			batches[root.Name] = make([]string, 0)
			contentsErr := semAnBatchContents(child, batches, root.Name)
			if contentsErr != nil {
				return contentsErr
			}
			break

		default:
			errString := fmt.Sprintf("Found node with unexpected type, node: %s\n", child)
			log.Printf(errString)
			return errors.New(errString)
		}
	}
	return nil
}
