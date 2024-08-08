package scriptParser

import "fmt"

func printTokens(tokens []ScriptItem) {
    for _, token := range tokens {
        fmt.Printf("%s\n", token)
    }
}
