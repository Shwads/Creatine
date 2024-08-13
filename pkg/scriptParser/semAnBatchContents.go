package scriptParser

import (
	"errors"
	"fmt"
	"log"
)

func semAnBatchContents(root *ASTNode, batches map[string][]string, batchName string) error {
	for _, child := range root.Children {
		switch child.Type {
		case File:
			batches[batchName] = append(batches[batchName], child.Name)
			break
		case Batch:
			if contents, ok := batches[child.Name]; ok {
				batches[batchName] = append(batches[batchName], contents...)
			} else {
				errString := fmt.Sprintf("syntax error: batch %s is not declared in this scope", batchName)
				log.Printf(errString)
				return errors.New(errString)
			}
			break
		default:
			errString := fmt.Sprintf("syntax error: batch %s should not contain: %s", batchName, child)
			log.Println(errString)
			return errors.New(errString)
		}
	}
	return nil
}
