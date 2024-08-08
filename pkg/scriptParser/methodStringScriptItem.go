package scriptParser

import "fmt"

func (item ScriptItem) String() string {
    var typ string

    switch item.Type {
    case FileName:
        typ = "FileName"
        break
    case BatchName:
        typ = "BatchName"
        break
    case Parentheses:
        typ = "Parentheses"
        break
    }

    return fmt.Sprintf("{ Type: %s, Val: %s }", typ, item.Val)
}
