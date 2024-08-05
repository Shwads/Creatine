package scriptParser

import (
	"bufio"
	"log"
	"os"
)

func ParseScript(nonIdempotent bool, name string) error {
    file, fileOpenErr := os.Open(name)
    if fileOpenErr != nil {
        log.Printf("Encountered error: %s when opening script file: %s\n", fileOpenErr, name)
        return fileOpenErr
    }

    fileScanner := bufio.NewScanner(file)

    for {

        if !fileScanner.Scan() {
            return nil
        }
    }
}
