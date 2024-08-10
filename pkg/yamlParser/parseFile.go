/*
This is incredibly hacky. I should've just used and tweaked someone elses yaml parser but I wanted to do it myself.
Maybe write a better one with a proper lexer and parser in a future project but for now this will have to do.
*/

/*
The map structure we get at the end of this is:
map:
    request1:
        method: ...
        url: ...

        headers: map[string][]string

        body:
            ....
            ....
    .....

essentially a 'map' version of the request file where:
    - 'request' denotes a map for an individual request
    - string to string mappings for simple tags like 'file', 'console', 'title' or 'method'
    - 'headers' maps strings to lists of strings
    - 'body' maps to a single string
*/

package fileParser

import (
	"bufio"
	"log"
	"os"
)

func ParseFile(fileName string) (map[string]interface{}, error) {
	file, fileOpenErr := os.Open(fileName)
	if fileOpenErr != nil {
		log.Printf("Encountered Error: %s. In function 'Parsefile'.", fileOpenErr)
		return nil, fileOpenErr
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	requests := make(map[string]interface{})

	if fileScanner.Text() == "" {
		fileScanner.Scan()
	}

	mainParserErr := mainParserThread(fileScanner, requests)
	if mainParserErr != nil {
		return nil, mainParserErr
	}

	// printMap(requests, 0)
	return requests, nil

}
