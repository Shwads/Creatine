/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
*/

package fileParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseFile(fileName string) (map[string]interface{}, error) {
	file, fileOpenErr := os.Open(fileName)
	if fileOpenErr != nil {
		log.Printf("Encountered Error: %s. In function 'Parsefile'.", fileOpenErr)
		return nil, fileOpenErr
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var requestTag string
	requestNum := 0

	requests := make(map[string]interface{})

	if fileScanner.Text() == "" {
		fileScanner.Scan()
	}

    mainParserErr := mainParserThread(fileScanner, requests)
    if mainParserErr != nil {
        return nil, mainParserErr
    }
    
    fmt.Print("\n\n\n")
    printMap(requests, 0)
    return requests, nil

	var currLineProcessed bool

	for {
		currLineProcessed = false

		line := strings.Split(fileScanner.Text(), ":")

		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 && len(line) < 3 {
			if line[1] == "" {
				line = []string{line[0]}
			}
		}

		if len(line) > 1 {
			line[1] = strings.TrimSpace(line[1])
			if multiRequest {
				//fmt.Printf("Trying to index map with requestTag: %s\n", requestTag)
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					//fmt.Printf("Creating key: %s with value: %s\n", line[0], line[1])
					line[1] = strings.Join(line[1:], ":")
					requestMap[line[0]] = line[1]
				} else {
					fmt.Printf("map did not have expected type. Actual type was %T\n", requests[requestTag])
				}
			}
		}

		switch line[0] {
		case "request":
			requestNum += 1
			multiRequest = true
			requestTag = fmt.Sprintf("request-%d", requestNum)
			requests[requestTag] = make(map[string]interface{})
			currLineProcessed = true
			break

		case "headers":
			if multiRequest {
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					requestMap["headers"] = make(map[string][]string)
					anotherRequest, headerParseErr := headerParser(fileScanner, requestMap)
					if headerParseErr != nil {
						return nil, headerParseErr
					}
					currLineProcessed = !anotherRequest
				}
			} else {
				requests["headers"] = make(map[string][]string)
				anotherRequest, headerParseErr := headerParser(fileScanner, requests)
				if headerParseErr != nil {
					return nil, headerParseErr
				}
				currLineProcessed = !anotherRequest
			}
			break

		case "body":
			if multiRequest {
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					requestMap["body"] = ""
					_, bodyParseErr := bodyParser(fileScanner, requestMap)
					if bodyParseErr != nil {
						return nil, bodyParseErr
					}
					currLineProcessed = false
				}
			} else {
				requests["body"] = ""
				_, bodyParseErr := bodyParser(fileScanner, requests)
				if bodyParseErr != nil {
					return nil, bodyParseErr
				}
				currLineProcessed = false
			}
			break
		}

		if currLineProcessed {
			if !fileScanner.Scan() {
				//fmt.Printf("\n\n\nPRINTING THE MAP\n\n\n")
				printMap(requests, 0)
				return requests, nil
			}
		}
	}
}
