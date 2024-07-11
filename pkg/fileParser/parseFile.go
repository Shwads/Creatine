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

	fileScanner.Scan()

	for {
		//fmt.Printf("%s\n", fileScanner.Text())

		line := strings.Split(fileScanner.Text(), ":")

		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 {
			if line[1] == "" {
				line = []string{line[0]}
			}
		}

		if len(line) > 1 {
			line[1] = strings.TrimSpace(line[1])
			if multiRequest {
				fmt.Printf("Trying to index map with requestTag: %s\n", requestTag)
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					fmt.Printf("Creating key: %s with value: %s\n", line[0], line[1])
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
			//fmt.Printf("requestTag is %s\n", requestTag)
			//fmt.Printf("length of split line is %d.\n", len(line))
			requests[requestTag] = make(map[string]interface{})
			//fmt.Printf("Type of requests['requestTag'] is %T", requests[requestTag])
			break

		case "headers":
			if multiRequest {
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					requestMap["headers"] = make(map[string][]string)
					headerParseErr := headerParser(fileScanner, requestMap)
					if headerParseErr != nil {
						return nil, headerParseErr
					}
				}
			} else {
				requests["headers"] = make(map[string][]string)
				headerParseErr := headerParser(fileScanner, requests)
				if headerParseErr != nil {
					return nil, headerParseErr
				}
			}
			break

		case "body":
			if multiRequest {
				if requestMap, ok := requests[requestTag].(map[string]interface{}); ok {
					requestMap["body"] = ""
					bodyParseErr := bodyParser(fileScanner, requestMap)
					if bodyParseErr != nil {
						return nil, bodyParseErr
					}
				}
			} else {
				requests["body"] = ""
				bodyParseErr := bodyParser(fileScanner, requests)
				if bodyParseErr != nil {
					return nil, bodyParseErr
				}
			}
			break
		}

		if !fileScanner.Scan() {
			fmt.Printf("\n\n\nPRINTING THE MAP\n\n\n")
			printMap(requests, 0)
			return requests, nil
		}
	}
}
