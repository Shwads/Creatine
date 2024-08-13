package scriptParser

import "fmt"

func printBatches(batches map[string][]string) {
    for key, list := range batches {
        fmt.Printf("%s:\n", key)

        for _, val := range list {
            fmt.Printf("\t- %s\n", val)
        }
    }
}
