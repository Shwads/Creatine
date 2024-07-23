/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
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
