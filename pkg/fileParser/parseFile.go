/*
This is incredibly hacky. Should've just used and tweaked someone elses yaml parser but I wanted to do it myself. Maybes you can write a better one in a future project
but for now this will have to do.
*/

package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var multiRequest bool = false

func printMap(requests map[string]interface{}, indent int) {
	line := ""

	for x := 0; x < indent+1; x++ {
		line = fmt.Sprintf(" %s", line)
	}

	for key, _ := range requests {
		fmt.Printf("%s: ", key)

		if tagString, ok := requests[key].(string); ok {
			fmt.Print(tagString)
		} else if tagList, ok := requests[key].([]string); ok {
			for _, val := range tagList {
				fmt.Printf("%s%s", line, val)
			}
		} else if requestMap, ok := requests[key].(map[string][]string); ok {
			for key, list := range requestMap {
				fmt.Printf("%s%s: \n", line, key)

				for _, item := range list {
					fmt.Printf("%s %s\n", line, item)
				}
			}
		} else if requestMap, ok := requests[key].(map[string]interface{}); ok {
			printMap(requestMap, indent+1)
		}
		fmt.Print("\n")

	}
}

func isTag(s string) bool {
	asRunes := []rune(s)

	for _, rn := range asRunes {
		if rn == ':' {
			return true
		}
	}
	return false
}

func bodyParser(fileScanner *bufio.Scanner, requests map[string]interface{}) error {
	fmt.Printf("\nEntered bodyParser\n\n")
	fileScanner.Scan()

	for {
		line := strings.TrimSpace(fileScanner.Text())
		if bodyString, ok := requests["body"].(string); ok {
			line = strings.TrimSpace(line)
			bodyString = fmt.Sprintf("%s%s", bodyString, line)
			requests["body"] = bodyString
		} else {
			fmt.Printf("improper formatting.  object had type %T\n", requests["body"])
			return errors.New("improperly formatted map object. map['body'] should have type string")
		}

		if !fileScanner.Scan() {
			fmt.Printf("\nExiting bodyParser\n\n")
			return nil
		}
	}
}

func headerParser(fileScanner *bufio.Scanner, requests map[string]interface{}) error {
	fmt.Printf("\nEntered headerParser\n\n")
	var currTag string

	fileScanner.Scan()

	for {
		isTag := false
		line := strings.Split(strings.TrimSpace(fileScanner.Text()), ":")
		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 {
			if line[1] == "" {
				isTag = true
				line = []string{line[0]}
			}
		}

		if line[0] == "body" {
			if len(line) > 1 {
				line[1] = strings.TrimSpace(line[1])
				if line[1] == "|" {
					fmt.Printf("calling bodyParser from headerParser\n")
					requests["body"] = ""
					bodyParser(fileScanner, requests)
					return nil
				} else {
					requests["body"] = line[1]
					return nil
				}
			} else {
				return errors.New("parser expects a body key to have a value, use a pipe | to signify a multi line value")
			}
		}

		if len(line) > 1 {
			fmt.Printf("length of line: %s was >1\n", fileScanner.Text())
			line[1] = strings.TrimSpace(line[1])
			if headerMap, ok := requests["headers"].(map[string][]string); ok {
				headerMap[line[0]] = []string{line[1]}
			} else {
				fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
			}
		} else {
			fmt.Printf("length of line: %s was not >1\n", fileScanner.Text())
			if !isTag {
				fmt.Printf("%s was judged to be not a tag\n", line[0])
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					fmt.Printf("length of requests['headers'] before is %d\n", len(headerMap[currTag]))
					headerMap[currTag] = append(headerMap[currTag], line[0])
					fmt.Printf("length of requests['headers'] after is %d\n", len(headerMap[currTag]))
					fmt.Printf("currTag is %s and its value is: %v\n", currTag, headerMap[currTag])
				} else {
					fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
				}

			} else {
				fmt.Printf("%s was judged to be a tag\n", line[0])
				currTag = line[0]
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					headerMap[currTag] = make([]string, 0)
				} else {
					fmt.Printf("map did not have expected type. Instead it had type %T\n", requests["headers"])
				}

			}
		}
		if !fileScanner.Scan() {
			fmt.Printf("\nExiting headerParser\n\n")
			return nil
		}
	}
}

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
