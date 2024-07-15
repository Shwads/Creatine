package fileParser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"
)

func headerParser(fileScanner *bufio.Scanner, requests map[string]interface{}) (bool, error) {
	var currTag string

	fmt.Printf("\nSkipping the line: %s\n\n", fileScanner.Text())
	fileScanner.Scan()

	for {
		isTag := false
		line := strings.Split(strings.TrimSpace(fileScanner.Text()), ":")
		//fmt.Printf("Hi from headerParser the line is: %s\n", fileScanner.Text())
		line[0] = strings.TrimSpace(line[0])

		if len(line) > 1 && len(line) < 3 {
			asRunes := []rune(line[0])
			if line[1] == "" && asRunes[0] != '-' {
				isTag = true
				line = []string{line[0]}
			}
		}

		if line[0] == "body" {
			if len(line) > 1 {
				line[1] = strings.TrimSpace(line[1])
				if line[1] == "|" {
					requests["body"] = ""
					anotherRequest, bodyParserErr := bodyParser(fileScanner, requests)
					if bodyParserErr != nil {
						return false, bodyParserErr
					}

					return anotherRequest, nil
				} else {
					requests["body"] = line[1]
					return false, nil
				}
			} else {
				return false, errors.New("parser expects a body key to have a value, use a pipe | to signify a multi line value")
			}
		}

		if len(line) > 1 {
			line[1] = strings.Join(line[1:], ":")
			line[1] = strings.TrimSpace(line[1])

			if headerMap, ok := requests["headers"].(map[string][]string); ok {
				headerMap[line[0]] = []string{line[1]}
			} else {
				log.Printf("where line > 1. requests['headers'] did not have expected type. Instead it had type %T\n", requests["headers"])
				return false, errors.New(fmt.Sprintf("map did not have expected type instead it had type %T\n", requests["headers"]))
			}

		} else {
			if !isTag {
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					headerMap[currTag] = append(headerMap[currTag], line[0])
				} else {
					log.Printf("where not a tag. requests['headers'] did not have expected type. Instead it had type %T\n", requests["headers"])
					return false, errors.New(fmt.Sprintf("map did not have expected type instead it had type %T\n", requests["headers"]))
				}

			} else {
				currTag = line[0]
				if headerMap, ok := requests["headers"].(map[string][]string); ok {
					headerMap[currTag] = make([]string, 0)
				} else {
					log.Printf("where setting tag. requests['headers'] did not have expected type. Instead it had type %T\n", requests["headers"])
					return false, errors.New(fmt.Sprintf("map did not have expected type. Instead it had type %T\n", requests["headers"]))
				}

			}
		}

		if !fileScanner.Scan() {
			return false, nil
		}
	}
}
