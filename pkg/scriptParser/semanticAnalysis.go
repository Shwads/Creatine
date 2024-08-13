package scriptParser

import (
	"errors"
	"fmt"
	"log"
)

/*
Right now the global node has only one Type of child node (NewBatch)
check the type with a switch to easily add new functionality to scripts.
- Add a settings tag?
- Add individual tags for configuring request behaviour? e.g. context?
*/
func semAn(root *ASTNode) (map[string][]string, error) {
    batches := make(map[string][]string)

    if root.Type != Global {
        log.Printf("getBatches expects the global node")
        return batches, errors.New("getBatches expects the global node")
    }

    for _, child := range root.Children {
        switch child.Type {
        case NewBatch:
            semAnErr := semAnNewBatch(child, batches)
            if semAnErr != nil {
                log.Printf("Encountered error: %s. When processing NewBatch\n", semAnErr)
                return batches, semAnErr
            }
            break

        default:
            errString := fmt.Sprintf("Found node with unexpected type, node: %s\n", child)
            log.Printf(errString)
            return batches, errors.New(errString)
        } 
    }

    return batches, nil
}
