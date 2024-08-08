package scriptParser

import "fmt"

func (node ASTNode) String() string {
    var typ string

    switch node.Type {
    case Global:
        typ = "Global"
        break
    case NewBatch:
        typ = "NewBatch"
        break
    case BatchContents:
        typ = "BatchContents"
        break
    case Batch:
        typ = "Batch"
        break
    case File:
        typ = "File"
        break
    }

    return fmt.Sprintf("{ NodeType: %s\tNodeVal: %s }", typ, node.Name)
}
