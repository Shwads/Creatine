package scriptParser

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ParseScript(executeAll bool, name string) error {
	file, fileOpenErr := os.Open(name)
	if fileOpenErr != nil {
		log.Printf("Encountered error: %s when opening script file: %s\n", fileOpenErr, name)
		return fileOpenErr
	}
	defer file.Close()

	buffer := make([]byte, 256)

	fileReader := bufio.NewReader(file)

	tokens := make([]ScriptItem, 0)
	currLexeme := make([]rune, 0)

	for {
		numBytesRead, fileReadErr := fileReader.Read(buffer)

		var currRune rune

		for x := 0; x < numBytesRead; x++ {
			currRune = rune(buffer[x])

			if !isDelimiter(currRune) {
				currLexeme = append(currLexeme, currRune)
			} else {
				if len(currLexeme) > 0 {
					if isFileName(string(currLexeme)) {
						token := ScriptItem{
							Type: FileName,
							Val:  string(currLexeme),
						}

						tokens = append(tokens, token)
					} else {
						token := ScriptItem{
							Type: BatchName,
							Val:  string(currLexeme),
						}

						tokens = append(tokens, token)
					}

					currLexeme = make([]rune, 0)
				}

				if isParentheses(currRune) {
					token := ScriptItem{
						Type: Parentheses,
						Val:  string(currRune),
					}

					tokens = append(tokens, token)
				}
			}
		}

		if fileReadErr != nil {
			if fileReadErr != io.EOF {
				log.Printf("Encountered error: %s. When attempting to read from file %s\n", fileReadErr, name)
				return fileReadErr
			}

			if len(currLexeme) > 0 {
				var token ScriptItem

				if isFileName(string(currLexeme)) {
					token = ScriptItem{
						Type: FileName,
						Val:  string(currLexeme),
					}
				} else {
					token = ScriptItem{
						Type: BatchName,
						Val:  string(currLexeme),
					}
				}

				tokens = append(tokens, token)
			}

			break
		}
	}

	printTokens(tokens)

	return nil
}
