package scriptParser

import (
	"bufio"
	"io"
	"log"
	"os"
)

// The lexer for our script file. Iterate over lexemes and tokenise them for parsing.
func lexScript(file *os.File) ([]ScriptItem, error) {
    // Keep memory use small in case our script file is particularly long
	buffer := make([]byte, 256)

	fileReader := bufio.NewReader(file)

	tokens := make([]ScriptItem, 0)
	currLexeme := make([]rune, 0)

    // Keep filling the buffer until the whole file has been read
	for {
		numBytesRead, fileReadErr := fileReader.Read(buffer)

		var currRune rune

		for x := 0; x < numBytesRead; x++ {
			currRune = rune(buffer[x])

			if !isDelimiter(currRune) {
				currLexeme = append(currLexeme, currRune)
			} else {
                // When we reach a delimiting character we want to process and tokenise the last lexeme
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

                    // Empty the current lexeme
					currLexeme = make([]rune, 0)
				}

                // We need to check if the current character is a token to be added 
                // to our list before progressing
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
				log.Printf("Encountered error: %s. When attempting to read from file\n", fileReadErr)
				return tokens, fileReadErr
			}

            // We need to make sure we're not left with unprocessed text
            // once our Reader has completed
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

    return tokens, nil
}
